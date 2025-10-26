import { Octokit } from "@octokit/rest";
import { query } from "@anthropic-ai/claude-agent-sdk";
import fs from "fs";
import path from "path";

const githubToken = process.env.GITHUB_TOKEN!;
const claudeApiKey = process.env.CLAUDE_API_KEY!;
const octokit = new Octokit({ auth: githubToken });
const repo = process.env.GITHUB_REPOSITORY!;
const [owner, repoName] = repo.split("/");

async function getPRNumber(): Promise<number> {
  // If PR_NUMBER is explicitly provided, use it
  if (process.env.PR_NUMBER) {
    console.log("Using provided PR_NUMBER:", process.env.PR_NUMBER);
    return Number(process.env.PR_NUMBER);
  }

  // Otherwise, find PR by current branch
  const branchRef = process.env.GITHUB_REF || "";
  const branchName = branchRef.replace("refs/heads/", "");
  console.log("Auto-detecting PR for branch:", branchName);

  const { data: pulls } = await octokit.pulls.list({
    owner,
    repo: repoName,
    state: "open",
    head: `${owner}:${branchName}`,
  });

  if (pulls.length === 0) {
    throw new Error(`No open PR found for branch: ${branchName}`);
  }

  console.log("Found PR #", pulls[0].number);
  return pulls[0].number;
}

async function run() {
  const prNumber = await getPRNumber();

  // 1. diffを取得
  const { data: files } = await octokit.pulls.listFiles({
    owner,
    repo: repoName,
    pull_number: Number(prNumber),
  });

  const diffs = files
    .map((f) => `### ${f.filename}\n\n\`\`\`${f.patch}\`\`\``)
    .join("\n\n");


  // 3. Claudeにレビュー依頼
  const reviewPrompt = `
You are a professional code reviewer.
Refer to the following specification when reviewing:

Now review the following pull request diff and point out problems, improvements, and spec mismatches:

${diffs}
`;

  // 最新のAgent SDKのquery()関数を使用
  const stream = query({prompt: reviewPrompt});

  let reviewText = "";

  // ストリームからメッセージを取得
  for await (const message of stream) {
    if (message.type === "assistant") {
      // アシスタントのレスポンスを結合
      console.log("message.message", message.message)
      reviewText += message.message
      // for (const block of message.message) {
      //   if (block.type === "text") {
      //     reviewText += block.text;
      //   }
      // }
    } else if (message.type === "result") {
      // 最終結果メッセージ
      console.log("Review completed");
    }
  }
  console.log("reviewText", reviewText)
  if (!reviewText) {
    reviewText = "No comments generated.";
  }

  // 4. GitHubにコメント投稿
  await octokit.issues.createComment({
    owner,
    repo: repoName,
    issue_number: Number(prNumber),
    body: `🤖 **Claude PR Review**\n\n${reviewText}`,
  });

  console.log("✅ PR review comment posted.");
}

run().catch((e) => {
  console.error(e);
  process.exit(1);
});
