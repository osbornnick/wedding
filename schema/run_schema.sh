#!/usr/bin/env bash
set -euo pipefail

# Run Flyway in Docker against a local PostgreSQL instance using the local sql/ directory.
# Defaults can be overridden via environment variables (DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD, SQL_DIR).
# On Docker Desktop for Windows use host.docker.internal as DB_HOST so the container can reach the host Postgres.

DB_HOST="${DB_HOST:-host.docker.internal}"
DB_PORT="${DB_PORT:-5432}"
DB_NAME="${DB_NAME:-wedding}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
SQL_DIR="${SQL_DIR:-$(pwd)/sql}"

echo "Running Flyway migrations from: $SQL_DIR"
echo "Target: jdbc:postgresql://$DB_HOST:$DB_PORT/$DB_NAME (user=$DB_USER)"

docker run --rm \
  -v "$SQL_DIR":/flyway/sql \
  -e FLYWAY_URL="jdbc:postgresql://$DB_HOST:$DB_PORT/$DB_NAME" \
  -e FLYWAY_USER="$DB_USER" \
  -e FLYWAY_PASSWORD="$DB_PASSWORD" \
  flyway/flyway:latest migrate
