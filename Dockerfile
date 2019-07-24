FROM golang:stretch as builder
WORKDIR src/hpa/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v

FROM scratch
COPY --from=builder go/src/hpa/hpa .

EXPOSE 8080
ENTRYPOINT ["./hpa"]