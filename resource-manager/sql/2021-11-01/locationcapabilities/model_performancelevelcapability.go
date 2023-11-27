package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerformanceLevelCapability struct {
	Unit  *PerformanceLevelUnit `json:"unit,omitempty"`
	Value *float64              `json:"value,omitempty"`
}
