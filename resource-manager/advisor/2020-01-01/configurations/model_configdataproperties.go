package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigDataProperties struct {
	Digests         *[]DigestConfig `json:"digests,omitempty"`
	Exclude         *bool           `json:"exclude,omitempty"`
	LowCPUThreshold *CPUThreshold   `json:"lowCpuThreshold,omitempty"`
}
