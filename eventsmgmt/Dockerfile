# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11 AS builder



ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep



# Copy the local package files to the container's workspace.
COPY . /go/src/booking-app/eventsmgmt
WORKDIR /go/src/booking-app/eventsmgmt
RUN ls -ltr /go/src/booking-app/eventsmgmt/pkg

# Only install dependencies specified in dep vendor file
RUN dep ensure --vendor-only
COPY . ./

RUN groupadd -g 2000 go \
&& useradd -m -u 2001 -g go go
USER go

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 go install booking-app/eventsmgmt/

# Final image layer
# Using alpine in place of scratch
# for shell access
FROM alpine:3.8
RUN apk update && apk add curl
COPY --from=builder /go/bin/eventsmgmt ./app
ENTRYPOINT ["./app"]

# Document that the service listens on port 9000.
EXPOSE 9000
