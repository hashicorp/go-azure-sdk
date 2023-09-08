package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Agreement struct {
	Acceptances                       *[]AgreementAcceptance       `json:"acceptances,omitempty"`
	DisplayName                       *string                      `json:"displayName,omitempty"`
	File                              *AgreementFile               `json:"file,omitempty"`
	Files                             *[]AgreementFileLocalization `json:"files,omitempty"`
	Id                                *string                      `json:"id,omitempty"`
	IsPerDeviceAcceptanceRequired     *bool                        `json:"isPerDeviceAcceptanceRequired,omitempty"`
	IsViewingBeforeAcceptanceRequired *bool                        `json:"isViewingBeforeAcceptanceRequired,omitempty"`
	ODataType                         *string                      `json:"@odata.type,omitempty"`
	TermsExpiration                   *TermsExpiration             `json:"termsExpiration,omitempty"`
	UserReacceptRequiredFrequency     *string                      `json:"userReacceptRequiredFrequency,omitempty"`
}
