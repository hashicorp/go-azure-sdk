package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelHealthThreshold struct {
	DefaultHealthyThreshold   *int64  `json:"defaultHealthyThreshold,omitempty"`
	DefaultUnhealthyThreshold *int64  `json:"defaultUnhealthyThreshold,omitempty"`
	HealthyThreshold          *int64  `json:"healthyThreshold,omitempty"`
	Id                        *string `json:"id,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	UnhealthyThreshold        *int64  `json:"unhealthyThreshold,omitempty"`
}
