package adaptivenetworkhardenings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveNetworkHardeningEnforceRequest struct {
	NetworkSecurityGroups []string `json:"networkSecurityGroups"`
	Rules                 []Rule   `json:"rules"`
}
