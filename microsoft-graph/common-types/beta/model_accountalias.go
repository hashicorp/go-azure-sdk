package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountAlias struct {
	Id        *string `json:"id,omitempty"`
	IdType    *string `json:"idType,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
