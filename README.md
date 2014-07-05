# envd

Serves application environment variables over HTTP from a config directory.

## Usage

Endpoints:

```
GET /
GET /:service
GET /:service/:environment
```

To dynamically reload configs:

```
POST /reload
```