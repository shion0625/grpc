#!/bin/bash
set -e

# Wait for the MySQL server to be available
until mysql -h"querydb" -u"root" -p"${MYSQL_ROOT_PASSWORD}" -e "SELECT 1"; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

# Configure replication
mysql -h"querydb" -u"root" -p"${MYSQL_ROOT_PASSWORD}" <<-EOSQL
  CHANGE MASTER TO
    MASTER_HOST='commanddb',
    MASTER_USER='root',
    MASTER_PASSWORD='rootpassword',
    MASTER_PORT=3306,
    MASTER_LOG_FILE='mysql-bin.000001',
    MASTER_LOG_POS=4;
  START SLAVE;
EOSQL