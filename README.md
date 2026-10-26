# Google Chat Translator (Abang Translator)

[![Github Action](https://github.com/dyaskur/google-chat-translator/workflows/Go/badge.svg?branch=master)](https://github.com/dyaskur/google-chat-translator/actions)
[![Scrutinizer Quality Score](https://scrutinizer-ci.com/g/dyaskur/google-chat-translator/badges/quality-score.png?s=4023c984fc1163a44f4220cd7d57406643ced9f2)](https://scrutinizer-ci.com/g/dyaskur/google-chat-translator/)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=dyaskur_google-chat-translator&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=dyaskur_google-chat-translator)

Google Chat Translator is an add-on for Google Chat that enables users to seamlessly translate text between languages directly within the platform.

I initially created this application as a project to learn Golang. It has been installed more than 140,000 times.

I'm constantly looking for ways to improve my code and enhance the app. Your feedback and suggestions are always welcome!

For anyone who is interested, you can install my application using the following URL:
https://workspace.google.com/marketplace/app/abang_translator/963775668283

<a href="https://workspace.google.com/marketplace/app/abang_translator/963775668283" target="_blank" aria-label="Get it from the Google Workspace Marketplace">
  <img alt="Google Workspace Marketplace badge" alt-text="Get it from the Google Workspace Marketplace" src="https://workspace.google.com/static/img/marketplace/en/gwmBadge.svg?" style="height: 68px">
</a>


# How to deploy
Pushing to `main` deploys the function through the `Deploy` GitHub Actions workflow
(`.github/workflows/deploy.yml`). It can also be run by hand from the Actions tab.
All project-specific values are repository secrets, and the workflow hides gcloud
output, so the logs only show whether the deploy succeeded or failed.

## Repository secrets

| Secret | Required | Value |
| --- | --- | --- |
| `GCP_PROJECT_ID` | yes | Google Cloud project ID |
| `GCP_REGION` | yes | Function region, e.g. `asia-southeast1` |
| `GCP_FUNCTION_NAME` | yes | Cloud Function name |
| `REDIS_URI` | yes | Redis connection URI |
| `MYMEMORY_EMAIL` | no | Email sent to MyMemory for the higher daily quota |
| `GOOGLE_MONTHLY_CHAR_LIMIT` | no | Monthly Google Translate character budget |
| `GCP_WORKLOAD_IDENTITY_PROVIDER` | one auth option | Full provider name, `projects/NUMBER/locations/global/workloadIdentityPools/POOL/providers/PROVIDER` |
| `GCP_SERVICE_ACCOUNT` | with the option above | Deployer service account email |
| `GCP_CREDENTIALS_JSON` | other auth option | JSON key of the deployer service account |

Use either Workload Identity Federation (the first two auth secrets, no long-lived key) or
`GCP_CREDENTIALS_JSON`, not both.

## Google Cloud setup

1. Create a deployer service account and grant it `Cloud Functions Admin`, `Cloud Run Admin`
   and `Service Account User`.
2. For Workload Identity Federation, enable the IAM Service Account Credentials API, then:
   ```sh
   gcloud iam workload-identity-pools create github --location=global
   gcloud iam workload-identity-pools providers create-oidc github \
     --location=global --workload-identity-pool=github \
     --issuer-uri=https://token.actions.githubusercontent.com \
     --attribute-mapping=google.subject=assertion.sub,attribute.repository=assertion.repository \
     --attribute-condition="assertion.repository=='dyaskur/google-chat-translator'"
   gcloud iam service-accounts add-iam-policy-binding DEPLOYER_SA_EMAIL \
     --role=roles/iam.workloadIdentityUser \
     --member="principalSet://iam.googleapis.com/projects/PROJECT_NUMBER/locations/global/workloadIdentityPools/github/attribute.repository/dyaskur/google-chat-translator"
   ```
   Otherwise create a JSON key for the service account and store it as `GCP_CREDENTIALS_JSON`.
3. The function's runtime service account still needs the `Cloud Translation API User` role.

`deploy.sh.example` remains for deploying from a local machine.
