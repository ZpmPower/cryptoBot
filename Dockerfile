FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
    ca-certificates
ENV LANGUAGE="en"
COPY code/code .
RUN chmod +x code
EXPOSE 80/tcp
CMD [ "./code" ]