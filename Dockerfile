FROM golang:1.18-alpine3.17 AS build
ARG REFLEX_VERSION=v0.3.1
RUN apk add gcc musl-dev git
WORKDIR /src
RUN go install github.com/cespare/reflex@${REFLEX_VERSION}
COPY go.mod go.sum ./
RUN go mod download -x
COPY . .
RUN go build -o /app/im-database-manager ./cmd/serve

FROM alpine:3.17
RUN apk --no-cache -U upgrade \
    && apk add --no-cache postgresql-client
WORKDIR /app
COPY --from=build /app/im-database-manager .
COPY --from=build /src/swagger/swagger.yaml ./swagger/
USER guest
ENTRYPOINT ["/app/im-database-manager"]
