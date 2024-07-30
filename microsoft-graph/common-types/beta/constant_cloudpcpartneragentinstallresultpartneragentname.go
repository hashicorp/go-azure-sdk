package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResultPartnerAgentName string

const (
	CloudPCPartnerAgentInstallResultPartnerAgentName_Citrix CloudPCPartnerAgentInstallResultPartnerAgentName = "citrix"
	CloudPCPartnerAgentInstallResultPartnerAgentName_Hp     CloudPCPartnerAgentInstallResultPartnerAgentName = "hp"
	CloudPCPartnerAgentInstallResultPartnerAgentName_VMware CloudPCPartnerAgentInstallResultPartnerAgentName = "vMware"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultPartnerAgentName() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultPartnerAgentName_Citrix),
		string(CloudPCPartnerAgentInstallResultPartnerAgentName_Hp),
		string(CloudPCPartnerAgentInstallResultPartnerAgentName_VMware),
	}
}

func (s *CloudPCPartnerAgentInstallResultPartnerAgentName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentInstallResultPartnerAgentName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentInstallResultPartnerAgentName(input string) (*CloudPCPartnerAgentInstallResultPartnerAgentName, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultPartnerAgentName{
		"citrix": CloudPCPartnerAgentInstallResultPartnerAgentName_Citrix,
		"hp":     CloudPCPartnerAgentInstallResultPartnerAgentName_Hp,
		"vmware": CloudPCPartnerAgentInstallResultPartnerAgentName_VMware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultPartnerAgentName(input)
	return &out, nil
}
