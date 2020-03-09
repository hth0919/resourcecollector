package CRM

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func initmeticmap() *map[string]Metric {
	mm := map[string]Metric{
		"cpu_cfs_periods_total":                  {0},
		"cpu_cfs_throttled_periods_total":        {0},
		"cpu_cfs_throttled_seconds_total":        {0},
		"cpu_load_average_10s":                   {0},
		"cpu_system_seconds_total":               {0},
		"cpu_usage_seconds_total":                {0},
		"cpu_user_seconds_total":                 {0},
		"fs_inodes_free":                         {0},
		"fs_inodes_total":                        {0},
		"fs_io_current":                          {0},
		"fs_io_time_seconds_total":               {0},
		"fs_io_time_weighted_seconds_total":      {0},
		"fs_limit_bytes":                         {0},
		"fs_read_seconds_total":                  {0},
		"fs_reads_bytes_total":                   {0},
		"fs_reads_merged_total":                  {0},
		"fs_reads_total":                         {0},
		"fs_sector_reads_total":                  {0},
		"fs_sector_writes_total":                 {0},
		"fs_usage_bytes":                         {0},
		"fs_write_seconds_total":                 {0},
		"fs_writes_bytes_total":                  {0},
		"fs_writes_merged_total":                 {0},
		"fs_writes_total":                        {0},
		"last_seen":                              {0},
		"memory_cache":                           {0},
		"memory_failcnt":                         {0},
		"memory_failures_total":                  {0},
		"memory_mapped_file":                     {0},
		"memory_max_usage_bytes":                 {0},
		"memory_rss":                             {0},
		"memory_swap":                            {0},
		"memory_usage_bytes":                     {0},
		"memory_working_set_bytes":               {0},
		"network_receive_bytes_total":            {0},
		"network_receive_errors_total":           {0},
		"network_receive_packets_dropped_total":  {0},
		"network_receive_packets_total":          {0},
		"network_transmit_bytes_total":           {0},
		"network_transmit_errors_total":          {0},
		"network_transmit_packets_dropped_total": {0},
		"network_transmit_packets_total":         {0},
		"scrape_error 0":                         {0},
		"spec_cpu_period":                        {0},
		"spec_cpu_quota":                         {0},
		"spec_cpu_shares":                        {0},
		"spec_memory_limit_bytes":                {0},
		"spec_memory_reservation_limit_bytes":    {0},
		"spec_memory_swap_limit_bytes":           {0},
		"start_time_seconds":                     {0},
		"tasks_state":                            {0},
		"machine_cpu_cores":                      {0},
		"machine_memory_bytes":                   {0},
	}
	return &mm
}

func (ci *ClusterInfo)NewClusterClient(masterUri string) {
	var clientset ClientSet
	ci.MetricValue = []string{"container_cpu_cfs_periods_total", "container_cpu_cfs_throttled_periods_total", "container_cpu_cfs_throttled_seconds_total", "container_cpu_load_average_10s", "container_cpu_system_seconds_total", "container_cpu_usage_seconds_total", "container_cpu_user_seconds_total", "container_fs_inodes_free", "container_fs_inodes_total", "container_fs_io_current", "container_fs_io_time_seconds_total", "container_fs_io_time_weighted_seconds_total", "container_fs_limit_bytes", "container_fs_read_seconds_total", "container_fs_reads_bytes_total", "container_fs_reads_merged_total", "container_fs_reads_total", "container_fs_sector_reads_total", "container_fs_sector_writes_total", "container_fs_usage_bytes", "container_fs_write_seconds_total", "container_fs_writes_bytes_total", "container_fs_writes_merged_total", "container_fs_writes_total", "container_last_seen", "container_memory_cache", "container_memory_failcnt", "container_memory_failures_total", "container_memory_mapped_file", "container_memory_max_usage_bytes", "container_memory_rss", "container_memory_swap", "container_memory_usage_bytes", "container_memory_working_set_bytes", "container_network_receive_bytes_total", "container_network_receive_errors_total", "container_network_receive_packets_dropped_total", "container_network_receive_packets_total", "container_network_transmit_bytes_total", "container_network_transmit_errors_total", "container_network_transmit_packets_dropped_total", "container_network_transmit_packets_total", "container_scrape_error", "container_spec_cpu_period", "container_spec_cpu_quota", "container_spec_cpu_shares", "container_spec_memory_limit_bytes", "container_spec_memory_reservation_limit_bytes", "container_spec_memory_swap_limit_bytes", "container_start_time_seconds", "container_tasks_state", "machine_cpu_cores", "machine_memory_bytes"}
	ci.ClusterMetricSum =*initmeticmap()
	if home := homeDir(); home != "" {
		ci.KubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		ci.KubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(masterUri, *ci.KubeConfig)
	if err != nil {
		panic(err.Error())
	}
	ci.Host = config.Host
	clientset.clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secrets, _ := clientset.clientSet.CoreV1().Secrets(metav1.NamespaceAll).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	ci.AdminToken = string(secrets.Items[0].Data["token"])
	ci.AdminToken = strings.TrimSpace(ci.AdminToken)
	for i:=0;i<len(secrets.Items);i++ {
		if strings.Contains(secrets.Items[i].Name, "sa-admin") {
			ci.AdminToken = string(secrets.Items[i].Data["token"])
			ci.AdminToken = strings.TrimSpace(ci.AdminToken)
		}
	}

}

func (ci *ClusterInfo)NodeListInit() {
	var ni *NodeInfo
	var clientset ClientSet
	nodes, err :=clientset.clientSet.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ci.NodeList = make([]*NodeInfo,0,len(nodes.Items))
	for i:=0;i<len(nodes.Items);i++ {
		ni = &NodeInfo{
			NodeName:      "",
			PodList:       []*PodInfo{},
			NodeMetricSum: map[string]Metric{},
			CpuCores:      0,
			MemoryTotal:   0,
			ScrapeError:   0,
		}
		ni.NodeMetricSum = *initmeticmap()
		ci.NodeList = append(ci.NodeList, ni)
	}
	for i:=0;i<len(nodes.Items);i++ {
		ci.NodeList[i].NodeName = nodes.Items[i].Name
		responseTokenizer(ci,ci.NodeList[i].NodeName,i)
	}

}

func (ci *ClusterInfo)CalculateClusterMetricSum() map[string]Metric{
	var sum *Metric
	sum = &Metric{
		Value: 0,
	}

	for i:=0;i<len(ci.MetricValue);i++ {
		sum = &Metric{
			Value: 0,
		}
		for j := 0;j<len(ci.NodeList);j++ {
			sum.Value += ci.NodeList[j].NodeMetricSum[ci.MetricValue[i]].Value
			ci.ClusterMetricSum[ci.MetricValue[i]] = *sum
		}
	}
	return ci.ClusterMetricSum
}

func (ci *ClusterInfo)CalculateNodeMetricSum(index int) map[string]Metric {
	var sum *Metric
	sum = &Metric{
		Value:      0,
	}
	for i:=0;i<len(ci.MetricValue);i++ {
		sum = &Metric{
			Value:      0,
		}
		for j := 0;j<len(ci.NodeList[index].PodList);j++ {
			sum.Value += ci.NodeList[index].PodList[j].PodMetrics[ci.MetricValue[i]].Value
			ci.NodeList[index].NodeMetricSum[ci.MetricValue[i]] = *sum
		}
	}
	return ci.NodeList[index].NodeMetricSum
}

func FindOrMakePodInfo(ci *ClusterInfo, name string, nodeindex int) (*PodInfo, int) {
	var result PodInfo
	result = PodInfo{
		PodName:    name,
		PodMetrics: *initmeticmap(),
	}

	for i:=0;i<len(ci.NodeList[nodeindex].PodList);i++ {
		if ci.NodeList[nodeindex].PodList[i].PodName == name {
			return ci.NodeList[nodeindex].PodList[i], i
		}
	}
	return &result, -1
}

func responseTokenizer(ci *ClusterInfo, nodename string,indexnum int) {
	//var pil []*podInfo
	var pi *PodInfo
	var isfind int
	var met *Metric

	url := "" + ci.Host + "/api/v1/nodes/" + nodename + "/proxy/metrics/cadvisor"
	url = strings.TrimSpace(url)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("Authorization", "Bearer " + ci.AdminToken)

	// Client객체에서 Request 실행
	client := &http.Client{Transport:tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	// 결과 출력
	var temp []string

	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로

	replacer := strings.NewReplacer("{", " ",
		"cadvisorRevision=", "",
		"container=", "",
		"cadvisorVersion=", "",
		"container_name=", "",
		"dockerVersion=", "",
		"id=", "",
		"cpu=", "",
		"device=", "",
		"failure_type=", "",
		"kernelVersion=", "",
		"image=", "",
		"osVersion=", "",
		"name=", "",
		"interface=", "",
		"namespace=", "",
		"pod=", "",
		"pod_name=", "",
		"state=", "",
		"scope=", "",
		"\"", "",
		",", " ",
		"}", "")
	str = replacer.Replace(str)
	temp = strings.Split(str, "\n")

	for i:=0;i<len(temp);i++ {
		met = &Metric{
			Value : 0,
		}
		atom := strings.Split(temp[i]," ")

		if atom[0] == "#" || atom[0] == "cadvisor_version_info" || atom[0] == "" {

			continue
		}
		if valueCheck(atom[0]) == 1 {
			met.Value, _ = strconv.ParseFloat(atom[len(atom)-1],64)
		} else {
			met.Value, _ = strconv.ParseFloat(atom[len(atom)-2],64)
		}
		index := podNamePlace(atom[0])
		if index < 2 {
			switch index {
			case 0:
				continue
			case -1:
				ci.NodeList[indexnum].ScrapeError = met.Value
			case -2:
				ci.NodeList[indexnum].CpuCores = met.Value
			case -3:
				ci.NodeList[indexnum].MemoryTotal = met.Value
			}
		} else {
			pi, isfind = FindOrMakePodInfo(ci, atom[index], indexnum)
		}
		if isfind == -1 {
			pi.PodMetrics[atom[0]] = *met
			ci.NodeList[indexnum].PodList = append(ci.NodeList[indexnum].PodList,pi)
		}else {
			pi.PodMetrics[atom[0]] = *met
			ci.NodeList[indexnum].PodList[isfind].PodMetrics = pi.PodMetrics
		}

	}
}

func valueCheck(metricName string) int {

	switch metricName {
	case "container_start_time_seconds": return 1
	case "container_spec_memory_swap_limit_bytes": return 1
	case "container_spec_memory_reservation_limit_bytes": return 1
	case "container_spec_memory_limit_bytes": return 1
	case "container_spec_cpu_shares": return 1
	case "container_spec_cpu_quota": return 1
	case "container_spec_cpu_period": return 1
	case "container_scrape_error": return 1
	case "machine_cpu_cores": return 1
	case "machine_memory_bytes": return 1
	}
	return 0
}

func podNamePlace(metricName string) int {
	switch metricName {
	case "container_cpu_cfs_periods_total": return 7
	case "container_cpu_cfs_throttled_periods_total": return 7
	case "container_cpu_cfs_throttled_seconds_total": return 7
	case "container_cpu_load_average_10s": return 7
	case "container_cpu_system_seconds_total": return 7
	case "container_cpu_user_seconds_total": return 7
	case "container_last_seen": return 7
	case "container_memory_cache": return 7
	case "container_memory_failcnt": return 7
	case "container_memory_mapped_file": return 7
	case "container_memory_max_usage_bytes": return 7
	case "container_memory_rss": return 7
	case "container_memory_swap": return 7
	case "container_memory_usage_bytes": return 7
	case "container_memory_working_set_bytes": return 7
	case "container_spec_cpu_period": return 7
	case "container_spec_cpu_quota": return 7
	case "container_spec_cpu_shares": return 7
	case "container_spec_memory_limit_bytes": return 7
	case "container_spec_memory_reservation_limit_bytes": return 7
	case "container_spec_memory_swap_limit_bytes": return 7
	case "container_start_time_seconds": return 7
	case "container_tasks_state": return 7
	case "container_cpu_usage_seconds_total": return 8
	case "container_fs_inodes_free": return 8
	case "container_fs_inodes_total": return 8
	case "container_fs_io_current": return 8
	case "container_fs_io_time_seconds_total": return 8
	case "container_fs_io_time_weighted_seconds_total": return 8
	case "container_fs_limit_bytes": return 8
	case "container_fs_read_seconds_total": return 8
	case "container_fs_reads_bytes_total": return 8
	case "container_fs_reads_merged_total": return 8
	case "container_fs_reads_total": return 8
	case "container_fs_sector_reads_total": return 8
	case "container_fs_sector_writes_total": return 8
	case "container_fs_usage_bytes": return 8
	case "container_fs_write_seconds_total": return 8
	case "container_fs_writes_bytes_total": return 8
	case "container_fs_writes_merged_total": return 8
	case "container_fs_writes_total": return 8
	case "container_memory_failures_total": return 8
	case "container_network_receive_bytes_total": return 8
	case "container_network_receive_errors_total": return 8
	case "container_network_receive_packets_dropped_total": return 8
	case "container_network_receive_packets_total": return 8
	case "container_network_transmit_bytes_total": return 8
	case "container_network_transmit_errors_total": return 8
	case "container_network_transmit_packets_dropped_total": return 8
	case "container_network_transmit_packets_total": return 8
	case "cadvisor_version_info": return 0
	case "container_scrape_error": return -1
	case "machine_cpu_cores": return -2
	case "machine_memory_bytes": return -3
	}
	return 0
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

