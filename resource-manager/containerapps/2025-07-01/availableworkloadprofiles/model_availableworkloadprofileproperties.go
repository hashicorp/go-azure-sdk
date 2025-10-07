package availableworkloadprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableWorkloadProfileProperties struct {
	Applicability *Applicability `json:"applicability,omitempty"`
	Category      *string        `json:"category,omitempty"`
	Cores         *int64         `json:"cores,omitempty"`
	DisplayName   *string        `json:"displayName,omitempty"`
	Gpus          *int64         `json:"gpus,omitempty"`
	MemoryGiB     *int64         `json:"memoryGiB,omitempty"`
}
