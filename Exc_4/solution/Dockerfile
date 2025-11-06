# -------- Stage 1: Builder --------
FROM golang:1.25 AS builder
WORKDIR /app

# Copy all Go source files and modules
COPY . .

# Run the build script
RUN chmod +x scripts/build-application.sh && bash scripts/build-application.sh


# -------- Stage 2: Runner --------
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ordersystem .

# Command to run the service
CMD ["./ordersystem"]
