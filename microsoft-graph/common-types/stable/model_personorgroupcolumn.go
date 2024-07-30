package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonOrGroupColumn struct {
	AllowMultipleSelection *bool   `json:"allowMultipleSelection,omitempty"`
	ChooseFromType         *string `json:"chooseFromType,omitempty"`
	DisplayAs              *string `json:"displayAs,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
}
