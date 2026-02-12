package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobDefinition struct {
	Description          *string              `json:"description,omitempty"`
	JobProperties        SparkJobProperties   `json:"jobProperties"`
	Language             *string              `json:"language,omitempty"`
	RequiredSparkVersion *string              `json:"requiredSparkVersion,omitempty"`
	TargetBigDataPool    BigDataPoolReference `json:"targetBigDataPool"`
}
