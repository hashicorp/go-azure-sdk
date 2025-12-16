package devicecapacityinfos

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterCapacityViewData struct {
	Fqdn                        *string                `json:"fqdn,omitempty"`
	GpuCapacity                 *ClusterGpuCapacity    `json:"gpuCapacity,omitempty"`
	LastRefreshedTime           *string                `json:"lastRefreshedTime,omitempty"`
	MemoryCapacity              *ClusterMemoryCapacity `json:"memoryCapacity,omitempty"`
	TotalProvisionedNonHpnCores *int64                 `json:"totalProvisionedNonHpnCores,omitempty"`
}

func (o *ClusterCapacityViewData) GetLastRefreshedTimeAsTime() (*time.Time, error) {
	if o.LastRefreshedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRefreshedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ClusterCapacityViewData) SetLastRefreshedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRefreshedTime = &formatted
}
