module resourcecollector/crm

go 1.12

require (
	github.com/golang/protobuf v1.3.4
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	google.golang.org/grpc v1.27.1
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
	 k8s.io/api v0.0.0
        k8s.io/apimachinery v0.0.0
        k8s.io/client-go v0.0.0
        resourcecollector/crm v0.0.0
	k8s.io/utils v0.0.0-20200229041039-0a110f9eb7ab // indirect
)

replace (
        k8s.io/api => ./staging/src/k8s.io/api
        k8s.io/apimachinery => ./staging/src/k8s.io/apimachinery
        k8s.io/client-go => ./staging/src/k8s.io/client-go
        resourcecollector/crm => ./cmd/creator/crm
)
