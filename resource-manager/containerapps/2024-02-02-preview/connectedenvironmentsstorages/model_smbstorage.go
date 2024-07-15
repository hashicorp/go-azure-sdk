package connectedenvironmentsstorages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmbStorage struct {
	AccessMode *AccessMode `json:"accessMode,omitempty"`
	Domain     *string     `json:"domain,omitempty"`
	Host       *string     `json:"host,omitempty"`
	Password   *string     `json:"password,omitempty"`
	ShareName  *string     `json:"shareName,omitempty"`
	Username   *string     `json:"username,omitempty"`
}
