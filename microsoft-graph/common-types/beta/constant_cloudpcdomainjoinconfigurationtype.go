package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfigurationType string

const (
	CloudPCDomainJoinConfigurationType_AzureADJoin       CloudPCDomainJoinConfigurationType = "azureADJoin"
	CloudPCDomainJoinConfigurationType_HybridAzureADJoin CloudPCDomainJoinConfigurationType = "hybridAzureADJoin"
)

func PossibleValuesForCloudPCDomainJoinConfigurationType() []string {
	return []string{
		string(CloudPCDomainJoinConfigurationType_AzureADJoin),
		string(CloudPCDomainJoinConfigurationType_HybridAzureADJoin),
	}
}

func (s *CloudPCDomainJoinConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDomainJoinConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDomainJoinConfigurationType(input string) (*CloudPCDomainJoinConfigurationType, error) {
	vals := map[string]CloudPCDomainJoinConfigurationType{
		"azureadjoin":       CloudPCDomainJoinConfigurationType_AzureADJoin,
		"hybridazureadjoin": CloudPCDomainJoinConfigurationType_HybridAzureADJoin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDomainJoinConfigurationType(input)
	return &out, nil
}
