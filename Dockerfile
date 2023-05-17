FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /elk-api

FROM scratch
COPY --from=builder ./elk-api /elk-api
CMD ["/elk-api"]
