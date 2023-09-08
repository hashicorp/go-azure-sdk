package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceManagementAgents string

const (
	ComanagementEligibleDeviceManagementAgentsconfigurationManagerClient        ComanagementEligibleDeviceManagementAgents = "ConfigurationManagerClient"
	ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdm     ComanagementEligibleDeviceManagementAgents = "ConfigurationManagerClientMdm"
	ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdmEas  ComanagementEligibleDeviceManagementAgents = "ConfigurationManagerClientMdmEas"
	ComanagementEligibleDeviceManagementAgentseas                               ComanagementEligibleDeviceManagementAgents = "Eas"
	ComanagementEligibleDeviceManagementAgentseasIntuneClient                   ComanagementEligibleDeviceManagementAgents = "EasIntuneClient"
	ComanagementEligibleDeviceManagementAgentseasMdm                            ComanagementEligibleDeviceManagementAgents = "EasMdm"
	ComanagementEligibleDeviceManagementAgentsgoogleCloudDevicePolicyController ComanagementEligibleDeviceManagementAgents = "GoogleCloudDevicePolicyController"
	ComanagementEligibleDeviceManagementAgentsintuneAosp                        ComanagementEligibleDeviceManagementAgents = "IntuneAosp"
	ComanagementEligibleDeviceManagementAgentsintuneClient                      ComanagementEligibleDeviceManagementAgents = "IntuneClient"
	ComanagementEligibleDeviceManagementAgentsjamf                              ComanagementEligibleDeviceManagementAgents = "Jamf"
	ComanagementEligibleDeviceManagementAgentsmdm                               ComanagementEligibleDeviceManagementAgents = "Mdm"
	ComanagementEligibleDeviceManagementAgentsmicrosoft365ManagedMdm            ComanagementEligibleDeviceManagementAgents = "Microsoft365ManagedMdm"
	ComanagementEligibleDeviceManagementAgentsmsSense                           ComanagementEligibleDeviceManagementAgents = "MsSense"
	ComanagementEligibleDeviceManagementAgentsunknown                           ComanagementEligibleDeviceManagementAgents = "Unknown"
)

func PossibleValuesForComanagementEligibleDeviceManagementAgents() []string {
	return []string{
		string(ComanagementEligibleDeviceManagementAgentsconfigurationManagerClient),
		string(ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdm),
		string(ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdmEas),
		string(ComanagementEligibleDeviceManagementAgentseas),
		string(ComanagementEligibleDeviceManagementAgentseasIntuneClient),
		string(ComanagementEligibleDeviceManagementAgentseasMdm),
		string(ComanagementEligibleDeviceManagementAgentsgoogleCloudDevicePolicyController),
		string(ComanagementEligibleDeviceManagementAgentsintuneAosp),
		string(ComanagementEligibleDeviceManagementAgentsintuneClient),
		string(ComanagementEligibleDeviceManagementAgentsjamf),
		string(ComanagementEligibleDeviceManagementAgentsmdm),
		string(ComanagementEligibleDeviceManagementAgentsmicrosoft365ManagedMdm),
		string(ComanagementEligibleDeviceManagementAgentsmsSense),
		string(ComanagementEligibleDeviceManagementAgentsunknown),
	}
}

func parseComanagementEligibleDeviceManagementAgents(input string) (*ComanagementEligibleDeviceManagementAgents, error) {
	vals := map[string]ComanagementEligibleDeviceManagementAgents{
		"configurationmanagerclient":        ComanagementEligibleDeviceManagementAgentsconfigurationManagerClient,
		"configurationmanagerclientmdm":     ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ComanagementEligibleDeviceManagementAgentsconfigurationManagerClientMdmEas,
		"eas":                               ComanagementEligibleDeviceManagementAgentseas,
		"easintuneclient":                   ComanagementEligibleDeviceManagementAgentseasIntuneClient,
		"easmdm":                            ComanagementEligibleDeviceManagementAgentseasMdm,
		"googleclouddevicepolicycontroller": ComanagementEligibleDeviceManagementAgentsgoogleCloudDevicePolicyController,
		"intuneaosp":                        ComanagementEligibleDeviceManagementAgentsintuneAosp,
		"intuneclient":                      ComanagementEligibleDeviceManagementAgentsintuneClient,
		"jamf":                              ComanagementEligibleDeviceManagementAgentsjamf,
		"mdm":                               ComanagementEligibleDeviceManagementAgentsmdm,
		"microsoft365managedmdm":            ComanagementEligibleDeviceManagementAgentsmicrosoft365ManagedMdm,
		"mssense":                           ComanagementEligibleDeviceManagementAgentsmsSense,
		"unknown":                           ComanagementEligibleDeviceManagementAgentsunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceManagementAgents(input)
	return &out, nil
}
