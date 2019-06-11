# https://github.com/eldad87/go-boilerplate
FROM golang:1.11.1 as builder

ARG app_port
ENV APP_PORT $app_port

# Path
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
WORKDIR $GOPATH/src/github.com/vcita/helloworld/

# Copy our code
ADD src/ src/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /app ./src/main.go

# Reduce image size, copy our app
FROM scratch
COPY --from=builder /app /app

EXPOSE ${APP_PORT}
ENTRYPOINT ["/app"]