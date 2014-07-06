# envd

Serves application environment variables over HTTP from a config directory.

## Overview

Original idea was to fetch application environment variables for `.env` file from the network.
[Dotenv](https://github.com/bkeepers/dotenv) files are very useful for local development as 
well as for production deployments. Service could be used to fetch data for `.env` files 
per environment with deployment tools like Capistrano, Mina or Shuttle via curl/wget command.

All services and service environments are stored as directories under main config directory. 
Each environment variable is stored in a file that contains its value.

Here's the example config structure:

```
myapp (dir)
- production (dir)
-- rails_env (file)
-- redis_url (file)
-- s3_access_key (file)
-- s3_secret_key (file)
-- s3_bucket (file)
```

Request production environment of myapp with curl:

```
curl http://envdhost:5000/myapp/production
```

Example output:

```
RAILS_ENV=production
REDIS_URL=redis://localhost:6379
S3_ACCESS_KEY=foo
S3_SECRET_KEY=bar
S3_BUCKET=foobar
```

## Usage

Options:

```
Usage of ./envd:
  -c="": Path to config directory
  -p=3050: Port to listen on
  -t="": Authentication token
```

Start server:

```
envd -c ./examples -t foo
```

Make requests with `curl`:

``` bash
# Without authentication
curl http://localhost:3050/myapp/production

# With authentication
curl -f http://localhost:3050/myapp/production?token=foo
curl -f http://localhost:3050/myapp/production -H "Token: foo"
```

## Endpoints

Fetch all available services:

```
GET /
```

Fetch service environments:

```
GET /:service
```

Fetch service environment variables:

```
GET /:service/:environment
```

Reload configuration files:

```
POST /reload
```

## Compile

To compile application from source execute:

```
make deps
make build
```

## Test

Execute test suite with command:

```
make test
```

## License

The MIT License (MIT)

Copyright (c) 2014 Dan Sosedoff, <dan.sosedoff@gmail.com>