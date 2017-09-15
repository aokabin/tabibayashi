FROM golang:alpine

WORKDIR /app
COPY hoge.go /app

EXPOSE 8000

CMD ["go","run","hoge.go"]