package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAction struct {
	Name       *string         `json:"name,omitempty"`
	ODataType  *string         `json:"@odata.type,omitempty"`
	Properties *[]KeyValuePair `json:"properties,omitempty"`
}
