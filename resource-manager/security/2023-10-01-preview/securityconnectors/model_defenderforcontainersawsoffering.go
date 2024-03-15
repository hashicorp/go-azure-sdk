package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForContainersAwsOffering{}

type DefenderForContainersAwsOffering struct {
	AutoProvisioning                       *bool                                                                 `json:"autoProvisioning,omitempty"`
	CloudWatchToKinesis                    *DefenderForContainersAwsOfferingCloudWatchToKinesis                  `json:"cloudWatchToKinesis,omitempty"`
	ContainerVulnerabilityAssessment       *DefenderForContainersAwsOfferingContainerVulnerabilityAssessment     `json:"containerVulnerabilityAssessment,omitempty"`
	ContainerVulnerabilityAssessmentTask   *DefenderForContainersAwsOfferingContainerVulnerabilityAssessmentTask `json:"containerVulnerabilityAssessmentTask,omitempty"`
	EnableContainerVulnerabilityAssessment *bool                                                                 `json:"enableContainerVulnerabilityAssessment,omitempty"`
	KinesisToS3                            *DefenderForContainersAwsOfferingKinesisToS3                          `json:"kinesisToS3,omitempty"`
	KubeAuditRetentionTime                 *int64                                                                `json:"kubeAuditRetentionTime,omitempty"`
	KubernetesScubaReader                  *DefenderForContainersAwsOfferingKubernetesScubaReader                `json:"kubernetesScubaReader,omitempty"`
	KubernetesService                      *DefenderForContainersAwsOfferingKubernetesService                    `json:"kubernetesService,omitempty"`
	MdcContainersAgentlessDiscoveryK8s     *DefenderForContainersAwsOfferingMdcContainersAgentlessDiscoveryK8s   `json:"mdcContainersAgentlessDiscoveryK8s,omitempty"`
	MdcContainersImageAssessment           *DefenderForContainersAwsOfferingMdcContainersImageAssessment         `json:"mdcContainersImageAssessment,omitempty"`
	ScubaExternalId                        *string                                                               `json:"scubaExternalId,omitempty"`

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = DefenderForContainersAwsOffering{}

func (s DefenderForContainersAwsOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForContainersAwsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForContainersAwsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForContainersAwsOffering: %+v", err)
	}
	decoded["offeringType"] = "DefenderForContainersAws"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForContainersAwsOffering: %+v", err)
	}

	return encoded, nil
}
