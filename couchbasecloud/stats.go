package couchbasecloud

// ServiceStats encapsulates the services stats for a given node.
type ServiceStats struct {
	Services []string `json:"services"`
	NodeName string   `json:"nodeName"`
	Status   string   `json:"status"`
}

// BucketStats encapsulates the bucket stats for a given cluster.
type BucketStats struct {
	Status         string            `json:"status"`
	TotalCount     int               `json:"totalCount"`
	UnhealthyCount int               `json:"unhealthyCount"`
	HealthStats    map[string]string `json:"healthStats"`
}

// NodeStats encapsulates the node stats for a given cluster.
type NodeStats struct {
	Status          string         `json:"status"`
	TotalCount      int            `json:"totalCount"`
	FailedOverCount int            `json:"failedOverCount"`
	UnhealthyCount  int            `json:"unhealthyCount"`
	ServiceStats    []ServiceStats `json:"serviceStats"`
}

// ClusterHealth encapsulates the status of a cloud cluster.
type ClusterHealth struct {
	Status      string       `json:"status"`
	Health      string       `json:"health"`
	BucketStats *BucketStats `json:"bucketStats"`
	NodeStats   *NodeStats   `json:"nodeStats"`
}

