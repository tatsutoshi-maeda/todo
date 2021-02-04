FROM golang:1.15 as builder
WORKDIR /api
COPY api/. .
RUN GO111MODULE=on go build main.go 

FROM golang:1.15
WORKDIR /api
COPY --from=builder /api/main ./main
EXPOSE 8080
ENTRYPOINT [ "./main"]