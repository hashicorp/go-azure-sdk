package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRuleSchedule struct {
	NextRunDateTime *string `json:"nextRunDateTime,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	Period          *string `json:"period,omitempty"`
}
