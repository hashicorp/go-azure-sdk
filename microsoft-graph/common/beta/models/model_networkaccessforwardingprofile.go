package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingProfile struct {
	Associations          *[]NetworkaccessAssociation                          `json:"associations,omitempty"`
	Description           *string                                              `json:"description,omitempty"`
	Id                    *string                                              `json:"id,omitempty"`
	LastModifiedDateTime  *string                                              `json:"lastModifiedDateTime,omitempty"`
	Name                  *string                                              `json:"name,omitempty"`
	ODataType             *string                                              `json:"@odata.type,omitempty"`
	Policies              *[]NetworkaccessPolicyLink                           `json:"policies,omitempty"`
	Priority              *int64                                               `json:"priority,omitempty"`
	State                 *NetworkaccessForwardingProfileState                 `json:"state,omitempty"`
	TrafficForwardingType *NetworkaccessForwardingProfileTrafficForwardingType `json:"trafficForwardingType,omitempty"`
	Version               *string                                              `json:"version,omitempty"`
}
