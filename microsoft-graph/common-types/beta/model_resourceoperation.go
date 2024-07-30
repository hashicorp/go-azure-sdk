package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceOperation struct {
	ActionName                *string `json:"actionName,omitempty"`
	Description               *string `json:"description,omitempty"`
	EnabledForScopeValidation *bool   `json:"enabledForScopeValidation,omitempty"`
	Id                        *string `json:"id,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	Resource                  *string `json:"resource,omitempty"`
	ResourceName              *string `json:"resourceName,omitempty"`
}
