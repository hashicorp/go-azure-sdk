package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightsErrorKind struct {
	ErrorMessage string           `json:"errorMessage"`
	Kind         GetInsightsError `json:"kind"`
	QueryId      *string          `json:"queryId,omitempty"`
}
