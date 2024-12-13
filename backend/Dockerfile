
FROM golang:1.23.4-bullseye AS build

WORKDIR /app 

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM build AS dev

RUN go install github.com/air-verse/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

CMD ["air", "-c", ".air.toml"]

FROM build AS build-production

RUN useradd -u 1001 golang

COPY . .

RUN go build \
  -ldflags="-linkmode external -extldflags -static" \
  -tags netgo \
  -o data-collector

FROM scratch

ENV GIN_MODE=release

WORKDIR /

COPY --from=build-production /etc/passwd /etc/passwd

COPY --from=build-production /app/data-collector data-collector

USER golang

EXPOSE 8080

CMD ["/data-collector"]