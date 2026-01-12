#Задание №1
#FROM nginx:alpine
#COPY index.html /usr/share/nginx/html/

#Задание №2
FROM nginx:alpine
ARG MESSAGE="Hello, Docker!"
RUN echo "<!DOCTYPE html><html><head><title>Hello</title></head><body><h1>${MESSAGE}</h1></body></html>" > /usr/share/nginx/html/index.html
#\Learning_EM\docker-hello> docker build --no-cache --build-arg MESSAGE="Hi, just testing..." -t hello-docker .
#docker run -d -p 8080:80 --rm --name hello-nginx hello-docker