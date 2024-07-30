package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolationPolicyTip struct {
	ComplianceUrl                *string   `json:"complianceUrl,omitempty"`
	GeneralText                  *string   `json:"generalText,omitempty"`
	MatchedConditionDescriptions *[]string `json:"matchedConditionDescriptions,omitempty"`
	ODataType                    *string   `json:"@odata.type,omitempty"`
}
