package inferencedeltamodel

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeltaModelStatusResponse struct {
	ActualInstanceCount   *int64                               `json:"actualInstanceCount,omitempty"`
	DeltaModels           *map[string][]DeltaModelCurrentState `json:"deltaModels,omitempty"`
	ExpectedInstanceCount *int64                               `json:"expectedInstanceCount,omitempty"`
	RevisionId            *string                              `json:"revisionId,omitempty"`
	TargetBaseModel       *string                              `json:"targetBaseModel,omitempty"`
}
