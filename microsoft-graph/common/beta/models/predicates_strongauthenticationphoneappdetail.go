package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StrongAuthenticationPhoneAppDetailOperationPredicate struct {
	AuthenticationType               *string
	AuthenticatorFlavor              *string
	DeviceId                         *string
	DeviceName                       *string
	DeviceTag                        *string
	DeviceToken                      *string
	HashFunction                     *string
	Id                               *string
	LastAuthenticatedDateTime        *string
	NotificationType                 *string
	ODataType                        *string
	OathSecretKey                    *string
	OathTokenTimeDriftInSeconds      *int64
	PhoneAppVersion                  *string
	TenantDeviceId                   *string
	TokenGenerationIntervalInSeconds *int64
}

func (p StrongAuthenticationPhoneAppDetailOperationPredicate) Matches(input StrongAuthenticationPhoneAppDetail) bool {

	if p.AuthenticationType != nil && (input.AuthenticationType == nil || *p.AuthenticationType != *input.AuthenticationType) {
		return false
	}

	if p.AuthenticatorFlavor != nil && (input.AuthenticatorFlavor == nil || *p.AuthenticatorFlavor != *input.AuthenticatorFlavor) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.DeviceTag != nil && (input.DeviceTag == nil || *p.DeviceTag != *input.DeviceTag) {
		return false
	}

	if p.DeviceToken != nil && (input.DeviceToken == nil || *p.DeviceToken != *input.DeviceToken) {
		return false
	}

	if p.HashFunction != nil && (input.HashFunction == nil || *p.HashFunction != *input.HashFunction) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastAuthenticatedDateTime != nil && (input.LastAuthenticatedDateTime == nil || *p.LastAuthenticatedDateTime != *input.LastAuthenticatedDateTime) {
		return false
	}

	if p.NotificationType != nil && (input.NotificationType == nil || *p.NotificationType != *input.NotificationType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OathSecretKey != nil && (input.OathSecretKey == nil || *p.OathSecretKey != *input.OathSecretKey) {
		return false
	}

	if p.OathTokenTimeDriftInSeconds != nil && (input.OathTokenTimeDriftInSeconds == nil || *p.OathTokenTimeDriftInSeconds != *input.OathTokenTimeDriftInSeconds) {
		return false
	}

	if p.PhoneAppVersion != nil && (input.PhoneAppVersion == nil || *p.PhoneAppVersion != *input.PhoneAppVersion) {
		return false
	}

	if p.TenantDeviceId != nil && (input.TenantDeviceId == nil || *p.TenantDeviceId != *input.TenantDeviceId) {
		return false
	}

	if p.TokenGenerationIntervalInSeconds != nil && (input.TokenGenerationIntervalInSeconds == nil || *p.TokenGenerationIntervalInSeconds != *input.TokenGenerationIntervalInSeconds) {
		return false
	}

	return true
}
