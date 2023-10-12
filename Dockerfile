FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    curl \
    golang=2:1.21.3-1ubuntu1

# Set Go environment variables
ENV PATH="/usr/lib/go-1.21.3/bin:${PATH}"
ENV GOPATH="/go"

ENV LANGUAGE="en"
RUN git clone https://github.com/ZpmPower/cryptoBot.git

# Set the working directory to the cloned repository
WORKDIR /cryptoBot

# Initialize the Go module
RUN go mod init bot

# Install Go dependencies
RUN go get github.com/Syfaro/telegram-bot-api
RUN go get github.com/lib/pq

# Build the Go application
RUN go build -o code crypto.go news.go

# Expose the necessary port
EXPOSE 80

# Specify the command to run your Go application
CMD ["./code"]
