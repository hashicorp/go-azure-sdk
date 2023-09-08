package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementAgent string

const (
	ManagedDeviceManagementAgentconfigurationManagerClient        ManagedDeviceManagementAgent = "ConfigurationManagerClient"
	ManagedDeviceManagementAgentconfigurationManagerClientMdm     ManagedDeviceManagementAgent = "ConfigurationManagerClientMdm"
	ManagedDeviceManagementAgentconfigurationManagerClientMdmEas  ManagedDeviceManagementAgent = "ConfigurationManagerClientMdmEas"
	ManagedDeviceManagementAgenteas                               ManagedDeviceManagementAgent = "Eas"
	ManagedDeviceManagementAgenteasIntuneClient                   ManagedDeviceManagementAgent = "EasIntuneClient"
	ManagedDeviceManagementAgenteasMdm                            ManagedDeviceManagementAgent = "EasMdm"
	ManagedDeviceManagementAgentgoogleCloudDevicePolicyController ManagedDeviceManagementAgent = "GoogleCloudDevicePolicyController"
	ManagedDeviceManagementAgentintuneAosp                        ManagedDeviceManagementAgent = "IntuneAosp"
	ManagedDeviceManagementAgentintuneClient                      ManagedDeviceManagementAgent = "IntuneClient"
	ManagedDeviceManagementAgentjamf                              ManagedDeviceManagementAgent = "Jamf"
	ManagedDeviceManagementAgentmdm                               ManagedDeviceManagementAgent = "Mdm"
	ManagedDeviceManagementAgentmicrosoft365ManagedMdm            ManagedDeviceManagementAgent = "Microsoft365ManagedMdm"
	ManagedDeviceManagementAgentmsSense                           ManagedDeviceManagementAgent = "MsSense"
	ManagedDeviceManagementAgentunknown                           ManagedDeviceManagementAgent = "Unknown"
)

func PossibleValuesForManagedDeviceManagementAgent() []string {
	return []string{
		string(ManagedDeviceManagementAgentconfigurationManagerClient),
		string(ManagedDeviceManagementAgentconfigurationManagerClientMdm),
		string(ManagedDeviceManagementAgentconfigurationManagerClientMdmEas),
		string(ManagedDeviceManagementAgenteas),
		string(ManagedDeviceManagementAgenteasIntuneClient),
		string(ManagedDeviceManagementAgenteasMdm),
		string(ManagedDeviceManagementAgentgoogleCloudDevicePolicyController),
		string(ManagedDeviceManagementAgentintuneAosp),
		string(ManagedDeviceManagementAgentintuneClient),
		string(ManagedDeviceManagementAgentjamf),
		string(ManagedDeviceManagementAgentmdm),
		string(ManagedDeviceManagementAgentmicrosoft365ManagedMdm),
		string(ManagedDeviceManagementAgentmsSense),
		string(ManagedDeviceManagementAgentunknown),
	}
}

func parseManagedDeviceManagementAgent(input string) (*ManagedDeviceManagementAgent, error) {
	vals := map[string]ManagedDeviceManagementAgent{
		"configurationmanagerclient":        ManagedDeviceManagementAgentconfigurationManagerClient,
		"configurationmanagerclientmdm":     ManagedDeviceManagementAgentconfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ManagedDeviceManagementAgentconfigurationManagerClientMdmEas,
		"eas":                               ManagedDeviceManagementAgenteas,
		"easintuneclient":                   ManagedDeviceManagementAgenteasIntuneClient,
		"easmdm":                            ManagedDeviceManagementAgenteasMdm,
		"googleclouddevicepolicycontroller": ManagedDeviceManagementAgentgoogleCloudDevicePolicyController,
		"intuneaosp":                        ManagedDeviceManagementAgentintuneAosp,
		"intuneclient":                      ManagedDeviceManagementAgentintuneClient,
		"jamf":                              ManagedDeviceManagementAgentjamf,
		"mdm":                               ManagedDeviceManagementAgentmdm,
		"microsoft365managedmdm":            ManagedDeviceManagementAgentmicrosoft365ManagedMdm,
		"mssense":                           ManagedDeviceManagementAgentmsSense,
		"unknown":                           ManagedDeviceManagementAgentunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementAgent(input)
	return &out, nil
}
