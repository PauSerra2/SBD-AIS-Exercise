# -------- Stage 1: Builder --------
FROM golang:1.25 AS builder
WORKDIR /app

# Copy the script folder and make the build script executable
COPY skeleton/scripts ./scripts
RUN chmod +x scripts/build-application.sh && bash scripts/build-application.sh

# -------- Stage 2: Runner --------
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/ordersystem ./ordersystem

# Expose the port (use 3000 for consistency with docker-compose)
EXPOSE 3000

# Set environment variables (using DB_HOST as 'postgres' to match docker-compose config)
ENV POSTGRES_DB=order \
    POSTGRES_USER=docker \
    POSTGRES_PASSWORD=docker \
    POSTGRES_TCP_PORT=5432 \
    DB_HOST=postgres

# Command to run the application
CMD ["./ordersystem"]
