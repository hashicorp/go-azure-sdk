package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReputationCategory struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Vendor      *string `json:"vendor,omitempty"`
}
