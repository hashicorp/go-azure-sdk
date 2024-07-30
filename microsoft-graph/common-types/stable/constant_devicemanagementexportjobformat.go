package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobFormat string

const (
	DeviceManagementExportJobFormat_Csv  DeviceManagementExportJobFormat = "csv"
	DeviceManagementExportJobFormat_Json DeviceManagementExportJobFormat = "json"
	DeviceManagementExportJobFormat_Pdf  DeviceManagementExportJobFormat = "pdf"
)

func PossibleValuesForDeviceManagementExportJobFormat() []string {
	return []string{
		string(DeviceManagementExportJobFormat_Csv),
		string(DeviceManagementExportJobFormat_Json),
		string(DeviceManagementExportJobFormat_Pdf),
	}
}

func (s *DeviceManagementExportJobFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExportJobFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExportJobFormat(input string) (*DeviceManagementExportJobFormat, error) {
	vals := map[string]DeviceManagementExportJobFormat{
		"csv":  DeviceManagementExportJobFormat_Csv,
		"json": DeviceManagementExportJobFormat_Json,
		"pdf":  DeviceManagementExportJobFormat_Pdf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobFormat(input)
	return &out, nil
}
