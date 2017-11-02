FROM alpine:3.4

WORKDIR /app

RUN apk add --no-cache --update ca-certificates lame sox

COPY sound /app/sound
COPY main /app

EXPOSE 1323

CMD ["./main"]