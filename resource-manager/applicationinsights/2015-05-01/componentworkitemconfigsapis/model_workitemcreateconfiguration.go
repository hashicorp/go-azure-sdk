package componentworkitemconfigsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkItemCreateConfiguration struct {
	ConnectorDataConfiguration *string            `json:"ConnectorDataConfiguration,omitempty"`
	ConnectorId                *string            `json:"ConnectorId,omitempty"`
	ValidateOnly               *bool              `json:"ValidateOnly,omitempty"`
	WorkItemProperties         *map[string]string `json:"WorkItemProperties,omitempty"`
}
