FROM mysql:8.0

COPY ../internal/app/queryservice/infra/config/my.cnf /etc/mysql/my.cnf
COPY ../internal/app/queryservice/infra/config/init.sh /docker-entrypoint-initdb.d/init-querydb.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-querydb.sh