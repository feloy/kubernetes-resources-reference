FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y \
        make \
        xsltproc \
        docbook-xsl \
        fop \
    && rm -rf /var/lib/apt/lists/*
