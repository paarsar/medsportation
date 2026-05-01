# Medsportation Backend (Go Cloud Function)

This directory contains the Go source code for the "Request a Quote" backend, deployed as a serverless Google Cloud Function.

## Architecture
- **Language:** Go 1.22+
- **Infrastructure:** Google Cloud Functions (2nd Gen)
- **Protocol:** SMTP (Standard Email protocol)
- **Compatibility:** Works with GoDaddy (Microsoft 365), Gmail (App Passwords), and other SMTP providers.

## Configuration

The backend uses environment variables for security. Create a `.env` file for local testing (see `.env.example`).

| Variable | Description | Example (GoDaddy/M365) |
| :--- | :--- | :--- |
| `SMTP_HOST` | The address of the mail server | `smtp.office365.com` |
| `SMTP_PORT` | The server port (usually TLS) | `587` |
| `SMTP_EMAIL` | The account sending the mail | `info@yourdomain.com` |
| `SMTP_PASSWORD` | The account password (or App Password) | `your-secure-password` |
| `NOTIFICATION_RECIPIENT` | Where the quote requests are sent | `quotes@medsportation.com` |

### Note on GoDaddy / Microsoft 365
If your GoDaddy email is powered by Microsoft 365:
1. **MFA Enabled:** You **must** generate an "App Password" from your Microsoft security dashboard and use it as your `SMTP_PASSWORD`.
2. **SMTP AUTH:** Microsoft often disables SMTP by default. Ensure "Authenticated SMTP" is enabled in the Microsoft 365 Admin Center for that specific user mailbox.

## Local Testing

1.  **Start the server:**
    ```bash
    cd medsportation-be
    go run cmd/main.go
    ```
2.  **The endpoint will be available at:** `http://localhost:8080/RequestQuote`
3.  **Frontend Integration:** Ensure your Angular `QuoteService` points to this local URL during development.

## Deployment

To deploy the function to Google Cloud, run the following command from this directory:

```bash
gcloud functions deploy RequestQuote \
  --gen2 \
  --runtime=go122 \
  --region=us-central1 \
  --source=. \
  --entry-point=RequestQuote \
  --trigger-http \
  --allow-unauthenticated \
  --project 4420256990 \
  --set-env-vars SMTP_HOST=smtp.office365.com,SMTP_PORT=587,SMTP_EMAIL=your-email@yourdomain.com,SMTP_PASSWORD=your-password,NOTIFICATION_RECIPIENT=quotes@medsportation.com
```

### Post-Deployment
1.  Copy the **URL** provided by the `gcloud` command.
2.  Update `medsportation-wb/src/app/services/quote.ts` with this URL.
3.  Redeploy the frontend via `make deploy`.
