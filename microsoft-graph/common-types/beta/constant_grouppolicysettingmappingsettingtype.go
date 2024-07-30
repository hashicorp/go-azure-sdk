package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingSettingType string

const (
	GroupPolicySettingMappingSettingType_Account                 GroupPolicySettingMappingSettingType = "account"
	GroupPolicySettingMappingSettingType_AppLockerRuleCollection GroupPolicySettingMappingSettingType = "appLockerRuleCollection"
	GroupPolicySettingMappingSettingType_AuditSetting            GroupPolicySettingMappingSettingType = "auditSetting"
	GroupPolicySettingMappingSettingType_DataSourcesSettings     GroupPolicySettingMappingSettingType = "dataSourcesSettings"
	GroupPolicySettingMappingSettingType_DevicesSettings         GroupPolicySettingMappingSettingType = "devicesSettings"
	GroupPolicySettingMappingSettingType_DriveMapSettings        GroupPolicySettingMappingSettingType = "driveMapSettings"
	GroupPolicySettingMappingSettingType_EnvironmentVariables    GroupPolicySettingMappingSettingType = "environmentVariables"
	GroupPolicySettingMappingSettingType_FilesSettings           GroupPolicySettingMappingSettingType = "filesSettings"
	GroupPolicySettingMappingSettingType_FolderOptions           GroupPolicySettingMappingSettingType = "folderOptions"
	GroupPolicySettingMappingSettingType_Folders                 GroupPolicySettingMappingSettingType = "folders"
	GroupPolicySettingMappingSettingType_IniFiles                GroupPolicySettingMappingSettingType = "iniFiles"
	GroupPolicySettingMappingSettingType_InternetOptions         GroupPolicySettingMappingSettingType = "internetOptions"
	GroupPolicySettingMappingSettingType_LocalUsersAndGroups     GroupPolicySettingMappingSettingType = "localUsersAndGroups"
	GroupPolicySettingMappingSettingType_NetworkOptions          GroupPolicySettingMappingSettingType = "networkOptions"
	GroupPolicySettingMappingSettingType_NetworkShares           GroupPolicySettingMappingSettingType = "networkShares"
	GroupPolicySettingMappingSettingType_NtServices              GroupPolicySettingMappingSettingType = "ntServices"
	GroupPolicySettingMappingSettingType_Policy                  GroupPolicySettingMappingSettingType = "policy"
	GroupPolicySettingMappingSettingType_PowerOptions            GroupPolicySettingMappingSettingType = "powerOptions"
	GroupPolicySettingMappingSettingType_Printers                GroupPolicySettingMappingSettingType = "printers"
	GroupPolicySettingMappingSettingType_RegionalOptionsSettings GroupPolicySettingMappingSettingType = "regionalOptionsSettings"
	GroupPolicySettingMappingSettingType_RegistrySettings        GroupPolicySettingMappingSettingType = "registrySettings"
	GroupPolicySettingMappingSettingType_ScheduledTasks          GroupPolicySettingMappingSettingType = "scheduledTasks"
	GroupPolicySettingMappingSettingType_SecurityOptions         GroupPolicySettingMappingSettingType = "securityOptions"
	GroupPolicySettingMappingSettingType_ShortcutSettings        GroupPolicySettingMappingSettingType = "shortcutSettings"
	GroupPolicySettingMappingSettingType_StartMenuSettings       GroupPolicySettingMappingSettingType = "startMenuSettings"
	GroupPolicySettingMappingSettingType_Unknown                 GroupPolicySettingMappingSettingType = "unknown"
	GroupPolicySettingMappingSettingType_UserRightsAssignment    GroupPolicySettingMappingSettingType = "userRightsAssignment"
	GroupPolicySettingMappingSettingType_WindowsFirewallSettings GroupPolicySettingMappingSettingType = "windowsFirewallSettings"
)

func PossibleValuesForGroupPolicySettingMappingSettingType() []string {
	return []string{
		string(GroupPolicySettingMappingSettingType_Account),
		string(GroupPolicySettingMappingSettingType_AppLockerRuleCollection),
		string(GroupPolicySettingMappingSettingType_AuditSetting),
		string(GroupPolicySettingMappingSettingType_DataSourcesSettings),
		string(GroupPolicySettingMappingSettingType_DevicesSettings),
		string(GroupPolicySettingMappingSettingType_DriveMapSettings),
		string(GroupPolicySettingMappingSettingType_EnvironmentVariables),
		string(GroupPolicySettingMappingSettingType_FilesSettings),
		string(GroupPolicySettingMappingSettingType_FolderOptions),
		string(GroupPolicySettingMappingSettingType_Folders),
		string(GroupPolicySettingMappingSettingType_IniFiles),
		string(GroupPolicySettingMappingSettingType_InternetOptions),
		string(GroupPolicySettingMappingSettingType_LocalUsersAndGroups),
		string(GroupPolicySettingMappingSettingType_NetworkOptions),
		string(GroupPolicySettingMappingSettingType_NetworkShares),
		string(GroupPolicySettingMappingSettingType_NtServices),
		string(GroupPolicySettingMappingSettingType_Policy),
		string(GroupPolicySettingMappingSettingType_PowerOptions),
		string(GroupPolicySettingMappingSettingType_Printers),
		string(GroupPolicySettingMappingSettingType_RegionalOptionsSettings),
		string(GroupPolicySettingMappingSettingType_RegistrySettings),
		string(GroupPolicySettingMappingSettingType_ScheduledTasks),
		string(GroupPolicySettingMappingSettingType_SecurityOptions),
		string(GroupPolicySettingMappingSettingType_ShortcutSettings),
		string(GroupPolicySettingMappingSettingType_StartMenuSettings),
		string(GroupPolicySettingMappingSettingType_Unknown),
		string(GroupPolicySettingMappingSettingType_UserRightsAssignment),
		string(GroupPolicySettingMappingSettingType_WindowsFirewallSettings),
	}
}

func (s *GroupPolicySettingMappingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicySettingMappingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicySettingMappingSettingType(input string) (*GroupPolicySettingMappingSettingType, error) {
	vals := map[string]GroupPolicySettingMappingSettingType{
		"account":                 GroupPolicySettingMappingSettingType_Account,
		"applockerrulecollection": GroupPolicySettingMappingSettingType_AppLockerRuleCollection,
		"auditsetting":            GroupPolicySettingMappingSettingType_AuditSetting,
		"datasourcessettings":     GroupPolicySettingMappingSettingType_DataSourcesSettings,
		"devicessettings":         GroupPolicySettingMappingSettingType_DevicesSettings,
		"drivemapsettings":        GroupPolicySettingMappingSettingType_DriveMapSettings,
		"environmentvariables":    GroupPolicySettingMappingSettingType_EnvironmentVariables,
		"filessettings":           GroupPolicySettingMappingSettingType_FilesSettings,
		"folderoptions":           GroupPolicySettingMappingSettingType_FolderOptions,
		"folders":                 GroupPolicySettingMappingSettingType_Folders,
		"inifiles":                GroupPolicySettingMappingSettingType_IniFiles,
		"internetoptions":         GroupPolicySettingMappingSettingType_InternetOptions,
		"localusersandgroups":     GroupPolicySettingMappingSettingType_LocalUsersAndGroups,
		"networkoptions":          GroupPolicySettingMappingSettingType_NetworkOptions,
		"networkshares":           GroupPolicySettingMappingSettingType_NetworkShares,
		"ntservices":              GroupPolicySettingMappingSettingType_NtServices,
		"policy":                  GroupPolicySettingMappingSettingType_Policy,
		"poweroptions":            GroupPolicySettingMappingSettingType_PowerOptions,
		"printers":                GroupPolicySettingMappingSettingType_Printers,
		"regionaloptionssettings": GroupPolicySettingMappingSettingType_RegionalOptionsSettings,
		"registrysettings":        GroupPolicySettingMappingSettingType_RegistrySettings,
		"scheduledtasks":          GroupPolicySettingMappingSettingType_ScheduledTasks,
		"securityoptions":         GroupPolicySettingMappingSettingType_SecurityOptions,
		"shortcutsettings":        GroupPolicySettingMappingSettingType_ShortcutSettings,
		"startmenusettings":       GroupPolicySettingMappingSettingType_StartMenuSettings,
		"unknown":                 GroupPolicySettingMappingSettingType_Unknown,
		"userrightsassignment":    GroupPolicySettingMappingSettingType_UserRightsAssignment,
		"windowsfirewallsettings": GroupPolicySettingMappingSettingType_WindowsFirewallSettings,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingSettingType(input)
	return &out, nil
}
