FROM debian:buster-20190708

RUN cp sources.list /etc/apt/sources.list \
    && apt update \
    && apt install -y python3.7 curl python3-distutils python3.7-dev \
    && apt install -y git python-setuptools vim wget \
    && curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py \
    && python3.7 get-pip.py \
    && pip config set global.index-url http://mirrors.aliyuncs.com/pypi/simple/ \
    && pip config set global.trusted-host mirrors.aliyuncs.com \
