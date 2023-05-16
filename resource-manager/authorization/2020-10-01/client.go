package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/eligiblechildresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentschedulerequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityschedulerequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/rolemanagementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/rolemanagementpolicyassignments"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	EligibleChildResources           *eligiblechildresources.EligibleChildResourcesClient
	RoleAssignmentScheduleInstances  *roleassignmentscheduleinstances.RoleAssignmentScheduleInstancesClient
	RoleAssignmentScheduleRequests   *roleassignmentschedulerequests.RoleAssignmentScheduleRequestsClient
	RoleAssignmentSchedules          *roleassignmentschedules.RoleAssignmentSchedulesClient
	RoleEligibilityScheduleInstances *roleeligibilityscheduleinstances.RoleEligibilityScheduleInstancesClient
	RoleEligibilityScheduleRequests  *roleeligibilityschedulerequests.RoleEligibilityScheduleRequestsClient
	RoleEligibilitySchedules         *roleeligibilityschedules.RoleEligibilitySchedulesClient
	RoleManagementPolicies           *rolemanagementpolicies.RoleManagementPoliciesClient
	RoleManagementPolicyAssignments  *rolemanagementpolicyassignments.RoleManagementPolicyAssignmentsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	eligibleChildResourcesClient, err := eligiblechildresources.NewEligibleChildResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EligibleChildResources client: %+v", err)
	}
	configureFunc(eligibleChildResourcesClient.Client)

	roleAssignmentScheduleInstancesClient, err := roleassignmentscheduleinstances.NewRoleAssignmentScheduleInstancesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentScheduleInstances client: %+v", err)
	}
	configureFunc(roleAssignmentScheduleInstancesClient.Client)

	roleAssignmentScheduleRequestsClient, err := roleassignmentschedulerequests.NewRoleAssignmentScheduleRequestsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentScheduleRequests client: %+v", err)
	}
	configureFunc(roleAssignmentScheduleRequestsClient.Client)

	roleAssignmentSchedulesClient, err := roleassignmentschedules.NewRoleAssignmentSchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleAssignmentSchedules client: %+v", err)
	}
	configureFunc(roleAssignmentSchedulesClient.Client)

	roleEligibilityScheduleInstancesClient, err := roleeligibilityscheduleinstances.NewRoleEligibilityScheduleInstancesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleEligibilityScheduleInstances client: %+v", err)
	}
	configureFunc(roleEligibilityScheduleInstancesClient.Client)

	roleEligibilityScheduleRequestsClient, err := roleeligibilityschedulerequests.NewRoleEligibilityScheduleRequestsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleEligibilityScheduleRequests client: %+v", err)
	}
	configureFunc(roleEligibilityScheduleRequestsClient.Client)

	roleEligibilitySchedulesClient, err := roleeligibilityschedules.NewRoleEligibilitySchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleEligibilitySchedules client: %+v", err)
	}
	configureFunc(roleEligibilitySchedulesClient.Client)

	roleManagementPoliciesClient, err := rolemanagementpolicies.NewRoleManagementPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicies client: %+v", err)
	}
	configureFunc(roleManagementPoliciesClient.Client)

	roleManagementPolicyAssignmentsClient, err := rolemanagementpolicyassignments.NewRoleManagementPolicyAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementPolicyAssignments client: %+v", err)
	}
	configureFunc(roleManagementPolicyAssignmentsClient.Client)

	return &Client{
		EligibleChildResources:           eligibleChildResourcesClient,
		RoleAssignmentScheduleInstances:  roleAssignmentScheduleInstancesClient,
		RoleAssignmentScheduleRequests:   roleAssignmentScheduleRequestsClient,
		RoleAssignmentSchedules:          roleAssignmentSchedulesClient,
		RoleEligibilityScheduleInstances: roleEligibilityScheduleInstancesClient,
		RoleEligibilityScheduleRequests:  roleEligibilityScheduleRequestsClient,
		RoleEligibilitySchedules:         roleEligibilitySchedulesClient,
		RoleManagementPolicies:           roleManagementPoliciesClient,
		RoleManagementPolicyAssignments:  roleManagementPolicyAssignmentsClient,
	}, nil
}
