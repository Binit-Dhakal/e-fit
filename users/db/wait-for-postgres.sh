#!/bin/sh

set -e

host="$1"
shift

# Wait for the port to be open using nc (netcat)
until nc -z "$host" 5432; do
  echo "Waiting for Postgres ($host:5432) to be reachable..."
  sleep 1
done

echo "Postgres is reachable!"
exec "$@"
