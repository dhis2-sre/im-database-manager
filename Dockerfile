FROM golang:1.16-alpine AS build
RUN apk -U upgrade && \
    apk add gcc musl-dev git
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/im-database-manager ./cmd/serve

FROM alpine:3.13
RUN apk --no-cache -U upgrade \
    && apk add --no-cache postgresql-client
WORKDIR /app
COPY --from=build /app/im-database-manager .
COPY --from=build /src/swagger/swagger.yaml ./swagger/
USER guest
CMD ["/app/im-database-manager"]
