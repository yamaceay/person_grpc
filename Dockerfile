# Build stage
FROM golang:latest as build

# Current directory: /app
WORKDIR /app

# Save go.mod and go.sum inside /app
# Initialize Go project
COPY go.mod .
COPY go.sum .
RUN go mod download

# Save other files inside /app too
COPY . .

# Save executable of server inside app as ./person-server
RUN go build -o person-server ./server

# Distroless stage
FROM gcr.io/distroless/base-debian10

# Current directory: /
WORKDIR /

# Save the executable content from the build stage as ./server
COPY --from=build ./app/person-server ./server

# Copy config.yaml
COPY ./config.yaml ./config.yaml

# Copy bash shell commands in bin directory
COPY --from=busybox:1.35.0-uclibc /bin /bin

# Specify run command using ./server
CMD [ "./server" , "-f", "config.yaml"]