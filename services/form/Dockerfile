FROM golang:1.22.2-bookworm

WORKDIR /services/form

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/zzwx/fresh@latest

RUN chmod +x ./start.sh

EXPOSE 8081

CMD ["./start.sh"]