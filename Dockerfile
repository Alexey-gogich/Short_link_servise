FROM golang:latest

WORKDIR /app
RUN mkdir /cmd
COPY ./cmd /app/cmd
RUN mkdir /internal
COPY ./internal /app/internal

RUN go mod init short_link_servise
RUN go mod tidy

RUN go build -o /app/build ./cmd/Short_link_servise/

CMD [ "/app/build" ]
EXPOSE 5248