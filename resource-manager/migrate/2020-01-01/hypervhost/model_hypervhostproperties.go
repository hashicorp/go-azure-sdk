package hypervhost

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVHostProperties struct {
	CreatedTimestamp *string               `json:"createdTimestamp,omitempty"`
	Errors           *[]HealthErrorDetails `json:"errors,omitempty"`
	Fqdn             *string               `json:"fqdn,omitempty"`
	RunAsAccountId   *string               `json:"runAsAccountId,omitempty"`
	UpdatedTimestamp *string               `json:"updatedTimestamp,omitempty"`
	Version          *string               `json:"version,omitempty"`
}
