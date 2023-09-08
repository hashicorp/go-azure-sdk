package mesecurityinformationprotectionsensitivitylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultRequest struct {
	ClassificationResults *[]SecurityClassificationResult `json:"classificationResults,omitempty"`
	ContentInfo           *ContentInfo                    `json:"contentInfo,omitempty"`
}
