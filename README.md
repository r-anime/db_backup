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
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOBIN) v2.11.2
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
-d, --docker-container string   the docker (not compose) container name for the db (default "modbot-db")
-n, --db-name string            the name of the database (default "r_anime")
-t, --backup-type string        the backup series label (yearly, monthly, weekly, daily, manual) (default "manual")
-c, --compression-level int8    the compression level (1 - 19) (default 3)
-b, --backup-dir string         the directory the backups are located in (default "manual")
-m, --min-save-size uint16      the minimum file size to consider for saving in MiB (default 2048)

### Examples

```bash
# Normal dev like run
go run . -d stage-modbot-db -n r_anime_staging -t manual

# Using build and running with default values
go build -o ./bin/db_backup
./bin/db_backup

# Using docker compose (has some issues with volume management)
docker compose run --build --rm db_backup -d stage-modbot-db -n r_anime_staging -t manual
```

## Releasing

Install GoReleaser (This installs a metric fuckton of shit it seems)
`go install github.com/goreleaser/goreleaser/v2@latest`

For CI, just push a git tag and it'll automatically make a release

````

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
