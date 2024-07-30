package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSignInInsight struct {
	Id                     *string `json:"id,omitempty"`
	InsightCreatedDateTime *string `json:"insightCreatedDateTime,omitempty"`
	LastSignInDateTime     *string `json:"lastSignInDateTime,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
}
