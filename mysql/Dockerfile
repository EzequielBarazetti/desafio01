FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD root
ENV LANG=C.UTF-8

COPY script.sql docker-entrypoint-initdb.d/

VOLUME ["/var/lib/mysql"]

ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 3306
CMD ["mysqld"]