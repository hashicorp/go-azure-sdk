package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneAuthenticationMethod struct {
	Id             *string                                  `json:"id,omitempty"`
	ODataType      *string                                  `json:"@odata.type,omitempty"`
	PhoneNumber    *string                                  `json:"phoneNumber,omitempty"`
	PhoneType      *PhoneAuthenticationMethodPhoneType      `json:"phoneType,omitempty"`
	SmsSignInState *PhoneAuthenticationMethodSmsSignInState `json:"smsSignInState,omitempty"`
}
