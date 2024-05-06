FROM golang:1.22.2-bookworm

WORKDIR /services/logger

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/zzwx/fresh@latest

RUN chmod +x ./start.sh

EXPOSE 8082

CMD ["./start.sh"]