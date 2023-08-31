FROM debian:12.1-slim

RUN apt update 
RUN apt install nginx -y 

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]