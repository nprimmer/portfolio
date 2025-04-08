#!/bin/bash

# Generate a random username
USERNAME="user$(openssl rand -hex 4)"

# Generate a random 8-character password
PASSWORD=$(openssl rand -base64 6)

# Generate a random flag between 8 and 24 characters long
FLAG_LENGTH=$((8 + RANDOM % 17))
FLAG=$(openssl rand -base64 $((FLAG_LENGTH * 3 / 4)))

# Print the generated values
echo "Username: $USERNAME"
echo "Password: $PASSWORD"
echo "Flag: $FLAG"

# Create the flag using the generated username and password
curl -u "$USERNAME:$PASSWORD" -X POST http://localhost:9090/MakeFlag -H "Content-Type: application/json" -d "{\"flag\": \"$FLAG\"}"

# Retrieve the flags using the same username and password
curl -u "$USERNAME:$PASSWORD" -X GET http://localhost:9090/GetFlags

