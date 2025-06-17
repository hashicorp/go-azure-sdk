package siteinformationprotectionsensitivitylabelsublabel

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceRequest struct {
	DelegatedUserEmail      nullable.Type[string]    `json:"delegatedUserEmail,omitempty"`
	Locale                  nullable.Type[string]    `json:"locale,omitempty"`
	ProtectedContents       *[]beta.ProtectedContent `json:"protectedContents,omitempty"`
	SupportedContentFormats *[]string                `json:"supportedContentFormats,omitempty"`
}
