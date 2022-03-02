#Stage 1: BUILD
FROM golang:alpine AS BUILD
ENV TZ=Asia/Ulaanbaatar
ENV GO111MODULE=on
RUN apk add bash ca-certificates git gcc g++ libc-dev make
RUN apk add --update tzdata
WORKDIR /go/src/github.com/johandui/domain-automation
COPY . .
RUN go mod tidy
ENV TZ=Asia/Ulaanbaatar
RUN CGO_ENABLED=0 GOOS=linux go build -a -gcflags='-N -l' -installsuffix cgo -o main .

# Stage 2: RUN
FROM busybox:stable-musl
ENV TZ=Asia/Ulaanbaatar
WORKDIR /home
COPY --from=BUILD /go/src/github.com/johandui/domain-automation/main /home/
EXPOSE 3000
ENTRYPOINT ["/home/main"]