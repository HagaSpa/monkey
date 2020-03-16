FROM golang:1.13.6

# Disable SSL verification
RUN git config --global http.sslverify false

RUN go get -v golang.org/x/tools/gopls
RUN go get -v github.com/go-delve/delve/cmd/dlv