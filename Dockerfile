FROM golang:1.13

WORKDIR /app

COPY . .

RUN go build -i -o structify main.go

# Run at least once conversion now to speed up conversions when the server starts
# There is a race for the first build that causes a crash. Just retry for now
# TODO: Find a better way to speed up the first conversion and replace this
RUN /app/structify convert examples/cm.yaml || /app/structify convert examples/cm.yaml

CMD ["/app/structify", "serve"]
