package informationprotectionpolicylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionPolicyLabelEvaluateRemovalRequest struct {
	ContentInfo            *ContentInfo            `json:"contentInfo,omitempty"`
	DowngradeJustification *DowngradeJustification `json:"downgradeJustification,omitempty"`
}
