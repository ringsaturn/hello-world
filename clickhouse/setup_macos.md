# Setup Clickhouse on macOS

```sh
# download binary
curl -O 'https://builds.clickhouse.tech/master/macos/clickhouse' && chmod a+x ./clickhouse

# download server config
wget "https://raw.githubusercontent.com/ClickHouse/ClickHouse/master/programs/server/users.xml"

# download client config
wget "https://raw.githubusercontent.com/ClickHouse/ClickHouse/master/programs/server/config.xml"

# start server
sudo ./clickhouse server --config-file=config.xml

# start client CLI
sudo ./clickhouse client --config-file=users.xml
```
