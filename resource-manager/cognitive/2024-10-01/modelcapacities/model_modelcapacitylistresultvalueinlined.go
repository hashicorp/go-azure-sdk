package modelcapacities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelCapacityListResultValueInlined struct {
	Id         *string                     `json:"id,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *ModelSkuCapacityProperties `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
