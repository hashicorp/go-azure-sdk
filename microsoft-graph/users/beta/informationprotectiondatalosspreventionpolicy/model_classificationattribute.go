package informationprotectiondatalosspreventionpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationAttribute struct {
	Confidence *int64  `json:"confidence,omitempty"`
	Count      *int64  `json:"count,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
