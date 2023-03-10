package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/eligiblechildresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentschedulerequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleassignmentschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityscheduleinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityschedulerequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/roleeligibilityschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/rolemanagementpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2020-10-01/rolemanagementpolicyassignments"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	eligibleChildResourcesClient := eligiblechildresources.NewEligibleChildResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&eligibleChildResourcesClient.Client)

	roleAssignmentScheduleInstancesClient := roleassignmentscheduleinstances.NewRoleAssignmentScheduleInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&roleAssignmentScheduleInstancesClient.Client)

	roleAssignmentScheduleRequestsClient := roleassignmentschedulerequests.NewRoleAssignmentScheduleRequestsClientWithBaseURI(endpoint)
	configureAuthFunc(&roleAssignmentScheduleRequestsClient.Client)

	roleAssignmentSchedulesClient := roleassignmentschedules.NewRoleAssignmentSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&roleAssignmentSchedulesClient.Client)

	roleEligibilityScheduleInstancesClient := roleeligibilityscheduleinstances.NewRoleEligibilityScheduleInstancesClientWithBaseURI(endpoint)
	configureAuthFunc(&roleEligibilityScheduleInstancesClient.Client)

	roleEligibilityScheduleRequestsClient := roleeligibilityschedulerequests.NewRoleEligibilityScheduleRequestsClientWithBaseURI(endpoint)
	configureAuthFunc(&roleEligibilityScheduleRequestsClient.Client)

	roleEligibilitySchedulesClient := roleeligibilityschedules.NewRoleEligibilitySchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&roleEligibilitySchedulesClient.Client)

	roleManagementPoliciesClient := rolemanagementpolicies.NewRoleManagementPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&roleManagementPoliciesClient.Client)

	roleManagementPolicyAssignmentsClient := rolemanagementpolicyassignments.NewRoleManagementPolicyAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&roleManagementPolicyAssignmentsClient.Client)

	return Client{
		EligibleChildResources:           &eligibleChildResourcesClient,
		RoleAssignmentScheduleInstances:  &roleAssignmentScheduleInstancesClient,
		RoleAssignmentScheduleRequests:   &roleAssignmentScheduleRequestsClient,
		RoleAssignmentSchedules:          &roleAssignmentSchedulesClient,
		RoleEligibilityScheduleInstances: &roleEligibilityScheduleInstancesClient,
		RoleEligibilityScheduleRequests:  &roleEligibilityScheduleRequestsClient,
		RoleEligibilitySchedules:         &roleEligibilitySchedulesClient,
		RoleManagementPolicies:           &roleManagementPoliciesClient,
		RoleManagementPolicyAssignments:  &roleManagementPolicyAssignmentsClient,
	}
}
