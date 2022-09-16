package sqlpoolssensitivitylabels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelUpdateProperties struct {
	Column           string                     `json:"column"`
	Op               SensitivityLabelUpdateKind `json:"op"`
	Schema           string                     `json:"schema"`
	SensitivityLabel *SensitivityLabel          `json:"sensitivityLabel,omitempty"`
	Table            string                     `json:"table"`
}
