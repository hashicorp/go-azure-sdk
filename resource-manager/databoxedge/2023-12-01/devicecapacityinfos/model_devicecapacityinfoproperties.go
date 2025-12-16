package devicecapacityinfos

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityInfoProperties struct {
	ClusterComputeCapacityInfo *ClusterCapacityViewData `json:"clusterComputeCapacityInfo,omitempty"`
	ClusterStorageCapacityInfo *ClusterStorageViewData  `json:"clusterStorageCapacityInfo,omitempty"`
	NodeCapacityInfos          *map[string]HostCapacity `json:"nodeCapacityInfos,omitempty"`
	TimeStamp                  *string                  `json:"timeStamp,omitempty"`
}

func (o *DeviceCapacityInfoProperties) GetTimeStampAsTime() (*time.Time, error) {
	if o.TimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *DeviceCapacityInfoProperties) SetTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeStamp = &formatted
}
