FROM golang:alpine
ADD . /go/src/web-app-run-python
RUN go install web-app-run-python

FROM alpine
COPY --from=0 /go/bin/web-app-run-python .
COPY --from=0 /go/src/web-app-run-python .
RUN apk add python3
CMD ["./web-app-run-python"]

