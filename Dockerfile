FROM golang:1.17-alpine
# Set destination for COPY
WORKDIR /go/src/app
RUN echo $GOPATH
# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /hubble
EXPOSE 8080
CMD [ "/hubble" ]