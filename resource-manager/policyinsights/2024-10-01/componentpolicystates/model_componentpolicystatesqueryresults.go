package componentpolicystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentPolicyStatesQueryResults struct {
	OdataContext *string                 `json:"@odata.context,omitempty"`
	OdataCount   *int64                  `json:"@odata.count,omitempty"`
	Value        *[]ComponentPolicyState `json:"value,omitempty"`
}
