package watcher

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WatcherProperties struct {
	CreationTime                *string            `json:"creationTime,omitempty"`
	Description                 *string            `json:"description,omitempty"`
	ExecutionFrequencyInSeconds *int64             `json:"executionFrequencyInSeconds,omitempty"`
	LastModifiedBy              *string            `json:"lastModifiedBy,omitempty"`
	LastModifiedTime            *string            `json:"lastModifiedTime,omitempty"`
	ScriptName                  *string            `json:"scriptName,omitempty"`
	ScriptParameters            *map[string]string `json:"scriptParameters,omitempty"`
	ScriptRunOn                 *string            `json:"scriptRunOn,omitempty"`
	Status                      *string            `json:"status,omitempty"`
}
