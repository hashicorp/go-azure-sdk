package inferencedeltamodel

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeltaModelModifyRequest struct {
	AddDeltaModels    *[]string `json:"addDeltaModels,omitempty"`
	RemoveDeltaModels *[]string `json:"removeDeltaModels,omitempty"`
	TargetBaseModel   *string   `json:"targetBaseModel,omitempty"`
}
