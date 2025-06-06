# AWS SQS Example App for Meetup

## Building the application

You can build the `producer` application with the following command.

```bash
go build ./cmd/producer/
```

You can build the `consumer` application with the following command.

```bash
go build ./cmd/consumer/
```

## Running the application

You can run the application with:

```bash
./consumer
```

```bash
./producer
```

or, if you are on Windows:

```pwsh
./consumer.exe
```

```pwsh
./producer.exe
```

## Configuring access and pointing to the queue

Make sure your AWS CLI is correctly configured with your credentials or use the `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` environment variables to configure access.
Also, the `SQS_QUEUE_NAME` environment needs to be set to the name of the queue. Also make sure the AWS region is set, either in your `~/.aws/config` or through `AWS_REGION`.
