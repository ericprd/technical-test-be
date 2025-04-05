# Project Setup Instructions

## Documentation

[POSTMAN](https://documenter.getpostman.com/view/43758505/2sB2cUAiAo)

## Environment Variables

Before you can run the project, you need to configure the following environment variables in your `.env` file. Please make sure to fill in all the fields below:

```env
POSTGRES_HOST=           # Set the PostgreSQL host
POSTGRES_USER=           # Set the PostgreSQL user
POSTGRES_PASS=           # Set the PostgreSQL password
POSTGRES_DB=             # Set the PostgreSQL database name
POSTGRES_PORT=           # Set the PostgreSQL port (default: 5432)

REDIS_HOST=              # Set the Redis host
REDIS_PORT=              # Set the Redis port (default: 6379)
REDIS_PASS=              # Set the Redis password
REDIS_DB=                # Set the Redis database number
REDIS_DISABLE_ID=        # Set whether to disable Redis ID (true/false)
REDIS_LIFESPAN=          # Set the lifespan for Redis keys

SECRET_KEY=              # Set the secret key for encryption and sessions
```

## Running The Project

Once you have filled out the .env file with the correct values, you can build and run the project using the following commands.

1. Build Project

To build the project and prepare the binary, run the following command:

```bash
make build
```

This will generate the binary files and place them in the `bin` directory.

2. Run the Project

After the project is built, you can run it using the following command:

```bash
make run
```

This will start the project and run it using the built binary.

Note: If you encounter any issues, make sure your .env file is correctly configured with the necessary environment variables.


```
This markdown document clearly outlines the steps for configuring the environment, building the project, and running it, while also explaining where to place the environment variables. Let me know if you need any changes!
```