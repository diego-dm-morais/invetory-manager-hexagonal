FROM golang:1.18.2-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
#RUN go mod tidy
RUN go mod download
COPY application ./application
COPY framework ./framework
COPY main.go ./

ARG MONGO_DATA_BASE_URL
ARG PORT_SERVER

ENV MONGO_DATA_BASE_URL=${MONGO_DATA_BASE_URL}
ENV PORT_SERVER=${PORT_SERVER}

RUN go build -o /inventory-manager
##
## Deploy
##
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /inventory-manager /inventory-manager
EXPOSE 1323
USER nonroot:nonroot
ENTRYPOINT ["/inventory-manager"]