package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScaleRule struct {
	AzureQueue *QueueScaleRule  `json:"azureQueue"`
	Custom     *CustomScaleRule `json:"custom"`
	HTTP       *HTTPScaleRule   `json:"http"`
	Name       *string          `json:"name,omitempty"`
}
