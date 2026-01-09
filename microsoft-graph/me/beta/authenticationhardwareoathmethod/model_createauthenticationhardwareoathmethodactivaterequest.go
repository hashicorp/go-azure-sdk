package authenticationhardwareoathmethod

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationHardwareOathMethodActivateRequest struct {
	DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
	VerificationCode nullable.Type[string] `json:"verificationCode,omitempty"`
}
