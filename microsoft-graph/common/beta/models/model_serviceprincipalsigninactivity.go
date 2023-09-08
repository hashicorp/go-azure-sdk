package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalSignInActivity struct {
	AppId                                           *string         `json:"appId,omitempty"`
	ApplicationAuthenticationClientSignInActivity   *SignInActivity `json:"applicationAuthenticationClientSignInActivity,omitempty"`
	ApplicationAuthenticationResourceSignInActivity *SignInActivity `json:"applicationAuthenticationResourceSignInActivity,omitempty"`
	DelegatedClientSignInActivity                   *SignInActivity `json:"delegatedClientSignInActivity,omitempty"`
	DelegatedResourceSignInActivity                 *SignInActivity `json:"delegatedResourceSignInActivity,omitempty"`
	Id                                              *string         `json:"id,omitempty"`
	LastSignInActivity                              *SignInActivity `json:"lastSignInActivity,omitempty"`
	ODataType                                       *string         `json:"@odata.type,omitempty"`
}
