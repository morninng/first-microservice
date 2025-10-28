# PR Review Bot Scripts

This directory contains scripts for automated PR reviews using Claude AI.

## Setup

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Create a `.env` file:**
   Copy the example file and fill in your credentials:
   ```bash
   cp .env.example .env
   ```

3. **Configure environment variables in `.env`:**
   - `GITHUB_TOKEN`: Your GitHub Personal Access Token
     - Create at: https://github.com/settings/tokens
     - Required scopes: `repo` (Full control of private repositories)

   - `GITHUB_REPOSITORY`: Your repository in format `owner/repo`
     - Example: `morninng/first-microservice`

   - `PR_NUMBER`: (Optional) The PR number to review
     - If not provided, will auto-detect from `GITHUB_REF`

   - `ANTHROPIC_API_KEY`: Your Anthropic API key
     - Get at: https://console.anthropic.com/

## Usage

### Running from localhost

```bash
npm start
```

Or directly:
```bash
npx tsx index.ts
```

### Running directly (without index.ts)

```bash
npx tsx pr_review_bot.ts
```

## Files

- `index.ts`: Main entry point that loads environment variables and calls the PR review bot
- `pr_review_bot.ts`: Core logic for fetching PR diffs and posting Claude reviews
- `.env`: Your local environment variables (git-ignored)
- `.env.example`: Template for required environment variables

## GitHub Actions

The bot is also configured to run in GitHub Actions. See `.github/workflows/` for the workflow configuration.
