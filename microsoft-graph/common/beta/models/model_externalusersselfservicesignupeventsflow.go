package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalUsersSelfServiceSignUpEventsFlow struct {
	Conditions                      *AuthenticationConditions               `json:"conditions,omitempty"`
	Description                     *string                                 `json:"description,omitempty"`
	DisplayName                     *string                                 `json:"displayName,omitempty"`
	Id                              *string                                 `json:"id,omitempty"`
	ODataType                       *string                                 `json:"@odata.type,omitempty"`
	OnAttributeCollection           *OnAttributeCollectionHandler           `json:"onAttributeCollection,omitempty"`
	OnAuthenticationMethodLoadStart *OnAuthenticationMethodLoadStartHandler `json:"onAuthenticationMethodLoadStart,omitempty"`
	OnInteractiveAuthFlowStart      *OnInteractiveAuthFlowStartHandler      `json:"onInteractiveAuthFlowStart,omitempty"`
	OnUserCreateStart               *OnUserCreateStartHandler               `json:"onUserCreateStart,omitempty"`
	Priority                        *int64                                  `json:"priority,omitempty"`
}
