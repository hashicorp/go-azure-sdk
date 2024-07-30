package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadDetail struct {
	Coachmarks  *[]PayloadCoachmark `json:"coachmarks,omitempty"`
	Content     *string             `json:"content,omitempty"`
	ODataType   *string             `json:"@odata.type,omitempty"`
	PhishingUrl *string             `json:"phishingUrl,omitempty"`
}
