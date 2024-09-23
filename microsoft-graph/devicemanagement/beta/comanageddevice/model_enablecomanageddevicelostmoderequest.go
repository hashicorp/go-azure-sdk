package comanageddevice

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableComanagedDeviceLostModeRequest struct {
	Footer      nullable.Type[string] `json:"footer,omitempty"`
	Message     nullable.Type[string] `json:"message,omitempty"`
	PhoneNumber nullable.Type[string] `json:"phoneNumber,omitempty"`
}
