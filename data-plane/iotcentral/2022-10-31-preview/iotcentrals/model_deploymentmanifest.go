package iotcentrals

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentManifest struct {
	Data          interface{} `json:"data"`
	DisplayName   *string     `json:"displayName,omitempty"`
	Etag          *string     `json:"etag,omitempty"`
	Id            *string     `json:"id,omitempty"`
	Organizations *[]string   `json:"organizations,omitempty"`
}
