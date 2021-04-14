FROM nginx:1.12

#  default conf for proxy service
COPY ./mybooksshelf.conf /etc/nginx/conf.d/mybooksshelf.conf
COPY ./users.conf /etc/nginx/conf.d/users.conf

# TLS and HTTPS certificate
COPY ./certs/ /etc/nginx/certs

# Not found 
COPY ./backend-not-found.html /var/www/html/backend-not-found.html
