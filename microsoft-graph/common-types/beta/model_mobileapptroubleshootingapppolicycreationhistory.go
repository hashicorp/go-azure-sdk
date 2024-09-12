package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppTroubleshootingHistoryItem = MobileAppTroubleshootingAppPolicyCreationHistory{}

type MobileAppTroubleshootingAppPolicyCreationHistory struct {
	// Error code for the failure, empty if no failure.
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// Indicates the type of execution status of the device management script.
	RunState *RunState `json:"runState,omitempty"`

	// Fields inherited from MobileAppTroubleshootingHistoryItem

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time when the history item occurred.
	OccurrenceDateTime *string `json:"occurrenceDateTime,omitempty"`

	// Object containing detailed information about the error and its remediation.
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`
}

func (s MobileAppTroubleshootingAppPolicyCreationHistory) MobileAppTroubleshootingHistoryItem() BaseMobileAppTroubleshootingHistoryItemImpl {
	return BaseMobileAppTroubleshootingHistoryItemImpl{
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
		OccurrenceDateTime:          s.OccurrenceDateTime,
		TroubleshootingErrorDetails: s.TroubleshootingErrorDetails,
	}
}

var _ json.Marshaler = MobileAppTroubleshootingAppPolicyCreationHistory{}

func (s MobileAppTroubleshootingAppPolicyCreationHistory) MarshalJSON() ([]byte, error) {
	type wrapper MobileAppTroubleshootingAppPolicyCreationHistory
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MobileAppTroubleshootingAppPolicyCreationHistory: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MobileAppTroubleshootingAppPolicyCreationHistory: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.mobileAppTroubleshootingAppPolicyCreationHistory"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MobileAppTroubleshootingAppPolicyCreationHistory: %+v", err)
	}

	return encoded, nil
}
