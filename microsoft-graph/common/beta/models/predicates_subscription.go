package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionOperationPredicate struct {
	ApplicationId             *string
	ChangeType                *string
	ClientState               *string
	CreatorId                 *string
	EncryptionCertificate     *string
	EncryptionCertificateId   *string
	ExpirationDateTime        *string
	Id                        *string
	IncludeResourceData       *bool
	LatestSupportedTlsVersion *string
	LifecycleNotificationUrl  *string
	NotificationContentType   *string
	NotificationQueryOptions  *string
	NotificationUrl           *string
	NotificationUrlAppId      *string
	ODataType                 *string
	Resource                  *string
}

func (p SubscriptionOperationPredicate) Matches(input Subscription) bool {

	if p.ApplicationId != nil && (input.ApplicationId == nil || *p.ApplicationId != *input.ApplicationId) {
		return false
	}

	if p.ChangeType != nil && (input.ChangeType == nil || *p.ChangeType != *input.ChangeType) {
		return false
	}

	if p.ClientState != nil && (input.ClientState == nil || *p.ClientState != *input.ClientState) {
		return false
	}

	if p.CreatorId != nil && (input.CreatorId == nil || *p.CreatorId != *input.CreatorId) {
		return false
	}

	if p.EncryptionCertificate != nil && (input.EncryptionCertificate == nil || *p.EncryptionCertificate != *input.EncryptionCertificate) {
		return false
	}

	if p.EncryptionCertificateId != nil && (input.EncryptionCertificateId == nil || *p.EncryptionCertificateId != *input.EncryptionCertificateId) {
		return false
	}

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IncludeResourceData != nil && (input.IncludeResourceData == nil || *p.IncludeResourceData != *input.IncludeResourceData) {
		return false
	}

	if p.LatestSupportedTlsVersion != nil && (input.LatestSupportedTlsVersion == nil || *p.LatestSupportedTlsVersion != *input.LatestSupportedTlsVersion) {
		return false
	}

	if p.LifecycleNotificationUrl != nil && (input.LifecycleNotificationUrl == nil || *p.LifecycleNotificationUrl != *input.LifecycleNotificationUrl) {
		return false
	}

	if p.NotificationContentType != nil && (input.NotificationContentType == nil || *p.NotificationContentType != *input.NotificationContentType) {
		return false
	}

	if p.NotificationQueryOptions != nil && (input.NotificationQueryOptions == nil || *p.NotificationQueryOptions != *input.NotificationQueryOptions) {
		return false
	}

	if p.NotificationUrl != nil && (input.NotificationUrl == nil || *p.NotificationUrl != *input.NotificationUrl) {
		return false
	}

	if p.NotificationUrlAppId != nil && (input.NotificationUrlAppId == nil || *p.NotificationUrlAppId != *input.NotificationUrlAppId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Resource != nil && (input.Resource == nil || *p.Resource != *input.Resource) {
		return false
	}

	return true
}
