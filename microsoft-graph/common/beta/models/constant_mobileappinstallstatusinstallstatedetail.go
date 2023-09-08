package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusInstallStateDetail string

const (
	MobileAppInstallStatusInstallStateDetailappRemovedBySupersedence            MobileAppInstallStatusInstallStateDetail = "AppRemovedBySupersedence"
	MobileAppInstallStatusInstallStateDetailautoInstallDisabled                 MobileAppInstallStatusInstallStateDetail = "AutoInstallDisabled"
	MobileAppInstallStatusInstallStateDetailcontentDownloaded                   MobileAppInstallStatusInstallStateDetail = "ContentDownloaded"
	MobileAppInstallStatusInstallStateDetaildependencyFailedToInstall           MobileAppInstallStatusInstallStateDetail = "DependencyFailedToInstall"
	MobileAppInstallStatusInstallStateDetaildependencyPendingReboot             MobileAppInstallStatusInstallStateDetail = "DependencyPendingReboot"
	MobileAppInstallStatusInstallStateDetaildependencyWithAutoInstallDisabled   MobileAppInstallStatusInstallStateDetail = "DependencyWithAutoInstallDisabled"
	MobileAppInstallStatusInstallStateDetaildependencyWithRequirementsNotMet    MobileAppInstallStatusInstallStateDetail = "DependencyWithRequirementsNotMet"
	MobileAppInstallStatusInstallStateDetailfileSystemRequirementNotMet         MobileAppInstallStatusInstallStateDetail = "FileSystemRequirementNotMet"
	MobileAppInstallStatusInstallStateDetailinstallingDependencies              MobileAppInstallStatusInstallStateDetail = "InstallingDependencies"
	MobileAppInstallStatusInstallStateDetailiosAppStoreUpdateFailedToInstall    MobileAppInstallStatusInstallStateDetail = "IosAppStoreUpdateFailedToInstall"
	MobileAppInstallStatusInstallStateDetailmanagedAppNoLongerPresent           MobileAppInstallStatusInstallStateDetail = "ManagedAppNoLongerPresent"
	MobileAppInstallStatusInstallStateDetailminimumCpuSpeedNotMet               MobileAppInstallStatusInstallStateDetail = "MinimumCpuSpeedNotMet"
	MobileAppInstallStatusInstallStateDetailminimumDiskSpaceNotMet              MobileAppInstallStatusInstallStateDetail = "MinimumDiskSpaceNotMet"
	MobileAppInstallStatusInstallStateDetailminimumLogicalProcessorCountNotMet  MobileAppInstallStatusInstallStateDetail = "MinimumLogicalProcessorCountNotMet"
	MobileAppInstallStatusInstallStateDetailminimumOsVersionNotMet              MobileAppInstallStatusInstallStateDetail = "MinimumOsVersionNotMet"
	MobileAppInstallStatusInstallStateDetailminimumPhysicalMemoryNotMet         MobileAppInstallStatusInstallStateDetail = "MinimumPhysicalMemoryNotMet"
	MobileAppInstallStatusInstallStateDetailnoAdditionalDetails                 MobileAppInstallStatusInstallStateDetail = "NoAdditionalDetails"
	MobileAppInstallStatusInstallStateDetailpendingReboot                       MobileAppInstallStatusInstallStateDetail = "PendingReboot"
	MobileAppInstallStatusInstallStateDetailplatformNotApplicable               MobileAppInstallStatusInstallStateDetail = "PlatformNotApplicable"
	MobileAppInstallStatusInstallStateDetailpowerShellScriptRequirementNotMet   MobileAppInstallStatusInstallStateDetail = "PowerShellScriptRequirementNotMet"
	MobileAppInstallStatusInstallStateDetailprocessorArchitectureNotApplicable  MobileAppInstallStatusInstallStateDetail = "ProcessorArchitectureNotApplicable"
	MobileAppInstallStatusInstallStateDetailregistryRequirementNotMet           MobileAppInstallStatusInstallStateDetail = "RegistryRequirementNotMet"
	MobileAppInstallStatusInstallStateDetailremovingSupersededApps              MobileAppInstallStatusInstallStateDetail = "RemovingSupersededApps"
	MobileAppInstallStatusInstallStateDetailseeInstallErrorCode                 MobileAppInstallStatusInstallStateDetail = "SeeInstallErrorCode"
	MobileAppInstallStatusInstallStateDetailseeUninstallErrorCode               MobileAppInstallStatusInstallStateDetail = "SeeUninstallErrorCode"
	MobileAppInstallStatusInstallStateDetailsupersededAppUninstallFailed        MobileAppInstallStatusInstallStateDetail = "SupersededAppUninstallFailed"
	MobileAppInstallStatusInstallStateDetailsupersededAppUninstallPendingReboot MobileAppInstallStatusInstallStateDetail = "SupersededAppUninstallPendingReboot"
	MobileAppInstallStatusInstallStateDetailsupersededAppsDetected              MobileAppInstallStatusInstallStateDetail = "SupersededAppsDetected"
	MobileAppInstallStatusInstallStateDetailsupersedingAppsDetected             MobileAppInstallStatusInstallStateDetail = "SupersedingAppsDetected"
	MobileAppInstallStatusInstallStateDetailsupersedingAppsNotApplicable        MobileAppInstallStatusInstallStateDetail = "SupersedingAppsNotApplicable"
	MobileAppInstallStatusInstallStateDetailuninstallPendingReboot              MobileAppInstallStatusInstallStateDetail = "UninstallPendingReboot"
	MobileAppInstallStatusInstallStateDetailuntargetedSupersedingAppsDetected   MobileAppInstallStatusInstallStateDetail = "UntargetedSupersedingAppsDetected"
	MobileAppInstallStatusInstallStateDetailuserIsNotLoggedIntoAppStore         MobileAppInstallStatusInstallStateDetail = "UserIsNotLoggedIntoAppStore"
	MobileAppInstallStatusInstallStateDetailuserRejectedInstall                 MobileAppInstallStatusInstallStateDetail = "UserRejectedInstall"
	MobileAppInstallStatusInstallStateDetailuserRejectedUpdate                  MobileAppInstallStatusInstallStateDetail = "UserRejectedUpdate"
	MobileAppInstallStatusInstallStateDetailvppAppHasUpdateAvailable            MobileAppInstallStatusInstallStateDetail = "VppAppHasUpdateAvailable"
)

func PossibleValuesForMobileAppInstallStatusInstallStateDetail() []string {
	return []string{
		string(MobileAppInstallStatusInstallStateDetailappRemovedBySupersedence),
		string(MobileAppInstallStatusInstallStateDetailautoInstallDisabled),
		string(MobileAppInstallStatusInstallStateDetailcontentDownloaded),
		string(MobileAppInstallStatusInstallStateDetaildependencyFailedToInstall),
		string(MobileAppInstallStatusInstallStateDetaildependencyPendingReboot),
		string(MobileAppInstallStatusInstallStateDetaildependencyWithAutoInstallDisabled),
		string(MobileAppInstallStatusInstallStateDetaildependencyWithRequirementsNotMet),
		string(MobileAppInstallStatusInstallStateDetailfileSystemRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetailinstallingDependencies),
		string(MobileAppInstallStatusInstallStateDetailiosAppStoreUpdateFailedToInstall),
		string(MobileAppInstallStatusInstallStateDetailmanagedAppNoLongerPresent),
		string(MobileAppInstallStatusInstallStateDetailminimumCpuSpeedNotMet),
		string(MobileAppInstallStatusInstallStateDetailminimumDiskSpaceNotMet),
		string(MobileAppInstallStatusInstallStateDetailminimumLogicalProcessorCountNotMet),
		string(MobileAppInstallStatusInstallStateDetailminimumOsVersionNotMet),
		string(MobileAppInstallStatusInstallStateDetailminimumPhysicalMemoryNotMet),
		string(MobileAppInstallStatusInstallStateDetailnoAdditionalDetails),
		string(MobileAppInstallStatusInstallStateDetailpendingReboot),
		string(MobileAppInstallStatusInstallStateDetailplatformNotApplicable),
		string(MobileAppInstallStatusInstallStateDetailpowerShellScriptRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetailprocessorArchitectureNotApplicable),
		string(MobileAppInstallStatusInstallStateDetailregistryRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetailremovingSupersededApps),
		string(MobileAppInstallStatusInstallStateDetailseeInstallErrorCode),
		string(MobileAppInstallStatusInstallStateDetailseeUninstallErrorCode),
		string(MobileAppInstallStatusInstallStateDetailsupersededAppUninstallFailed),
		string(MobileAppInstallStatusInstallStateDetailsupersededAppUninstallPendingReboot),
		string(MobileAppInstallStatusInstallStateDetailsupersededAppsDetected),
		string(MobileAppInstallStatusInstallStateDetailsupersedingAppsDetected),
		string(MobileAppInstallStatusInstallStateDetailsupersedingAppsNotApplicable),
		string(MobileAppInstallStatusInstallStateDetailuninstallPendingReboot),
		string(MobileAppInstallStatusInstallStateDetailuntargetedSupersedingAppsDetected),
		string(MobileAppInstallStatusInstallStateDetailuserIsNotLoggedIntoAppStore),
		string(MobileAppInstallStatusInstallStateDetailuserRejectedInstall),
		string(MobileAppInstallStatusInstallStateDetailuserRejectedUpdate),
		string(MobileAppInstallStatusInstallStateDetailvppAppHasUpdateAvailable),
	}
}

func parseMobileAppInstallStatusInstallStateDetail(input string) (*MobileAppInstallStatusInstallStateDetail, error) {
	vals := map[string]MobileAppInstallStatusInstallStateDetail{
		"appremovedbysupersedence":            MobileAppInstallStatusInstallStateDetailappRemovedBySupersedence,
		"autoinstalldisabled":                 MobileAppInstallStatusInstallStateDetailautoInstallDisabled,
		"contentdownloaded":                   MobileAppInstallStatusInstallStateDetailcontentDownloaded,
		"dependencyfailedtoinstall":           MobileAppInstallStatusInstallStateDetaildependencyFailedToInstall,
		"dependencypendingreboot":             MobileAppInstallStatusInstallStateDetaildependencyPendingReboot,
		"dependencywithautoinstalldisabled":   MobileAppInstallStatusInstallStateDetaildependencyWithAutoInstallDisabled,
		"dependencywithrequirementsnotmet":    MobileAppInstallStatusInstallStateDetaildependencyWithRequirementsNotMet,
		"filesystemrequirementnotmet":         MobileAppInstallStatusInstallStateDetailfileSystemRequirementNotMet,
		"installingdependencies":              MobileAppInstallStatusInstallStateDetailinstallingDependencies,
		"iosappstoreupdatefailedtoinstall":    MobileAppInstallStatusInstallStateDetailiosAppStoreUpdateFailedToInstall,
		"managedappnolongerpresent":           MobileAppInstallStatusInstallStateDetailmanagedAppNoLongerPresent,
		"minimumcpuspeednotmet":               MobileAppInstallStatusInstallStateDetailminimumCpuSpeedNotMet,
		"minimumdiskspacenotmet":              MobileAppInstallStatusInstallStateDetailminimumDiskSpaceNotMet,
		"minimumlogicalprocessorcountnotmet":  MobileAppInstallStatusInstallStateDetailminimumLogicalProcessorCountNotMet,
		"minimumosversionnotmet":              MobileAppInstallStatusInstallStateDetailminimumOsVersionNotMet,
		"minimumphysicalmemorynotmet":         MobileAppInstallStatusInstallStateDetailminimumPhysicalMemoryNotMet,
		"noadditionaldetails":                 MobileAppInstallStatusInstallStateDetailnoAdditionalDetails,
		"pendingreboot":                       MobileAppInstallStatusInstallStateDetailpendingReboot,
		"platformnotapplicable":               MobileAppInstallStatusInstallStateDetailplatformNotApplicable,
		"powershellscriptrequirementnotmet":   MobileAppInstallStatusInstallStateDetailpowerShellScriptRequirementNotMet,
		"processorarchitecturenotapplicable":  MobileAppInstallStatusInstallStateDetailprocessorArchitectureNotApplicable,
		"registryrequirementnotmet":           MobileAppInstallStatusInstallStateDetailregistryRequirementNotMet,
		"removingsupersededapps":              MobileAppInstallStatusInstallStateDetailremovingSupersededApps,
		"seeinstallerrorcode":                 MobileAppInstallStatusInstallStateDetailseeInstallErrorCode,
		"seeuninstallerrorcode":               MobileAppInstallStatusInstallStateDetailseeUninstallErrorCode,
		"supersededappuninstallfailed":        MobileAppInstallStatusInstallStateDetailsupersededAppUninstallFailed,
		"supersededappuninstallpendingreboot": MobileAppInstallStatusInstallStateDetailsupersededAppUninstallPendingReboot,
		"supersededappsdetected":              MobileAppInstallStatusInstallStateDetailsupersededAppsDetected,
		"supersedingappsdetected":             MobileAppInstallStatusInstallStateDetailsupersedingAppsDetected,
		"supersedingappsnotapplicable":        MobileAppInstallStatusInstallStateDetailsupersedingAppsNotApplicable,
		"uninstallpendingreboot":              MobileAppInstallStatusInstallStateDetailuninstallPendingReboot,
		"untargetedsupersedingappsdetected":   MobileAppInstallStatusInstallStateDetailuntargetedSupersedingAppsDetected,
		"userisnotloggedintoappstore":         MobileAppInstallStatusInstallStateDetailuserIsNotLoggedIntoAppStore,
		"userrejectedinstall":                 MobileAppInstallStatusInstallStateDetailuserRejectedInstall,
		"userrejectedupdate":                  MobileAppInstallStatusInstallStateDetailuserRejectedUpdate,
		"vppapphasupdateavailable":            MobileAppInstallStatusInstallStateDetailvppAppHasUpdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusInstallStateDetail(input)
	return &out, nil
}
