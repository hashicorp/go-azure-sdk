package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectivityCriterion struct {
	Type  string    `json:"type"`
	Value *[]string `json:"value,omitempty"`
}
