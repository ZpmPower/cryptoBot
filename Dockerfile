FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    curl \
    tar
    
RUN curl -O https://golang.org/dl/go1.17.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.2.linux-amd64.tar.gz && \
    rm go1.17.2.linux-amd64.tar.gz

ENV LANGUAGE="en"
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH
 
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
