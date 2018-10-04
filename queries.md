# Cadvisor
## Monitoring container:
* CPU total: `sum by(cpu) (rate(container_cpu_usage_seconds_total[2m]))`
* RAM: `sum by(instance) (container_memory_usage_bytes)`
* Network
    - Transmit: `sum by(instance, interface) (rate(container_network_transmit_bytes_total{interface="eth0"}[1m]))`
    - Receive: `sum by(instance, interface) (rate(container_network_receive_bytes_total{interface="eth0"}[1m]))`

# Node exporter
* Disk IO: `(max(avg(irate(node_disk_io_time_seconds_total[10m])) by (instance, device)) by (instance))/10`
* Disk Space: `max((node_filesystem_files{fstype="ext4"} - node_filesystem_files_free{fstype="ext4"}) / node_filesystem_files{fstype="ext4"} * 100) by(instance)`
* Memory Used: `((node_memory_MemTotal_bytes - node_memory_MemFree_bytes) / node_memory_MemTotal_bytes) * 100
*CPU: (1 - avg(irate(node_cpu_seconds_total{mode="idle"}[10m])) by (instance)) * 100`

# Webserver
## Nginx
* HTTP Connection State: `nginx_http_connections`
* Letancy: `sum by(le)(rate(nginx_http_request_duration_seconds_bucket[1m]))`
* HTTP Status code per sercond: `rate(nginx_http_requests_total[1m])`
## Custom Prometheus Exporter
* xxx page per second during 1 mitute: `rate(page_counter{page="xx_page"}[1m])`
* yy page per second during 1 mitute: `rate(page_counter{page="yy_page"}[1m])`

# Postgres
* Insert: `rate(pg_stat_database_tup_inserted[1m])`
* Update: `rate(pg_stat_database_tup_updated[1m])`
* Delete: `rate(pg_stat_database_tup_deleted[1m])`
* Fetch: `rate(pg_stat_database_tup_fetched[1m])`



# References
* https://www.dogsbodytechnology.com/blog/turning-prometheus-data-into-metrics-for-alerting/