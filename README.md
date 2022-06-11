# Memo

An application for storing and retrieving memos, definitions, etc.

Access to the application functionality will be realized via telegram.

SQLite will be used as storage, Gin as http web framework .

## Development plan

- [x] Set up Gin
- [x] Add linting
- [x] Init database structure
- [x] Create telegram bot
- [x] Add the ability to add new memos
- [x] Add "rand" command
- [x] Add cron for polling_bot
- [x] Add the ability to receive random memos on daily basis
- [x] Link memos to telegram accounts
- [x] Add the ability to search memos by keywords
- [ ] Add visibility for memos. Private memos only for owner, public will be used for subscriptions.
- [ ] Add subscriptions. Subscriber will receive notifications about new publisher's memos. Publisher's memos will be used in daily bot postings.
- [ ] Link internal users with telegram accounts via id not name