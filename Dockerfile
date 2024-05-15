FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./careerhub/review_service/ ./careerhub/review_service/
RUN ls --recursive ./

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./careerhub/review_service/

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/myapp /myapp

USER nonroot:nonroot

ENTRYPOINT ["/myapp"]
