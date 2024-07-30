package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotification struct {
	CreatedBy            *EmailIdentity                       `json:"createdBy,omitempty"`
	CreatedDateTime      *string                              `json:"createdDateTime,omitempty"`
	Description          *string                              `json:"description,omitempty"`
	Details              *[]EndUserNotificationDetail         `json:"details,omitempty"`
	DisplayName          *string                              `json:"displayName,omitempty"`
	Id                   *string                              `json:"id,omitempty"`
	LastModifiedBy       *EmailIdentity                       `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                              `json:"lastModifiedDateTime,omitempty"`
	NotificationType     *EndUserNotificationNotificationType `json:"notificationType,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	Source               *EndUserNotificationSource           `json:"source,omitempty"`
	Status               *EndUserNotificationStatus           `json:"status,omitempty"`
	SupportedLocales     *[]string                            `json:"supportedLocales,omitempty"`
}
