package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnSettingsKerberosSignOnMappingAttributeType string

const (
	KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesSAMAccountName        KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "onPremisesSAMAccountName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName     KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "onPremisesUserPrincipalName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "onPremisesUserPrincipalUsername"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalName               KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "userPrincipalName"
	KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalUsername           KerberosSignOnSettingsKerberosSignOnMappingAttributeType = "userPrincipalUsername"
)

func PossibleValuesForKerberosSignOnSettingsKerberosSignOnMappingAttributeType() []string {
	return []string{
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesSAMAccountName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalName),
		string(KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalUsername),
	}
}

func (s *KerberosSignOnSettingsKerberosSignOnMappingAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKerberosSignOnSettingsKerberosSignOnMappingAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKerberosSignOnSettingsKerberosSignOnMappingAttributeType(input string) (*KerberosSignOnSettingsKerberosSignOnMappingAttributeType, error) {
	vals := map[string]KerberosSignOnSettingsKerberosSignOnMappingAttributeType{
		"onpremisessamaccountname":        KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesSAMAccountName,
		"onpremisesuserprincipalname":     KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName,
		"onpremisesuserprincipalusername": KerberosSignOnSettingsKerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername,
		"userprincipalname":               KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalName,
		"userprincipalusername":           KerberosSignOnSettingsKerberosSignOnMappingAttributeType_UserPrincipalUsername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KerberosSignOnSettingsKerberosSignOnMappingAttributeType(input)
	return &out, nil
}
