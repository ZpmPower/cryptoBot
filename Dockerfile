FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates \
    git \
    curl 
# Download Go 1.2.2 and install it to /usr/local/go
RUN curl -s https://storage.googleapis.com/golang/go1.2.2.linux-amd64.tar.gz| tar -v -C /usr/local -xz
ENV LANGUAGE="en"
RUN git clone https://github.com/ZpmPower/cryptoBot.git .
RUN cd cryptoBot
RUN go mod init bot
RUN go get github.com/Syfaro/telegram-bot-api
RUN go get github.com/lib/pq
RUN go build -o code crypto.go news.go
EXPOSE 80/tcp
CMD [ "./cryptoBot/code" ]
