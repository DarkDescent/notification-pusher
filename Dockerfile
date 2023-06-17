FROM golang:1.20.5-alpine3.18
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/notification-pusher

# ENV SERVER_RUN_MODE=release
# ENV SERVER_PORT=8080
# ENV DATABASE_CONN_STRING=postgres://test:test@localhost:5432/postgres?sslmode=disable
# ENV FIREBASE_CREDENTIALS=

EXPOSE 8080

CMD ["notification-pusher"]