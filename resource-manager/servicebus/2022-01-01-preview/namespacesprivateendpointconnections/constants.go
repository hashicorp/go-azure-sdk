package namespacesprivateendpointconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

func PossibleValuesForEndPointProvisioningState() []string {
	return []string{
		string(EndPointProvisioningStateCanceled),
		string(EndPointProvisioningStateCreating),
		string(EndPointProvisioningStateDeleting),
		string(EndPointProvisioningStateFailed),
		string(EndPointProvisioningStateSucceeded),
		string(EndPointProvisioningStateUpdating),
	}
}

func (s *EndPointProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForEndPointProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = EndPointProvisioningState(decoded)
	return nil
}

type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkConnectionStatus() []string {
	return []string{
		string(PrivateLinkConnectionStatusApproved),
		string(PrivateLinkConnectionStatusDisconnected),
		string(PrivateLinkConnectionStatusPending),
		string(PrivateLinkConnectionStatusRejected),
	}
}

func (s *PrivateLinkConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPrivateLinkConnectionStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PrivateLinkConnectionStatus(decoded)
	return nil
}
