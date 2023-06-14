package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperTargetConnectionsInfo struct {
	Connection         *MapperConnection    `json:"connection,omitempty"`
	DataMapperMappings *[]DataMapperMapping `json:"dataMapperMappings,omitempty"`
	Relationships      *[]interface{}       `json:"relationships,omitempty"`
	TargetEntities     *[]MapperTable       `json:"targetEntities,omitempty"`
}
