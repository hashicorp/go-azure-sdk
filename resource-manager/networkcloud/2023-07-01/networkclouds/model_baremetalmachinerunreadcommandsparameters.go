package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineRunReadCommandsParameters struct {
	Commands         []BareMetalMachineCommandSpecification `json:"commands"`
	LimitTimeSeconds int64                                  `json:"limitTimeSeconds"`
}
