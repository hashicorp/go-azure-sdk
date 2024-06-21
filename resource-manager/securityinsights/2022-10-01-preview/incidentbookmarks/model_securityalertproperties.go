package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertProperties struct {
	AdditionalData        *map[string]interface{}                            `json:"additionalData,omitempty"`
	AlertDisplayName      *string                                            `json:"alertDisplayName,omitempty"`
	AlertLink             *string                                            `json:"alertLink,omitempty"`
	AlertType             *string                                            `json:"alertType,omitempty"`
	CompromisedEntity     *string                                            `json:"compromisedEntity,omitempty"`
	ConfidenceLevel       *ConfidenceLevel                                   `json:"confidenceLevel,omitempty"`
	ConfidenceReasons     *[]SecurityAlertPropertiesConfidenceReasonsInlined `json:"confidenceReasons,omitempty"`
	ConfidenceScore       *float64                                           `json:"confidenceScore,omitempty"`
	ConfidenceScoreStatus *ConfidenceScoreStatus                             `json:"confidenceScoreStatus,omitempty"`
	Description           *string                                            `json:"description,omitempty"`
	EndTimeUtc            *string                                            `json:"endTimeUtc,omitempty"`
	FriendlyName          *string                                            `json:"friendlyName,omitempty"`
	Intent                *KillChainIntent                                   `json:"intent,omitempty"`
	ProcessingEndTime     *string                                            `json:"processingEndTime,omitempty"`
	ProductComponentName  *string                                            `json:"productComponentName,omitempty"`
	ProductName           *string                                            `json:"productName,omitempty"`
	ProductVersion        *string                                            `json:"productVersion,omitempty"`
	ProviderAlertId       *string                                            `json:"providerAlertId,omitempty"`
	RemediationSteps      *[]string                                          `json:"remediationSteps,omitempty"`
	ResourceIdentifiers   *[]interface{}                                     `json:"resourceIdentifiers,omitempty"`
	Severity              *AlertSeverity                                     `json:"severity,omitempty"`
	StartTimeUtc          *string                                            `json:"startTimeUtc,omitempty"`
	Status                *AlertStatus                                       `json:"status,omitempty"`
	SystemAlertId         *string                                            `json:"systemAlertId,omitempty"`
	Tactics               *[]AttackTactic                                    `json:"tactics,omitempty"`
	TimeGenerated         *string                                            `json:"timeGenerated,omitempty"`
	VendorName            *string                                            `json:"vendorName,omitempty"`
}
