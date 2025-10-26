import { Octokit } from "@octokit/rest";
import { query } from "@anthropic-ai/claude-agent-sdk";
import fs from "fs";
import path from "path";

const githubToken = process.env.GITHUB_TOKEN!;
const claudeApiKey = process.env.CLAUDE_API_KEY!;
const octokit = new Octokit({ auth: githubToken });
const repo = process.env.GITHUB_REPOSITORY!;
const [owner, repoName] = repo.split("/");
const prNumber = process.env.PR_NUMBER || process.env.GITHUB_REF?.split("/").pop();

async function run() {
  if (!prNumber) throw new Error("PR number not found");

  // 1. diffを取得
  const { data: files } = await octokit.pulls.listFiles({
    owner,
    repo: repoName,
    pull_number: Number(prNumber),
  });

  const diffs = files
    .map((f) => `### ${f.filename}\n\n\`\`\`${f.patch}\`\`\``)
    .join("\n\n");

  // 2. specリポジトリ内の仕様書を読み込む
  const specDir = path.resolve("specs/specs");
  const specFiles = fs
    .readdirSync(specDir)
    .filter((f) => f.endsWith(".md"))
    .map((f) => fs.readFileSync(path.join(specDir, f), "utf-8"))
    .join("\n\n");

  // 3. Claudeにレビュー依頼
  const reviewPrompt = `
You are a professional code reviewer.
Refer to the following specification when reviewing:

--- SPEC START ---
${specFiles}
--- SPEC END ---

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
      for (const block of message.message) {
        if (block.type === "text") {
          reviewText += block.text;
        }
      }
    } else if (message.type === "result") {
      // 最終結果メッセージ
      console.log("Review completed");
    }
  }

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