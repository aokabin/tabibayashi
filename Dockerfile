FROM alpine:3.4

WORKDIR /app
COPY app /app

EXPOSE 1323

CMD ["./app"]