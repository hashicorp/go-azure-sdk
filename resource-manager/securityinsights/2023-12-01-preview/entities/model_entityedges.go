package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityEdges struct {
	AdditionalData *interface{} `json:"additionalData,omitempty"`
	TargetEntityId *string      `json:"targetEntityId,omitempty"`
}
