# latest golang base image
FROM golang:latest

# Add Maintainer info
LABEL maintainer="Quique <kp@email.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy Go modules dependency requirements file
COPY go.mod .

# Copy Go modules expected hashes file
COPY go.sum .

RUN go mod download

# Copy all the app sources (recursively copies files and directories from the into the container)
COPY . .

ENV PORT 5000

RUN go build

# Remove source files
RUN find . -name "*.go" -type f -delete

# make port 5000 available to the outside the conatiner
EXPOSE $PORT

CMD ["./go-gin-project-001"]


