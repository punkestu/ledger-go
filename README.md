# Ledger App in Terminal

With this app, you can push and track you expense with google spreadsheet integration from you terminal.

## How to use it?
1. You need to create a desktop OAuth2 client secret from your google console and save it in this directory with name `client-secret.json`. Then build the app with `make all` or you can build it like you build other golang projects.
2. Get your ledger spreadsheet with column ID, timestamp, kredit, debit, note, and total
3. Take the spreadsheet id and save it in your environtment variable with name LEDGER_KEY
4. Then when you open the app at the first time, it will be show a OAuth2 url so you can login with your google account.
5. After you login, you will see a url in the url bar in your browser. Somethings like `http://localhost/?state=state-token&code=<code>....`. You can take the code and copy it into the app prompt.