# db_backup
A Go script to backup the DB

## Installation

[Install Go](https://go.dev/doc/install)

```bash
# Clone the repository
git clone https://github.com/r-anime/db_backup.git
cd db_backup

# Fetch dependencies
go mod tidy

# Build the CLI
go build -o ./bin/db_backup
```

Install the linter by following the instructions
https://golangci-lint.run/docs/welcome/install/local/

which should suggest something like
```bash
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOBIN) v2.7.2
```

Run the linter by running

```BASH
golangci-lint run
```

Cry if these instructions don't work, cause you're on your own.

## Usage

```bash
# Basic usage
./db_backup -s stage_db -d r_anime_staging -b daily -c 5

# Flags
-s, --db-service string        the docker compose service name for the db (default "db")
-d, --db-name string           the name of the database (default "r_anime")
-b, --backup-type string       the backup series label (yearly monthly weekly daily manual) (default "manual")
-c, --compression-level int8   the compression level (1 - 19) (default 3)
```

### Examples

```bash
# Normal dev like run
go run . -s stage_db -d r_anime_staging -b weekly

# Using build and running with default values
go build -o ./bin/db_backup
./bin/db_backup

# Using docker compose
docker compose run --build --rm db_backup -s stage_db -d r_anime_staging -b weekly
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
