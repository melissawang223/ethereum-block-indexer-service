FROM golang:1.16

ENV TZ=Asia/Taipei

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update \
    && apt-get -y upgrade

RUN apt-get install -y apt-utils \
    supervisor \
    net-tools \
    vim

RUN mkdir -p /var/www/html/blockchain \
    && mkdir -p /var/log/supervisor

COPY supervisor/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

WORKDIR /var/www/html/rf-stock-member

EXPOSE 80 443

CMD ["/usr/bin/supervisord"]
