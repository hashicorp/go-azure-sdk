package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSignalRecommendationsResponse struct {
	RecommendedConfigurations []SignalConfiguration `json:"recommendedConfigurations"`
	RecommendedSignals        []SignalConfiguration `json:"recommendedSignals"`
}
