import dotenv from 'dotenv';
import { runLocalTest } from './pr_review_bot';

// Load environment variables from .env file
dotenv.config();

async function main() {
  console.log('Starting PR Review Bot from localhost...\n');

  // Validate required environment variables
  const requiredEnvVars = [
    'GITHUB_TOKEN',
    'GITHUB_REPOSITORY',
    'ANTHROPIC_API_KEY',
  ];

  const missingVars = requiredEnvVars.filter(varName => !process.env[varName]);

  if (missingVars.length > 0) {
    console.error('Missing required environment variables:');
    missingVars.forEach(varName => console.error(`   - ${varName}`));
    console.error(
      'Please create a .env file in the scripts directory with these variables.'
    );
    console.error('See .env.example for reference.\n');
    process.exit(1);
  }

  if (!process.env.PR_NUMBER) {
    console.warn(
      'PR_NUMBER not set. Will attempt to auto-detect from GITHUB_REF.\n'
    );
  }

  try {
    await runLocalTest();
    console.log('\n Successfully completed PR review!');
  } catch (error) {
    console.error('\n Error running PR review:');
    console.error(error);
    process.exit(1);
  }
}

main();
