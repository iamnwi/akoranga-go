FROM golang:1.10.3-alpine3.8 as builder

COPY main.go ./akoranga-go.go
RUN go build akoranga-go.go

# RUN
FROM alpine:edge

ENV PORT 80

EXPOSE $PORT

COPY --from=builder /go/akoranga-go /
CMD ["/akoranga-go"]
