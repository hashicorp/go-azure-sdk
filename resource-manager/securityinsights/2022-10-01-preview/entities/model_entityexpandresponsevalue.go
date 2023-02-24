package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityExpandResponseValue struct {
	Edges    *[]EntityEdges `json:"edges,omitempty"`
	Entities *[]Entity      `json:"entities,omitempty"`
}
