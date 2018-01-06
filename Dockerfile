FROM ubuntu
WORKDIR /root
ADD ./cmd/main .
ADD ./public ./public
ADD ./view ./view

EXPOSE 8000

# Run it
ENTRYPOINT ["./main"]
