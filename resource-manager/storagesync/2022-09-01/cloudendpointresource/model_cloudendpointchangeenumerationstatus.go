package cloudendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointChangeEnumerationStatus struct {
	Activity              *CloudEndpointChangeEnumerationActivity   `json:"activity,omitempty"`
	LastEnumerationStatus *CloudEndpointLastChangeEnumerationStatus `json:"lastEnumerationStatus,omitempty"`
	LastUpdatedTimestamp  *string                                   `json:"lastUpdatedTimestamp,omitempty"`
}
