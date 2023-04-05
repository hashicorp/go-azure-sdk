package pipelineruns

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunSourceType string

const (
	PipelineRunSourceTypeAzureStorageBlob PipelineRunSourceType = "AzureStorageBlob"
)

func PossibleValuesForPipelineRunSourceType() []string {
	return []string{
		string(PipelineRunSourceTypeAzureStorageBlob),
	}
}

func (s *PipelineRunSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPipelineRunSourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PipelineRunSourceType(decoded)
	return nil
}

type PipelineRunTargetType string

const (
	PipelineRunTargetTypeAzureStorageBlob PipelineRunTargetType = "AzureStorageBlob"
)

func PossibleValuesForPipelineRunTargetType() []string {
	return []string{
		string(PipelineRunTargetTypeAzureStorageBlob),
	}
}

func (s *PipelineRunTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPipelineRunTargetType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PipelineRunTargetType(decoded)
	return nil
}

type PipelineSourceType string

const (
	PipelineSourceTypeAzureStorageBlobContainer PipelineSourceType = "AzureStorageBlobContainer"
)

func PossibleValuesForPipelineSourceType() []string {
	return []string{
		string(PipelineSourceTypeAzureStorageBlobContainer),
	}
}

func (s *PipelineSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPipelineSourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PipelineSourceType(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
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
