package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagementAgent string

const (
	WindowsManagedDeviceManagementAgent_ConfigurationManagerClient        WindowsManagedDeviceManagementAgent = "configurationManagerClient"
	WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdm     WindowsManagedDeviceManagementAgent = "configurationManagerClientMdm"
	WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas  WindowsManagedDeviceManagementAgent = "configurationManagerClientMdmEas"
	WindowsManagedDeviceManagementAgent_Eas                               WindowsManagedDeviceManagementAgent = "eas"
	WindowsManagedDeviceManagementAgent_EasIntuneClient                   WindowsManagedDeviceManagementAgent = "easIntuneClient"
	WindowsManagedDeviceManagementAgent_EasMdm                            WindowsManagedDeviceManagementAgent = "easMdm"
	WindowsManagedDeviceManagementAgent_GoogleCloudDevicePolicyController WindowsManagedDeviceManagementAgent = "googleCloudDevicePolicyController"
	WindowsManagedDeviceManagementAgent_IntuneAosp                        WindowsManagedDeviceManagementAgent = "intuneAosp"
	WindowsManagedDeviceManagementAgent_IntuneClient                      WindowsManagedDeviceManagementAgent = "intuneClient"
	WindowsManagedDeviceManagementAgent_Jamf                              WindowsManagedDeviceManagementAgent = "jamf"
	WindowsManagedDeviceManagementAgent_Mdm                               WindowsManagedDeviceManagementAgent = "mdm"
	WindowsManagedDeviceManagementAgent_Microsoft365ManagedMdm            WindowsManagedDeviceManagementAgent = "microsoft365ManagedMdm"
	WindowsManagedDeviceManagementAgent_MsSense                           WindowsManagedDeviceManagementAgent = "msSense"
	WindowsManagedDeviceManagementAgent_Unknown                           WindowsManagedDeviceManagementAgent = "unknown"
)

func PossibleValuesForWindowsManagedDeviceManagementAgent() []string {
	return []string{
		string(WindowsManagedDeviceManagementAgent_ConfigurationManagerClient),
		string(WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdm),
		string(WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas),
		string(WindowsManagedDeviceManagementAgent_Eas),
		string(WindowsManagedDeviceManagementAgent_EasIntuneClient),
		string(WindowsManagedDeviceManagementAgent_EasMdm),
		string(WindowsManagedDeviceManagementAgent_GoogleCloudDevicePolicyController),
		string(WindowsManagedDeviceManagementAgent_IntuneAosp),
		string(WindowsManagedDeviceManagementAgent_IntuneClient),
		string(WindowsManagedDeviceManagementAgent_Jamf),
		string(WindowsManagedDeviceManagementAgent_Mdm),
		string(WindowsManagedDeviceManagementAgent_Microsoft365ManagedMdm),
		string(WindowsManagedDeviceManagementAgent_MsSense),
		string(WindowsManagedDeviceManagementAgent_Unknown),
	}
}

func (s *WindowsManagedDeviceManagementAgent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceManagementAgent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceManagementAgent(input string) (*WindowsManagedDeviceManagementAgent, error) {
	vals := map[string]WindowsManagedDeviceManagementAgent{
		"configurationmanagerclient":        WindowsManagedDeviceManagementAgent_ConfigurationManagerClient,
		"configurationmanagerclientmdm":     WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  WindowsManagedDeviceManagementAgent_ConfigurationManagerClientMdmEas,
		"eas":                               WindowsManagedDeviceManagementAgent_Eas,
		"easintuneclient":                   WindowsManagedDeviceManagementAgent_EasIntuneClient,
		"easmdm":                            WindowsManagedDeviceManagementAgent_EasMdm,
		"googleclouddevicepolicycontroller": WindowsManagedDeviceManagementAgent_GoogleCloudDevicePolicyController,
		"intuneaosp":                        WindowsManagedDeviceManagementAgent_IntuneAosp,
		"intuneclient":                      WindowsManagedDeviceManagementAgent_IntuneClient,
		"jamf":                              WindowsManagedDeviceManagementAgent_Jamf,
		"mdm":                               WindowsManagedDeviceManagementAgent_Mdm,
		"microsoft365managedmdm":            WindowsManagedDeviceManagementAgent_Microsoft365ManagedMdm,
		"mssense":                           WindowsManagedDeviceManagementAgent_MsSense,
		"unknown":                           WindowsManagedDeviceManagementAgent_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagementAgent(input)
	return &out, nil
}
