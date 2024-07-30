package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesAzureADDeviceRegistrationError struct {
	ODataType *string                                             `json:"@odata.type,omitempty"`
	Reason    *WindowsUpdatesAzureADDeviceRegistrationErrorReason `json:"reason,omitempty"`
}
