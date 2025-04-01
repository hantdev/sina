# SINA Email Agent

SINA Email Agent is used for sending emails. It wraps basic SMTP features and
provides a simple API that SINA services can use to send email notifications.

## Configuration

SINA Email Agent is configured using the following configuration parameters:

| Parameter                           | Description                                                             |
| ----------------------------------- | ----------------------------------------------------------------------- |
| SINA_EMAIL_HOST                       | Mail server host                                                        |
| SINA_EMAIL_PORT                       | Mail server port                                                        |
| SINA_EMAIL_USERNAME                   | Mail server username                                                    |
| SINA_EMAIL_PASSWORD                   | Mail server password                                                    |
| SINA_EMAIL_FROM_ADDRESS               | Email "from" address                                                    |
| SINA_EMAIL_FROM_NAME                  | Email "from" name                                                       |
| SINA_EMAIL_TEMPLATE                   | Email template for sending notification emails                          |

There are two authentication methods supported: Basic Auth and CRAM-MD5.
If `SINA_EMAIL_USERNAME` is empty, no authentication will be used.
