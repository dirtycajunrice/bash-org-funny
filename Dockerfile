FROM golang:1.16 as builder

ENV GO11MODULES=on

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /app/main .

CMD ["/main"]
