package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OsProfileForVMInstance struct {
	AdminPassword   *string `json:"adminPassword,omitempty"`
	ComputerName    *string `json:"computerName,omitempty"`
	DomainName      *string `json:"domainName,omitempty"`
	DomainPassword  *string `json:"domainPassword,omitempty"`
	DomainUsername  *string `json:"domainUsername,omitempty"`
	OsSku           *string `json:"osSku,omitempty"`
	OsType          *OsType `json:"osType,omitempty"`
	OsVersion       *string `json:"osVersion,omitempty"`
	ProductKey      *string `json:"productKey,omitempty"`
	RunOnceCommands *string `json:"runOnceCommands,omitempty"`
	Timezone        *int64  `json:"timezone,omitempty"`
	Workgroup       *string `json:"workgroup,omitempty"`
}
