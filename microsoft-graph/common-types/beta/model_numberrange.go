package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberRange struct {
	LowerNumber *int64  `json:"lowerNumber,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	UpperNumber *int64  `json:"upperNumber,omitempty"`
}
