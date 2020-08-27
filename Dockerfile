FROM centos:8

WORKDIR /

COPY ./bin/random /random

CMD ["./random"]