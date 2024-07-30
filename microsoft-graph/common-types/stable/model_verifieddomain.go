package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedDomain struct {
	Capabilities *string `json:"capabilities,omitempty"`
	IsDefault    *bool   `json:"isDefault,omitempty"`
	IsInitial    *bool   `json:"isInitial,omitempty"`
	Name         *string `json:"name,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Type         *string `json:"type,omitempty"`
}
