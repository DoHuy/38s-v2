FROM golang:1.14.1 as builder
WORKDIR /app
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*
COPY go.mod /app
COPY go.sum /app
RUN eval $(ssh-agent -s) && \
    mkdir ~/.ssh && \
    ssh-keyscan gitlab.id.vin >> ~/.ssh/known_hosts && \
    cat /opt/private/gitlab_runner.priv > ~/.ssh/id_rsa && \
    chmod 400 ~/.ssh/id_rsa
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o iam ./src/auth/cmd/app.go

FROM golang:1.14.1
WORKDIR /app
COPY --from=builder /app/src/38s/template /app/template
COPY --from=builder /app/src/38s/static /app/static
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["/app/iam"]
