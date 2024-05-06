package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsConfiguration struct {
	AutoLogon          *bool     `json:"autoLogon,omitempty"`
	AutoLogonCount     *int64    `json:"autoLogonCount,omitempty"`
	DomainName         *string   `json:"domainName,omitempty"`
	DomainUserPassword *string   `json:"domainUserPassword,omitempty"`
	DomainUsername     *string   `json:"domainUsername,omitempty"`
	FirstLogonCommands *[]string `json:"firstLogonCommands,omitempty"`
	FullName           *string   `json:"fullName,omitempty"`
	OrgName            *string   `json:"orgName,omitempty"`
	ProductId          *string   `json:"productId,omitempty"`
	TimeZone           *string   `json:"timeZone,omitempty"`
	WorkGroupName      *string   `json:"workGroupName,omitempty"`
}
