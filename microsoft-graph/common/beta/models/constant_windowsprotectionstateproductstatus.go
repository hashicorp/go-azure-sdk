package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsProtectionStateProductStatus string

const (
	WindowsProtectionStateProductStatusasSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "AsSignaturesOutOfDate"
	WindowsProtectionStateProductStatusavSignaturesOutOfDate                           WindowsProtectionStateProductStatus = "AvSignaturesOutOfDate"
	WindowsProtectionStateProductStatusnoFullScanHappenedForSpecifiedPeriod            WindowsProtectionStateProductStatus = "NoFullScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatusnoQuickScanHappenedForSpecifiedPeriod           WindowsProtectionStateProductStatus = "NoQuickScanHappenedForSpecifiedPeriod"
	WindowsProtectionStateProductStatusnoStatus                                        WindowsProtectionStateProductStatus = "NoStatus"
	WindowsProtectionStateProductStatusnoStatusFlagsSet                                WindowsProtectionStateProductStatus = "NoStatusFlagsSet"
	WindowsProtectionStateProductStatusofflineScanRequired                             WindowsProtectionStateProductStatus = "OfflineScanRequired"
	WindowsProtectionStateProductStatuspendingFullScanDueToThreatAction                WindowsProtectionStateProductStatus = "PendingFullScanDueToThreatAction"
	WindowsProtectionStateProductStatuspendingManualStepsDueToThreatAction             WindowsProtectionStateProductStatus = "PendingManualStepsDueToThreatAction"
	WindowsProtectionStateProductStatuspendingRebootDueToThreatAction                  WindowsProtectionStateProductStatus = "PendingRebootDueToThreatAction"
	WindowsProtectionStateProductStatusplatformAboutToBeOutdated                       WindowsProtectionStateProductStatus = "PlatformAboutToBeOutdated"
	WindowsProtectionStateProductStatusplatformOutOfDate                               WindowsProtectionStateProductStatus = "PlatformOutOfDate"
	WindowsProtectionStateProductStatusplatformUpdateInProgress                        WindowsProtectionStateProductStatus = "PlatformUpdateInProgress"
	WindowsProtectionStateProductStatusproductExpired                                  WindowsProtectionStateProductStatus = "ProductExpired"
	WindowsProtectionStateProductStatusproductRunningInEvaluationMode                  WindowsProtectionStateProductStatus = "ProductRunningInEvaluationMode"
	WindowsProtectionStateProductStatusproductRunningInNonGenuineMode                  WindowsProtectionStateProductStatus = "ProductRunningInNonGenuineMode"
	WindowsProtectionStateProductStatussamplesPendingSubmission                        WindowsProtectionStateProductStatus = "SamplesPendingSubmission"
	WindowsProtectionStateProductStatusserviceNotRunning                               WindowsProtectionStateProductStatus = "ServiceNotRunning"
	WindowsProtectionStateProductStatusserviceShutdownAsPartOfSystemShutdown           WindowsProtectionStateProductStatus = "ServiceShutdownAsPartOfSystemShutdown"
	WindowsProtectionStateProductStatusserviceStartedWithoutMalwareProtection          WindowsProtectionStateProductStatus = "ServiceStartedWithoutMalwareProtection"
	WindowsProtectionStateProductStatussignatureOrPlatformEndOfLifeIsPastOrIsImpending WindowsProtectionStateProductStatus = "SignatureOrPlatformEndOfLifeIsPastOrIsImpending"
	WindowsProtectionStateProductStatussystemInitiatedCleanInProgress                  WindowsProtectionStateProductStatus = "SystemInitiatedCleanInProgress"
	WindowsProtectionStateProductStatussystemInitiatedScanInProgress                   WindowsProtectionStateProductStatus = "SystemInitiatedScanInProgress"
	WindowsProtectionStateProductStatusthreatRemediationFailedCritically               WindowsProtectionStateProductStatus = "ThreatRemediationFailedCritically"
	WindowsProtectionStateProductStatusthreatRemediationFailedNonCritically            WindowsProtectionStateProductStatus = "ThreatRemediationFailedNonCritically"
	WindowsProtectionStateProductStatuswindowsSModeSignaturesInUseOnNonWin10SInstall   WindowsProtectionStateProductStatus = "WindowsSModeSignaturesInUseOnNonWin10SInstall"
)

func PossibleValuesForWindowsProtectionStateProductStatus() []string {
	return []string{
		string(WindowsProtectionStateProductStatusasSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatusavSignaturesOutOfDate),
		string(WindowsProtectionStateProductStatusnoFullScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatusnoQuickScanHappenedForSpecifiedPeriod),
		string(WindowsProtectionStateProductStatusnoStatus),
		string(WindowsProtectionStateProductStatusnoStatusFlagsSet),
		string(WindowsProtectionStateProductStatusofflineScanRequired),
		string(WindowsProtectionStateProductStatuspendingFullScanDueToThreatAction),
		string(WindowsProtectionStateProductStatuspendingManualStepsDueToThreatAction),
		string(WindowsProtectionStateProductStatuspendingRebootDueToThreatAction),
		string(WindowsProtectionStateProductStatusplatformAboutToBeOutdated),
		string(WindowsProtectionStateProductStatusplatformOutOfDate),
		string(WindowsProtectionStateProductStatusplatformUpdateInProgress),
		string(WindowsProtectionStateProductStatusproductExpired),
		string(WindowsProtectionStateProductStatusproductRunningInEvaluationMode),
		string(WindowsProtectionStateProductStatusproductRunningInNonGenuineMode),
		string(WindowsProtectionStateProductStatussamplesPendingSubmission),
		string(WindowsProtectionStateProductStatusserviceNotRunning),
		string(WindowsProtectionStateProductStatusserviceShutdownAsPartOfSystemShutdown),
		string(WindowsProtectionStateProductStatusserviceStartedWithoutMalwareProtection),
		string(WindowsProtectionStateProductStatussignatureOrPlatformEndOfLifeIsPastOrIsImpending),
		string(WindowsProtectionStateProductStatussystemInitiatedCleanInProgress),
		string(WindowsProtectionStateProductStatussystemInitiatedScanInProgress),
		string(WindowsProtectionStateProductStatusthreatRemediationFailedCritically),
		string(WindowsProtectionStateProductStatusthreatRemediationFailedNonCritically),
		string(WindowsProtectionStateProductStatuswindowsSModeSignaturesInUseOnNonWin10SInstall),
	}
}

func parseWindowsProtectionStateProductStatus(input string) (*WindowsProtectionStateProductStatus, error) {
	vals := map[string]WindowsProtectionStateProductStatus{
		"assignaturesoutofdate":                           WindowsProtectionStateProductStatusasSignaturesOutOfDate,
		"avsignaturesoutofdate":                           WindowsProtectionStateProductStatusavSignaturesOutOfDate,
		"nofullscanhappenedforspecifiedperiod":            WindowsProtectionStateProductStatusnoFullScanHappenedForSpecifiedPeriod,
		"noquickscanhappenedforspecifiedperiod":           WindowsProtectionStateProductStatusnoQuickScanHappenedForSpecifiedPeriod,
		"nostatus":                                        WindowsProtectionStateProductStatusnoStatus,
		"nostatusflagsset":                                WindowsProtectionStateProductStatusnoStatusFlagsSet,
		"offlinescanrequired":                             WindowsProtectionStateProductStatusofflineScanRequired,
		"pendingfullscanduetothreataction":                WindowsProtectionStateProductStatuspendingFullScanDueToThreatAction,
		"pendingmanualstepsduetothreataction":             WindowsProtectionStateProductStatuspendingManualStepsDueToThreatAction,
		"pendingrebootduetothreataction":                  WindowsProtectionStateProductStatuspendingRebootDueToThreatAction,
		"platformabouttobeoutdated":                       WindowsProtectionStateProductStatusplatformAboutToBeOutdated,
		"platformoutofdate":                               WindowsProtectionStateProductStatusplatformOutOfDate,
		"platformupdateinprogress":                        WindowsProtectionStateProductStatusplatformUpdateInProgress,
		"productexpired":                                  WindowsProtectionStateProductStatusproductExpired,
		"productrunninginevaluationmode":                  WindowsProtectionStateProductStatusproductRunningInEvaluationMode,
		"productrunninginnongenuinemode":                  WindowsProtectionStateProductStatusproductRunningInNonGenuineMode,
		"samplespendingsubmission":                        WindowsProtectionStateProductStatussamplesPendingSubmission,
		"servicenotrunning":                               WindowsProtectionStateProductStatusserviceNotRunning,
		"serviceshutdownaspartofsystemshutdown":           WindowsProtectionStateProductStatusserviceShutdownAsPartOfSystemShutdown,
		"servicestartedwithoutmalwareprotection":          WindowsProtectionStateProductStatusserviceStartedWithoutMalwareProtection,
		"signatureorplatformendoflifeispastorisimpending": WindowsProtectionStateProductStatussignatureOrPlatformEndOfLifeIsPastOrIsImpending,
		"systeminitiatedcleaninprogress":                  WindowsProtectionStateProductStatussystemInitiatedCleanInProgress,
		"systeminitiatedscaninprogress":                   WindowsProtectionStateProductStatussystemInitiatedScanInProgress,
		"threatremediationfailedcritically":               WindowsProtectionStateProductStatusthreatRemediationFailedCritically,
		"threatremediationfailednoncritically":            WindowsProtectionStateProductStatusthreatRemediationFailedNonCritically,
		"windowssmodesignaturesinuseonnonwin10sinstall":   WindowsProtectionStateProductStatuswindowsSModeSignaturesInUseOnNonWin10SInstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsProtectionStateProductStatus(input)
	return &out, nil
}
