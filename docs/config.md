# Config and Inputs

## Process Inputs
### Environment Variables
| Env Variable |Type | Descriptiom|
| --- | --- | --- |
| `LINKS_PORT`| `str(int)`| Port to start Server on |
| `LINKS_FRONTEND_PATH`| `str`| Path to find the admin server's static files |
| `LINKS_BACKEND`| `str`| Types of backends `(memory, postgres)` |
| `LINKS_DB_HOST`| `str`| Port to connect to Postgres Database |
| `LINKS_DB_USER` | `str` | User to use for DB connection |
| `LINKS_DB_PASSWORD` | `str` | Password for User

~ OKTA Secrets ~


### Flag Parameters

| Flag |Type | Descriptiom|
| --- | --- | --- |
| `-c,--config`| `str`| yaml config file to load


## Static Config

```yaml
# example links_config.yaml

# Worker Count?
# Database Connection Count?
# Cache Size?
# Okta config?
```