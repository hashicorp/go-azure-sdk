package databaserecommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionStateInfo struct {
	ActionInitiatedBy *RecommendedActionInitiatedBy `json:"actionInitiatedBy,omitempty"`
	CurrentValue      RecommendedActionCurrentState `json:"currentValue"`
	LastModified      *string                       `json:"lastModified,omitempty"`
}
