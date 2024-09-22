# Use Golang official image as a base
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . .

# Install dependencies and build the api
RUN go mod tidy && go build -o merinio-api .

# Expose the api's port
EXPOSE 8080

# Command to run the api
CMD ["./merinio-api"]
