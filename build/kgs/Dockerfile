FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/kgs

FROM scratch

WORKDIR /app

COPY --from=builder /app/kgs ./

ENTRYPOINT ["./kgs"]