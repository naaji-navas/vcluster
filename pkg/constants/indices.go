package constants

import "time"

const (
	IndexByPhysicalName  = "IndexByPhysicalName"
	IndexByVirtualName   = "IndexByVirtualName"
	IndexByAssigned      = "IndexByAssigned"
	IndexByStorageClass  = "IndexByStorageClass"
	IndexByIngressSecret = "IndexByIngressSecret"
	IndexByPodSecret     = "IndexByPodSecret"
	IndexByConfigMap     = "IndexByConfigMap"
	// IndexByHostName is used to map rewritten hostnames(advertised as node addresses) to nodenames
	IndexByHostName = "IndexByHostName"

	IndexByClusterIP = "IndexByClusterIP"
)

const DefaultCacheSyncTimeout = time.Minute * 15
