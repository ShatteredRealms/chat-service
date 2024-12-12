# Build application
FROM golang:1.23 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG APP_VERSION=v0.0.1
RUN go build \
	-ldflags="-X 'github.com/ShatteredRealms/chat-service/pkg/config/default.Version=${APP_VERSION}'" \
	-o /out/chat ./cmd/chat-service

# Run server
FROM alpine:3.15.0
WORKDIR /app
COPY --from=build /out/chat ./
COPY ./migrations ./migrations
EXPOSE 8081
ENTRYPOINT [ "./chat" ]

