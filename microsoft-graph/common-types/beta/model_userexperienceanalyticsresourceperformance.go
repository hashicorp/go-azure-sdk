package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsResourcePerformance struct {
	AverageSpikeTimeScore           *int64   `json:"averageSpikeTimeScore,omitempty"`
	CpuSpikeTimePercentage          *float64 `json:"cpuSpikeTimePercentage,omitempty"`
	CpuSpikeTimePercentageThreshold *float64 `json:"cpuSpikeTimePercentageThreshold,omitempty"`
	CpuSpikeTimeScore               *int64   `json:"cpuSpikeTimeScore,omitempty"`
	DeviceCount                     *int64   `json:"deviceCount,omitempty"`
	DeviceId                        *string  `json:"deviceId,omitempty"`
	DeviceName                      *string  `json:"deviceName,omitempty"`
	DeviceResourcePerformanceScore  *int64   `json:"deviceResourcePerformanceScore,omitempty"`
	Id                              *string  `json:"id,omitempty"`
	Manufacturer                    *string  `json:"manufacturer,omitempty"`
	Model                           *string  `json:"model,omitempty"`
	ODataType                       *string  `json:"@odata.type,omitempty"`
	RamSpikeTimePercentage          *float64 `json:"ramSpikeTimePercentage,omitempty"`
	RamSpikeTimePercentageThreshold *float64 `json:"ramSpikeTimePercentageThreshold,omitempty"`
	RamSpikeTimeScore               *int64   `json:"ramSpikeTimeScore,omitempty"`
}
