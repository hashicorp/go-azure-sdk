package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineRunCommandParameters struct {
	Arguments        *[]string `json:"arguments,omitempty"`
	LimitTimeSeconds int64     `json:"limitTimeSeconds"`
	Script           string    `json:"script"`
}
