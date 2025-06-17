package authenticationhardwareoathmethod

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignAuthenticationHardwareOathMethodsAndActivateRequest struct {
	Device           *beta.HardwareOathTokenAuthenticationMethodDevice `json:"device,omitempty"`
	DisplayName      nullable.Type[string]                             `json:"displayName,omitempty"`
	VerificationCode nullable.Type[string]                             `json:"verificationCode,omitempty"`
}
