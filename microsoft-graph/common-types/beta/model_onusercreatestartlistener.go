package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnUserCreateStartListener struct {
	AuthenticationEventsFlowId *string                   `json:"authenticationEventsFlowId,omitempty"`
	Conditions                 *AuthenticationConditions `json:"conditions,omitempty"`
	Handler                    *OnUserCreateStartHandler `json:"handler,omitempty"`
	Id                         *string                   `json:"id,omitempty"`
	ODataType                  *string                   `json:"@odata.type,omitempty"`
	Priority                   *int64                    `json:"priority,omitempty"`
}
