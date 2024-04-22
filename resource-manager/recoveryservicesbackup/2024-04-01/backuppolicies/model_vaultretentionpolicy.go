package backuppolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultRetentionPolicy struct {
	SnapshotRetentionInDays int64           `json:"snapshotRetentionInDays"`
	VaultRetention          RetentionPolicy `json:"vaultRetention"`
}

var _ json.Unmarshaler = &VaultRetentionPolicy{}

func (s *VaultRetentionPolicy) UnmarshalJSON(bytes []byte) error {
	type alias VaultRetentionPolicy
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into VaultRetentionPolicy: %+v", err)
	}

	s.SnapshotRetentionInDays = decoded.SnapshotRetentionInDays

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VaultRetentionPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["vaultRetention"]; ok {
		impl, err := unmarshalRetentionPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'VaultRetention' for 'VaultRetentionPolicy': %+v", err)
		}
		s.VaultRetention = impl
	}
	return nil
}
