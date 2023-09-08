package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResultPartnerAgentName string

const (
	CloudPCPartnerAgentInstallResultPartnerAgentNamecitrix CloudPCPartnerAgentInstallResultPartnerAgentName = "Citrix"
	CloudPCPartnerAgentInstallResultPartnerAgentNamevMware CloudPCPartnerAgentInstallResultPartnerAgentName = "VMware"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultPartnerAgentName() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultPartnerAgentNamecitrix),
		string(CloudPCPartnerAgentInstallResultPartnerAgentNamevMware),
	}
}

func parseCloudPCPartnerAgentInstallResultPartnerAgentName(input string) (*CloudPCPartnerAgentInstallResultPartnerAgentName, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultPartnerAgentName{
		"citrix": CloudPCPartnerAgentInstallResultPartnerAgentNamecitrix,
		"vmware": CloudPCPartnerAgentInstallResultPartnerAgentNamevMware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultPartnerAgentName(input)
	return &out, nil
}
