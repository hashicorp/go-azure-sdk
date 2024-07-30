package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRecommendedHuntingQuery struct {
	KqlText   *string `json:"kqlText,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
