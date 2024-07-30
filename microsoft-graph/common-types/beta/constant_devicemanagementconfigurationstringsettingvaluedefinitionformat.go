package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueDefinitionFormat string

const (
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Base64     DeviceManagementConfigurationStringSettingValueDefinitionFormat = "base64"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_BashScript DeviceManagementConfigurationStringSettingValueDefinitionFormat = "bashScript"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Binary     DeviceManagementConfigurationStringSettingValueDefinitionFormat = "binary"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Date       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "date"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_DateTime   DeviceManagementConfigurationStringSettingValueDefinitionFormat = "dateTime"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Email      DeviceManagementConfigurationStringSettingValueDefinitionFormat = "email"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Guid       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "guid"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Ip         DeviceManagementConfigurationStringSettingValueDefinitionFormat = "ip"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Json       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "json"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_None       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "none"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_RegEx      DeviceManagementConfigurationStringSettingValueDefinitionFormat = "regEx"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_SurfaceHub DeviceManagementConfigurationStringSettingValueDefinitionFormat = "surfaceHub"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Time       DeviceManagementConfigurationStringSettingValueDefinitionFormat = "time"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Url        DeviceManagementConfigurationStringSettingValueDefinitionFormat = "url"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Version    DeviceManagementConfigurationStringSettingValueDefinitionFormat = "version"
	DeviceManagementConfigurationStringSettingValueDefinitionFormat_Xml        DeviceManagementConfigurationStringSettingValueDefinitionFormat = "xml"
)

func PossibleValuesForDeviceManagementConfigurationStringSettingValueDefinitionFormat() []string {
	return []string{
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Base64),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_BashScript),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Binary),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Date),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_DateTime),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Email),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Guid),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Ip),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Json),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_None),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_RegEx),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_SurfaceHub),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Time),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Url),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Version),
		string(DeviceManagementConfigurationStringSettingValueDefinitionFormat_Xml),
	}
}

func (s *DeviceManagementConfigurationStringSettingValueDefinitionFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationStringSettingValueDefinitionFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationStringSettingValueDefinitionFormat(input string) (*DeviceManagementConfigurationStringSettingValueDefinitionFormat, error) {
	vals := map[string]DeviceManagementConfigurationStringSettingValueDefinitionFormat{
		"base64":     DeviceManagementConfigurationStringSettingValueDefinitionFormat_Base64,
		"bashscript": DeviceManagementConfigurationStringSettingValueDefinitionFormat_BashScript,
		"binary":     DeviceManagementConfigurationStringSettingValueDefinitionFormat_Binary,
		"date":       DeviceManagementConfigurationStringSettingValueDefinitionFormat_Date,
		"datetime":   DeviceManagementConfigurationStringSettingValueDefinitionFormat_DateTime,
		"email":      DeviceManagementConfigurationStringSettingValueDefinitionFormat_Email,
		"guid":       DeviceManagementConfigurationStringSettingValueDefinitionFormat_Guid,
		"ip":         DeviceManagementConfigurationStringSettingValueDefinitionFormat_Ip,
		"json":       DeviceManagementConfigurationStringSettingValueDefinitionFormat_Json,
		"none":       DeviceManagementConfigurationStringSettingValueDefinitionFormat_None,
		"regex":      DeviceManagementConfigurationStringSettingValueDefinitionFormat_RegEx,
		"surfacehub": DeviceManagementConfigurationStringSettingValueDefinitionFormat_SurfaceHub,
		"time":       DeviceManagementConfigurationStringSettingValueDefinitionFormat_Time,
		"url":        DeviceManagementConfigurationStringSettingValueDefinitionFormat_Url,
		"version":    DeviceManagementConfigurationStringSettingValueDefinitionFormat_Version,
		"xml":        DeviceManagementConfigurationStringSettingValueDefinitionFormat_Xml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationStringSettingValueDefinitionFormat(input)
	return &out, nil
}
