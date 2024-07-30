package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus string

const (
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_AndroidForWorkApiError AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "androidForWorkApiError"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_CredentialsNotValid    AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "credentialsNotValid"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_ManagementServiceError AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "managementServiceError"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_None                   AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "none"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_Success                AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "success"
	AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_UnknownError           AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus = "unknownError"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_AndroidForWorkApiError),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_CredentialsNotValid),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_ManagementServiceError),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_None),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_Success),
		string(AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_UnknownError),
	}
}

func (s *AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus(input string) (*AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus{
		"androidforworkapierror": AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_AndroidForWorkApiError,
		"credentialsnotvalid":    AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_CredentialsNotValid,
		"managementserviceerror": AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_ManagementServiceError,
		"none":                   AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_None,
		"success":                AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_Success,
		"unknownerror":           AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsLastAppSyncStatus(input)
	return &out, nil
}
