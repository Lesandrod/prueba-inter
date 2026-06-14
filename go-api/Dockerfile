FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /app/go-api ./main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app/go-api /app/go-api
WORKDIR /app
ENV NODE_STATS_URL=http://node-api:3000/stats
ENV JWT_SECRET=secret
EXPOSE 8080
CMD ["/app/go-api"]