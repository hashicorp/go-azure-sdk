package memanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeManagedDeviceByIdDeprovisionRequest struct {
	DeprovisionReason *string `json:"deprovisionReason,omitempty"`
}
