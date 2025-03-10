# My website

Source code for my website

## How to run locally

1. Install Go 1.22
2. Install and configure PostgreSQL (create user and db for project, grant permissions to user)
3. Clone repository
4. Add .env with PORT, APP_ENV, DB_USERNAME, DB_PASSWORD, DB_PORT, DB_HOST, DB_NAME
5. Run with commands below:

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

## Useful commands:

Connect to psql as postgres:
```bash
sudo -u postgres psql
```

Then:
```sql
CREATE DATABASE my_website_db;
CREATE DATABASE my_website_db_test;
CREATE USER my_website_admin WITH PASSWORD '123';
GRANT ALL PRIVILEGES ON DATABASE my_website_db TO my_website_admin;
GRANT ALL PRIVILEGES ON DATABASE my_website_db_test TO my_website_admin;
ALTER USER my_website_admin WITH SUPERUSER;
```


Connect to my_website_db as my_website_admin:
```bash
psql -h localhost -d my_website_db -U my_website_admin -p 5432
```