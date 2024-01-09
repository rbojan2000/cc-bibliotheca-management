FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
EXPOSE 9080
CMD ["go", "run", "."]