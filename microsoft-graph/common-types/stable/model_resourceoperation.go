package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceOperation struct {
	ActionName   *string `json:"actionName,omitempty"`
	Description  *string `json:"description,omitempty"`
	Id           *string `json:"id,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
}
