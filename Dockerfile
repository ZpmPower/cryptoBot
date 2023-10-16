FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    curl \
    tar

# Add the PPA for golang-backports
RUN add-apt-repository ppa:longsleep/golang-backports

# Update the package list again after adding the PPA
RUN apt-get update

# Install golang from the PPA
RUN apt-get install -y golang

ENV LANGUAGE="en"
RUN git clone https://github.com/ZpmPower/cryptoBot.git

# Set the working directory to the cloned repository
WORKDIR /cryptoBot

# Initialize the Go module
RUN go mod init bot

# Install Go dependencies
RUN go get github.com/Syfaro/telegram-bot-api
RUN go get github.com/lib/pq
RUN go get github.com/PuerkitoBio/goquery

# Build the Go application
RUN go build -o code crypto.go news.go

# Expose the necessary port
EXPOSE 80

# Specify the command to run your Go application
CMD ["./code"]
