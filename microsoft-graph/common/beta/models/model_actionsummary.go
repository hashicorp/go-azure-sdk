package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionSummary struct {
	Assigned  *int64  `json:"assigned,omitempty"`
	Available *int64  `json:"available,omitempty"`
	Exercised *int64  `json:"exercised,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
