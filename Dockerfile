FROM ubuntu:16.04
WORKDIR /kqm
COPY ./*.sh ./
RUN chmod +x ./*.sh
RUN ./install.sh
RUN ./test.sh