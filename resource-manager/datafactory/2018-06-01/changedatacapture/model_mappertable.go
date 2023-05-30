package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperTable struct {
	Name       *string                `json:"name,omitempty"`
	Properties *MapperTableProperties `json:"properties,omitempty"`
}
