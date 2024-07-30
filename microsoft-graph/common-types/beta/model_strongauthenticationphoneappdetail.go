package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StrongAuthenticationPhoneAppDetail struct {
	AuthenticationType               *string            `json:"authenticationType,omitempty"`
	AuthenticatorFlavor              *string            `json:"authenticatorFlavor,omitempty"`
	DeviceId                         *string            `json:"deviceId,omitempty"`
	DeviceName                       *string            `json:"deviceName,omitempty"`
	DeviceTag                        *string            `json:"deviceTag,omitempty"`
	DeviceToken                      *string            `json:"deviceToken,omitempty"`
	HashFunction                     *string            `json:"hashFunction,omitempty"`
	Id                               *string            `json:"id,omitempty"`
	LastAuthenticatedDateTime        *string            `json:"lastAuthenticatedDateTime,omitempty"`
	NotificationType                 *string            `json:"notificationType,omitempty"`
	ODataType                        *string            `json:"@odata.type,omitempty"`
	OathSecretKey                    *string            `json:"oathSecretKey,omitempty"`
	OathTokenMetadata                *OathTokenMetadata `json:"oathTokenMetadata,omitempty"`
	OathTokenTimeDriftInSeconds      *int64             `json:"oathTokenTimeDriftInSeconds,omitempty"`
	PhoneAppVersion                  *string            `json:"phoneAppVersion,omitempty"`
	TenantDeviceId                   *string            `json:"tenantDeviceId,omitempty"`
	TokenGenerationIntervalInSeconds *int64             `json:"tokenGenerationIntervalInSeconds,omitempty"`
}
