package authenticationqrcodepinmethodpin

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAuthenticationQrCodePinMethodPinPinRequest struct {
	CurrentPin nullable.Type[string] `json:"currentPin,omitempty"`
	NewPin     nullable.Type[string] `json:"newPin,omitempty"`
}
