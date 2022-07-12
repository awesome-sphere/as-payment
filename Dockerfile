FROM golang:1.18-alpine3.16 as build-stage

COPY ./ /go/src/as-payment
WORKDIR /go/src/as-payment

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o as-payment

FROM alpine:latest as production-stage

RUN apk --no-cache add ca-certificates

COPY --from=build-stage /go/src/as-payment /as-payment
WORKDIR /as-payment

CMD ["./as-payment"]