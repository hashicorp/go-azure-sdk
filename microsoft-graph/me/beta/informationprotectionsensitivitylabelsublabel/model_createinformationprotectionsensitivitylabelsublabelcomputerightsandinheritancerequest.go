package informationprotectionsensitivitylabelsublabel

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceRequest struct {
	DelegatedUserEmail      nullable.Type[string]    `json:"delegatedUserEmail,omitempty"`
	Locale                  nullable.Type[string]    `json:"locale,omitempty"`
	ProtectedContents       *[]beta.ProtectedContent `json:"protectedContents,omitempty"`
	SupportedContentFormats *[]string                `json:"supportedContentFormats,omitempty"`
}
