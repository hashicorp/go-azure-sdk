package inferencedeltamodel

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeltaModelStatusRequest struct {
	DeltaModels     *[]string `json:"deltaModels,omitempty"`
	TargetBaseModel *string   `json:"targetBaseModel,omitempty"`
}
