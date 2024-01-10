FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
EXPOSE 9081
CMD ["go", "run", "."]