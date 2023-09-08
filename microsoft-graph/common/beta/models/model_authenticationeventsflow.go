package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlow struct {
	Conditions  *AuthenticationConditions `json:"conditions,omitempty"`
	Description *string                   `json:"description,omitempty"`
	DisplayName *string                   `json:"displayName,omitempty"`
	Id          *string                   `json:"id,omitempty"`
	ODataType   *string                   `json:"@odata.type,omitempty"`
	Priority    *int64                    `json:"priority,omitempty"`
}
