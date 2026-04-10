package provider

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteConfigPropertiesDictionary struct {
	JavaVersion           *string `json:"javaVersion,omitempty"`
	LinuxFxVersion        *string `json:"linuxFxVersion,omitempty"`
	PowerShellVersion     *string `json:"powerShellVersion,omitempty"`
	Use32BitWorkerProcess *bool   `json:"use32BitWorkerProcess,omitempty"`
}
