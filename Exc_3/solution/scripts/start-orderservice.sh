#!/bin/bash
# Build and run the Order Service container

docker build -t orderservice .
docker run -d \
  --name orderservice \
  --env-file debug.env \
  --network host \
  orderservice
