package componentworkitemconfigsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkItemConfiguration struct {
	ConfigDisplayName *string `json:"ConfigDisplayName,omitempty"`
	ConfigProperties  *string `json:"ConfigProperties,omitempty"`
	ConnectorId       *string `json:"ConnectorId,omitempty"`
	Id                *string `json:"Id,omitempty"`
	IsDefault         *bool   `json:"IsDefault,omitempty"`
}
