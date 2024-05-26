FROM golang:1.22 AS build

WORKDIR /app

# pinned according to go.mod
RUN go install github.com/a-h/templ/cmd/templ@v0.2.697

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN templ generate

RUN go build -o out

FROM debian:bookworm as run

RUN useradd -m runner

COPY --from=build --chown=runner:runner /app/out /home/runner/out
COPY --from=build --chown=runner:runner /app/static /home/runner/static

EXPOSE 8080

USER runner:runner

WORKDIR /home/runner

RUN ./out -seed

ENTRYPOINT ["/home/runner/out"]

