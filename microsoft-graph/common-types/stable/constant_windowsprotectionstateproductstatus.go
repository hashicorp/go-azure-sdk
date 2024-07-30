package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsProtectionStateProductStatus string

const (
	WindowsProtectionStateProductStatus_AsSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "asSignaturesOutOfDate"
	WindowsProtectionStateProductStatus_AvSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "avSignaturesOutOfDate"
	WindowsProtectionStateProductStatus_NoFullScanHappenedForSpecifiedPeriod            WindowsProtectionStateProductStatus = "noFullScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatus_NoQuickScanHappenedForSpecifiedPeriod           WindowsProtectionStateProductStatus = "noQuickScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatus_NoStatus                                        WindowsProtectionStateProductStatus = "noStatus"
	WindowsProtectionStateProductStatus_NoStatusFlagsSet                                WindowsProtectionStateProductStatus = "noStatusFlagsSet"
	WindowsProtectionStateProductStatus_OfflineScanRequired                             WindowsProtectionStateProductStatus = "offlineScanRequired"
	WindowsProtectionStateProductStatus_PendingFullScanDueToThreatAction                WindowsProtectionStateProductStatus = "pendingFullScanDueToThreatAction"
	WindowsProtectionStateProductStatus_PendingManualStepsDueToThreatAction             WindowsProtectionStateProductStatus = "pendingManualStepsDueToThreatAction"
	WindowsProtectionStateProductStatus_PendingRebootDueToThreatAction                  WindowsProtectionStateProductStatus = "pendingRebootDueToThreatAction"
	WindowsProtectionStateProductStatus_PlatformAboutToBeOutdated                       WindowsProtectionStateProductStatus = "platformAboutToBeOutdated"
	WindowsProtectionStateProductStatus_PlatformOutOfDate                               WindowsProtectionStateProductStatus = "platformOutOfDate"
	WindowsProtectionStateProductStatus_PlatformUpdateInProgress                        WindowsProtectionStateProductStatus = "platformUpdateInProgress"
	WindowsProtectionStateProductStatus_ProductExpired                                  WindowsProtectionStateProductStatus = "productExpired"
	WindowsProtectionStateProductStatus_ProductRunningInEvaluationMode                  WindowsProtectionStateProductStatus = "productRunningInEvaluationMode"
	WindowsProtectionStateProductStatus_ProductRunningInNonGenuineMode                  WindowsProtectionStateProductStatus = "productRunningInNonGenuineMode"
	WindowsProtectionStateProductStatus_SamplesPendingSubmission                        WindowsProtectionStateProductStatus = "samplesPendingSubmission"
	WindowsProtectionStateProductStatus_ServiceNotRunning                               WindowsProtectionStateProductStatus = "serviceNotRunning"
	WindowsProtectionStateProductStatus_ServiceShutdownAsPartOfSystemShutdown           WindowsProtectionStateProductStatus = "serviceShutdownAsPartOfSystemShutdown"
	WindowsProtectionStateProductStatus_ServiceStartedWithoutMalwareProtection          WindowsProtectionStateProductStatus = "serviceStartedWithoutMalwareProtection"
	WindowsProtectionStateProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending WindowsProtectionStateProductStatus = "signatureOrPlatformEndOfLifeIsPastOrIsImpending"
	WindowsProtectionStateProductStatus_SystemInitiatedCleanInProgress                  WindowsProtectionStateProductStatus = "systemInitiatedCleanInProgress"
	WindowsProtectionStateProductStatus_SystemInitiatedScanInProgress                   WindowsProtectionStateProductStatus = "systemInitiatedScanInProgress"
	WindowsProtectionStateProductStatus_ThreatRemediationFailedCritically               WindowsProtectionStateProductStatus = "threatRemediationFailedCritically"
	WindowsProtectionStateProductStatus_ThreatRemediationFailedNonCritically            WindowsProtectionStateProductStatus = "threatRemediationFailedNonCritically"
	WindowsProtectionStateProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall   WindowsProtectionStateProductStatus = "windowsSModeSignaturesInUseOnNonWin10SInstall"
)

func PossibleValuesForWindowsProtectionStateProductStatus() []string {
	return []string{
		string(WindowsProtectionStateProductStatus_AsSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatus_AvSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatus_NoFullScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatus_NoQuickScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatus_NoStatus),
		string(WindowsProtectionStateProductStatus_NoStatusFlagsSet),
		string(WindowsProtectionStateProductStatus_OfflineScanRequired),
		string(WindowsProtectionStateProductStatus_PendingFullScanDueToThreatAction),
		string(WindowsProtectionStateProductStatus_PendingManualStepsDueToThreatAction),
		string(WindowsProtectionStateProductStatus_PendingRebootDueToThreatAction),
		string(WindowsProtectionStateProductStatus_PlatformAboutToBeOutdated),
		string(WindowsProtectionStateProductStatus_PlatformOutOfDate),
		string(WindowsProtectionStateProductStatus_PlatformUpdateInProgress),
		string(WindowsProtectionStateProductStatus_ProductExpired),
		string(WindowsProtectionStateProductStatus_ProductRunningInEvaluationMode),
		string(WindowsProtectionStateProductStatus_ProductRunningInNonGenuineMode),
		string(WindowsProtectionStateProductStatus_SamplesPendingSubmission),
		string(WindowsProtectionStateProductStatus_ServiceNotRunning),
		string(WindowsProtectionStateProductStatus_ServiceShutdownAsPartOfSystemShutdown),
		string(WindowsProtectionStateProductStatus_ServiceStartedWithoutMalwareProtection),
		string(WindowsProtectionStateProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending),
		string(WindowsProtectionStateProductStatus_SystemInitiatedCleanInProgress),
		string(WindowsProtectionStateProductStatus_SystemInitiatedScanInProgress),
		string(WindowsProtectionStateProductStatus_ThreatRemediationFailedCritically),
		string(WindowsProtectionStateProductStatus_ThreatRemediationFailedNonCritically),
		string(WindowsProtectionStateProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall),
	}
}

func (s *WindowsProtectionStateProductStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsProtectionStateProductStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsProtectionStateProductStatus(input string) (*WindowsProtectionStateProductStatus, error) {
	vals := map[string]WindowsProtectionStateProductStatus{
		"assignaturesoutofdate":                           WindowsProtectionStateProductStatus_AsSignaturesOutOfDate,
		"avsignaturesoutofdate":                           WindowsProtectionStateProductStatus_AvSignaturesOutOfDate,
		"nofullscanhappenedforspecifiedperiod":            WindowsProtectionStateProductStatus_NoFullScanHappenedForSpecifiedPeriod,
		"noquickscanhappenedforspecifiedperiod":           WindowsProtectionStateProductStatus_NoQuickScanHappenedForSpecifiedPeriod,
		"nostatus":                                        WindowsProtectionStateProductStatus_NoStatus,
		"nostatusflagsset":                                WindowsProtectionStateProductStatus_NoStatusFlagsSet,
		"offlinescanrequired":                             WindowsProtectionStateProductStatus_OfflineScanRequired,
		"pendingfullscanduetothreataction":                WindowsProtectionStateProductStatus_PendingFullScanDueToThreatAction,
		"pendingmanualstepsduetothreataction":             WindowsProtectionStateProductStatus_PendingManualStepsDueToThreatAction,
		"pendingrebootduetothreataction":                  WindowsProtectionStateProductStatus_PendingRebootDueToThreatAction,
		"platformabouttobeoutdated":                       WindowsProtectionStateProductStatus_PlatformAboutToBeOutdated,
		"platformoutofdate":                               WindowsProtectionStateProductStatus_PlatformOutOfDate,
		"platformupdateinprogress":                        WindowsProtectionStateProductStatus_PlatformUpdateInProgress,
		"productexpired":                                  WindowsProtectionStateProductStatus_ProductExpired,
		"productrunninginevaluationmode":                  WindowsProtectionStateProductStatus_ProductRunningInEvaluationMode,
		"productrunninginnongenuinemode":                  WindowsProtectionStateProductStatus_ProductRunningInNonGenuineMode,
		"samplespendingsubmission":                        WindowsProtectionStateProductStatus_SamplesPendingSubmission,
		"servicenotrunning":                               WindowsProtectionStateProductStatus_ServiceNotRunning,
		"serviceshutdownaspartofsystemshutdown":           WindowsProtectionStateProductStatus_ServiceShutdownAsPartOfSystemShutdown,
		"servicestartedwithoutmalwareprotection":          WindowsProtectionStateProductStatus_ServiceStartedWithoutMalwareProtection,
		"signatureorplatformendoflifeispastorisimpending": WindowsProtectionStateProductStatus_SignatureOrPlatformEndOfLifeIsPastOrIsImpending,
		"systeminitiatedcleaninprogress":                  WindowsProtectionStateProductStatus_SystemInitiatedCleanInProgress,
		"systeminitiatedscaninprogress":                   WindowsProtectionStateProductStatus_SystemInitiatedScanInProgress,
		"threatremediationfailedcritically":               WindowsProtectionStateProductStatus_ThreatRemediationFailedCritically,
		"threatremediationfailednoncritically":            WindowsProtectionStateProductStatus_ThreatRemediationFailedNonCritically,
		"windowssmodesignaturesinuseonnonwin10sinstall":   WindowsProtectionStateProductStatus_WindowsSModeSignaturesInUseOnNonWin10SInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsProtectionStateProductStatus(input)
	return &out, nil
}
