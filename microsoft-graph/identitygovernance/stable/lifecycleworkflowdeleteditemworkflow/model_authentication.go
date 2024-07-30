package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Authentication struct {
	EmailMethods                   *[]EmailAuthenticationMethod                   `json:"emailMethods,omitempty"`
	Fido2Methods                   *[]Fido2AuthenticationMethod                   `json:"fido2Methods,omitempty"`
	Id                             *string                                        `json:"id,omitempty"`
	Methods                        *[]AuthenticationMethod                        `json:"methods,omitempty"`
	MicrosoftAuthenticatorMethods  *[]MicrosoftAuthenticatorAuthenticationMethod  `json:"microsoftAuthenticatorMethods,omitempty"`
	ODataType                      *string                                        `json:"@odata.type,omitempty"`
	Operations                     *[]LongRunningOperation                        `json:"operations,omitempty"`
	PasswordMethods                *[]PasswordAuthenticationMethod                `json:"passwordMethods,omitempty"`
	PhoneMethods                   *[]PhoneAuthenticationMethod                   `json:"phoneMethods,omitempty"`
	SoftwareOathMethods            *[]SoftwareOathAuthenticationMethod            `json:"softwareOathMethods,omitempty"`
	TemporaryAccessPassMethods     *[]TemporaryAccessPassAuthenticationMethod     `json:"temporaryAccessPassMethods,omitempty"`
	WindowsHelloForBusinessMethods *[]WindowsHelloForBusinessAuthenticationMethod `json:"windowsHelloForBusinessMethods,omitempty"`
}
