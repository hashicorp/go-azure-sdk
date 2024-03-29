package autoscaleapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveAutoscalePolicy struct {
	ScaleLookAheadTime *string                            `json:"scaleLookAheadTime,omitempty"`
	ScaleMode          PredictiveAutoscalePolicyScaleMode `json:"scaleMode"`
}
