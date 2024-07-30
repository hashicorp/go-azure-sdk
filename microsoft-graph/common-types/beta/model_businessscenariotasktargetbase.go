package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BusinessScenarioTaskTargetBase struct {
	ODataType      *string                                       `json:"@odata.type,omitempty"`
	TaskTargetKind *BusinessScenarioTaskTargetBaseTaskTargetKind `json:"taskTargetKind,omitempty"`
}
