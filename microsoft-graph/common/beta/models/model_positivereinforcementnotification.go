package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositiveReinforcementNotification struct {
	DefaultLanguage     *string                                              `json:"defaultLanguage,omitempty"`
	DeliveryPreference  *PositiveReinforcementNotificationDeliveryPreference `json:"deliveryPreference,omitempty"`
	EndUserNotification *EndUserNotification                                 `json:"endUserNotification,omitempty"`
	ODataType           *string                                              `json:"@odata.type,omitempty"`
}
