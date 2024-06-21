package prediction

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictionResponseProperties struct {
	Category           *Category         `json:"category,omitempty"`
	ExtendedProperties *interface{}      `json:"extendedProperties,omitempty"`
	Impact             *Impact           `json:"impact,omitempty"`
	ImpactedField      *string           `json:"impactedField,omitempty"`
	LastUpdated        *string           `json:"lastUpdated,omitempty"`
	PredictionType     *PredictionType   `json:"predictionType,omitempty"`
	ShortDescription   *ShortDescription `json:"shortDescription,omitempty"`
}
