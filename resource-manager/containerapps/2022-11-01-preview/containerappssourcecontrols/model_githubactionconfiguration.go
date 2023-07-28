package containerappssourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GithubActionConfiguration struct {
	AzureCredentials *AzureCredentials `json:"azureCredentials,omitempty"`
	ContextPath      *string           `json:"contextPath,omitempty"`
	Image            *string           `json:"image,omitempty"`
	Os               *string           `json:"os,omitempty"`
	PublishType      *string           `json:"publishType,omitempty"`
	RegistryInfo     *RegistryInfo     `json:"registryInfo,omitempty"`
	RuntimeStack     *string           `json:"runtimeStack,omitempty"`
	RuntimeVersion   *string           `json:"runtimeVersion,omitempty"`
}
