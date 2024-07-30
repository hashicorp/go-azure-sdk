package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResultInstallStatus string

const (
	CloudPCPartnerAgentInstallResultInstallStatus_InstallFailed   CloudPCPartnerAgentInstallResultInstallStatus = "installFailed"
	CloudPCPartnerAgentInstallResultInstallStatus_Installed       CloudPCPartnerAgentInstallResultInstallStatus = "installed"
	CloudPCPartnerAgentInstallResultInstallStatus_Installing      CloudPCPartnerAgentInstallResultInstallStatus = "installing"
	CloudPCPartnerAgentInstallResultInstallStatus_Licensed        CloudPCPartnerAgentInstallResultInstallStatus = "licensed"
	CloudPCPartnerAgentInstallResultInstallStatus_UninstallFailed CloudPCPartnerAgentInstallResultInstallStatus = "uninstallFailed"
	CloudPCPartnerAgentInstallResultInstallStatus_Uninstalling    CloudPCPartnerAgentInstallResultInstallStatus = "uninstalling"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultInstallStatus() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultInstallStatus_InstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatus_Installed),
		string(CloudPCPartnerAgentInstallResultInstallStatus_Installing),
		string(CloudPCPartnerAgentInstallResultInstallStatus_Licensed),
		string(CloudPCPartnerAgentInstallResultInstallStatus_UninstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatus_Uninstalling),
	}
}

func (s *CloudPCPartnerAgentInstallResultInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentInstallResultInstallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentInstallResultInstallStatus(input string) (*CloudPCPartnerAgentInstallResultInstallStatus, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultInstallStatus{
		"installfailed":   CloudPCPartnerAgentInstallResultInstallStatus_InstallFailed,
		"installed":       CloudPCPartnerAgentInstallResultInstallStatus_Installed,
		"installing":      CloudPCPartnerAgentInstallResultInstallStatus_Installing,
		"licensed":        CloudPCPartnerAgentInstallResultInstallStatus_Licensed,
		"uninstallfailed": CloudPCPartnerAgentInstallResultInstallStatus_UninstallFailed,
		"uninstalling":    CloudPCPartnerAgentInstallResultInstallStatus_Uninstalling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultInstallStatus(input)
	return &out, nil
}
