package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageRepositoryCredentials struct {
	Password    string `json:"password"`
	RegistryUrl string `json:"registryUrl"`
	Username    string `json:"username"`
}
