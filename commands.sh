#!/bin/bash

# Stop the running container
docker stop ascii-art-web-con

# Remove unused data
docker system prune -f

# Remove the old image (corrected typo in command)
docker rmi ascii-art-web-img

# Build a new image
docker build -f Dockerfile -t ascii-art-web-img .

# Run a new container
docker run -d -p 8080:8080 --name ascii-art-web-con ascii-art-web-img

# Execute an interactive bash shell in the running container
docker exec -it ascii-art-web-con /bin/bash