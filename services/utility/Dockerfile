FROM golang:1.22.2-bookworm

WORKDIR /services/utility

RUN git clone https://github.com/mrspec7er/license-request-utility .

COPY .env ./

RUN go mod download

RUN chmod +x ./start.sh

EXPOSE 8000

CMD ["./start.sh"]