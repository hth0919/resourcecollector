package CRM




type ClusterInfo struct {
	MetricValue          []string           `protobuf:"bytes,1,rep,name=MetricValue,proto3" json:"MetricValue,omitempty"`
	Clustername          string             `protobuf:"bytes,2,opt,name=Clustername,proto3" json:"Clustername,omitempty"`
	KubeConfig           string             `protobuf:"bytes,3,opt,name=KubeConfig,proto3" json:"KubeConfig,omitempty"`
	AdminToken           string             `protobuf:"bytes,4,opt,name=AdminToken,proto3" json:"AdminToken,omitempty"`
	NodeList             []*NodeInfo        `protobuf:"bytes,5,rep,name=NodeList,proto3" json:"NodeList,omitempty"`
	ClusterMetricSum     map[string]float64 `protobuf:"bytes,6,rep,name=ClusterMetricSum,proto3" json:"ClusterMetricSum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Host                 string             `protobuf:"bytes,7,opt,name=Host,proto3" json:"Host,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
type NodeInfo struct {
	NodeName             string             `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	PodList              []*PodInfo         `protobuf:"bytes,2,rep,name=PodList,proto3" json:"PodList,omitempty"`
	NodeMetricSum        map[string]float64 `protobuf:"bytes,3,rep,name=NodeMetricSum,proto3" json:"NodeMetricSum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	CpuCores             float64            `protobuf:"fixed64,4,opt,name=CpuCores,proto3" json:"CpuCores,omitempty"`
	MemoryTotal          float64            `protobuf:"fixed64,5,opt,name=MemoryTotal,proto3" json:"MemoryTotal,omitempty"`
	ScrapeError          float64            `protobuf:"fixed64,6,opt,name=ScrapeError,proto3" json:"ScrapeError,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
type PodInfo struct {
	PodName              string             `protobuf:"bytes,1,opt,name=PodName,proto3" json:"PodName,omitempty"`
	PodMetrics           map[string]float64 `protobuf:"bytes,2,rep,name=PodMetrics,proto3" json:"PodMetrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}
type TimeTick struct {
	Tick                 int64    `protobuf:"varint,1,opt,name=Tick,proto3" json:"Tick,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
