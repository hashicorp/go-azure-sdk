package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Device struct {
	DeploymentManifest *DeploymentManifest `json:"deploymentManifest,omitempty"`
	DisplayName        *string             `json:"displayName,omitempty"`
	Enabled            *bool               `json:"enabled,omitempty"`
	Etag               *string             `json:"etag,omitempty"`
	Id                 *string             `json:"id,omitempty"`
	Organizations      *[]string           `json:"organizations,omitempty"`
	Provisioned        *bool               `json:"provisioned,omitempty"`
	Simulated          *bool               `json:"simulated,omitempty"`
	Template           *string             `json:"template,omitempty"`
	Type               *[]DeviceType       `json:"type,omitempty"`
}
