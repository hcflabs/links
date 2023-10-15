# Build Image
FROM golang:1.21.3 AS BUILD
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /links

## Final Image
FROM alpine:3.18.4 

COPY --from=BUILD /links /links

# Run
CMD ["/links"]