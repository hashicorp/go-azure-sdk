package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudApplicationEntityProperties struct {
	AdditionalData *interface{} `json:"additionalData,omitempty"`
	AppId          *int64       `json:"appId,omitempty"`
	AppName        *string      `json:"appName,omitempty"`
	FriendlyName   *string      `json:"friendlyName,omitempty"`
	InstanceName   *string      `json:"instanceName,omitempty"`
}
