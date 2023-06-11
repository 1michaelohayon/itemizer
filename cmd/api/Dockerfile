FROM golang:1.20.3-alpine as build-stage

WORKDIR /api
COPY . .
RUN go mod download && \
	go build -o bin/api ./cmd/api	

############################
FROM alpine
COPY --from=build-stage /api/bin/api .

CMD ./api
