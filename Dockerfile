FROM golang:alpine AS dependencies
WORKDIR /app
COPY go.mod go.sum .
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

FROM dependencies AS builder
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -buildvcs=false -ldflags '-s -w -extldflags "-static"' -o redirect

FROM alpine AS final
COPY --from=builder /app/redirect .
ENTRYPOINT ["/redirect"]
