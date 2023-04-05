package importpipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineOptions string

const (
	PipelineOptionsContinueOnErrors          PipelineOptions = "ContinueOnErrors"
	PipelineOptionsDeleteSourceBlobOnSuccess PipelineOptions = "DeleteSourceBlobOnSuccess"
	PipelineOptionsOverwriteBlobs            PipelineOptions = "OverwriteBlobs"
	PipelineOptionsOverwriteTags             PipelineOptions = "OverwriteTags"
)

func PossibleValuesForPipelineOptions() []string {
	return []string{
		string(PipelineOptionsContinueOnErrors),
		string(PipelineOptionsDeleteSourceBlobOnSuccess),
		string(PipelineOptionsOverwriteBlobs),
		string(PipelineOptionsOverwriteTags),
	}
}

func (s *PipelineOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForPipelineOptions() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = PipelineOptions(decoded)
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

type TriggerStatus string

const (
	TriggerStatusDisabled TriggerStatus = "Disabled"
	TriggerStatusEnabled  TriggerStatus = "Enabled"
)

func PossibleValuesForTriggerStatus() []string {
	return []string{
		string(TriggerStatusDisabled),
		string(TriggerStatusEnabled),
	}
}

func (s *TriggerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForTriggerStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = TriggerStatus(decoded)
	return nil
}
