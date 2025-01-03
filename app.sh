#!/bin/bash

echo "Starting production app ..."

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

# Default PORT
PORT="3000"

# Parse named parameters
while [[ "$#" -gt 0 ]]; do
  case $1 in
    --port) PORT="$2"; shift ;;
    *) echo "Unknown parameter passed: $1"; exit 1 ;;
  esac
  shift
done

# Function to handle cleanup on exit
cleanup() {
  echo "Stopping frontend, backend, and indexer..."
  pkill -P $$
  wait
  echo "Stopped."
}

# Trap SIGINT (ctrl + c) and call cleanup
trap cleanup SIGINT

cd ./frontend/
pnpm i
pnpm run build
pnpm run preview --port $PORT &

echo "Starting frontend project in port: :${PORT}"

cd ./../backend/
go mod tidy
go run . &

echo "Starting backend project in port: :3333"

cd ./../indexer/
ZINC_FIRST_ADMIN_USER=${ZINCSEARCH_ADMIN_USER} ZINC_FIRST_ADMIN_PASSWORD=${ZINCSEARCH_ADMIN_PASSWORD} ./zincsearch &

echo "Starting indexer service"

# Wait for all background processes to finish
wait
