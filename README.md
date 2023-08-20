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
- [x] In "add" command check if payload is empty
- [x] In response for "add" command return a new memo id
- [x] In "rand" command with memo also send it's id
- [x] Add "help" command. Which will return all available commands with descriptions.
- [x] Add "delete" command. Which will delete the memo with a given id.
- [x] Add env config for shutdowning polling bot
- [ ] Add "update" command. Which will update the memo with a given id.
- [ ] Add "get" command. Which will return the memo with a given id.
- [ ] Add visibility for memos. Private memos only for owner, public will be used for subscriptions.
- [ ] Add subscriptions. Subscriber will receive notifications about new publisher's memos. Publisher's memos will be used in daily bot postings.
- [ ] Link internal users with telegram accounts via id not name
- [ ] Store images localy, if there are urls to png or jpg files in new memos.