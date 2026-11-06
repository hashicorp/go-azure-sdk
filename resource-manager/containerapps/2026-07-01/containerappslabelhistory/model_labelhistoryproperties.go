package containerappslabelhistory

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelHistoryProperties struct {
	Records *[]LabelHistoryRecordItem `json:"records,omitempty"`
}
