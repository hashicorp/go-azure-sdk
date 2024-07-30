package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Phone struct {
	Language  *string    `json:"language,omitempty"`
	Number    *string    `json:"number,omitempty"`
	ODataType *string    `json:"@odata.type,omitempty"`
	Region    *string    `json:"region,omitempty"`
	Type      *PhoneType `json:"type,omitempty"`
}
