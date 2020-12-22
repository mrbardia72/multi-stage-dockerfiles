# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

# or
#COPY go.mod go.sum ./
# update and upgrade linux on container
#RUN apk update && apk upgrade
# download pkg to container golang
#RUN go get -v \
#    github.com/labstack/echo \
#    github.com/stretchr/testify/assert

RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/httpserver /app/
EXPOSE 8080
ENTRYPOINT ["/app/httpserver"]