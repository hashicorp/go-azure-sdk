package informationprotectionpolicylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionPolicyLabelEvaluateClassificationResultRequest struct {
	ClassificationResults *[]ClassificationResult `json:"classificationResults,omitempty"`
	ContentInfo           *ContentInfo            `json:"contentInfo,omitempty"`
}
