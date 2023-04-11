package adaptivenetworkhardenings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EffectiveNetworkSecurityGroups struct {
	NetworkInterface      *string   `json:"networkInterface,omitempty"`
	NetworkSecurityGroups *[]string `json:"networkSecurityGroups,omitempty"`
}
