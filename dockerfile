FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go mod init helloworld
RUN go build -o /out/example .

FROM scratch AS bin
COPY --from=build /out/example /

CMD ["go","run","main.go"]