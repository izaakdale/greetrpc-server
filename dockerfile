FROM golang:1.22-alpine as builder
WORKDIR /
COPY . ./
RUN go mod download


RUN go build -o /greetrpc-server


FROM alpine
COPY --from=builder /greetrpc-server .


EXPOSE 80
CMD [ "/greetrpc-server" ]