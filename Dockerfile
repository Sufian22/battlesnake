FROM golang:1.16.4-alpine3.13 as builder

RUN apk update && \
    apk add --no-cache --update make

WORKDIR /app
ENV GO111MODULE=on
COPY . .

RUN rm -rf build && mkdir build
RUN make build

FROM alpine:3.13

EXPOSE 3000

COPY --from=builder /app/build/* /build/
COPY --from=builder /app/config.json .

CMD [ "./build/battlesnake", "-config", "config.json" ]