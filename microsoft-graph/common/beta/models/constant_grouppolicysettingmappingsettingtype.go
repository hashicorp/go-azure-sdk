package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingMappingSettingType string

const (
	GroupPolicySettingMappingSettingTypeaccount                 GroupPolicySettingMappingSettingType = "Account"
	GroupPolicySettingMappingSettingTypeappLockerRuleCollection GroupPolicySettingMappingSettingType = "AppLockerRuleCollection"
	GroupPolicySettingMappingSettingTypeauditSetting            GroupPolicySettingMappingSettingType = "AuditSetting"
	GroupPolicySettingMappingSettingTypedataSourcesSettings     GroupPolicySettingMappingSettingType = "DataSourcesSettings"
	GroupPolicySettingMappingSettingTypedevicesSettings         GroupPolicySettingMappingSettingType = "DevicesSettings"
	GroupPolicySettingMappingSettingTypedriveMapSettings        GroupPolicySettingMappingSettingType = "DriveMapSettings"
	GroupPolicySettingMappingSettingTypeenvironmentVariables    GroupPolicySettingMappingSettingType = "EnvironmentVariables"
	GroupPolicySettingMappingSettingTypefilesSettings           GroupPolicySettingMappingSettingType = "FilesSettings"
	GroupPolicySettingMappingSettingTypefolderOptions           GroupPolicySettingMappingSettingType = "FolderOptions"
	GroupPolicySettingMappingSettingTypefolders                 GroupPolicySettingMappingSettingType = "Folders"
	GroupPolicySettingMappingSettingTypeiniFiles                GroupPolicySettingMappingSettingType = "IniFiles"
	GroupPolicySettingMappingSettingTypeinternetOptions         GroupPolicySettingMappingSettingType = "InternetOptions"
	GroupPolicySettingMappingSettingTypelocalUsersAndGroups     GroupPolicySettingMappingSettingType = "LocalUsersAndGroups"
	GroupPolicySettingMappingSettingTypenetworkOptions          GroupPolicySettingMappingSettingType = "NetworkOptions"
	GroupPolicySettingMappingSettingTypenetworkShares           GroupPolicySettingMappingSettingType = "NetworkShares"
	GroupPolicySettingMappingSettingTypentServices              GroupPolicySettingMappingSettingType = "NtServices"
	GroupPolicySettingMappingSettingTypepolicy                  GroupPolicySettingMappingSettingType = "Policy"
	GroupPolicySettingMappingSettingTypepowerOptions            GroupPolicySettingMappingSettingType = "PowerOptions"
	GroupPolicySettingMappingSettingTypeprinters                GroupPolicySettingMappingSettingType = "Printers"
	GroupPolicySettingMappingSettingTyperegionalOptionsSettings GroupPolicySettingMappingSettingType = "RegionalOptionsSettings"
	GroupPolicySettingMappingSettingTyperegistrySettings        GroupPolicySettingMappingSettingType = "RegistrySettings"
	GroupPolicySettingMappingSettingTypescheduledTasks          GroupPolicySettingMappingSettingType = "ScheduledTasks"
	GroupPolicySettingMappingSettingTypesecurityOptions         GroupPolicySettingMappingSettingType = "SecurityOptions"
	GroupPolicySettingMappingSettingTypeshortcutSettings        GroupPolicySettingMappingSettingType = "ShortcutSettings"
	GroupPolicySettingMappingSettingTypestartMenuSettings       GroupPolicySettingMappingSettingType = "StartMenuSettings"
	GroupPolicySettingMappingSettingTypeunknown                 GroupPolicySettingMappingSettingType = "Unknown"
	GroupPolicySettingMappingSettingTypeuserRightsAssignment    GroupPolicySettingMappingSettingType = "UserRightsAssignment"
	GroupPolicySettingMappingSettingTypewindowsFirewallSettings GroupPolicySettingMappingSettingType = "WindowsFirewallSettings"
)

func PossibleValuesForGroupPolicySettingMappingSettingType() []string {
	return []string{
		string(GroupPolicySettingMappingSettingTypeaccount),
		string(GroupPolicySettingMappingSettingTypeappLockerRuleCollection),
		string(GroupPolicySettingMappingSettingTypeauditSetting),
		string(GroupPolicySettingMappingSettingTypedataSourcesSettings),
		string(GroupPolicySettingMappingSettingTypedevicesSettings),
		string(GroupPolicySettingMappingSettingTypedriveMapSettings),
		string(GroupPolicySettingMappingSettingTypeenvironmentVariables),
		string(GroupPolicySettingMappingSettingTypefilesSettings),
		string(GroupPolicySettingMappingSettingTypefolderOptions),
		string(GroupPolicySettingMappingSettingTypefolders),
		string(GroupPolicySettingMappingSettingTypeiniFiles),
		string(GroupPolicySettingMappingSettingTypeinternetOptions),
		string(GroupPolicySettingMappingSettingTypelocalUsersAndGroups),
		string(GroupPolicySettingMappingSettingTypenetworkOptions),
		string(GroupPolicySettingMappingSettingTypenetworkShares),
		string(GroupPolicySettingMappingSettingTypentServices),
		string(GroupPolicySettingMappingSettingTypepolicy),
		string(GroupPolicySettingMappingSettingTypepowerOptions),
		string(GroupPolicySettingMappingSettingTypeprinters),
		string(GroupPolicySettingMappingSettingTyperegionalOptionsSettings),
		string(GroupPolicySettingMappingSettingTyperegistrySettings),
		string(GroupPolicySettingMappingSettingTypescheduledTasks),
		string(GroupPolicySettingMappingSettingTypesecurityOptions),
		string(GroupPolicySettingMappingSettingTypeshortcutSettings),
		string(GroupPolicySettingMappingSettingTypestartMenuSettings),
		string(GroupPolicySettingMappingSettingTypeunknown),
		string(GroupPolicySettingMappingSettingTypeuserRightsAssignment),
		string(GroupPolicySettingMappingSettingTypewindowsFirewallSettings),
	}
}

func parseGroupPolicySettingMappingSettingType(input string) (*GroupPolicySettingMappingSettingType, error) {
	vals := map[string]GroupPolicySettingMappingSettingType{
		"account":                 GroupPolicySettingMappingSettingTypeaccount,
		"applockerrulecollection": GroupPolicySettingMappingSettingTypeappLockerRuleCollection,
		"auditsetting":            GroupPolicySettingMappingSettingTypeauditSetting,
		"datasourcessettings":     GroupPolicySettingMappingSettingTypedataSourcesSettings,
		"devicessettings":         GroupPolicySettingMappingSettingTypedevicesSettings,
		"drivemapsettings":        GroupPolicySettingMappingSettingTypedriveMapSettings,
		"environmentvariables":    GroupPolicySettingMappingSettingTypeenvironmentVariables,
		"filessettings":           GroupPolicySettingMappingSettingTypefilesSettings,
		"folderoptions":           GroupPolicySettingMappingSettingTypefolderOptions,
		"folders":                 GroupPolicySettingMappingSettingTypefolders,
		"inifiles":                GroupPolicySettingMappingSettingTypeiniFiles,
		"internetoptions":         GroupPolicySettingMappingSettingTypeinternetOptions,
		"localusersandgroups":     GroupPolicySettingMappingSettingTypelocalUsersAndGroups,
		"networkoptions":          GroupPolicySettingMappingSettingTypenetworkOptions,
		"networkshares":           GroupPolicySettingMappingSettingTypenetworkShares,
		"ntservices":              GroupPolicySettingMappingSettingTypentServices,
		"policy":                  GroupPolicySettingMappingSettingTypepolicy,
		"poweroptions":            GroupPolicySettingMappingSettingTypepowerOptions,
		"printers":                GroupPolicySettingMappingSettingTypeprinters,
		"regionaloptionssettings": GroupPolicySettingMappingSettingTyperegionalOptionsSettings,
		"registrysettings":        GroupPolicySettingMappingSettingTyperegistrySettings,
		"scheduledtasks":          GroupPolicySettingMappingSettingTypescheduledTasks,
		"securityoptions":         GroupPolicySettingMappingSettingTypesecurityOptions,
		"shortcutsettings":        GroupPolicySettingMappingSettingTypeshortcutSettings,
		"startmenusettings":       GroupPolicySettingMappingSettingTypestartMenuSettings,
		"unknown":                 GroupPolicySettingMappingSettingTypeunknown,
		"userrightsassignment":    GroupPolicySettingMappingSettingTypeuserRightsAssignment,
		"windowsfirewallsettings": GroupPolicySettingMappingSettingTypewindowsFirewallSettings,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingMappingSettingType(input)
	return &out, nil
}
