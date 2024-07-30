package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Alert struct {
	ActivityGroupName           *string                       `json:"activityGroupName,omitempty"`
	AlertDetections             *[]AlertDetection             `json:"alertDetections,omitempty"`
	AssignedTo                  *string                       `json:"assignedTo,omitempty"`
	AzureSubscriptionId         *string                       `json:"azureSubscriptionId,omitempty"`
	AzureTenantId               *string                       `json:"azureTenantId,omitempty"`
	Category                    *string                       `json:"category,omitempty"`
	ClosedDateTime              *string                       `json:"closedDateTime,omitempty"`
	CloudAppStates              *[]CloudAppSecurityState      `json:"cloudAppStates,omitempty"`
	Comments                    *[]string                     `json:"comments,omitempty"`
	Confidence                  *int64                        `json:"confidence,omitempty"`
	CreatedDateTime             *string                       `json:"createdDateTime,omitempty"`
	Description                 *string                       `json:"description,omitempty"`
	DetectionIds                *[]string                     `json:"detectionIds,omitempty"`
	EventDateTime               *string                       `json:"eventDateTime,omitempty"`
	Feedback                    *AlertFeedback                `json:"feedback,omitempty"`
	FileStates                  *[]FileSecurityState          `json:"fileStates,omitempty"`
	HistoryStates               *[]AlertHistoryState          `json:"historyStates,omitempty"`
	HostStates                  *[]HostSecurityState          `json:"hostStates,omitempty"`
	Id                          *string                       `json:"id,omitempty"`
	IncidentIds                 *[]string                     `json:"incidentIds,omitempty"`
	InvestigationSecurityStates *[]InvestigationSecurityState `json:"investigationSecurityStates,omitempty"`
	LastEventDateTime           *string                       `json:"lastEventDateTime,omitempty"`
	LastModifiedDateTime        *string                       `json:"lastModifiedDateTime,omitempty"`
	MalwareStates               *[]MalwareState               `json:"malwareStates,omitempty"`
	MessageSecurityStates       *[]MessageSecurityState       `json:"messageSecurityStates,omitempty"`
	NetworkConnections          *[]NetworkConnection          `json:"networkConnections,omitempty"`
	ODataType                   *string                       `json:"@odata.type,omitempty"`
	Processes                   *[]Process                    `json:"processes,omitempty"`
	RecommendedActions          *[]string                     `json:"recommendedActions,omitempty"`
	RegistryKeyStates           *[]RegistryKeyState           `json:"registryKeyStates,omitempty"`
	SecurityResources           *[]SecurityResource           `json:"securityResources,omitempty"`
	Severity                    *AlertSeverity                `json:"severity,omitempty"`
	SourceMaterials             *[]string                     `json:"sourceMaterials,omitempty"`
	Status                      *AlertStatus                  `json:"status,omitempty"`
	Tags                        *[]string                     `json:"tags,omitempty"`
	Title                       *string                       `json:"title,omitempty"`
	Triggers                    *[]AlertTrigger               `json:"triggers,omitempty"`
	UriClickSecurityStates      *[]UriClickSecurityState      `json:"uriClickSecurityStates,omitempty"`
	UserStates                  *[]UserSecurityState          `json:"userStates,omitempty"`
	VendorInformation           *SecurityVendorInformation    `json:"vendorInformation,omitempty"`
	VulnerabilityStates         *[]VulnerabilityState         `json:"vulnerabilityStates,omitempty"`
}
