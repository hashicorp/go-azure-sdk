package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceManagementAgent string

const (
	RetireScheduledManagedDeviceManagementAgentconfigurationManagerClient        RetireScheduledManagedDeviceManagementAgent = "ConfigurationManagerClient"
	RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdm     RetireScheduledManagedDeviceManagementAgent = "ConfigurationManagerClientMdm"
	RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdmEas  RetireScheduledManagedDeviceManagementAgent = "ConfigurationManagerClientMdmEas"
	RetireScheduledManagedDeviceManagementAgenteas                               RetireScheduledManagedDeviceManagementAgent = "Eas"
	RetireScheduledManagedDeviceManagementAgenteasIntuneClient                   RetireScheduledManagedDeviceManagementAgent = "EasIntuneClient"
	RetireScheduledManagedDeviceManagementAgenteasMdm                            RetireScheduledManagedDeviceManagementAgent = "EasMdm"
	RetireScheduledManagedDeviceManagementAgentgoogleCloudDevicePolicyController RetireScheduledManagedDeviceManagementAgent = "GoogleCloudDevicePolicyController"
	RetireScheduledManagedDeviceManagementAgentintuneAosp                        RetireScheduledManagedDeviceManagementAgent = "IntuneAosp"
	RetireScheduledManagedDeviceManagementAgentintuneClient                      RetireScheduledManagedDeviceManagementAgent = "IntuneClient"
	RetireScheduledManagedDeviceManagementAgentjamf                              RetireScheduledManagedDeviceManagementAgent = "Jamf"
	RetireScheduledManagedDeviceManagementAgentmdm                               RetireScheduledManagedDeviceManagementAgent = "Mdm"
	RetireScheduledManagedDeviceManagementAgentmicrosoft365ManagedMdm            RetireScheduledManagedDeviceManagementAgent = "Microsoft365ManagedMdm"
	RetireScheduledManagedDeviceManagementAgentmsSense                           RetireScheduledManagedDeviceManagementAgent = "MsSense"
	RetireScheduledManagedDeviceManagementAgentunknown                           RetireScheduledManagedDeviceManagementAgent = "Unknown"
)

func PossibleValuesForRetireScheduledManagedDeviceManagementAgent() []string {
	return []string{
		string(RetireScheduledManagedDeviceManagementAgentconfigurationManagerClient),
		string(RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdm),
		string(RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdmEas),
		string(RetireScheduledManagedDeviceManagementAgenteas),
		string(RetireScheduledManagedDeviceManagementAgenteasIntuneClient),
		string(RetireScheduledManagedDeviceManagementAgenteasMdm),
		string(RetireScheduledManagedDeviceManagementAgentgoogleCloudDevicePolicyController),
		string(RetireScheduledManagedDeviceManagementAgentintuneAosp),
		string(RetireScheduledManagedDeviceManagementAgentintuneClient),
		string(RetireScheduledManagedDeviceManagementAgentjamf),
		string(RetireScheduledManagedDeviceManagementAgentmdm),
		string(RetireScheduledManagedDeviceManagementAgentmicrosoft365ManagedMdm),
		string(RetireScheduledManagedDeviceManagementAgentmsSense),
		string(RetireScheduledManagedDeviceManagementAgentunknown),
	}
}

func parseRetireScheduledManagedDeviceManagementAgent(input string) (*RetireScheduledManagedDeviceManagementAgent, error) {
	vals := map[string]RetireScheduledManagedDeviceManagementAgent{
		"configurationmanagerclient":        RetireScheduledManagedDeviceManagementAgentconfigurationManagerClient,
		"configurationmanagerclientmdm":     RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  RetireScheduledManagedDeviceManagementAgentconfigurationManagerClientMdmEas,
		"eas":                               RetireScheduledManagedDeviceManagementAgenteas,
		"easintuneclient":                   RetireScheduledManagedDeviceManagementAgenteasIntuneClient,
		"easmdm":                            RetireScheduledManagedDeviceManagementAgenteasMdm,
		"googleclouddevicepolicycontroller": RetireScheduledManagedDeviceManagementAgentgoogleCloudDevicePolicyController,
		"intuneaosp":                        RetireScheduledManagedDeviceManagementAgentintuneAosp,
		"intuneclient":                      RetireScheduledManagedDeviceManagementAgentintuneClient,
		"jamf":                              RetireScheduledManagedDeviceManagementAgentjamf,
		"mdm":                               RetireScheduledManagedDeviceManagementAgentmdm,
		"microsoft365managedmdm":            RetireScheduledManagedDeviceManagementAgentmicrosoft365ManagedMdm,
		"mssense":                           RetireScheduledManagedDeviceManagementAgentmsSense,
		"unknown":                           RetireScheduledManagedDeviceManagementAgentunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceManagementAgent(input)
	return &out, nil
}
