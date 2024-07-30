package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceManagementAgent string

const (
	RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClient        RetireScheduledManagedDeviceManagementAgent = "configurationManagerClient"
	RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdm     RetireScheduledManagedDeviceManagementAgent = "configurationManagerClientMdm"
	RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas  RetireScheduledManagedDeviceManagementAgent = "configurationManagerClientMdmEas"
	RetireScheduledManagedDeviceManagementAgent_Eas                               RetireScheduledManagedDeviceManagementAgent = "eas"
	RetireScheduledManagedDeviceManagementAgent_EasIntuneClient                   RetireScheduledManagedDeviceManagementAgent = "easIntuneClient"
	RetireScheduledManagedDeviceManagementAgent_EasMdm                            RetireScheduledManagedDeviceManagementAgent = "easMdm"
	RetireScheduledManagedDeviceManagementAgent_GoogleCloudDevicePolicyController RetireScheduledManagedDeviceManagementAgent = "googleCloudDevicePolicyController"
	RetireScheduledManagedDeviceManagementAgent_IntuneAosp                        RetireScheduledManagedDeviceManagementAgent = "intuneAosp"
	RetireScheduledManagedDeviceManagementAgent_IntuneClient                      RetireScheduledManagedDeviceManagementAgent = "intuneClient"
	RetireScheduledManagedDeviceManagementAgent_Jamf                              RetireScheduledManagedDeviceManagementAgent = "jamf"
	RetireScheduledManagedDeviceManagementAgent_Mdm                               RetireScheduledManagedDeviceManagementAgent = "mdm"
	RetireScheduledManagedDeviceManagementAgent_Microsoft365ManagedMdm            RetireScheduledManagedDeviceManagementAgent = "microsoft365ManagedMdm"
	RetireScheduledManagedDeviceManagementAgent_MsSense                           RetireScheduledManagedDeviceManagementAgent = "msSense"
	RetireScheduledManagedDeviceManagementAgent_Unknown                           RetireScheduledManagedDeviceManagementAgent = "unknown"
)

func PossibleValuesForRetireScheduledManagedDeviceManagementAgent() []string {
	return []string{
		string(RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClient),
		string(RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdm),
		string(RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas),
		string(RetireScheduledManagedDeviceManagementAgent_Eas),
		string(RetireScheduledManagedDeviceManagementAgent_EasIntuneClient),
		string(RetireScheduledManagedDeviceManagementAgent_EasMdm),
		string(RetireScheduledManagedDeviceManagementAgent_GoogleCloudDevicePolicyController),
		string(RetireScheduledManagedDeviceManagementAgent_IntuneAosp),
		string(RetireScheduledManagedDeviceManagementAgent_IntuneClient),
		string(RetireScheduledManagedDeviceManagementAgent_Jamf),
		string(RetireScheduledManagedDeviceManagementAgent_Mdm),
		string(RetireScheduledManagedDeviceManagementAgent_Microsoft365ManagedMdm),
		string(RetireScheduledManagedDeviceManagementAgent_MsSense),
		string(RetireScheduledManagedDeviceManagementAgent_Unknown),
	}
}

func (s *RetireScheduledManagedDeviceManagementAgent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetireScheduledManagedDeviceManagementAgent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetireScheduledManagedDeviceManagementAgent(input string) (*RetireScheduledManagedDeviceManagementAgent, error) {
	vals := map[string]RetireScheduledManagedDeviceManagementAgent{
		"configurationmanagerclient":        RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClient,
		"configurationmanagerclientmdm":     RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  RetireScheduledManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas,
		"eas":                               RetireScheduledManagedDeviceManagementAgent_Eas,
		"easintuneclient":                   RetireScheduledManagedDeviceManagementAgent_EasIntuneClient,
		"easmdm":                            RetireScheduledManagedDeviceManagementAgent_EasMdm,
		"googleclouddevicepolicycontroller": RetireScheduledManagedDeviceManagementAgent_GoogleCloudDevicePolicyController,
		"intuneaosp":                        RetireScheduledManagedDeviceManagementAgent_IntuneAosp,
		"intuneclient":                      RetireScheduledManagedDeviceManagementAgent_IntuneClient,
		"jamf":                              RetireScheduledManagedDeviceManagementAgent_Jamf,
		"mdm":                               RetireScheduledManagedDeviceManagementAgent_Mdm,
		"microsoft365managedmdm":            RetireScheduledManagedDeviceManagementAgent_Microsoft365ManagedMdm,
		"mssense":                           RetireScheduledManagedDeviceManagementAgent_MsSense,
		"unknown":                           RetireScheduledManagedDeviceManagementAgent_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceManagementAgent(input)
	return &out, nil
}
