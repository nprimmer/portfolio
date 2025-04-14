#!/bin/bash

# Define the API endpoint
API_ENDPOINT="http://pagemaster.mentats.org:80/validate"

# Check if the user provided a URL as an argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <URL>"
    exit 1
fi

# Get the URL from the command line arguments
URL=$1

# Create a JSON payload
PAYLOAD=$(jq -n --arg url "$URL" '{url: $url}')

# Submit the request to the API
RESPONSE=$(curl -s -w "\nHTTP_STATUS_CODE:%{http_code}\n" -X POST -H "Content-Type: application/json" -d "$PAYLOAD" "$API_ENDPOINT")

# Extract the HTTP status code
HTTP_STATUS=$(echo "$RESPONSE" | grep "HTTP_STATUS_CODE" | awk -F: '{print $2}')

# Extract the body from the response
BODY=$(echo "$RESPONSE" | sed -e 's/HTTP_STATUS_CODE.*//g')

# Print the response body
echo "Response Body: $BODY"

# Handle different HTTP status codes
case $HTTP_STATUS in
    200)
        echo "Success!"
        ;;
    406)
        echo "Error: $BODY"
        ;;
    400)
        echo "Bad Request: $BODY"
        ;;
    500)
        echo "Internal Server Error: $BODY"
        ;;
    *)
        echo "Unexpected HTTP status: $HTTP_STATUS"
        ;;
esac
