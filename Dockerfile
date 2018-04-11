FROM golang
WORKDIR /app
COPY go-workspace /app/go-workspace
RUN export GOPATH=/app/go-workspace && export GOBIN=/app && go install /app/go-workspace/src/github.com/mongo-exporter/cmd/mongo-exporter/main.go

FROM mongo:latest

COPY --from=0 /app/main /main
CMD ["/main"]