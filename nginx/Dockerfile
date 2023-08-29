FROM debian:12.1-slim

EXPOSE 80

RUN apt update 
RUN apt install nginx -y 

COPY ./nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]