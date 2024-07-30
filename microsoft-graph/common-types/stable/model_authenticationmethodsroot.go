package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRoot struct {
	Id                      *string                    `json:"id,omitempty"`
	ODataType               *string                    `json:"@odata.type,omitempty"`
	UserRegistrationDetails *[]UserRegistrationDetails `json:"userRegistrationDetails,omitempty"`
}
