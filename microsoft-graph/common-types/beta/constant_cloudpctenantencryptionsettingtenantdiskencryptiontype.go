package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCTenantEncryptionSettingTenantDiskEncryptionType string

const (
	CloudPCTenantEncryptionSettingTenantDiskEncryptionType_CustomerManagedKey CloudPCTenantEncryptionSettingTenantDiskEncryptionType = "customerManagedKey"
	CloudPCTenantEncryptionSettingTenantDiskEncryptionType_PlatformManagedKey CloudPCTenantEncryptionSettingTenantDiskEncryptionType = "platformManagedKey"
)

func PossibleValuesForCloudPCTenantEncryptionSettingTenantDiskEncryptionType() []string {
	return []string{
		string(CloudPCTenantEncryptionSettingTenantDiskEncryptionType_CustomerManagedKey),
		string(CloudPCTenantEncryptionSettingTenantDiskEncryptionType_PlatformManagedKey),
	}
}

func (s *CloudPCTenantEncryptionSettingTenantDiskEncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCTenantEncryptionSettingTenantDiskEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCTenantEncryptionSettingTenantDiskEncryptionType(input string) (*CloudPCTenantEncryptionSettingTenantDiskEncryptionType, error) {
	vals := map[string]CloudPCTenantEncryptionSettingTenantDiskEncryptionType{
		"customermanagedkey": CloudPCTenantEncryptionSettingTenantDiskEncryptionType_CustomerManagedKey,
		"platformmanagedkey": CloudPCTenantEncryptionSettingTenantDiskEncryptionType_PlatformManagedKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCTenantEncryptionSettingTenantDiskEncryptionType(input)
	return &out, nil
}
