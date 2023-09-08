package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosLobAppProvisioningConfigurationPolicySetItemStatus string

const (
	IosLobAppProvisioningConfigurationPolicySetItemStatuserror          IosLobAppProvisioningConfigurationPolicySetItemStatus = "Error"
	IosLobAppProvisioningConfigurationPolicySetItemStatusnotAssigned    IosLobAppProvisioningConfigurationPolicySetItemStatus = "NotAssigned"
	IosLobAppProvisioningConfigurationPolicySetItemStatuspartialSuccess IosLobAppProvisioningConfigurationPolicySetItemStatus = "PartialSuccess"
	IosLobAppProvisioningConfigurationPolicySetItemStatussuccess        IosLobAppProvisioningConfigurationPolicySetItemStatus = "Success"
	IosLobAppProvisioningConfigurationPolicySetItemStatusunknown        IosLobAppProvisioningConfigurationPolicySetItemStatus = "Unknown"
	IosLobAppProvisioningConfigurationPolicySetItemStatusvalidating     IosLobAppProvisioningConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForIosLobAppProvisioningConfigurationPolicySetItemStatus() []string {
	return []string{
		string(IosLobAppProvisioningConfigurationPolicySetItemStatuserror),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatusnotAssigned),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatuspartialSuccess),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatussuccess),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatusunknown),
		string(IosLobAppProvisioningConfigurationPolicySetItemStatusvalidating),
	}
}

func parseIosLobAppProvisioningConfigurationPolicySetItemStatus(input string) (*IosLobAppProvisioningConfigurationPolicySetItemStatus, error) {
	vals := map[string]IosLobAppProvisioningConfigurationPolicySetItemStatus{
		"error":          IosLobAppProvisioningConfigurationPolicySetItemStatuserror,
		"notassigned":    IosLobAppProvisioningConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": IosLobAppProvisioningConfigurationPolicySetItemStatuspartialSuccess,
		"success":        IosLobAppProvisioningConfigurationPolicySetItemStatussuccess,
		"unknown":        IosLobAppProvisioningConfigurationPolicySetItemStatusunknown,
		"validating":     IosLobAppProvisioningConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosLobAppProvisioningConfigurationPolicySetItemStatus(input)
	return &out, nil
}
