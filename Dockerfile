FROM golang:1.14-alpine AS build

LABEL Author LeviLugato

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch as runner
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]