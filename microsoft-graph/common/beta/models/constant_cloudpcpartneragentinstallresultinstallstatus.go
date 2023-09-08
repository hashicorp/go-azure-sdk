package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallResultInstallStatus string

const (
	CloudPCPartnerAgentInstallResultInstallStatusinstallFailed   CloudPCPartnerAgentInstallResultInstallStatus = "InstallFailed"
	CloudPCPartnerAgentInstallResultInstallStatusinstalled       CloudPCPartnerAgentInstallResultInstallStatus = "Installed"
	CloudPCPartnerAgentInstallResultInstallStatusinstalling      CloudPCPartnerAgentInstallResultInstallStatus = "Installing"
	CloudPCPartnerAgentInstallResultInstallStatuslicensed        CloudPCPartnerAgentInstallResultInstallStatus = "Licensed"
	CloudPCPartnerAgentInstallResultInstallStatusuninstallFailed CloudPCPartnerAgentInstallResultInstallStatus = "UninstallFailed"
	CloudPCPartnerAgentInstallResultInstallStatusuninstalling    CloudPCPartnerAgentInstallResultInstallStatus = "Uninstalling"
)

func PossibleValuesForCloudPCPartnerAgentInstallResultInstallStatus() []string {
	return []string{
		string(CloudPCPartnerAgentInstallResultInstallStatusinstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatusinstalled),
		string(CloudPCPartnerAgentInstallResultInstallStatusinstalling),
		string(CloudPCPartnerAgentInstallResultInstallStatuslicensed),
		string(CloudPCPartnerAgentInstallResultInstallStatusuninstallFailed),
		string(CloudPCPartnerAgentInstallResultInstallStatusuninstalling),
	}
}

func parseCloudPCPartnerAgentInstallResultInstallStatus(input string) (*CloudPCPartnerAgentInstallResultInstallStatus, error) {
	vals := map[string]CloudPCPartnerAgentInstallResultInstallStatus{
		"installfailed":   CloudPCPartnerAgentInstallResultInstallStatusinstallFailed,
		"installed":       CloudPCPartnerAgentInstallResultInstallStatusinstalled,
		"installing":      CloudPCPartnerAgentInstallResultInstallStatusinstalling,
		"licensed":        CloudPCPartnerAgentInstallResultInstallStatuslicensed,
		"uninstallfailed": CloudPCPartnerAgentInstallResultInstallStatusuninstallFailed,
		"uninstalling":    CloudPCPartnerAgentInstallResultInstallStatusuninstalling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallResultInstallStatus(input)
	return &out, nil
}
