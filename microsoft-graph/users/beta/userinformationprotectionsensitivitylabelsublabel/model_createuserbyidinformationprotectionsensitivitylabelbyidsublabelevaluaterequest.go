package userinformationprotectionsensitivitylabelsublabel

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdInformationProtectionSensitivityLabelByIdSublabelEvaluateRequest struct {
	CurrentLabel             *CurrentLabel              `json:"currentLabel,omitempty"`
	DiscoveredSensitiveTypes *[]DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
}
