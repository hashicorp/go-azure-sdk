package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordProfile struct {
	ForceChangePasswordNextSignIn        *bool   `json:"forceChangePasswordNextSignIn,omitempty"`
	ForceChangePasswordNextSignInWithMfa *bool   `json:"forceChangePasswordNextSignInWithMfa,omitempty"`
	ODataType                            *string `json:"@odata.type,omitempty"`
	Password                             *string `json:"password,omitempty"`
}
