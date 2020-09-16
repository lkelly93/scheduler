FROM ubuntu:20.04


###############################
## Install needed Software ####
###############################
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

###############################
## Create Secure File System ##
###############################

# Must have info copied
RUN mkdir securefs \
securefs/bin \
securefs/sbin \
securefs/lib \
securefs/lib64 \
securefs/dev \
securefs/usr \
securefs/etc \
securefs/var

# Don't need to copy anything
RUN mkdir securefs/boot \
securefs/home \
securefs/media \
securefs/mnt \
securefs/root \
securefs/srv \
securefs/tmp 

# Mounted/Used in Scheduler
RUN mkdir securefs/proc \
securefs/runner_files \ 
securefs/sys

# Copy all the needed info
RUN cp -r /bin/* /securefs/bin/
RUN cp -r /sbin/* /securefs/sbin/
RUN cp -r /lib/* /securefs/lib/
RUN cp -r /lib64/*  /securefs/lib64/
RUN cp -r /dev/*  /securefs/dev/
RUN cp -r /usr/*  /securefs/usr/
RUN cp -r /etc/*  /securefs/etc/
RUN cp -r /var/*  /securefs/var/




###############################
### Copy over required files ##
###############################

#Install the scheduler
RUN go get github.com/lkelly93/scheduler
RUN go get github.com/lkelly93/scheduler/pkg/executable_container
ENV PATH="$PATH:/root/go/bin"

EXPOSE 3000

CMD ["scheduler"]
