package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertServiceSource string

const (
	SecurityAlertServiceSource_AzureAdIdentityProtection     SecurityAlertServiceSource = "azureAdIdentityProtection"
	SecurityAlertServiceSource_DataLossPrevention            SecurityAlertServiceSource = "dataLossPrevention"
	SecurityAlertServiceSource_Microsoft365Defender          SecurityAlertServiceSource = "microsoft365Defender"
	SecurityAlertServiceSource_MicrosoftAppGovernance        SecurityAlertServiceSource = "microsoftAppGovernance"
	SecurityAlertServiceSource_MicrosoftDefenderForCloud     SecurityAlertServiceSource = "microsoftDefenderForCloud"
	SecurityAlertServiceSource_MicrosoftDefenderForCloudApps SecurityAlertServiceSource = "microsoftDefenderForCloudApps"
	SecurityAlertServiceSource_MicrosoftDefenderForEndpoint  SecurityAlertServiceSource = "microsoftDefenderForEndpoint"
	SecurityAlertServiceSource_MicrosoftDefenderForIdentity  SecurityAlertServiceSource = "microsoftDefenderForIdentity"
	SecurityAlertServiceSource_MicrosoftDefenderForOffice365 SecurityAlertServiceSource = "microsoftDefenderForOffice365"
	SecurityAlertServiceSource_MicrosoftSentinel             SecurityAlertServiceSource = "microsoftSentinel"
	SecurityAlertServiceSource_Unknown                       SecurityAlertServiceSource = "unknown"
)

func PossibleValuesForSecurityAlertServiceSource() []string {
	return []string{
		string(SecurityAlertServiceSource_AzureAdIdentityProtection),
		string(SecurityAlertServiceSource_DataLossPrevention),
		string(SecurityAlertServiceSource_Microsoft365Defender),
		string(SecurityAlertServiceSource_MicrosoftAppGovernance),
		string(SecurityAlertServiceSource_MicrosoftDefenderForCloud),
		string(SecurityAlertServiceSource_MicrosoftDefenderForCloudApps),
		string(SecurityAlertServiceSource_MicrosoftDefenderForEndpoint),
		string(SecurityAlertServiceSource_MicrosoftDefenderForIdentity),
		string(SecurityAlertServiceSource_MicrosoftDefenderForOffice365),
		string(SecurityAlertServiceSource_MicrosoftSentinel),
		string(SecurityAlertServiceSource_Unknown),
	}
}

func (s *SecurityAlertServiceSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertServiceSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertServiceSource(input string) (*SecurityAlertServiceSource, error) {
	vals := map[string]SecurityAlertServiceSource{
		"azureadidentityprotection":     SecurityAlertServiceSource_AzureAdIdentityProtection,
		"datalossprevention":            SecurityAlertServiceSource_DataLossPrevention,
		"microsoft365defender":          SecurityAlertServiceSource_Microsoft365Defender,
		"microsoftappgovernance":        SecurityAlertServiceSource_MicrosoftAppGovernance,
		"microsoftdefenderforcloud":     SecurityAlertServiceSource_MicrosoftDefenderForCloud,
		"microsoftdefenderforcloudapps": SecurityAlertServiceSource_MicrosoftDefenderForCloudApps,
		"microsoftdefenderforendpoint":  SecurityAlertServiceSource_MicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":  SecurityAlertServiceSource_MicrosoftDefenderForIdentity,
		"microsoftdefenderforoffice365": SecurityAlertServiceSource_MicrosoftDefenderForOffice365,
		"microsoftsentinel":             SecurityAlertServiceSource_MicrosoftSentinel,
		"unknown":                       SecurityAlertServiceSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertServiceSource(input)
	return &out, nil
}
