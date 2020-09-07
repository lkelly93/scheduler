FROM ubuntu:20.04

RUN apt-get update -y

#Install needed packages
RUN apt-get install software-properties-common -y
RUN apt-get install python3 -y
RUN apt-get update -y 
RUN apt-get install python3-pip -y
RUN apt-get install default-jre -y
RUN apt-get install golang -y 
RUN apt-get install git -y 

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