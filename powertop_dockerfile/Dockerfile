# Dockerfile

#the base image is fedora
FROM fedora:latest

# Build the image as root user
USER root

#installing powertop
RUN yum -y update && \
    yum install -y powertop

#starting powertop
CMD powertop --csv=powertop_report.txt --time=20s