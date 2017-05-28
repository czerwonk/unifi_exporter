package common

import "github.com/prometheus/client_golang/prometheus"

var (
	AccessPointUpDesc      *prometheus.Desc
	AccessPointClientsDesc *prometheus.Desc
)

func init() {
	labels := make([]string, 0)
	labels = append(labels, "site", "ap_name")
	AccessPointUpDesc = prometheus.NewDesc("wifi_ap_up", "0 = AP is down, 1 = AP is running", labels, nil)

	labels = append(labels, "radio", "ssid")
	AccessPointClientsDesc = prometheus.NewDesc("wifi_ap_clients", "Number of clients on AP by radio and ssid", labels, nil)
}
