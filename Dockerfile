FROM ubuntu:20.04

RUN apt-get update -y

#Configure tzdata
ARG DEBIAN_FRONTEND="noninteractive" 
ENV TZ=America/Tijuana
RUN apt-get install -y tzdata

#Install needed packages
RUN apt-get install -y \ 
python3 \
default-jre \
golang \
git 

RUN apt-get update -y 

RUN apt-get install -y \
python3-pip

#Install language dependacies
    #Python
    RUN pip3 install numpy

#Reduce VM size
RUN rm -rf /var/lib/apt/lists/*
RUN mkdir runner_files

#Install the scheduler
RUN go get github.com/lkelly93/scheduler
ENV PATH="$PATH:/root/go/bin"

EXPOSE 3000

CMD ["scheduler"]