FROM alpine:3.4

WORKDIR /app
COPY main /app

EXPOSE 1323

CMD ["./main"]