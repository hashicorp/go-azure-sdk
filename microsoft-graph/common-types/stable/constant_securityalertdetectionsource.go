package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertDetectionSource string

const (
	SecurityAlertDetectionSource_Antivirus                                    SecurityAlertDetectionSource = "antivirus"
	SecurityAlertDetectionSource_AppGovernanceDetection                       SecurityAlertDetectionSource = "appGovernanceDetection"
	SecurityAlertDetectionSource_AppGovernancePolicy                          SecurityAlertDetectionSource = "appGovernancePolicy"
	SecurityAlertDetectionSource_AutomatedInvestigation                       SecurityAlertDetectionSource = "automatedInvestigation"
	SecurityAlertDetectionSource_AzureAdIdentityProtection                    SecurityAlertDetectionSource = "azureAdIdentityProtection"
	SecurityAlertDetectionSource_BuiltInMl                                    SecurityAlertDetectionSource = "builtInMl"
	SecurityAlertDetectionSource_CloudAppSecurity                             SecurityAlertDetectionSource = "cloudAppSecurity"
	SecurityAlertDetectionSource_CustomDetection                              SecurityAlertDetectionSource = "customDetection"
	SecurityAlertDetectionSource_CustomTi                                     SecurityAlertDetectionSource = "customTi"
	SecurityAlertDetectionSource_Manual                                       SecurityAlertDetectionSource = "manual"
	SecurityAlertDetectionSource_Microsoft365Defender                         SecurityAlertDetectionSource = "microsoft365Defender"
	SecurityAlertDetectionSource_MicrosoftDataLossPrevention                  SecurityAlertDetectionSource = "microsoftDataLossPrevention"
	SecurityAlertDetectionSource_MicrosoftDefenderForApiManagement            SecurityAlertDetectionSource = "microsoftDefenderForApiManagement"
	SecurityAlertDetectionSource_MicrosoftDefenderForAppService               SecurityAlertDetectionSource = "microsoftDefenderForAppService"
	SecurityAlertDetectionSource_MicrosoftDefenderForCloud                    SecurityAlertDetectionSource = "microsoftDefenderForCloud"
	SecurityAlertDetectionSource_MicrosoftDefenderForContainers               SecurityAlertDetectionSource = "microsoftDefenderForContainers"
	SecurityAlertDetectionSource_MicrosoftDefenderForDNS                      SecurityAlertDetectionSource = "microsoftDefenderForDNS"
	SecurityAlertDetectionSource_MicrosoftDefenderForDatabases                SecurityAlertDetectionSource = "microsoftDefenderForDatabases"
	SecurityAlertDetectionSource_MicrosoftDefenderForEndpoint                 SecurityAlertDetectionSource = "microsoftDefenderForEndpoint"
	SecurityAlertDetectionSource_MicrosoftDefenderForIdentity                 SecurityAlertDetectionSource = "microsoftDefenderForIdentity"
	SecurityAlertDetectionSource_MicrosoftDefenderForIoT                      SecurityAlertDetectionSource = "microsoftDefenderForIoT"
	SecurityAlertDetectionSource_MicrosoftDefenderForKeyVault                 SecurityAlertDetectionSource = "microsoftDefenderForKeyVault"
	SecurityAlertDetectionSource_MicrosoftDefenderForNetwork                  SecurityAlertDetectionSource = "microsoftDefenderForNetwork"
	SecurityAlertDetectionSource_MicrosoftDefenderForOffice365                SecurityAlertDetectionSource = "microsoftDefenderForOffice365"
	SecurityAlertDetectionSource_MicrosoftDefenderForResourceManager          SecurityAlertDetectionSource = "microsoftDefenderForResourceManager"
	SecurityAlertDetectionSource_MicrosoftDefenderForServers                  SecurityAlertDetectionSource = "microsoftDefenderForServers"
	SecurityAlertDetectionSource_MicrosoftDefenderForStorage                  SecurityAlertDetectionSource = "microsoftDefenderForStorage"
	SecurityAlertDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics SecurityAlertDetectionSource = "microsoftDefenderThreatIntelligenceAnalytics"
	SecurityAlertDetectionSource_MicrosoftSentinel                            SecurityAlertDetectionSource = "microsoftSentinel"
	SecurityAlertDetectionSource_MicrosoftThreatExperts                       SecurityAlertDetectionSource = "microsoftThreatExperts"
	SecurityAlertDetectionSource_NrtAlerts                                    SecurityAlertDetectionSource = "nrtAlerts"
	SecurityAlertDetectionSource_ScheduledAlerts                              SecurityAlertDetectionSource = "scheduledAlerts"
	SecurityAlertDetectionSource_SmartScreen                                  SecurityAlertDetectionSource = "smartScreen"
	SecurityAlertDetectionSource_Unknown                                      SecurityAlertDetectionSource = "unknown"
)

func PossibleValuesForSecurityAlertDetectionSource() []string {
	return []string{
		string(SecurityAlertDetectionSource_Antivirus),
		string(SecurityAlertDetectionSource_AppGovernanceDetection),
		string(SecurityAlertDetectionSource_AppGovernancePolicy),
		string(SecurityAlertDetectionSource_AutomatedInvestigation),
		string(SecurityAlertDetectionSource_AzureAdIdentityProtection),
		string(SecurityAlertDetectionSource_BuiltInMl),
		string(SecurityAlertDetectionSource_CloudAppSecurity),
		string(SecurityAlertDetectionSource_CustomDetection),
		string(SecurityAlertDetectionSource_CustomTi),
		string(SecurityAlertDetectionSource_Manual),
		string(SecurityAlertDetectionSource_Microsoft365Defender),
		string(SecurityAlertDetectionSource_MicrosoftDataLossPrevention),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForApiManagement),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForAppService),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForCloud),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForContainers),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForDNS),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForDatabases),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForEndpoint),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForIdentity),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForIoT),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForKeyVault),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForNetwork),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForOffice365),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForResourceManager),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForServers),
		string(SecurityAlertDetectionSource_MicrosoftDefenderForStorage),
		string(SecurityAlertDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics),
		string(SecurityAlertDetectionSource_MicrosoftSentinel),
		string(SecurityAlertDetectionSource_MicrosoftThreatExperts),
		string(SecurityAlertDetectionSource_NrtAlerts),
		string(SecurityAlertDetectionSource_ScheduledAlerts),
		string(SecurityAlertDetectionSource_SmartScreen),
		string(SecurityAlertDetectionSource_Unknown),
	}
}

func (s *SecurityAlertDetectionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertDetectionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertDetectionSource(input string) (*SecurityAlertDetectionSource, error) {
	vals := map[string]SecurityAlertDetectionSource{
		"antivirus":                                    SecurityAlertDetectionSource_Antivirus,
		"appgovernancedetection":                       SecurityAlertDetectionSource_AppGovernanceDetection,
		"appgovernancepolicy":                          SecurityAlertDetectionSource_AppGovernancePolicy,
		"automatedinvestigation":                       SecurityAlertDetectionSource_AutomatedInvestigation,
		"azureadidentityprotection":                    SecurityAlertDetectionSource_AzureAdIdentityProtection,
		"builtinml":                                    SecurityAlertDetectionSource_BuiltInMl,
		"cloudappsecurity":                             SecurityAlertDetectionSource_CloudAppSecurity,
		"customdetection":                              SecurityAlertDetectionSource_CustomDetection,
		"customti":                                     SecurityAlertDetectionSource_CustomTi,
		"manual":                                       SecurityAlertDetectionSource_Manual,
		"microsoft365defender":                         SecurityAlertDetectionSource_Microsoft365Defender,
		"microsoftdatalossprevention":                  SecurityAlertDetectionSource_MicrosoftDataLossPrevention,
		"microsoftdefenderforapimanagement":            SecurityAlertDetectionSource_MicrosoftDefenderForApiManagement,
		"microsoftdefenderforappservice":               SecurityAlertDetectionSource_MicrosoftDefenderForAppService,
		"microsoftdefenderforcloud":                    SecurityAlertDetectionSource_MicrosoftDefenderForCloud,
		"microsoftdefenderforcontainers":               SecurityAlertDetectionSource_MicrosoftDefenderForContainers,
		"microsoftdefenderfordns":                      SecurityAlertDetectionSource_MicrosoftDefenderForDNS,
		"microsoftdefenderfordatabases":                SecurityAlertDetectionSource_MicrosoftDefenderForDatabases,
		"microsoftdefenderforendpoint":                 SecurityAlertDetectionSource_MicrosoftDefenderForEndpoint,
		"microsoftdefenderforidentity":                 SecurityAlertDetectionSource_MicrosoftDefenderForIdentity,
		"microsoftdefenderforiot":                      SecurityAlertDetectionSource_MicrosoftDefenderForIoT,
		"microsoftdefenderforkeyvault":                 SecurityAlertDetectionSource_MicrosoftDefenderForKeyVault,
		"microsoftdefenderfornetwork":                  SecurityAlertDetectionSource_MicrosoftDefenderForNetwork,
		"microsoftdefenderforoffice365":                SecurityAlertDetectionSource_MicrosoftDefenderForOffice365,
		"microsoftdefenderforresourcemanager":          SecurityAlertDetectionSource_MicrosoftDefenderForResourceManager,
		"microsoftdefenderforservers":                  SecurityAlertDetectionSource_MicrosoftDefenderForServers,
		"microsoftdefenderforstorage":                  SecurityAlertDetectionSource_MicrosoftDefenderForStorage,
		"microsoftdefenderthreatintelligenceanalytics": SecurityAlertDetectionSource_MicrosoftDefenderThreatIntelligenceAnalytics,
		"microsoftsentinel":                            SecurityAlertDetectionSource_MicrosoftSentinel,
		"microsoftthreatexperts":                       SecurityAlertDetectionSource_MicrosoftThreatExperts,
		"nrtalerts":                                    SecurityAlertDetectionSource_NrtAlerts,
		"scheduledalerts":                              SecurityAlertDetectionSource_ScheduledAlerts,
		"smartscreen":                                  SecurityAlertDetectionSource_SmartScreen,
		"unknown":                                      SecurityAlertDetectionSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertDetectionSource(input)
	return &out, nil
}
