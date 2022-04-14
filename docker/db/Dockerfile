FROM mysql:5.7

# https://hub.docker.com/_/mysql

# root password
ENV MYSQL_ROOT_PASSWORD 123456

# add a database
ENV MYSQL_DATABASE rf

# set timezone
ENV TZ=Asia/Taipei
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
