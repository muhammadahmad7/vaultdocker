FROM ubuntu
MAINTAINER Muhammad Ahmad <ahmadmalik752@gmail.com>
ADD vaultd vaultd
EXPOSE 8080 8081
ENTRYPOINT ["/vaultd"]
