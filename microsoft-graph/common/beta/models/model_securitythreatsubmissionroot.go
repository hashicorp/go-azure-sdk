package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionRoot struct {
	EmailThreatSubmissionPolicies *[]SecurityEmailThreatSubmissionPolicy `json:"emailThreatSubmissionPolicies,omitempty"`
	EmailThreats                  *[]SecurityEmailThreatSubmission       `json:"emailThreats,omitempty"`
	FileThreats                   *[]SecurityFileThreatSubmission        `json:"fileThreats,omitempty"`
	Id                            *string                                `json:"id,omitempty"`
	ODataType                     *string                                `json:"@odata.type,omitempty"`
	UrlThreats                    *[]SecurityUrlThreatSubmission         `json:"urlThreats,omitempty"`
}
