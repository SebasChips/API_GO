FROM mysql:9

ENV MYSQL_ROOT_PASSWORD=0000
ENV MYSQL_DATABASE=controlEscolar

COPY ./database/controlEscolar.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
