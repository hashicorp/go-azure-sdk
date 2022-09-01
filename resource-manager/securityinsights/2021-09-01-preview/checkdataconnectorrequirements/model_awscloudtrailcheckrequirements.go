package checkdataconnectorrequirements

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataConnectorsCheckRequirements = AwsCloudTrailCheckRequirements{}

type AwsCloudTrailCheckRequirements struct {

	// Fields inherited from DataConnectorsCheckRequirements
}

var _ json.Marshaler = AwsCloudTrailCheckRequirements{}

func (s AwsCloudTrailCheckRequirements) MarshalJSON() ([]byte, error) {
	type wrapper AwsCloudTrailCheckRequirements
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsCloudTrailCheckRequirements: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsCloudTrailCheckRequirements: %+v", err)
	}
	decoded["kind"] = "AmazonWebServicesCloudTrail"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsCloudTrailCheckRequirements: %+v", err)
	}

	return encoded, nil
}
