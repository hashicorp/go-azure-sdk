package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveValue struct {
	TimeStamp string  `json:"timeStamp"`
	Value     float64 `json:"value"`
}
