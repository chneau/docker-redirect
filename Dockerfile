FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -buildvcs=false -ldflags '-s -w -extldflags "-static"' -o redirect

FROM alpine AS final
COPY --from=builder /app/redirect .
ENTRYPOINT ["/redirect"]
