package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetadataValue struct {
	Name  *LocalizableString `json:"name,omitempty"`
	Value *string            `json:"value,omitempty"`
}
