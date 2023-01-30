package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProperties struct {
	EncryptedPassword *AsymmetricEncryptedSecret `json:"encryptedPassword,omitempty"`
	ShareAccessRights *[]ShareAccessRight        `json:"shareAccessRights,omitempty"`
	UserType          UserType                   `json:"userType"`
}
