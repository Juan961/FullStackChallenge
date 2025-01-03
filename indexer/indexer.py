import os
import sys
import json
from email import policy
from email.parser import BytesParser


items = []

# 'enron_mail_20110402/maildir'
DATA_PATH = sys.argv[1]
INDEX_NAME = 'Emails'


def main():
    # Get folder from data path
    for user_folder in os.listdir(DATA_PATH):
        try:
            # Get all the files in the folder
            for file in os.listdir(DATA_PATH + '/' + user_folder + "/_sent_mail"):
                # Open the file
                with open(DATA_PATH + '/' + user_folder + "/_sent_mail/" + file, 'rb') as f:
                    # Parse the email
                    msg = BytesParser(policy=policy.default).parse(f)
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
        except FileNotFoundError:
            print(f"No _sent_mail folder for user: {user_folder}")
            continue

    with open('indexer/index.ndjson', 'w') as f:
        for item in items:
            f.write(json.dumps({ "index" : { "_index" : INDEX_NAME } } ) + '\n')
            f.write(json.dumps(item) + '\n')


if __name__ == '__main__':
    main()
