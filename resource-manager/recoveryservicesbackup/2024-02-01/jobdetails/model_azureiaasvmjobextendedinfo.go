package jobdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureIaaSVMJobExtendedInfo struct {
	DynamicErrorMessage        *string                      `json:"dynamicErrorMessage,omitempty"`
	EstimatedRemainingDuration *string                      `json:"estimatedRemainingDuration,omitempty"`
	InternalPropertyBag        *map[string]string           `json:"internalPropertyBag,omitempty"`
	ProgressPercentage         *float64                     `json:"progressPercentage,omitempty"`
	PropertyBag                *map[string]string           `json:"propertyBag,omitempty"`
	TasksList                  *[]AzureIaaSVMJobTaskDetails `json:"tasksList,omitempty"`
}
