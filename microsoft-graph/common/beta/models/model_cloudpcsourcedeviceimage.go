package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSourceDeviceImage struct {
	DisplayName             *string `json:"displayName,omitempty"`
	Id                      *string `json:"id,omitempty"`
	ODataType               *string `json:"@odata.type,omitempty"`
	SubscriptionDisplayName *string `json:"subscriptionDisplayName,omitempty"`
	SubscriptionId          *string `json:"subscriptionId,omitempty"`
}
