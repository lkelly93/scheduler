FROM ubuntu:20.04

RUN apt-get update -y

#Install needed packages
RUN apt-get install python3 -y
RUN apt-get install default-jre -y
RUN apt-get install golang -y

#Reduce VM size
RUN rm -rf /var/lib/apt/lists/*

WORKDIR /secure

CMD bash