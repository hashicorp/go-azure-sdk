package providerinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SslPreference string

const (
	SslPreferenceDisabled          SslPreference = "Disabled"
	SslPreferenceRootCertificate   SslPreference = "RootCertificate"
	SslPreferenceServerCertificate SslPreference = "ServerCertificate"
)

func PossibleValuesForSslPreference() []string {
	return []string{
		string(SslPreferenceDisabled),
		string(SslPreferenceRootCertificate),
		string(SslPreferenceServerCertificate),
	}
}

func (s *SslPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSslPreference() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SslPreference(decoded)
	return nil
}

type WorkloadMonitorProvisioningState string

const (
	WorkloadMonitorProvisioningStateAccepted  WorkloadMonitorProvisioningState = "Accepted"
	WorkloadMonitorProvisioningStateCreating  WorkloadMonitorProvisioningState = "Creating"
	WorkloadMonitorProvisioningStateDeleting  WorkloadMonitorProvisioningState = "Deleting"
	WorkloadMonitorProvisioningStateFailed    WorkloadMonitorProvisioningState = "Failed"
	WorkloadMonitorProvisioningStateMigrating WorkloadMonitorProvisioningState = "Migrating"
	WorkloadMonitorProvisioningStateSucceeded WorkloadMonitorProvisioningState = "Succeeded"
	WorkloadMonitorProvisioningStateUpdating  WorkloadMonitorProvisioningState = "Updating"
)

func PossibleValuesForWorkloadMonitorProvisioningState() []string {
	return []string{
		string(WorkloadMonitorProvisioningStateAccepted),
		string(WorkloadMonitorProvisioningStateCreating),
		string(WorkloadMonitorProvisioningStateDeleting),
		string(WorkloadMonitorProvisioningStateFailed),
		string(WorkloadMonitorProvisioningStateMigrating),
		string(WorkloadMonitorProvisioningStateSucceeded),
		string(WorkloadMonitorProvisioningStateUpdating),
	}
}

func (s *WorkloadMonitorProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForWorkloadMonitorProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = WorkloadMonitorProvisioningState(decoded)
	return nil
}
