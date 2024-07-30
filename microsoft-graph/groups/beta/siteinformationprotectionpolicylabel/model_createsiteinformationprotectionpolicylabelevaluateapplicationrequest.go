package siteinformationprotectionpolicylabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionPolicyLabelEvaluateApplicationRequest struct {
	ContentInfo     *ContentInfo     `json:"contentInfo,omitempty"`
	LabelingOptions *LabelingOptions `json:"labelingOptions,omitempty"`
}
