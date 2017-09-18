FROM alpine:3.4

WORKDIR /app
COPY main /app

RUN apk add --no-cache --update ca-certificates lame sox

EXPOSE 1323

CMD ["./main"]