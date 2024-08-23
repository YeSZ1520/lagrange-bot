FROM ubuntu:22.04

WORKDIR /app

COPY github.com/YeSZ1520/lagrange-bot .

RUN chmod +x /app/github.com/YeSZ1520/lagrange-bot

CMD ["./github.com/YeSZ1520/lagrange-bot"]