package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DistributionConfiguration = Ray{}

type Ray struct {
	Address                  *string `json:"address,omitempty"`
	DashboardPort            *int64  `json:"dashboardPort,omitempty"`
	HeadNodeAdditionalArgs   *string `json:"headNodeAdditionalArgs,omitempty"`
	IncludeDashboard         *bool   `json:"includeDashboard,omitempty"`
	Port                     *int64  `json:"port,omitempty"`
	WorkerNodeAdditionalArgs *string `json:"workerNodeAdditionalArgs,omitempty"`

	// Fields inherited from DistributionConfiguration
}

var _ json.Marshaler = Ray{}

func (s Ray) MarshalJSON() ([]byte, error) {
	type wrapper Ray
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Ray: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Ray: %+v", err)
	}
	decoded["distributionType"] = "Ray"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Ray: %+v", err)
	}

	return encoded, nil
}
