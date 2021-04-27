FROM golang:1.16 as builder

ENV GO11MODULES=on

WORKDIR /app

COPY . .

RUN go build .

FROM scratch

WORKDIR /app

COPY --from=builder /app/bash-org-funny /app

ENTRYPOINT ["/app/bash-org-funny"]
