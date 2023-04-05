package sourcecontrolsyncjob

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCompleted ProvisioningState = "Completed"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateRunning   ProvisioningState = "Running"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCompleted),
		string(ProvisioningStateFailed),
		string(ProvisioningStateRunning),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type SyncType string

const (
	SyncTypeFullSync    SyncType = "FullSync"
	SyncTypePartialSync SyncType = "PartialSync"
)

func PossibleValuesForSyncType() []string {
	return []string{
		string(SyncTypeFullSync),
		string(SyncTypePartialSync),
	}
}

func (s *SyncType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSyncType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SyncType(decoded)
	return nil
}
