package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceManagementAgents string

const (
	ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClient        ComanagementEligibleDeviceManagementAgents = "configurationManagerClient"
	ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdm     ComanagementEligibleDeviceManagementAgents = "configurationManagerClientMdm"
	ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdmEas  ComanagementEligibleDeviceManagementAgents = "configurationManagerClientMdmEas"
	ComanagementEligibleDeviceManagementAgents_Eas                               ComanagementEligibleDeviceManagementAgents = "eas"
	ComanagementEligibleDeviceManagementAgents_EasIntuneClient                   ComanagementEligibleDeviceManagementAgents = "easIntuneClient"
	ComanagementEligibleDeviceManagementAgents_EasMdm                            ComanagementEligibleDeviceManagementAgents = "easMdm"
	ComanagementEligibleDeviceManagementAgents_GoogleCloudDevicePolicyController ComanagementEligibleDeviceManagementAgents = "googleCloudDevicePolicyController"
	ComanagementEligibleDeviceManagementAgents_IntuneAosp                        ComanagementEligibleDeviceManagementAgents = "intuneAosp"
	ComanagementEligibleDeviceManagementAgents_IntuneClient                      ComanagementEligibleDeviceManagementAgents = "intuneClient"
	ComanagementEligibleDeviceManagementAgents_Jamf                              ComanagementEligibleDeviceManagementAgents = "jamf"
	ComanagementEligibleDeviceManagementAgents_Mdm                               ComanagementEligibleDeviceManagementAgents = "mdm"
	ComanagementEligibleDeviceManagementAgents_Microsoft365ManagedMdm            ComanagementEligibleDeviceManagementAgents = "microsoft365ManagedMdm"
	ComanagementEligibleDeviceManagementAgents_MsSense                           ComanagementEligibleDeviceManagementAgents = "msSense"
	ComanagementEligibleDeviceManagementAgents_Unknown                           ComanagementEligibleDeviceManagementAgents = "unknown"
)

func PossibleValuesForComanagementEligibleDeviceManagementAgents() []string {
	return []string{
		string(ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClient),
		string(ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdm),
		string(ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdmEas),
		string(ComanagementEligibleDeviceManagementAgents_Eas),
		string(ComanagementEligibleDeviceManagementAgents_EasIntuneClient),
		string(ComanagementEligibleDeviceManagementAgents_EasMdm),
		string(ComanagementEligibleDeviceManagementAgents_GoogleCloudDevicePolicyController),
		string(ComanagementEligibleDeviceManagementAgents_IntuneAosp),
		string(ComanagementEligibleDeviceManagementAgents_IntuneClient),
		string(ComanagementEligibleDeviceManagementAgents_Jamf),
		string(ComanagementEligibleDeviceManagementAgents_Mdm),
		string(ComanagementEligibleDeviceManagementAgents_Microsoft365ManagedMdm),
		string(ComanagementEligibleDeviceManagementAgents_MsSense),
		string(ComanagementEligibleDeviceManagementAgents_Unknown),
	}
}

func (s *ComanagementEligibleDeviceManagementAgents) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceManagementAgents(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceManagementAgents(input string) (*ComanagementEligibleDeviceManagementAgents, error) {
	vals := map[string]ComanagementEligibleDeviceManagementAgents{
		"configurationmanagerclient":        ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClient,
		"configurationmanagerclientmdm":     ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ComanagementEligibleDeviceManagementAgents_ConfigurationManagerClientMdmEas,
		"eas":                               ComanagementEligibleDeviceManagementAgents_Eas,
		"easintuneclient":                   ComanagementEligibleDeviceManagementAgents_EasIntuneClient,
		"easmdm":                            ComanagementEligibleDeviceManagementAgents_EasMdm,
		"googleclouddevicepolicycontroller": ComanagementEligibleDeviceManagementAgents_GoogleCloudDevicePolicyController,
		"intuneaosp":                        ComanagementEligibleDeviceManagementAgents_IntuneAosp,
		"intuneclient":                      ComanagementEligibleDeviceManagementAgents_IntuneClient,
		"jamf":                              ComanagementEligibleDeviceManagementAgents_Jamf,
		"mdm":                               ComanagementEligibleDeviceManagementAgents_Mdm,
		"microsoft365managedmdm":            ComanagementEligibleDeviceManagementAgents_Microsoft365ManagedMdm,
		"mssense":                           ComanagementEligibleDeviceManagementAgents_MsSense,
		"unknown":                           ComanagementEligibleDeviceManagementAgents_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceManagementAgents(input)
	return &out, nil
}
