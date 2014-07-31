FROM ubuntu:14.04
MAINTAINER Dan Sosedoff "dan.sosedoff@gmail.com"

# Download envd binary from github releases
ADD https://github.com/sosedoff/envd/releases/download/v0.4.0/envd_linux_amd64 /usr/local/bin/envd
RUN chmod +x /usr/local/bin/envd

# Expose default port
EXPOSE 3050

# Define start command
CMD ["/usr/local/bin/envd", "-c", "/usr/local/etc/envd"]