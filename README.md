# Discord PR - Pull Request Notifier Bot

A Discord bot that sends notifications for opened pull requests.

## Usage

```bash
git clone https://github.com/xySaad/discord-pr
```

```bash
cd discord-pr
```

```bash
touch .env
```

fill the environement variables in .env file ([example](.example.env))

```bash
go mod tidy
```

```bash
go run main.go
```
## Webhook Setup

To connect GitHub with the bot:

1. Go to your repository on GitHub.
2. Navigate to **Settings > Webhooks**.
3. Click **Add webhook**.
4. In **Payload URL**, enter your server’s public URL and webhook path. For example:

   ```
   http://<your-server-ip>:8080/webhook
   ```
5. Set **Content type** to `application/json`.
6. In **Secret**, enter the same value you used in your `.env` file under `WEBHOOK_SECRET`.
7. Select the event type:

   * Choose **Let me select individual events** and check **Pull requests**.
8. Click **Add webhook**.

Now, whenever a pull request is opened, your bot will post a notification in your configured Discord forum channel.

## Features

#### - Opens a forum thread for every pull request

![Bot Screenshot](images/thread_overview.png)

#### - Provides a simple overview message about the pull request

![Bot Screenshot](images/pr_overview.png)

#### - Allows filtering using tags: `open` / `closed` or branch name

![Bot Screenshot](images/filter_by_tag.png)
