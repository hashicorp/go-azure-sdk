package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsLastAppSyncStatus string

const (
	AndroidForWorkSettingsLastAppSyncStatus_AndroidForWorkApiError AndroidForWorkSettingsLastAppSyncStatus = "androidForWorkApiError"
	AndroidForWorkSettingsLastAppSyncStatus_CredentialsNotValid    AndroidForWorkSettingsLastAppSyncStatus = "credentialsNotValid"
	AndroidForWorkSettingsLastAppSyncStatus_ManagementServiceError AndroidForWorkSettingsLastAppSyncStatus = "managementServiceError"
	AndroidForWorkSettingsLastAppSyncStatus_None                   AndroidForWorkSettingsLastAppSyncStatus = "none"
	AndroidForWorkSettingsLastAppSyncStatus_Success                AndroidForWorkSettingsLastAppSyncStatus = "success"
	AndroidForWorkSettingsLastAppSyncStatus_UnknownError           AndroidForWorkSettingsLastAppSyncStatus = "unknownError"
)

func PossibleValuesForAndroidForWorkSettingsLastAppSyncStatus() []string {
	return []string{
		string(AndroidForWorkSettingsLastAppSyncStatus_AndroidForWorkApiError),
		string(AndroidForWorkSettingsLastAppSyncStatus_CredentialsNotValid),
		string(AndroidForWorkSettingsLastAppSyncStatus_ManagementServiceError),
		string(AndroidForWorkSettingsLastAppSyncStatus_None),
		string(AndroidForWorkSettingsLastAppSyncStatus_Success),
		string(AndroidForWorkSettingsLastAppSyncStatus_UnknownError),
	}
}

func (s *AndroidForWorkSettingsLastAppSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkSettingsLastAppSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkSettingsLastAppSyncStatus(input string) (*AndroidForWorkSettingsLastAppSyncStatus, error) {
	vals := map[string]AndroidForWorkSettingsLastAppSyncStatus{
		"androidforworkapierror": AndroidForWorkSettingsLastAppSyncStatus_AndroidForWorkApiError,
		"credentialsnotvalid":    AndroidForWorkSettingsLastAppSyncStatus_CredentialsNotValid,
		"managementserviceerror": AndroidForWorkSettingsLastAppSyncStatus_ManagementServiceError,
		"none":                   AndroidForWorkSettingsLastAppSyncStatus_None,
		"success":                AndroidForWorkSettingsLastAppSyncStatus_Success,
		"unknownerror":           AndroidForWorkSettingsLastAppSyncStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsLastAppSyncStatus(input)
	return &out, nil
}
