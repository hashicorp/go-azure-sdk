package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogSizeCapability struct {
	Limit *int64       `json:"limit,omitempty"`
	Unit  *LogSizeUnit `json:"unit,omitempty"`
}
