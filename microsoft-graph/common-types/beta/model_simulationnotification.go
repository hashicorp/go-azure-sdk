package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationNotification struct {
	DefaultLanguage     *string                                  `json:"defaultLanguage,omitempty"`
	EndUserNotification *EndUserNotification                     `json:"endUserNotification,omitempty"`
	ODataType           *string                                  `json:"@odata.type,omitempty"`
	TargettedUserType   *SimulationNotificationTargettedUserType `json:"targettedUserType,omitempty"`
}
