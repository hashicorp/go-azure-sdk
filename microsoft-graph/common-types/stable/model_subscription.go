package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Subscription struct {
	ApplicationId             *string `json:"applicationId,omitempty"`
	ChangeType                *string `json:"changeType,omitempty"`
	ClientState               *string `json:"clientState,omitempty"`
	CreatorId                 *string `json:"creatorId,omitempty"`
	EncryptionCertificate     *string `json:"encryptionCertificate,omitempty"`
	EncryptionCertificateId   *string `json:"encryptionCertificateId,omitempty"`
	ExpirationDateTime        *string `json:"expirationDateTime,omitempty"`
	Id                        *string `json:"id,omitempty"`
	IncludeResourceData       *bool   `json:"includeResourceData,omitempty"`
	LatestSupportedTlsVersion *string `json:"latestSupportedTlsVersion,omitempty"`
	LifecycleNotificationUrl  *string `json:"lifecycleNotificationUrl,omitempty"`
	NotificationQueryOptions  *string `json:"notificationQueryOptions,omitempty"`
	NotificationUrl           *string `json:"notificationUrl,omitempty"`
	NotificationUrlAppId      *string `json:"notificationUrlAppId,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	Resource                  *string `json:"resource,omitempty"`
}
