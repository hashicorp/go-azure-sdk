package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressionInputObject struct {
	Definition *ObjectDefinition           `json:"definition,omitempty"`
	ODataType  *string                     `json:"@odata.type,omitempty"`
	Properties *[]StringKeyObjectValuePair `json:"properties,omitempty"`
}
