package virtualmachineimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageDeprecationStatus struct {
	AlternativeOption        *AlternativeOption `json:"alternativeOption,omitempty"`
	ImageState               *ImageState        `json:"imageState,omitempty"`
	ScheduledDeprecationTime *string            `json:"scheduledDeprecationTime,omitempty"`
}
