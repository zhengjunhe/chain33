FROM ubuntu:16.04

WORKDIR /root
COPY dplatform dplatform
COPY dplatform-cli dplatform-cli
COPY dplatform.toml dplatform-solo.toml ./

RUN ./dplatform-cli cert --host=127.0.0.1

CMD ["/root/dplatform", "-f", "/root/dplatform.toml"]
