FROM golang:1.22.4 AS builder

ARG VERSION=dev
WORKDIR /app

# {{ENV}}
ENV AUTH_SECRET=your_secret
ENV AUTH_EXPIRES_IN=1d
ENV MAIL_HOST=smtp.mailtrap.io
ENV MAIL_PORT=2525
ENV MAIL_USER=your_user
ENV MAIL_PASS=your_pass
# {{END ENV}}

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION}" -o main .
CMD ["./main"]