package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyIntegerValuePair struct {
	Key       *string `json:"key,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Value     *int64  `json:"value,omitempty"`
}
