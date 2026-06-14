# Multi-stage build for Go API
FROM golang:1.20-alpine AS build
WORKDIR /src

# cache modules
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN go mod download

# copy source and build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-api ./main.go

# final image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app/go-api /app/go-api
WORKDIR /app
ENV NODE_STATS_URL=http://node-api:3000/stats
ENV JWT_SECRET=secret
EXPOSE 8080
CMD ["/app/go-api"]
