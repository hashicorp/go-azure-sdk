package jobdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpmJobExtendedInfo struct {
	DynamicErrorMessage *string              `json:"dynamicErrorMessage,omitempty"`
	PropertyBag         *map[string]string   `json:"propertyBag,omitempty"`
	TasksList           *[]DpmJobTaskDetails `json:"tasksList,omitempty"`
}
