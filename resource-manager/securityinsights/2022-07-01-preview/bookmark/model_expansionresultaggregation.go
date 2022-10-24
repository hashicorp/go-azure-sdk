package bookmark

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpansionResultAggregation struct {
	AggregationType *string    `json:"aggregationType,omitempty"`
	Count           int64      `json:"count"`
	DisplayName     *string    `json:"displayName,omitempty"`
	EntityKind      EntityKind `json:"entityKind"`
}
