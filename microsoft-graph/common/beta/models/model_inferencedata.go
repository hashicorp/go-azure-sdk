package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceData struct {
	ODataType               *string `json:"@odata.type,omitempty"`
	UserHasVerifiedAccuracy *bool   `json:"userHasVerifiedAccuracy,omitempty"`
}
