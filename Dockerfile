# Use a specific version of Go with Alpine as the base image
FROM golang:1.22.3-alpine

# Install bash
RUN apk add bash

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Add metadata to the image
LABEL project="ascii-art-web" \
      version="1.0" \
      repo="https://learn.zone01oujda.ma/git/hmaach/ascii-art-web-dockerize"

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "."]