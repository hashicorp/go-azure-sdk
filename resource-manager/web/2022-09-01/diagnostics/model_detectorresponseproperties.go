package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorResponseProperties struct {
	DataProvidersMetadata *[]DataProviderMetadata `json:"dataProvidersMetadata,omitempty"`
	Dataset               *[]DiagnosticData       `json:"dataset,omitempty"`
	Metadata              *DetectorInfo           `json:"metadata,omitempty"`
	Status                *Status                 `json:"status,omitempty"`
	SuggestedUtterances   *QueryUtterancesResults `json:"suggestedUtterances,omitempty"`
}
