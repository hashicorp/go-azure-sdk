package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Condition struct {
	Message   *string `json:"message,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}
