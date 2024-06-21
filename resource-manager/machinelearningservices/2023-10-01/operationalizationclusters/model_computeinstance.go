package operationalizationclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Compute = ComputeInstance{}

type ComputeInstance struct {
	Properties *ComputeInstanceProperties `json:"properties,omitempty"`

	// Fields inherited from Compute
	ComputeLocation    *string            `json:"computeLocation,omitempty"`
	CreatedOn          *string            `json:"createdOn,omitempty"`
	Description        *string            `json:"description,omitempty"`
	DisableLocalAuth   *bool              `json:"disableLocalAuth,omitempty"`
	IsAttachedCompute  *bool              `json:"isAttachedCompute,omitempty"`
	ModifiedOn         *string            `json:"modifiedOn,omitempty"`
	ProvisioningErrors *[]ErrorResponse   `json:"provisioningErrors,omitempty"`
	ProvisioningState  *ProvisioningState `json:"provisioningState,omitempty"`
	ResourceId         *string            `json:"resourceId,omitempty"`
}

var _ json.Marshaler = ComputeInstance{}

func (s ComputeInstance) MarshalJSON() ([]byte, error) {
	type wrapper ComputeInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ComputeInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ComputeInstance: %+v", err)
	}
	decoded["computeType"] = "ComputeInstance"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ComputeInstance: %+v", err)
	}

	return encoded, nil
}
