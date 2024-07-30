package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppInstallStatusInstallStateDetail string

const (
	MobileAppInstallStatusInstallStateDetail_AppRemovedBySupersedence            MobileAppInstallStatusInstallStateDetail = "appRemovedBySupersedence"
	MobileAppInstallStatusInstallStateDetail_AutoInstallDisabled                 MobileAppInstallStatusInstallStateDetail = "autoInstallDisabled"
	MobileAppInstallStatusInstallStateDetail_ContentDownloaded                   MobileAppInstallStatusInstallStateDetail = "contentDownloaded"
	MobileAppInstallStatusInstallStateDetail_DependencyFailedToInstall           MobileAppInstallStatusInstallStateDetail = "dependencyFailedToInstall"
	MobileAppInstallStatusInstallStateDetail_DependencyPendingReboot             MobileAppInstallStatusInstallStateDetail = "dependencyPendingReboot"
	MobileAppInstallStatusInstallStateDetail_DependencyWithAutoInstallDisabled   MobileAppInstallStatusInstallStateDetail = "dependencyWithAutoInstallDisabled"
	MobileAppInstallStatusInstallStateDetail_DependencyWithRequirementsNotMet    MobileAppInstallStatusInstallStateDetail = "dependencyWithRequirementsNotMet"
	MobileAppInstallStatusInstallStateDetail_FileSystemRequirementNotMet         MobileAppInstallStatusInstallStateDetail = "fileSystemRequirementNotMet"
	MobileAppInstallStatusInstallStateDetail_InstallingDependencies              MobileAppInstallStatusInstallStateDetail = "installingDependencies"
	MobileAppInstallStatusInstallStateDetail_IosAppStoreUpdateFailedToInstall    MobileAppInstallStatusInstallStateDetail = "iosAppStoreUpdateFailedToInstall"
	MobileAppInstallStatusInstallStateDetail_ManagedAppNoLongerPresent           MobileAppInstallStatusInstallStateDetail = "managedAppNoLongerPresent"
	MobileAppInstallStatusInstallStateDetail_MinimumCpuSpeedNotMet               MobileAppInstallStatusInstallStateDetail = "minimumCpuSpeedNotMet"
	MobileAppInstallStatusInstallStateDetail_MinimumDiskSpaceNotMet              MobileAppInstallStatusInstallStateDetail = "minimumDiskSpaceNotMet"
	MobileAppInstallStatusInstallStateDetail_MinimumLogicalProcessorCountNotMet  MobileAppInstallStatusInstallStateDetail = "minimumLogicalProcessorCountNotMet"
	MobileAppInstallStatusInstallStateDetail_MinimumOsVersionNotMet              MobileAppInstallStatusInstallStateDetail = "minimumOsVersionNotMet"
	MobileAppInstallStatusInstallStateDetail_MinimumPhysicalMemoryNotMet         MobileAppInstallStatusInstallStateDetail = "minimumPhysicalMemoryNotMet"
	MobileAppInstallStatusInstallStateDetail_NoAdditionalDetails                 MobileAppInstallStatusInstallStateDetail = "noAdditionalDetails"
	MobileAppInstallStatusInstallStateDetail_PendingReboot                       MobileAppInstallStatusInstallStateDetail = "pendingReboot"
	MobileAppInstallStatusInstallStateDetail_PlatformNotApplicable               MobileAppInstallStatusInstallStateDetail = "platformNotApplicable"
	MobileAppInstallStatusInstallStateDetail_PowerShellScriptRequirementNotMet   MobileAppInstallStatusInstallStateDetail = "powerShellScriptRequirementNotMet"
	MobileAppInstallStatusInstallStateDetail_ProcessorArchitectureNotApplicable  MobileAppInstallStatusInstallStateDetail = "processorArchitectureNotApplicable"
	MobileAppInstallStatusInstallStateDetail_RegistryRequirementNotMet           MobileAppInstallStatusInstallStateDetail = "registryRequirementNotMet"
	MobileAppInstallStatusInstallStateDetail_RemovingSupersededApps              MobileAppInstallStatusInstallStateDetail = "removingSupersededApps"
	MobileAppInstallStatusInstallStateDetail_SeeInstallErrorCode                 MobileAppInstallStatusInstallStateDetail = "seeInstallErrorCode"
	MobileAppInstallStatusInstallStateDetail_SeeUninstallErrorCode               MobileAppInstallStatusInstallStateDetail = "seeUninstallErrorCode"
	MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallFailed        MobileAppInstallStatusInstallStateDetail = "supersededAppUninstallFailed"
	MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallPendingReboot MobileAppInstallStatusInstallStateDetail = "supersededAppUninstallPendingReboot"
	MobileAppInstallStatusInstallStateDetail_SupersededAppsDetected              MobileAppInstallStatusInstallStateDetail = "supersededAppsDetected"
	MobileAppInstallStatusInstallStateDetail_SupersedingAppsDetected             MobileAppInstallStatusInstallStateDetail = "supersedingAppsDetected"
	MobileAppInstallStatusInstallStateDetail_SupersedingAppsNotApplicable        MobileAppInstallStatusInstallStateDetail = "supersedingAppsNotApplicable"
	MobileAppInstallStatusInstallStateDetail_UninstallPendingReboot              MobileAppInstallStatusInstallStateDetail = "uninstallPendingReboot"
	MobileAppInstallStatusInstallStateDetail_UntargetedSupersedingAppsDetected   MobileAppInstallStatusInstallStateDetail = "untargetedSupersedingAppsDetected"
	MobileAppInstallStatusInstallStateDetail_UserIsNotLoggedIntoAppStore         MobileAppInstallStatusInstallStateDetail = "userIsNotLoggedIntoAppStore"
	MobileAppInstallStatusInstallStateDetail_UserRejectedInstall                 MobileAppInstallStatusInstallStateDetail = "userRejectedInstall"
	MobileAppInstallStatusInstallStateDetail_UserRejectedUpdate                  MobileAppInstallStatusInstallStateDetail = "userRejectedUpdate"
	MobileAppInstallStatusInstallStateDetail_VppAppHasUpdateAvailable            MobileAppInstallStatusInstallStateDetail = "vppAppHasUpdateAvailable"
)

func PossibleValuesForMobileAppInstallStatusInstallStateDetail() []string {
	return []string{
		string(MobileAppInstallStatusInstallStateDetail_AppRemovedBySupersedence),
		string(MobileAppInstallStatusInstallStateDetail_AutoInstallDisabled),
		string(MobileAppInstallStatusInstallStateDetail_ContentDownloaded),
		string(MobileAppInstallStatusInstallStateDetail_DependencyFailedToInstall),
		string(MobileAppInstallStatusInstallStateDetail_DependencyPendingReboot),
		string(MobileAppInstallStatusInstallStateDetail_DependencyWithAutoInstallDisabled),
		string(MobileAppInstallStatusInstallStateDetail_DependencyWithRequirementsNotMet),
		string(MobileAppInstallStatusInstallStateDetail_FileSystemRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetail_InstallingDependencies),
		string(MobileAppInstallStatusInstallStateDetail_IosAppStoreUpdateFailedToInstall),
		string(MobileAppInstallStatusInstallStateDetail_ManagedAppNoLongerPresent),
		string(MobileAppInstallStatusInstallStateDetail_MinimumCpuSpeedNotMet),
		string(MobileAppInstallStatusInstallStateDetail_MinimumDiskSpaceNotMet),
		string(MobileAppInstallStatusInstallStateDetail_MinimumLogicalProcessorCountNotMet),
		string(MobileAppInstallStatusInstallStateDetail_MinimumOsVersionNotMet),
		string(MobileAppInstallStatusInstallStateDetail_MinimumPhysicalMemoryNotMet),
		string(MobileAppInstallStatusInstallStateDetail_NoAdditionalDetails),
		string(MobileAppInstallStatusInstallStateDetail_PendingReboot),
		string(MobileAppInstallStatusInstallStateDetail_PlatformNotApplicable),
		string(MobileAppInstallStatusInstallStateDetail_PowerShellScriptRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetail_ProcessorArchitectureNotApplicable),
		string(MobileAppInstallStatusInstallStateDetail_RegistryRequirementNotMet),
		string(MobileAppInstallStatusInstallStateDetail_RemovingSupersededApps),
		string(MobileAppInstallStatusInstallStateDetail_SeeInstallErrorCode),
		string(MobileAppInstallStatusInstallStateDetail_SeeUninstallErrorCode),
		string(MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallFailed),
		string(MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallPendingReboot),
		string(MobileAppInstallStatusInstallStateDetail_SupersededAppsDetected),
		string(MobileAppInstallStatusInstallStateDetail_SupersedingAppsDetected),
		string(MobileAppInstallStatusInstallStateDetail_SupersedingAppsNotApplicable),
		string(MobileAppInstallStatusInstallStateDetail_UninstallPendingReboot),
		string(MobileAppInstallStatusInstallStateDetail_UntargetedSupersedingAppsDetected),
		string(MobileAppInstallStatusInstallStateDetail_UserIsNotLoggedIntoAppStore),
		string(MobileAppInstallStatusInstallStateDetail_UserRejectedInstall),
		string(MobileAppInstallStatusInstallStateDetail_UserRejectedUpdate),
		string(MobileAppInstallStatusInstallStateDetail_VppAppHasUpdateAvailable),
	}
}

func (s *MobileAppInstallStatusInstallStateDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppInstallStatusInstallStateDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppInstallStatusInstallStateDetail(input string) (*MobileAppInstallStatusInstallStateDetail, error) {
	vals := map[string]MobileAppInstallStatusInstallStateDetail{
		"appremovedbysupersedence":            MobileAppInstallStatusInstallStateDetail_AppRemovedBySupersedence,
		"autoinstalldisabled":                 MobileAppInstallStatusInstallStateDetail_AutoInstallDisabled,
		"contentdownloaded":                   MobileAppInstallStatusInstallStateDetail_ContentDownloaded,
		"dependencyfailedtoinstall":           MobileAppInstallStatusInstallStateDetail_DependencyFailedToInstall,
		"dependencypendingreboot":             MobileAppInstallStatusInstallStateDetail_DependencyPendingReboot,
		"dependencywithautoinstalldisabled":   MobileAppInstallStatusInstallStateDetail_DependencyWithAutoInstallDisabled,
		"dependencywithrequirementsnotmet":    MobileAppInstallStatusInstallStateDetail_DependencyWithRequirementsNotMet,
		"filesystemrequirementnotmet":         MobileAppInstallStatusInstallStateDetail_FileSystemRequirementNotMet,
		"installingdependencies":              MobileAppInstallStatusInstallStateDetail_InstallingDependencies,
		"iosappstoreupdatefailedtoinstall":    MobileAppInstallStatusInstallStateDetail_IosAppStoreUpdateFailedToInstall,
		"managedappnolongerpresent":           MobileAppInstallStatusInstallStateDetail_ManagedAppNoLongerPresent,
		"minimumcpuspeednotmet":               MobileAppInstallStatusInstallStateDetail_MinimumCpuSpeedNotMet,
		"minimumdiskspacenotmet":              MobileAppInstallStatusInstallStateDetail_MinimumDiskSpaceNotMet,
		"minimumlogicalprocessorcountnotmet":  MobileAppInstallStatusInstallStateDetail_MinimumLogicalProcessorCountNotMet,
		"minimumosversionnotmet":              MobileAppInstallStatusInstallStateDetail_MinimumOsVersionNotMet,
		"minimumphysicalmemorynotmet":         MobileAppInstallStatusInstallStateDetail_MinimumPhysicalMemoryNotMet,
		"noadditionaldetails":                 MobileAppInstallStatusInstallStateDetail_NoAdditionalDetails,
		"pendingreboot":                       MobileAppInstallStatusInstallStateDetail_PendingReboot,
		"platformnotapplicable":               MobileAppInstallStatusInstallStateDetail_PlatformNotApplicable,
		"powershellscriptrequirementnotmet":   MobileAppInstallStatusInstallStateDetail_PowerShellScriptRequirementNotMet,
		"processorarchitecturenotapplicable":  MobileAppInstallStatusInstallStateDetail_ProcessorArchitectureNotApplicable,
		"registryrequirementnotmet":           MobileAppInstallStatusInstallStateDetail_RegistryRequirementNotMet,
		"removingsupersededapps":              MobileAppInstallStatusInstallStateDetail_RemovingSupersededApps,
		"seeinstallerrorcode":                 MobileAppInstallStatusInstallStateDetail_SeeInstallErrorCode,
		"seeuninstallerrorcode":               MobileAppInstallStatusInstallStateDetail_SeeUninstallErrorCode,
		"supersededappuninstallfailed":        MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallFailed,
		"supersededappuninstallpendingreboot": MobileAppInstallStatusInstallStateDetail_SupersededAppUninstallPendingReboot,
		"supersededappsdetected":              MobileAppInstallStatusInstallStateDetail_SupersededAppsDetected,
		"supersedingappsdetected":             MobileAppInstallStatusInstallStateDetail_SupersedingAppsDetected,
		"supersedingappsnotapplicable":        MobileAppInstallStatusInstallStateDetail_SupersedingAppsNotApplicable,
		"uninstallpendingreboot":              MobileAppInstallStatusInstallStateDetail_UninstallPendingReboot,
		"untargetedsupersedingappsdetected":   MobileAppInstallStatusInstallStateDetail_UntargetedSupersedingAppsDetected,
		"userisnotloggedintoappstore":         MobileAppInstallStatusInstallStateDetail_UserIsNotLoggedIntoAppStore,
		"userrejectedinstall":                 MobileAppInstallStatusInstallStateDetail_UserRejectedInstall,
		"userrejectedupdate":                  MobileAppInstallStatusInstallStateDetail_UserRejectedUpdate,
		"vppapphasupdateavailable":            MobileAppInstallStatusInstallStateDetail_VppAppHasUpdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppInstallStatusInstallStateDetail(input)
	return &out, nil
}
