# Docker-Env CLI Tool

Docker-Env is a command-line interface (CLI) tool designed to simplify the process of injecting environment variables into Dockerfiles from `.env` files. It streamlines the development and deployment workflow by automating the tedious task of manually updating environment variables in Docker configurations.

## Features

- **Easy Integration**: Seamlessly integrates with existing Docker and Go projects.
- **Automatic Environment Variable Injection**: Automatically updates your Dockerfile with environment variables from a specified `.env` file.
- **Cross-Platform Compatibility**: Builds available for Linux, Windows, and macOS.

## Installation Methods

### Install using Go CLI:

#### Prerequisites

- Go 1.22.4 or later

#### Installation

Run the following `go install` command to install and keep this cli-tool in `$GOPATH/bin'. 
```sh
go install github.com/SRV332003/docker-env
```
### Install and Build from source

#### Prerequisites

- Go 1.22.4 or later
- Docker

#### Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/SRV332003/docker-env.git
cd docker-env
```
Build the project using the provided Makefile:
```sh
make build
```
This will compile the Docker-Env CLI tool and place the executable in your `$GOPATH/bin`.

## Usage
To inject environment variables into your Dockerfile, first place a `# {{ENV}` placeholder in your Dockerfile where you want the environment variables to be injected. Then, run the following command:
```sh
docker-env <path-to-envfile> <path-to-Dockerfile>
```
## Example
Given a .env file with the following content:

```sh
AUTH_SECRET=your_secret
AUTH_EXPIRES_IN=1d
MAIL_HOST=smtp.mailtrap.io
MAIL_PORT=2525
MAIL_USER=your_user
MAIL_PASS=your_pass
```
And a Dockerfile with placeholder `# {{ENV}}` for environment variables:

```dockerfile
FROM golang:1.22.4 AS builder

ARG VERSION=dev
WORKDIR /app

# {{ENV}}
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION}" -o main .
CMD ["./main"]
```

These have the following directory structure:

```sh
.
├── .env
├── Dockerfile
└── README.MD
```

Running the following command will inject the environment variables from the `.env` file into the Dockerfile:

```sh
docker-env .env Dockerfile
```

The updated Dockerfile will look like this:

```dockerfile   
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
```

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any bugs or feature requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
