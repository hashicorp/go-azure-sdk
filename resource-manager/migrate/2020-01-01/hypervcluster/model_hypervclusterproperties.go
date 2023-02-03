package hypervcluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVClusterProperties struct {
	CreatedTimestamp *string               `json:"createdTimestamp,omitempty"`
	Errors           *[]HealthErrorDetails `json:"errors,omitempty"`
	Fqdn             *string               `json:"fqdn,omitempty"`
	FunctionalLevel  *int64                `json:"functionalLevel,omitempty"`
	HostFqdnList     *[]string             `json:"hostFqdnList,omitempty"`
	RunAsAccountId   *string               `json:"runAsAccountId,omitempty"`
	Status           *string               `json:"status,omitempty"`
	UpdatedTimestamp *string               `json:"updatedTimestamp,omitempty"`
}
