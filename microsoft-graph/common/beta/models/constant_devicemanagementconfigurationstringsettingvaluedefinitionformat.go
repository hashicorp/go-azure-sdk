package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueDefinitionFormat string

const (
	DeviceManagementConfigurationStringSettingValueDefinitionFormatbase64     DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Base64"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatbashScript DeviceManagementConfigurationStringSettingValueDefinitionFormat = "BashScript"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatbinary     DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Binary"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatdate       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Date"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatdateTime   DeviceManagementConfigurationStringSettingValueDefinitionFormat = "DateTime"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatemail      DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Email"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatguid       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Guid"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatip         DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Ip"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatjson       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Json"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatnone       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "None"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatregEx      DeviceManagementConfigurationStringSettingValueDefinitionFormat = "RegEx"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatsurfaceHub DeviceManagementConfigurationStringSettingValueDefinitionFormat = "SurfaceHub"
	DeviceManagementConfigurationStringSettingValueDefinitionFormattime       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Time"
	DeviceManagementConfigurationStringSettingValueDefinitionFormaturl        DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Url"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatversion    DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Version"
	DeviceManagementConfigurationStringSettingValueDefinitionFormatxml        DeviceManagementConfigurationStringSettingValueDefinitionFormat = "Xml"
)

func PossibleValuesForDeviceManagementConfigurationStringSettingValueDefinitionFormat() []string {
	return []string{
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatbase64),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatbashScript),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatbinary),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatdate),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatdateTime),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatemail),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatguid),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatip),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatjson),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatnone),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatregEx),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatsurfaceHub),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormattime),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormaturl),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatversion),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormatxml),
	}
}

func parseDeviceManagementConfigurationStringSettingValueDefinitionFormat(input string) (*DeviceManagementConfigurationStringSettingValueDefinitionFormat, error) {
	vals := map[string]DeviceManagementConfigurationStringSettingValueDefinitionFormat{
		"base64":     DeviceManagementConfigurationStringSettingValueDefinitionFormatbase64,
		"bashscript": DeviceManagementConfigurationStringSettingValueDefinitionFormatbashScript,
		"binary":     DeviceManagementConfigurationStringSettingValueDefinitionFormatbinary,
		"date":       DeviceManagementConfigurationStringSettingValueDefinitionFormatdate,
		"datetime":   DeviceManagementConfigurationStringSettingValueDefinitionFormatdateTime,
		"email":      DeviceManagementConfigurationStringSettingValueDefinitionFormatemail,
		"guid":       DeviceManagementConfigurationStringSettingValueDefinitionFormatguid,
		"ip":         DeviceManagementConfigurationStringSettingValueDefinitionFormatip,
		"json":       DeviceManagementConfigurationStringSettingValueDefinitionFormatjson,
		"none":       DeviceManagementConfigurationStringSettingValueDefinitionFormatnone,
		"regex":      DeviceManagementConfigurationStringSettingValueDefinitionFormatregEx,
		"surfacehub": DeviceManagementConfigurationStringSettingValueDefinitionFormatsurfaceHub,
		"time":       DeviceManagementConfigurationStringSettingValueDefinitionFormattime,
		"url":        DeviceManagementConfigurationStringSettingValueDefinitionFormaturl,
		"version":    DeviceManagementConfigurationStringSettingValueDefinitionFormatversion,
		"xml":        DeviceManagementConfigurationStringSettingValueDefinitionFormatxml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationStringSettingValueDefinitionFormat(input)
	return &out, nil
}
