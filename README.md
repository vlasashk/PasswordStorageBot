# Telegram password manager BOT
## Build
#### Prerequisites
- docker

1. Clone project:
```
git clone https://github.com/vlasashk/PasswordStorageBot.git
cd PasswordStorageBot
```
2. Input your Telegram bot token in `docker-compose.yml` file, in field `TELEGRAM_BOT_TOKEN`.
3. Run bot:
```
docker compose up --build
```
<sup>*first build might take a while (1-2 minutes)*</sup>
## Project information
Telegram bot to store and manage passwords for different services
#### Functionality
- `/start` - used to initialize bot and display buttons for interaction.
- `/help` - used to display information about available commands.
- `/set` - used to add new service and it's credentials for user.
- `/get` - used to display available credentials for chosen service.
- `/del` - used to remove service and it's credentials from database.

User can interact with the bot using buttons or commands
#### Limitations
- Each user is allowed to have no more than 20 services per account.
- Each service is allowed to have one set of credentials (login:password).
- Max length of service, login or password must be no longer than 50 characters.
#### Storage
- Gorm as ORM
- SQLite3 as database to store information locally
#### Security
- All messages with credentials are getting deleted from the chat after 1-minute delay.
- AES-256 Symmetric encryption for passwords.