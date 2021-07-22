
# =============== build and run ===============
# build:  docker build -t go-weibo-push:latest .
# run:    docker run -dit go-weibo-push:latest


# =============== build stage ===============
FROM golang:1.16.5-buster AS build
# env
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
# dependent
WORKDIR /app
COPY go.* ./
RUN go mod download -x all
# build
COPY . ./
# ldflags:
#  -s: disable symbol table
#  -w: disable DWARF generation
# run "go tool link -help" to get the full list of ldflags
RUN go env && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o go-weibo-push -v ./main.go



# =============== final stage ===============
FROM alpine:latest AS final
# resources
WORKDIR /app
COPY --from=build /app/go-weibo-push ./
COPY --from=build /app/conf/base.toml ./conf/base.toml
COPY --from=build /app/conf/prod ./conf/prod
EXPOSE 8881
ENTRYPOINT ["env","GO_ENV=prod","/app/go-weibo-push", "-other", "flags"]
