package inputs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InputWatermarkProperties struct {
	WatermarkMode *InputWatermarkMode `json:"watermarkMode,omitempty"`
}
