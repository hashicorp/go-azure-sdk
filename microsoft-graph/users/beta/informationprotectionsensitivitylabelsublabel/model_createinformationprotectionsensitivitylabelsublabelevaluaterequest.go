package informationprotectionsensitivitylabelsublabel

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateInformationProtectionSensitivityLabelSublabelEvaluateRequest struct {
	CurrentLabel             *beta.CurrentLabel              `json:"currentLabel,omitempty"`
	DiscoveredSensitiveTypes *[]beta.DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
}
