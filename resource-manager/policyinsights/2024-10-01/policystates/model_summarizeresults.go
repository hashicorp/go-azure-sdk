package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummarizeResults struct {
	OdataContext *string    `json:"@odata.context,omitempty"`
	OdataCount   *int64     `json:"@odata.count,omitempty"`
	Value        *[]Summary `json:"value,omitempty"`
}
