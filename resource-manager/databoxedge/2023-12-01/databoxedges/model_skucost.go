package databoxedges

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkuCost struct {
	ExtendedUnit *string `json:"extendedUnit,omitempty"`
	MeterId      *string `json:"meterId,omitempty"`
	Quantity     *int64  `json:"quantity,omitempty"`
}
