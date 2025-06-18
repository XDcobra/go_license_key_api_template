FROM golang:1.24

WORKDIR /usr/src/app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
ADD . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /api_gateway

# Expose Port
EXPOSE 8000

#ENV DEV=FALSE

# Run
CMD ["/api_gateway"]