package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeNotification struct {
	ChangeType                     *ChangeNotificationChangeType       `json:"changeType,omitempty"`
	ClientState                    *string                             `json:"clientState,omitempty"`
	EncryptedContent               *ChangeNotificationEncryptedContent `json:"encryptedContent,omitempty"`
	Id                             *string                             `json:"id,omitempty"`
	LifecycleEvent                 *ChangeNotificationLifecycleEvent   `json:"lifecycleEvent,omitempty"`
	ODataType                      *string                             `json:"@odata.type,omitempty"`
	Resource                       *string                             `json:"resource,omitempty"`
	ResourceData                   *ResourceData                       `json:"resourceData,omitempty"`
	SubscriptionExpirationDateTime *string                             `json:"subscriptionExpirationDateTime,omitempty"`
	SubscriptionId                 *string                             `json:"subscriptionId,omitempty"`
	TenantId                       *string                             `json:"tenantId,omitempty"`
}
