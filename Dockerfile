FROM alpine:latest
WORKDIR /app
COPY  ./build/app .
CMD ["./app"]

