# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env
 
# Set environment variable
ENV APP_NAME http-server
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME
 
# Run Stage
FROM alpine:3.14
 
# Set environment variable
ENV APP_NAME http-server
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 8080
 
# Start app
CMD ./$APP_NAME