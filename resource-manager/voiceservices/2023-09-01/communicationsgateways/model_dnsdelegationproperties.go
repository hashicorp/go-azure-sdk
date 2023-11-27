package communicationsgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsDelegationProperties struct {
	Domain      *string   `json:"domain,omitempty"`
	NameServers *[]string `json:"nameServers,omitempty"`
}
