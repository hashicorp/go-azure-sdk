package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateAuthenticationMethodConfiguration struct {
	AuthenticationModeConfiguration *X509CertificateAuthenticationModeConfiguration        `json:"authenticationModeConfiguration,omitempty"`
	CertificateUserBindings         *[]X509CertificateUserBinding                          `json:"certificateUserBindings,omitempty"`
	ExcludeTargets                  *[]ExcludeTarget                                       `json:"excludeTargets,omitempty"`
	Id                              *string                                                `json:"id,omitempty"`
	IncludeTargets                  *[]AuthenticationMethodTarget                          `json:"includeTargets,omitempty"`
	ODataType                       *string                                                `json:"@odata.type,omitempty"`
	State                           *X509CertificateAuthenticationMethodConfigurationState `json:"state,omitempty"`
}
