package vcenter

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterProperties struct {
	CreatedTimestamp    *string               `json:"createdTimestamp,omitempty"`
	Errors              *[]HealthErrorDetails `json:"errors,omitempty"`
	Fqdn                *string               `json:"fqdn,omitempty"`
	InstanceUuid        *string               `json:"instanceUuid,omitempty"`
	PerfStatisticsLevel *string               `json:"perfStatisticsLevel,omitempty"`
	Port                *string               `json:"port,omitempty"`
	RunAsAccountId      *string               `json:"runAsAccountId,omitempty"`
	UpdatedTimestamp    *string               `json:"updatedTimestamp,omitempty"`
	Version             *string               `json:"version,omitempty"`
}
