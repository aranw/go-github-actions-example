FROM cgr.dev/chainguard/go:latest AS build

WORKDIR /usr/src/app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM cgr.dev/chainguard/glibc-dynamic:latest

COPY --from=build /run-app /usr/local/bin/
CMD ["run-app"]
