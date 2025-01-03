#!/bin/bash

echo "Starting indexing flow ..."

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | xargs)
fi

DATA_PATH="$1/maildir"
echo "Data path: $DATA_PATH"

# Check if DATA_PATH exists, otherwise download the data
if [ ! -d "$DATA_PATH" ]; then
  echo "Data path not found, downloading data..."
  wget http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz
  tar -xvf enron_mail_20110402.tgz
  DATA_PATH="./enron_mail_20110402/maildir"
  echo "Data downloaded and extracted to $DATA_PATH."
else
  echo "Data path already exists."
fi

# Check if zincsearch exists
if [ ! -f ./indexer/zincsearch ]; then
  echo "zincsearch not found, downloading..."
  wget https://github.com/zincsearch/zincsearch/releases/download/v0.4.10/zincsearch_0.4.10_linux_x86_64.tar.gz
  tar xzf zincsearch_0.4.10_linux_x86_64.tar.gz
  echo "zincsearch downloaded and extracted."
else
  echo "zincsearch already exists."
fi

cd ./indexer


echo "Using user ${ZINCSEARCH_ADMIN_USER} and password ${ZINCSEARCH_ADMIN_PASSWORD} to start ZincSearch"

# Start server
ZINC_FIRST_ADMIN_USER=${ZINCSEARCH_ADMIN_USER} ZINC_FIRST_ADMIN_PASSWORD=${ZINCSEARCH_ADMIN_PASSWORD} ./zincsearch &


# Wait for ZincSearch to start
echo "Waiting for ZincSearch to start..."
until curl -s http://localhost:4080/ > /dev/null; do
  sleep 1
done
echo "ZincSearch is up and running."

cd ..


# Check if the "Emails" index exists, and create it if it does not
if ! curl -s -o /dev/null -w "%{http_code}" -u ${ZINCSEARCH_ADMIN_USER}:${ZINCSEARCH_ADMIN_PASSWORD} http://localhost:4080/api/index/Emails | grep -q "200"; then
  echo "Creating 'Emails' index..."
  curl -X POST http://localhost:4080/api/index -u ${ZINCSEARCH_ADMIN_USER}:${ZINCSEARCH_ADMIN_PASSWORD} -H 'Content-Type: application/json' -d '{
    "name": "Emails",
    "storage_type": "disk",
    "mappings": {
      "properties": {
        "_id": {"type": "keyword"},
        "from": {"type": "keyword"},
        "to": {"type": "keyword"},
        "date": {"type": "date"},
        "subject": {"type": "text"},
        "body": {"type": "text"}
      }
    }
  }'
  echo "'Emails' index created."
else
  echo "'Emails' index already exists."
fi

# Run the indexer
python3 ./indexer/indexer.py "$DATA_PATH"

echo "Indexing process completed."

cd ./indexer

# Bulk the resulta data
curl http://localhost:4080/api/_bulk -i -u ${ZINCSEARCH_ADMIN_USER}:${ZINCSEARCH_ADMIN_PASSWORD} --data-binary "@index.ndjson"

# Stop ZincSearch server
echo "Stopping ZincSearch server..."
pkill -f zincsearch
echo "ZincSearch server stopped."
