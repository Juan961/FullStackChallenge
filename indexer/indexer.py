import os
import json
from email import policy
from email.parser import BytesParser


items = []


DATA_PATH = 'enron_mail_20110402/maildir'
INDEX_NAME = 'Emails'


def main():
    print("Starting indexing process...")
    # Get folder from data path
    for user_folder in os.listdir(DATA_PATH):
        try:
            print(f"Processing user folder: {user_folder}")
            # Get all the files in the folder
            for file in os.listdir(DATA_PATH + '/' + user_folder + "/_sent_mail"):
                print(f"Processing file: {file}")
                # Open the file
                with open(DATA_PATH + '/' + user_folder + "/_sent_mail/" + file, 'rb') as f:
                    content = f.read()    
                # Parse the email
                msg = BytesParser(policy=policy.default).parsebytes(content)
                # Add the email to the items dictionary
                item = {
                    "_id": msg['Message-ID'],
                    "from": msg['From'],
                    "to": msg['To'],
                    "date": msg['Date'],
                    "subject": msg['Subject'],
                    "body": None
                }

                # Body (plain text or HTML)
                if msg.is_multipart():
                    # Extract the plain text part
                    for part in msg.iter_parts():
                        if part.get_content_type() == 'text/plain':
                            item["body"] = part.get_content()
                            break
                else:
                    item["body"] = msg.get_content()

                items.append(item)
                print(f"Added email with ID: {item['_id']}")
            
        except FileNotFoundError:
            print("File not found.")
            continue

        except Exception as e:
            print(f"An error occurred: {e}")
            continue

    print("Writing to index file...")
    with open('indexer/index.ndjson', 'w') as f:
        for item in items:
            f.write(json.dumps({ "index" : { "_index" : INDEX_NAME } } ) + '\n')
            f.write(json.dumps(item) + '\n')
    print("Indexing process completed.")


if __name__ == '__main__':
    try:
        main()
    except Exception as e:
        print(f"An error occurred: {e}")
