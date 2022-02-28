FROM golang:1.14.3-alpine as localserver
WORKDIR /app
COPY . . 
RUN apk add --no-cache git
RUN go get -u github.com/shurcooL/goexec
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
# RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

FROM localserver
RUN GOOS=js GOARCH=wasm go build -o main.wasm
# RUN cp "./Web/index.html" "/var/www/html/"
CMD ["goexec", "http.ListenAndServe(`:80`, http.FileServer(http.Dir(`.`)))"]