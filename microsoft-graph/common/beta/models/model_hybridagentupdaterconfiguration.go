package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridAgentUpdaterConfiguration struct {
	AllowUpdateConfigurationOverride *bool         `json:"allowUpdateConfigurationOverride,omitempty"`
	DeferUpdateDateTime              *string       `json:"deferUpdateDateTime,omitempty"`
	ODataType                        *string       `json:"@odata.type,omitempty"`
	UpdateWindow                     *UpdateWindow `json:"updateWindow,omitempty"`
}
