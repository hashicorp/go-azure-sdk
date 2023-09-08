package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionType string

const (
	CloudPCOnPremisesConnectionTypeazureADJoin       CloudPCOnPremisesConnectionType = "AzureADJoin"
	CloudPCOnPremisesConnectionTypehybridAzureADJoin CloudPCOnPremisesConnectionType = "HybridAzureADJoin"
)

func PossibleValuesForCloudPCOnPremisesConnectionType() []string {
	return []string{
		string(CloudPCOnPremisesConnectionTypeazureADJoin),
		string(CloudPCOnPremisesConnectionTypehybridAzureADJoin),
	}
}

func parseCloudPCOnPremisesConnectionType(input string) (*CloudPCOnPremisesConnectionType, error) {
	vals := map[string]CloudPCOnPremisesConnectionType{
		"azureadjoin":       CloudPCOnPremisesConnectionTypeazureADJoin,
		"hybridazureadjoin": CloudPCOnPremisesConnectionTypehybridAzureADJoin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionType(input)
	return &out, nil
}
