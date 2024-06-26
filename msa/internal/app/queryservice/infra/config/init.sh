#!/bin/bash
set -e

# DDLスクリプトを実行
for script in /docker-entrypoint-initdb.d/ddl/*.sql; do
  echo "Running $script"
  mysql --host=commanddb --user=root --password=rootpassword --ssl-ca=/etc/mysql/ssl/ca-cert.pem --ssl-cert=/etc/mysql/ssl/server-cert.pem --ssl-key=/etc/mysql/ssl/server-key.pem < "$script"
done

# Get master status from commanddb
MASTER_STATUS=$(mysql -h "commanddb" -u "root" -p"${MYSQL_ROOT_PASSWORD}" -e "SHOW MASTER STATUS\G")
MASTER_LOG_FILE=$(echo "$MASTER_STATUS" | grep 'File:' | awk '{print $2}')
MASTER_LOG_POS=$(echo "$MASTER_STATUS" | grep 'Position:' | awk '{print $2}')

# Configure replication
mysql -u "root" -p"${MYSQL_ROOT_PASSWORD}" --ssl-ca=/etc/mysql/ssl/ca-cert.pem --ssl-cert=/etc/mysql/ssl/server-cert.pem --ssl-key=/etc/mysql/ssl/server-key.pem --ssl-mode=VERIFY_CA <<-EOSQL
  STOP SLAVE;
  CHANGE MASTER TO
    MASTER_HOST='commanddb',
    MASTER_USER='root',
    MASTER_PASSWORD='rootpassword',
    MASTER_PORT=3306,
    MASTER_LOG_FILE='$MASTER_LOG_FILE',
    MASTER_LOG_POS=$MASTER_LOG_POS;
  START SLAVE;
EOSQL