package templatedevicetemplate

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTemplateDeviceFromTemplateRequest struct {
	AccountEnabled         nullable.Type[bool]   `json:"accountEnabled,omitempty"`
	AlternativeNames       *[]string             `json:"alternativeNames,omitempty"`
	ExternalDeviceId       *string               `json:"externalDeviceId,omitempty"`
	ExternalSourceName     nullable.Type[string] `json:"externalSourceName,omitempty"`
	KeyCredential          *beta.KeyCredential   `json:"keyCredential,omitempty"`
	OperatingSystemVersion nullable.Type[string] `json:"operatingSystemVersion,omitempty"`
}
