FROM ubuntu:latest
EXPOSE 8080
ADD main /
CMD ["/main"]


