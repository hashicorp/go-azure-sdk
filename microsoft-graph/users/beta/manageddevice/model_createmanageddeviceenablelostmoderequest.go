package manageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceEnableLostModeRequest struct {
	Footer      *string `json:"footer,omitempty"`
	Message     *string `json:"message,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}
