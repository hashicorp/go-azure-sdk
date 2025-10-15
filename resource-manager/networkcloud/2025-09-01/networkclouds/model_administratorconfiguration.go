package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministratorConfiguration struct {
	AdminUsername *string         `json:"adminUsername,omitempty"`
	SshPublicKeys *[]SshPublicKey `json:"sshPublicKeys,omitempty"`
}
