FROM debian

CMD ["bash"]

RUN apt update
RUN apt install nginx -y

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]