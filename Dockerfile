## 1st stage: build golang ethereum runner
FROM golang:alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 go build -o backoffice cmd/server/main.go

## 2nd stage: prepare container to run node
FROM golang:alpine as runner

WORKDIR /home/backoffice

COPY --from=builder /usr/src/app/backoffice /home/backoffice
COPY --from=builder /usr/src/app/migrations /home/backoffice/migrations

CMD ["./backoffice"]

