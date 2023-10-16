FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    wget \
    tar

RUN wget https://golang.org/dl/go1.21.0.linux-386.tar.gz && \
    tar -C /usr/local -xzf go1.21.0.linux-386.tar.gz && \
    rm go1.21.0.linux-386.tar.gz

# Set Go environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH /go

RUN go version

ENV LANGUAGE="en"
RUN git clone https://github.com/ZpmPower/cryptoBot.git

# Set the working directory to the cloned repository
WORKDIR /cryptoBot

# Initialize the Go module
RUN go mod init bot

# Install Go dependencies
RUN export GO111MODULE=on
RUN go get -x github.com/go-telegram-bot-api/telegram-bot-api/v5
RUN go get github.com/lib/pq
RUN go get github.com/PuerkitoBio/goquery

# Build the Go application
RUN go build -o code crypto.go news.go

# Expose the necessary port
EXPOSE 80

# Specify the command to run your Go application
CMD ["./code"]
