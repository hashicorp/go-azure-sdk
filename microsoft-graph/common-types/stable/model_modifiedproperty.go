package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModifiedProperty struct {
	DisplayName *string `json:"displayName,omitempty"`
	NewValue    *string `json:"newValue,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	OldValue    *string `json:"oldValue,omitempty"`
}
