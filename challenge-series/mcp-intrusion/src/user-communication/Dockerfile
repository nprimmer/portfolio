FROM golang:alpine

COPY . .

# Remove the sensitive data from the image
RUN rm flag.txt

CMD ["go", "run", "main.go"]
