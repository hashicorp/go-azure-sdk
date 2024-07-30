package securityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest struct {
	ContentInfo            *ContentInfo                    `json:"contentInfo,omitempty"`
	DowngradeJustification *SecurityDowngradeJustification `json:"downgradeJustification,omitempty"`
}
