#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift

echo "Waiting for PostgreSQL at $host:5432..."

until nc -z "$host" 5432; do
  sleep 1
done

echo "PostgreSQL is up!"

exec "$@"
