package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContainerProperties struct {
	AuthType                    *AuthType              `json:"authType,omitempty"`
	CreatedTime                 *string                `json:"createdTime,omitempty"`
	EnvironmentVariables        *[]EnvironmentVariable `json:"environmentVariables,omitempty"`
	Image                       string                 `json:"image"`
	IsMain                      bool                   `json:"isMain"`
	LastModifiedTime            *string                `json:"lastModifiedTime,omitempty"`
	PasswordSecret              *string                `json:"passwordSecret,omitempty"`
	StartUpCommand              *string                `json:"startUpCommand,omitempty"`
	TargetPort                  *string                `json:"targetPort,omitempty"`
	UserManagedIdentityClientId *string                `json:"userManagedIdentityClientId,omitempty"`
	UserName                    *string                `json:"userName,omitempty"`
	VolumeMounts                *[]VolumeMount         `json:"volumeMounts,omitempty"`
}
