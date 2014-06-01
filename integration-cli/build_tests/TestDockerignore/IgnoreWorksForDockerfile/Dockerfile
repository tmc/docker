FROM busybox

ADD . /

RUN [ -f test_file ]
RUN [ ! -e Dockerfile ]
