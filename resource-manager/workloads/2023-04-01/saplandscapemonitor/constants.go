package saplandscapemonitor

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SapLandscapeMonitorProvisioningState string

const (
	SapLandscapeMonitorProvisioningStateAccepted  SapLandscapeMonitorProvisioningState = "Accepted"
	SapLandscapeMonitorProvisioningStateCanceled  SapLandscapeMonitorProvisioningState = "Canceled"
	SapLandscapeMonitorProvisioningStateCreated   SapLandscapeMonitorProvisioningState = "Created"
	SapLandscapeMonitorProvisioningStateFailed    SapLandscapeMonitorProvisioningState = "Failed"
	SapLandscapeMonitorProvisioningStateSucceeded SapLandscapeMonitorProvisioningState = "Succeeded"
)

func PossibleValuesForSapLandscapeMonitorProvisioningState() []string {
	return []string{
		string(SapLandscapeMonitorProvisioningStateAccepted),
		string(SapLandscapeMonitorProvisioningStateCanceled),
		string(SapLandscapeMonitorProvisioningStateCreated),
		string(SapLandscapeMonitorProvisioningStateFailed),
		string(SapLandscapeMonitorProvisioningStateSucceeded),
	}
}

func (s *SapLandscapeMonitorProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSapLandscapeMonitorProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SapLandscapeMonitorProvisioningState(decoded)
	return nil
}
