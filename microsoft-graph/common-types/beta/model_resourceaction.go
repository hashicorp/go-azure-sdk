package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAction struct {
	AllowedResourceActions    *[]string `json:"allowedResourceActions,omitempty"`
	NotAllowedResourceActions *[]string `json:"notAllowedResourceActions,omitempty"`
	ODataType                 *string   `json:"@odata.type,omitempty"`
}
