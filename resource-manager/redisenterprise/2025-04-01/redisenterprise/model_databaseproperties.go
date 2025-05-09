package redisenterprise

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseProperties struct {
	AccessKeysAuthentication *AccessKeysAuthentication         `json:"accessKeysAuthentication,omitempty"`
	ClientProtocol           *Protocol                         `json:"clientProtocol,omitempty"`
	ClusteringPolicy         *ClusteringPolicy                 `json:"clusteringPolicy,omitempty"`
	DeferUpgrade             *DeferUpgradeSetting              `json:"deferUpgrade,omitempty"`
	EvictionPolicy           *EvictionPolicy                   `json:"evictionPolicy,omitempty"`
	GeoReplication           *DatabasePropertiesGeoReplication `json:"geoReplication,omitempty"`
	Modules                  *[]Module                         `json:"modules,omitempty"`
	Persistence              *Persistence                      `json:"persistence,omitempty"`
	Port                     *int64                            `json:"port,omitempty"`
	ProvisioningState        *ProvisioningState                `json:"provisioningState,omitempty"`
	RedisVersion             *string                           `json:"redisVersion,omitempty"`
	ResourceState            *ResourceState                    `json:"resourceState,omitempty"`
}
