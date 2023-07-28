
# Start from the official Go image
FROM golang:1.20.6-alpine3.17 as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Create a new user with UID 10014
RUN addgroup -g 10020 choreo && \
    adduser  --disabled-password  --no-create-home --uid 10020 --ingroup choreo choreouser


# COPY the source code as the last step
COPY . .


# Build the binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app -buildvcs=false

FROM alpine
COPY --from=build-env /go/bin/app /go/bin/app

USER 10020
ENTRYPOINT ["/go/bin/app"]
