package mesecurityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest struct {
	ContentInfo            *ContentInfo                    `json:"contentInfo,omitempty"`
	DowngradeJustification *SecurityDowngradeJustification `json:"downgradeJustification,omitempty"`
}
