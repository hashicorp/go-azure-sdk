package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkServiceError struct {
	ErrorCode *string           `json:"errorCode,omitempty"`
	Message   *string           `json:"message,omitempty"`
	Source    *SparkErrorSource `json:"source,omitempty"`
}
