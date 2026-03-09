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

# On prod

# Normal manual backup
sudo db_backup

# or with some default flags explicitly listed that you might want to change
sudo db_backup -d modbot-db -n r_anime -t manual -c 3

# Help for options
db_backup --help
```

## Releasing

~~Install GoReleaser (This installs a metric fuckton of shit it seems)
`go install github.com/goreleaser/goreleaser/v2@latest`~~

Just push a git tag and it'll automatically make a release.

```bash
git tag -a 1.0.0 -m "1.0.0"
git push --tags
```

Then you can just download it to the server and unpack it and put it in a bin and use it.

```bash
VERSION=1.0.0 &&
wget -O db_backup.tar.gz "https://github.com/r-anime/db_backup/releases/download/${VERSION}/db_backup_${VERSION}_linux_amd64.tar.gz" &&
tar -xzf db_backup.tar.gz &&
sudo mv db_backup /usr/local/bin/ &&
sudo chmod +x /usr/local/bin/db_backup &&
rm db_backup.tar.gz
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
