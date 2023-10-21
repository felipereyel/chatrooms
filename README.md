# chatrooms

# How to Run

Use docker compose to run all required services:

```sh
    docker compose up
```

## Contents
This command will run the following containers:

- `Postgres`: to persist application data (users, rooms, ...)
- `RabbitMQ`: to broker messages between servers and bots.
- `ChatroomsServer`: to handle the main application logic (register, posts, ...)
- `ChatroomsBot`: to process commands decoupled from main application logic

## Notes
This application was developed with scaling of its individual components in mind:

- Commication between frontend client and server for real-time messaging uses
  the `pub sub` pattern. In this case there may be multiple instances of the
  server running (for scaling purposes) and the message broker will relay 
  realtime messages to those subscribed to that rooms topic.

- Command processing uses the `competing consumers` pattern. In this case there
  thera may be multiple instances of bots running (for scaling purposes) and
  only one bot will respond to the command. The bots also have to acknowledge
  that the command was processed otherwise the message will return to the queue
  so other bots can process it. This ensures that if a bot goes down for any
  reason the message will not be lost.

## Testing

Tests were written for the controllers, utils and modes. Run them with:

```sh
  go test ./... -v
```

Note that the tests for the user controller take about 4 seconds.   
That is a feature of `bcrypt` to prevent brute forcing passwords.