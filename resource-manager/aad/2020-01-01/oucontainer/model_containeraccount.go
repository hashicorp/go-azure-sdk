package oucontainer

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAccount struct {
	AccountName *string `json:"accountName,omitempty"`
	Password    *string `json:"password,omitempty"`
	Spn         *string `json:"spn,omitempty"`
}
