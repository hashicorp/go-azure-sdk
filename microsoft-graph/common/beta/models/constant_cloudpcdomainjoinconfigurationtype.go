package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfigurationType string

const (
	CloudPCDomainJoinConfigurationTypeazureADJoin       CloudPCDomainJoinConfigurationType = "AzureADJoin"
	CloudPCDomainJoinConfigurationTypehybridAzureADJoin CloudPCDomainJoinConfigurationType = "HybridAzureADJoin"
)

func PossibleValuesForCloudPCDomainJoinConfigurationType() []string {
	return []string{
		string(CloudPCDomainJoinConfigurationTypeazureADJoin),
		string(CloudPCDomainJoinConfigurationTypehybridAzureADJoin),
	}
}

func parseCloudPCDomainJoinConfigurationType(input string) (*CloudPCDomainJoinConfigurationType, error) {
	vals := map[string]CloudPCDomainJoinConfigurationType{
		"azureadjoin":       CloudPCDomainJoinConfigurationTypeazureADJoin,
		"hybridazureadjoin": CloudPCDomainJoinConfigurationTypehybridAzureADJoin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDomainJoinConfigurationType(input)
	return &out, nil
}
