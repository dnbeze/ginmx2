# ginmx2

ginmx2 is a simple note taking application utilizing Gin and SQLite


## Current functions

- Take notes
- Follow notes
- Delete an existing note
- Update an existing note

## Starting the project
**MakeFile will soon be available**

To start the Gin server, navigate into the projects root directory and run `go run main.go`

The SQLite Database will automatically be created after a few seconds.

Note: Many endpoints require JWT authentication, use the `/signup` endpoint to create an account, and `/login` to login.