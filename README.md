# Solution for the asessment quiz below

Used MongoDb for storing the users, and a golang map for Caching.

## Testing

After cloning the  repo, install dependences by runnung `go get`
Then run `make test` this will start a testing server, populate mongodb database with some dumb user data, then it will start sending requests to the server, After finishing it will show how many request were sent and how many request did hit the database.

## Development

For development environment, just start the server with `make run`. This will start a localhost server at 4096, you can access user information in this endpoint `http://localhost:4096/?id=2` id being any number between 1 and 150.
