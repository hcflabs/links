# Build Server Deps
FROM golang:1.21.6 AS SERVERBUILD
WORKDIR /app
COPY .. .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o /links

## Final Image
FROM alpine:3.18.4

COPY --from=SERVERBUILD /links /links
# Run
CMD ["/links"]