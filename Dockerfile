FROM golang:1.12-alpine AS build

RUN apk add --update --no-cache ca-certificates git

WORKDIR ./dosmj
COPY . ./

RUN GO111MODULE=on go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/dosmj

# This results in a single layer image
FROM scratch

# adding ca certificates
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# copying binary
COPY --from=build /bin/dosmj /bin/dosmj

ENTRYPOINT ["/bin/dosmj"]