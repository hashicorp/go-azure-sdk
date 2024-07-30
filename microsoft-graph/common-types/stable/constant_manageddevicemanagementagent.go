package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementAgent string

const (
	ManagedDeviceManagementAgent_ConfigurationManagerClient        ManagedDeviceManagementAgent = "configurationManagerClient"
	ManagedDeviceManagementAgent_ConfigurationManagerClientMdm     ManagedDeviceManagementAgent = "configurationManagerClientMdm"
	ManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas  ManagedDeviceManagementAgent = "configurationManagerClientMdmEas"
	ManagedDeviceManagementAgent_Eas                               ManagedDeviceManagementAgent = "eas"
	ManagedDeviceManagementAgent_EasIntuneClient                   ManagedDeviceManagementAgent = "easIntuneClient"
	ManagedDeviceManagementAgent_EasMdm                            ManagedDeviceManagementAgent = "easMdm"
	ManagedDeviceManagementAgent_GoogleCloudDevicePolicyController ManagedDeviceManagementAgent = "googleCloudDevicePolicyController"
	ManagedDeviceManagementAgent_IntuneClient                      ManagedDeviceManagementAgent = "intuneClient"
	ManagedDeviceManagementAgent_Jamf                              ManagedDeviceManagementAgent = "jamf"
	ManagedDeviceManagementAgent_Mdm                               ManagedDeviceManagementAgent = "mdm"
	ManagedDeviceManagementAgent_Microsoft365ManagedMdm            ManagedDeviceManagementAgent = "microsoft365ManagedMdm"
	ManagedDeviceManagementAgent_MsSense                           ManagedDeviceManagementAgent = "msSense"
	ManagedDeviceManagementAgent_Unknown                           ManagedDeviceManagementAgent = "unknown"
)

func PossibleValuesForManagedDeviceManagementAgent() []string {
	return []string{
		string(ManagedDeviceManagementAgent_ConfigurationManagerClient),
		string(ManagedDeviceManagementAgent_ConfigurationManagerClientMdm),
		string(ManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas),
		string(ManagedDeviceManagementAgent_Eas),
		string(ManagedDeviceManagementAgent_EasIntuneClient),
		string(ManagedDeviceManagementAgent_EasMdm),
		string(ManagedDeviceManagementAgent_GoogleCloudDevicePolicyController),
		string(ManagedDeviceManagementAgent_IntuneClient),
		string(ManagedDeviceManagementAgent_Jamf),
		string(ManagedDeviceManagementAgent_Mdm),
		string(ManagedDeviceManagementAgent_Microsoft365ManagedMdm),
		string(ManagedDeviceManagementAgent_MsSense),
		string(ManagedDeviceManagementAgent_Unknown),
	}
}

func (s *ManagedDeviceManagementAgent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementAgent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementAgent(input string) (*ManagedDeviceManagementAgent, error) {
	vals := map[string]ManagedDeviceManagementAgent{
		"configurationmanagerclient":        ManagedDeviceManagementAgent_ConfigurationManagerClient,
		"configurationmanagerclientmdm":     ManagedDeviceManagementAgent_ConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas,
		"eas":                               ManagedDeviceManagementAgent_Eas,
		"easintuneclient":                   ManagedDeviceManagementAgent_EasIntuneClient,
		"easmdm":                            ManagedDeviceManagementAgent_EasMdm,
		"googleclouddevicepolicycontroller": ManagedDeviceManagementAgent_GoogleCloudDevicePolicyController,
		"intuneclient":                      ManagedDeviceManagementAgent_IntuneClient,
		"jamf":                              ManagedDeviceManagementAgent_Jamf,
		"mdm":                               ManagedDeviceManagementAgent_Mdm,
		"microsoft365managedmdm":            ManagedDeviceManagementAgent_Microsoft365ManagedMdm,
		"mssense":                           ManagedDeviceManagementAgent_MsSense,
		"unknown":                           ManagedDeviceManagementAgent_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementAgent(input)
	return &out, nil
}
