package CRM




type ClusterInfo struct {
	MetricValue          []string           `protobuf:"bytes,1,rep,name=metricValue" json:"metricValue,omitempty"`
	Clustername          string            `protobuf:"bytes,2,opt,name=clustername" json:"clustername,omitempty"`
	KubeConfig           *string            `protobuf:"bytes,3,opt,name=kubeConfig" json:"kubeConfig,omitempty"`
	AdminToken           string            `protobuf:"bytes,4,opt,name=adminToken" json:"adminToken,omitempty"`
	NodeList             []*NodeInfo        `protobuf:"bytes,5,rep,name=NodeList" json:"NodeList,omitempty"`
	ClusterMetricSum     map[string]Metric `protobuf:"bytes,6,rep,name=ClusterMetricSum" json:"ClusterMetricSum,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Host                 string            `protobuf:"bytes,7,opt,name=host" json:"host,omitempty"`
}
type NodeInfo struct {
	NodeName             string            `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
	PodList              []*PodInfo         `protobuf:"bytes,2,rep,name=PodList" json:"PodList,omitempty"`
	NodeMetricSum        map[string]Metric `protobuf:"bytes,3,rep,name=NodeMetricSum" json:"NodeMetricSum,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CpuCores             float64           `protobuf:"fixed32,4,opt,name=cpuCores" json:"cpuCores,omitempty"`
	MemoryTotal          float64           `protobuf:"fixed32,5,opt,name=MemoryTotal" json:"MemoryTotal,omitempty"`
	ScrapeError          float64           `protobuf:"fixed32,6,opt,name=scrapeError" json:"scrapeError,omitempty"`
}
type PodInfo struct {
	PodName              string            `protobuf:"bytes,1,opt,name=PodName" json:"PodName,omitempty"`
	PodMetrics           map[string]Metric `protobuf:"bytes,2,rep,name=PodMetrics" json:"PodMetrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}
type Metric struct {
	Value                float64 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
}

