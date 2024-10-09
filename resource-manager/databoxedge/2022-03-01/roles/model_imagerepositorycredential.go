package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageRepositoryCredential struct {
	ImageRepositoryURL string                     `json:"imageRepositoryUrl"`
	Password           *AsymmetricEncryptedSecret `json:"password,omitempty"`
	UserName           string                     `json:"userName"`
}
