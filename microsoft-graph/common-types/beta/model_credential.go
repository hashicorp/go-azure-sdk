package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Credential struct {
	FieldId   *string `json:"fieldId,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Type      *string `json:"type,omitempty"`
	Value     *string `json:"value,omitempty"`
}
