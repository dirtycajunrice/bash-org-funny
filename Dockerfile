FROM golang:1.16 as builder

ENV GO11MODULES=on

WORKDIR /app

COPY . .

RUN go build .

FROM scratch

COPY --from=builder /app/bash-org-funny /

ENTRYPOINT ["/bash-org-funny"]
