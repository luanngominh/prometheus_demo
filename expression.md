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

## 
* xxx page:
* yy page:













# References
* https://www.dogsbodytechnology.com/blog/turning-prometheus-data-into-metrics-for-alerting/