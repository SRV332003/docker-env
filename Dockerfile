FROM golang:1.22.4 AS builder

ARG VERSION=dev
WORKDIR /app

# {{ENV}}
ENV AUTH_SECRET {AUTH_SECRET}
ENV AUTH_EXPIRES_IN {AUTH_EXPIRES_IN}
ENV MAIL_HOST {MAIL_HOST}
ENV MAIL_PORT {MAIL_PORT}
ENV MAIL_USER {MAIL_USER}
ENV MAIL_PASS {MAIL_PASS}
# {{END ENV}}


COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION}" -o main .
CMD ["./main"]