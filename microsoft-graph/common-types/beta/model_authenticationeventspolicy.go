package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsPolicy struct {
	Id            *string                   `json:"id,omitempty"`
	ODataType     *string                   `json:"@odata.type,omitempty"`
	OnSignupStart *[]AuthenticationListener `json:"onSignupStart,omitempty"`
}
