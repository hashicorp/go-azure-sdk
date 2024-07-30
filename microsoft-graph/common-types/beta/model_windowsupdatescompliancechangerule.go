package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesComplianceChangeRule struct {
	CreatedDateTime       *string `json:"createdDateTime,omitempty"`
	LastEvaluatedDateTime *string `json:"lastEvaluatedDateTime,omitempty"`
	LastModifiedDateTime  *string `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
