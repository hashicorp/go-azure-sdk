package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperTargetConnectionsInfo struct {
	Connection         *MapperConnection    `json:"Connection,omitempty"`
	DataMapperMappings *[]DataMapperMapping `json:"DataMapperMappings,omitempty"`
	Relationships      *[]interface{}       `json:"Relationships,omitempty"`
	TargetEntities     *[]MapperTable       `json:"TargetEntities,omitempty"`
}
