#!/bin/sh
# wait-for-postgres.sh

host="$1"
shift
until nc -z "$host" 5432; do
  echo "Waiting for PostgreSQL at $host:5432..."
  sleep 1
done

echo "PostgreSQL is up!"
exec "$@"
