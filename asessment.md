# Intro

Create a library for caching user information to increase database throughput
and availability.
Every request for user data should return the user information while simultaneously taking care of data throughput.

It means the library should take care of unnecesary requests for user data if some are pending.

Why:

Let's say we have a bunch of user information (identified by id) inside our main database.
Simultaneously there is a thousands requests per second for user information.
We can consider that the database is a bottleneck right.To avoid that we need a cache mechanism.

Notes:

You can use any language, library, database and framework - anything you think is best for this case.

Example:

If within 1000 requests there are 100 unique user ids then there should be only a maximum 100 requests into the database but all 1000 requests should get a response with a user data.
