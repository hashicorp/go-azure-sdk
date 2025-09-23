package v2024_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/linkedserver"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/redis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/rediscacheaccesspolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/rediscacheaccesspolicyassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/redisfirewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/redispatchschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redis/2024-11-01/redisresources"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	LinkedServer                      *linkedserver.LinkedServerClient
	PrivateEndpointConnections        *privateendpointconnections.PrivateEndpointConnectionsClient
	Redis                             *redis.RedisClient
	RedisCacheAccessPolicies          *rediscacheaccesspolicies.RedisCacheAccessPoliciesClient
	RedisCacheAccessPolicyAssignments *rediscacheaccesspolicyassignments.RedisCacheAccessPolicyAssignmentsClient
	RedisFirewallRules                *redisfirewallrules.RedisFirewallRulesClient
	RedisPatchSchedules               *redispatchschedules.RedisPatchSchedulesClient
	RedisResources                    *redisresources.RedisResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	linkedServerClient, err := linkedserver.NewLinkedServerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedServer client: %+v", err)
	}
	configureFunc(linkedServerClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	redisCacheAccessPoliciesClient, err := rediscacheaccesspolicies.NewRedisCacheAccessPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisCacheAccessPolicies client: %+v", err)
	}
	configureFunc(redisCacheAccessPoliciesClient.Client)

	redisCacheAccessPolicyAssignmentsClient, err := rediscacheaccesspolicyassignments.NewRedisCacheAccessPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisCacheAccessPolicyAssignments client: %+v", err)
	}
	configureFunc(redisCacheAccessPolicyAssignmentsClient.Client)

	redisClient, err := redis.NewRedisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Redis client: %+v", err)
	}
	configureFunc(redisClient.Client)

	redisFirewallRulesClient, err := redisfirewallrules.NewRedisFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisFirewallRules client: %+v", err)
	}
	configureFunc(redisFirewallRulesClient.Client)

	redisPatchSchedulesClient, err := redispatchschedules.NewRedisPatchSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisPatchSchedules client: %+v", err)
	}
	configureFunc(redisPatchSchedulesClient.Client)

	redisResourcesClient, err := redisresources.NewRedisResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RedisResources client: %+v", err)
	}
	configureFunc(redisResourcesClient.Client)

	return &Client{
		LinkedServer:                      linkedServerClient,
		PrivateEndpointConnections:        privateEndpointConnectionsClient,
		Redis:                             redisClient,
		RedisCacheAccessPolicies:          redisCacheAccessPoliciesClient,
		RedisCacheAccessPolicyAssignments: redisCacheAccessPolicyAssignmentsClient,
		RedisFirewallRules:                redisFirewallRulesClient,
		RedisPatchSchedules:               redisPatchSchedulesClient,
		RedisResources:                    redisResourcesClient,
	}, nil
}
