package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsExpiration struct {
	Frequency     *string `json:"frequency,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	StartDateTime *string `json:"startDateTime,omitempty"`
}
