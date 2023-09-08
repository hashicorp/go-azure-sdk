package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionHealthCheck struct {
	AdditionalDetails *string                                          `json:"additionalDetails,omitempty"`
	CorrelationId     *string                                          `json:"correlationId,omitempty"`
	DisplayName       *string                                          `json:"displayName,omitempty"`
	EndDateTime       *string                                          `json:"endDateTime,omitempty"`
	ErrorType         *CloudPCOnPremisesConnectionHealthCheckErrorType `json:"errorType,omitempty"`
	ODataType         *string                                          `json:"@odata.type,omitempty"`
	RecommendedAction *string                                          `json:"recommendedAction,omitempty"`
	StartDateTime     *string                                          `json:"startDateTime,omitempty"`
	Status            *CloudPCOnPremisesConnectionHealthCheckStatus    `json:"status,omitempty"`
}
