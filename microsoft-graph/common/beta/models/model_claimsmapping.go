package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsMapping struct {
	DisplayName *string `json:"displayName,omitempty"`
	Email       *string `json:"email,omitempty"`
	GivenName   *string `json:"givenName,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Surname     *string `json:"surname,omitempty"`
	UserId      *string `json:"userId,omitempty"`
}
