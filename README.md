# Memo

An application for storing and retrieving memos, definitions, etc.

Access to the application functionality will be realized via telegram.

SQLite will be used as storage, Gin as http web framework .

## Development plan

- [x] Set up Gin
- [x] Add linting
- [x] Init database structure
- [x] Create telegram bot
- [x] Add an ability to add new memos
- [x] Add "rand" command
- [ ] Add an ability to receive random memos on daily basis
- [ ] Link memos to telegram accounts
- [ ] Add an ability to search a memo by keywords
- [ ] Add visibility for memos. Private memos only for owner, public will be used for subscriptions.
- [ ] Add subscriptions. Subscriber will receive notifications about new publisher's memos. Publisher's memos will be used in daily bot postings.
