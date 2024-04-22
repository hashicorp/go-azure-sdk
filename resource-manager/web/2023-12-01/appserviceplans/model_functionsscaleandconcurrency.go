package appserviceplans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionsScaleAndConcurrency struct {
	AlwaysReady          *[]FunctionsAlwaysReadyConfig         `json:"alwaysReady,omitempty"`
	InstanceMemoryMB     *float64                              `json:"instanceMemoryMB,omitempty"`
	MaximumInstanceCount *float64                              `json:"maximumInstanceCount,omitempty"`
	Triggers             *FunctionsScaleAndConcurrencyTriggers `json:"triggers,omitempty"`
}
