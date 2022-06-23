package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetInsightsError struct {
	ErrorMessage string  `json:"errorMessage"`
	Kind         Kind    `json:"kind"`
	QueryId      *string `json:"queryId,omitempty"`
}
