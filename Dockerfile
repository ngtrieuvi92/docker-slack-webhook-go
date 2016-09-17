FROM caarlos0/alpine-go
MAINTAINER vi.nt "<vi.nt@geekup.vn">

# Install slack-go-webhoook package
RUN go get -v github.com/ashwanthkumar/slack-go-webhook

# Add source
RUN mkdir /data
WORKDIR /data
ADD ./notify.go /data/notify.go

# Build notify package
RUN go build notify.go

#  Set command to run notify
CMD ["/data/notify"]
