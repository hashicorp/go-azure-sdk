package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationMessageTemplate struct {
	BrandingOptions               *NotificationMessageTemplateBrandingOptions `json:"brandingOptions,omitempty"`
	DefaultLocale                 *string                                     `json:"defaultLocale,omitempty"`
	DisplayName                   *string                                     `json:"displayName,omitempty"`
	Id                            *string                                     `json:"id,omitempty"`
	LastModifiedDateTime          *string                                     `json:"lastModifiedDateTime,omitempty"`
	LocalizedNotificationMessages *[]LocalizedNotificationMessage             `json:"localizedNotificationMessages,omitempty"`
	ODataType                     *string                                     `json:"@odata.type,omitempty"`
	RoleScopeTagIds               *[]string                                   `json:"roleScopeTagIds,omitempty"`
}
