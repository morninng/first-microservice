import { Octokit } from "@octokit/rest";
import { query } from "@anthropic-ai/claude-agent-sdk";

async function getPRNumber(octokit: Octokit, owner: string, repoName: string): Promise<number> {
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

async function getPullRequestDiff(octokit: Octokit, owner: string, repoName: string, prNumber: number): Promise<string> {
  const { data: files } = await octokit.pulls.listFiles({
    owner,
    repo: repoName,
    pull_number: Number(prNumber),
  });
  const diffs = files
    .map((f) => `### ${f.filename}\n\n\`\`\`${f.patch}\`\`\``)
    .join("\n\n");

    return diffs;
}


async function getClaudeReviewComment(octokit: Octokit, diffs: string) {

  // 3. Claudeにレビュー依頼
  const reviewPrompt = `
You are a professional code reviewer.

Now review the following pull request diff and point out problems, improvements:

${diffs}
`;

  const stream = query({prompt: reviewPrompt});

  let reviewText = "";

  // ストリームからメッセージを取得
  for await (const message of stream) {
      console.log("message", message)
    if (message.type === "assistant") {
      // アシスタントのレスポンスを結合
      const content = message.message?.content;
      if (Array.isArray(content)) {
        for (const block of content) {
          if (block.type === "text") {
            reviewText += block.text;
          }
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
  return reviewText;
}

async function writeCommentOnPullRequest(octokit: Octokit, owner: string, repoName: string, prNumber: number, reviewText: string) {
  await octokit.issues.createComment({
    owner,
    repo: repoName,
    issue_number: Number(prNumber),
    body: `🤖 **Claude PR Review**\n\n${reviewText}`,
  });
}





export async function runLocalTest() {
  const githubToken = process.env.GITHUB_TOKEN;
  if (!githubToken) {
    throw new Error("GITHUB_TOKEN environment variable is required");
  }

  const repo = process.env.GITHUB_REPOSITORY;
  if (!repo) {
    throw new Error("GITHUB_REPOSITORY environment variable is required");
  }

  const [owner, repoName] = repo.split("/");
  const octokit = new Octokit({ auth: githubToken });

  const prNumber = await getPRNumber(octokit , owner, repoName )
  const diff = await getPullRequestDiff(octokit , owner, repoName, prNumber)

 const reviewComment = await getClaudeReviewComment(octokit, diff)
 console.log("------------- ")
 console.log(reviewComment)

}

async function runGithubActionReview() {

  const githubToken = process.env.GITHUB_TOKEN!;
  const repo = process.env.GITHUB_REPOSITORY!;
  const [owner, repoName] = repo.split("/");
  const octokit = new Octokit({ auth: githubToken });

  const prNumber = await getPRNumber(octokit , owner, repoName )
  const diff = await getPullRequestDiff(octokit , owner, repoName, prNumber)
  const reviewComment = await getClaudeReviewComment(octokit, diff)
  await writeCommentOnPullRequest(octokit,  owner, repoName, prNumber, reviewComment )

}


// Only run if this file is executed directly (not imported)
if (import.meta.url === `file://${process.argv[1]}`) {
  runGithubActionReview().catch((e: Error) => {
    console.error(e);
    process.exit(1);
  });
}
