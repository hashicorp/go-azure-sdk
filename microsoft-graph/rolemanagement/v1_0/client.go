package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagementdirectoryroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/v1_0/rolemanagemententitlementmanagementroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	RoleManagement                                                                   *rolemanagement.RoleManagementClient
	RoleManagementDirectory                                                          *rolemanagementdirectory.RoleManagementDirectoryClient
	RoleManagementDirectoryResourceNamespace                                         *rolemanagementdirectoryresourcenamespace.RoleManagementDirectoryResourceNamespaceClient
	RoleManagementDirectoryResourceNamespaceResourceAction                           *rolemanagementdirectoryresourcenamespaceresourceaction.RoleManagementDirectoryResourceNamespaceResourceActionClient
	RoleManagementDirectoryRoleAssignment                                            *rolemanagementdirectoryroleassignment.RoleManagementDirectoryRoleAssignmentClient
	RoleManagementDirectoryRoleAssignmentAppScope                                    *rolemanagementdirectoryroleassignmentappscope.RoleManagementDirectoryRoleAssignmentAppScopeClient
	RoleManagementDirectoryRoleAssignmentDirectoryScope                              *rolemanagementdirectoryroleassignmentdirectoryscope.RoleManagementDirectoryRoleAssignmentDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentPrincipal                                   *rolemanagementdirectoryroleassignmentprincipal.RoleManagementDirectoryRoleAssignmentPrincipalClient
	RoleManagementDirectoryRoleAssignmentRoleDefinition                              *rolemanagementdirectoryroleassignmentroledefinition.RoleManagementDirectoryRoleAssignmentRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentSchedule                                    *rolemanagementdirectoryroleassignmentschedule.RoleManagementDirectoryRoleAssignmentScheduleClient
	RoleManagementDirectoryRoleAssignmentScheduleActivatedUsing                      *rolemanagementdirectoryroleassignmentscheduleactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleAppScope                            *rolemanagementdirectoryroleassignmentscheduleappscope.RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleDirectoryScope                      *rolemanagementdirectoryroleassignmentscheduledirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstance                            *rolemanagementdirectoryroleassignmentscheduleinstance.RoleManagementDirectoryRoleAssignmentScheduleInstanceClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsing              *rolemanagementdirectoryroleassignmentscheduleinstanceactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScope                    *rolemanagementdirectoryroleassignmentscheduleinstanceappscope.RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScope              *rolemanagementdirectoryroleassignmentscheduleinstancedirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipal                   *rolemanagementdirectoryroleassignmentscheduleinstanceprincipal.RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinition              *rolemanagementdirectoryroleassignmentscheduleinstanceroledefinition.RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentSchedulePrincipal                           *rolemanagementdirectoryroleassignmentscheduleprincipal.RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleRequest                             *rolemanagementdirectoryroleassignmentschedulerequest.RoleManagementDirectoryRoleAssignmentScheduleRequestClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsing               *rolemanagementdirectoryroleassignmentschedulerequestactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestAppScope                     *rolemanagementdirectoryroleassignmentschedulerequestappscope.RoleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScope               *rolemanagementdirectoryroleassignmentschedulerequestdirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipal                    *rolemanagementdirectoryroleassignmentschedulerequestprincipal.RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinition               *rolemanagementdirectoryroleassignmentschedulerequestroledefinition.RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestTargetSchedule               *rolemanagementdirectoryroleassignmentschedulerequesttargetschedule.RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient
	RoleManagementDirectoryRoleAssignmentScheduleRoleDefinition                      *rolemanagementdirectoryroleassignmentscheduleroledefinition.RoleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient
	RoleManagementDirectoryRoleDefinition                                            *rolemanagementdirectoryroledefinition.RoleManagementDirectoryRoleDefinitionClient
	RoleManagementDirectoryRoleDefinitionInheritsPermissionsFrom                     *rolemanagementdirectoryroledefinitioninheritspermissionsfrom.RoleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient
	RoleManagementDirectoryRoleEligibilitySchedule                                   *rolemanagementdirectoryroleeligibilityschedule.RoleManagementDirectoryRoleEligibilityScheduleClient
	RoleManagementDirectoryRoleEligibilityScheduleAppScope                           *rolemanagementdirectoryroleeligibilityscheduleappscope.RoleManagementDirectoryRoleEligibilityScheduleAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleDirectoryScope                     *rolemanagementdirectoryroleeligibilityscheduledirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstance                           *rolemanagementdirectoryroleeligibilityscheduleinstance.RoleManagementDirectoryRoleEligibilityScheduleInstanceClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScope                   *rolemanagementdirectoryroleeligibilityscheduleinstanceappscope.RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScope             *rolemanagementdirectoryroleeligibilityscheduleinstancedirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipal                  *rolemanagementdirectoryroleeligibilityscheduleinstanceprincipal.RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinition             *rolemanagementdirectoryroleeligibilityscheduleinstanceroledefinition.RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient
	RoleManagementDirectoryRoleEligibilitySchedulePrincipal                          *rolemanagementdirectoryroleeligibilityscheduleprincipal.RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleRequest                            *rolemanagementdirectoryroleeligibilityschedulerequest.RoleManagementDirectoryRoleEligibilityScheduleRequestClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestAppScope                    *rolemanagementdirectoryroleeligibilityschedulerequestappscope.RoleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScope              *rolemanagementdirectoryroleeligibilityschedulerequestdirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipal                   *rolemanagementdirectoryroleeligibilityschedulerequestprincipal.RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinition              *rolemanagementdirectoryroleeligibilityschedulerequestroledefinition.RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestTargetSchedule              *rolemanagementdirectoryroleeligibilityschedulerequesttargetschedule.RoleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient
	RoleManagementDirectoryRoleEligibilityScheduleRoleDefinition                     *rolemanagementdirectoryroleeligibilityscheduleroledefinition.RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient
	RoleManagementEntitlementManagement                                              *rolemanagemententitlementmanagement.RoleManagementEntitlementManagementClient
	RoleManagementEntitlementManagementResourceNamespace                             *rolemanagemententitlementmanagementresourcenamespace.RoleManagementEntitlementManagementResourceNamespaceClient
	RoleManagementEntitlementManagementResourceNamespaceResourceAction               *rolemanagemententitlementmanagementresourcenamespaceresourceaction.RoleManagementEntitlementManagementResourceNamespaceResourceActionClient
	RoleManagementEntitlementManagementRoleAssignment                                *rolemanagemententitlementmanagementroleassignment.RoleManagementEntitlementManagementRoleAssignmentClient
	RoleManagementEntitlementManagementRoleAssignmentAppScope                        *rolemanagemententitlementmanagementroleassignmentappscope.RoleManagementEntitlementManagementRoleAssignmentAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentDirectoryScope                  *rolemanagemententitlementmanagementroleassignmentdirectoryscope.RoleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentPrincipal                       *rolemanagemententitlementmanagementroleassignmentprincipal.RoleManagementEntitlementManagementRoleAssignmentPrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentRoleDefinition                  *rolemanagemententitlementmanagementroleassignmentroledefinition.RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentSchedule                        *rolemanagemententitlementmanagementroleassignmentschedule.RoleManagementEntitlementManagementRoleAssignmentScheduleClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsing          *rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleAppScope                *rolemanagemententitlementmanagementroleassignmentscheduleappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScope          *rolemanagemententitlementmanagementroleassignmentscheduledirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstance                *rolemanagemententitlementmanagementroleassignmentscheduleinstance.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing  *rolemanagemententitlementmanagementroleassignmentscheduleinstanceactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScope        *rolemanagemententitlementmanagementroleassignmentscheduleinstanceappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope  *rolemanagemententitlementmanagementroleassignmentscheduleinstancedirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipal       *rolemanagemententitlementmanagementroleassignmentscheduleinstanceprincipal.RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition  *rolemanagemententitlementmanagementroleassignmentscheduleinstanceroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipal               *rolemanagemententitlementmanagementroleassignmentscheduleprincipal.RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequest                 *rolemanagemententitlementmanagementroleassignmentschedulerequest.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsing   *rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScope         *rolemanagemententitlementmanagementroleassignmentschedulerequestappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScope   *rolemanagemententitlementmanagementroleassignmentschedulerequestdirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipal        *rolemanagemententitlementmanagementroleassignmentschedulerequestprincipal.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinition   *rolemanagemententitlementmanagementroleassignmentschedulerequestroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule   *rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinition          *rolemanagemententitlementmanagementroleassignmentscheduleroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient
	RoleManagementEntitlementManagementRoleDefinition                                *rolemanagemententitlementmanagementroledefinition.RoleManagementEntitlementManagementRoleDefinitionClient
	RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFrom         *rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom.RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient
	RoleManagementEntitlementManagementRoleEligibilitySchedule                       *rolemanagemententitlementmanagementroleeligibilityschedule.RoleManagementEntitlementManagementRoleEligibilityScheduleClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleAppScope               *rolemanagemententitlementmanagementroleeligibilityscheduleappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScope         *rolemanagemententitlementmanagementroleeligibilityscheduledirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstance               *rolemanagemententitlementmanagementroleeligibilityscheduleinstance.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScope       *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope *rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipal      *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceprincipal.RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient
	RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipal              *rolemanagemententitlementmanagementroleeligibilityscheduleprincipal.RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequest                *rolemanagemententitlementmanagementroleeligibilityschedulerequest.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScope        *rolemanagemententitlementmanagementroleeligibilityschedulerequestappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope  *rolemanagemententitlementmanagementroleeligibilityschedulerequestdirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipal       *rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinition  *rolemanagemententitlementmanagementroleeligibilityschedulerequestroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetSchedule  *rolemanagemententitlementmanagementroleeligibilityschedulerequesttargetschedule.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinition         *rolemanagemententitlementmanagementroleeligibilityscheduleroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	roleManagementClient, err := rolemanagement.NewRoleManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagement client: %+v", err)
	}
	configureFunc(roleManagementClient.Client)

	roleManagementDirectoryClient, err := rolemanagementdirectory.NewRoleManagementDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectory client: %+v", err)
	}
	configureFunc(roleManagementDirectoryClient.Client)

	roleManagementDirectoryResourceNamespaceClient, err := rolemanagementdirectoryresourcenamespace.NewRoleManagementDirectoryResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementDirectoryResourceNamespaceClient.Client)

	roleManagementDirectoryResourceNamespaceResourceActionClient, err := rolemanagementdirectoryresourcenamespaceresourceaction.NewRoleManagementDirectoryResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementDirectoryResourceNamespaceResourceActionClient.Client)

	roleManagementDirectoryRoleAssignmentAppScopeClient, err := rolemanagementdirectoryroleassignmentappscope.NewRoleManagementDirectoryRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentAppScopeClient.Client)

	roleManagementDirectoryRoleAssignmentClient, err := rolemanagementdirectoryroleassignment.NewRoleManagementDirectoryRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentClient.Client)

	roleManagementDirectoryRoleAssignmentDirectoryScopeClient, err := rolemanagementdirectoryroleassignmentdirectoryscope.NewRoleManagementDirectoryRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentDirectoryScopeClient.Client)

	roleManagementDirectoryRoleAssignmentPrincipalClient, err := rolemanagementdirectoryroleassignmentprincipal.NewRoleManagementDirectoryRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentPrincipalClient.Client)

	roleManagementDirectoryRoleAssignmentRoleDefinitionClient, err := rolemanagementdirectoryroleassignmentroledefinition.NewRoleManagementDirectoryRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentRoleDefinitionClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient, err := rolemanagementdirectoryroleassignmentscheduleactivatedusing.NewRoleManagementDirectoryRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleAppScopeClient, err := rolemanagementdirectoryroleassignmentscheduleappscope.NewRoleManagementDirectoryRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleAppScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleClient, err := rolemanagementdirectoryroleassignmentschedule.NewRoleManagementDirectoryRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient, err := rolemanagementdirectoryroleassignmentscheduledirectoryscope.NewRoleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient, err := rolemanagementdirectoryroleassignmentscheduleinstanceactivatedusing.NewRoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient, err := rolemanagementdirectoryroleassignmentscheduleinstanceappscope.NewRoleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstanceClient, err := rolemanagementdirectoryroleassignmentscheduleinstance.NewRoleManagementDirectoryRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstanceClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient, err := rolemanagementdirectoryroleassignmentscheduleinstancedirectoryscope.NewRoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient, err := rolemanagementdirectoryroleassignmentscheduleinstanceprincipal.NewRoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient, err := rolemanagementdirectoryroleassignmentscheduleinstanceroledefinition.NewRoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	roleManagementDirectoryRoleAssignmentSchedulePrincipalClient, err := rolemanagementdirectoryroleassignmentscheduleprincipal.NewRoleManagementDirectoryRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentSchedulePrincipalClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient, err := rolemanagementdirectoryroleassignmentschedulerequestactivatedusing.NewRoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient, err := rolemanagementdirectoryroleassignmentschedulerequestappscope.NewRoleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestClient, err := rolemanagementdirectoryroleassignmentschedulerequest.NewRoleManagementDirectoryRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient, err := rolemanagementdirectoryroleassignmentschedulerequestdirectoryscope.NewRoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient, err := rolemanagementdirectoryroleassignmentschedulerequestprincipal.NewRoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient, err := rolemanagementdirectoryroleassignmentschedulerequestroledefinition.NewRoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient, err := rolemanagementdirectoryroleassignmentschedulerequesttargetschedule.NewRoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	roleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient, err := rolemanagementdirectoryroleassignmentscheduleroledefinition.NewRoleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient.Client)

	roleManagementDirectoryRoleDefinitionClient, err := rolemanagementdirectoryroledefinition.NewRoleManagementDirectoryRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleDefinitionClient.Client)

	roleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient, err := rolemanagementdirectoryroledefinitioninheritspermissionsfrom.NewRoleManagementDirectoryRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleAppScopeClient, err := rolemanagementdirectoryroleeligibilityscheduleappscope.NewRoleManagementDirectoryRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleAppScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleClient, err := rolemanagementdirectoryroleeligibilityschedule.NewRoleManagementDirectoryRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient, err := rolemanagementdirectoryroleeligibilityscheduledirectoryscope.NewRoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient, err := rolemanagementdirectoryroleeligibilityscheduleinstanceappscope.NewRoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleInstanceClient, err := rolemanagementdirectoryroleeligibilityscheduleinstance.NewRoleManagementDirectoryRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleInstanceClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient, err := rolemanagementdirectoryroleeligibilityscheduleinstancedirectoryscope.NewRoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient, err := rolemanagementdirectoryroleeligibilityscheduleinstanceprincipal.NewRoleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient, err := rolemanagementdirectoryroleeligibilityscheduleinstanceroledefinition.NewRoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	roleManagementDirectoryRoleEligibilitySchedulePrincipalClient, err := rolemanagementdirectoryroleeligibilityscheduleprincipal.NewRoleManagementDirectoryRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilitySchedulePrincipalClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient, err := rolemanagementdirectoryroleeligibilityschedulerequestappscope.NewRoleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestClient, err := rolemanagementdirectoryroleeligibilityschedulerequest.NewRoleManagementDirectoryRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient, err := rolemanagementdirectoryroleeligibilityschedulerequestdirectoryscope.NewRoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient, err := rolemanagementdirectoryroleeligibilityschedulerequestprincipal.NewRoleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient, err := rolemanagementdirectoryroleeligibilityschedulerequestroledefinition.NewRoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient, err := rolemanagementdirectoryroleeligibilityschedulerequesttargetschedule.NewRoleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	roleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient, err := rolemanagementdirectoryroleeligibilityscheduleroledefinition.NewRoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient.Client)

	roleManagementEntitlementManagementClient, err := rolemanagemententitlementmanagement.NewRoleManagementEntitlementManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagement client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementClient.Client)

	roleManagementEntitlementManagementResourceNamespaceClient, err := rolemanagemententitlementmanagementresourcenamespace.NewRoleManagementEntitlementManagementResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementResourceNamespaceClient.Client)

	roleManagementEntitlementManagementResourceNamespaceResourceActionClient, err := rolemanagemententitlementmanagementresourcenamespaceresourceaction.NewRoleManagementEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementResourceNamespaceResourceActionClient.Client)

	roleManagementEntitlementManagementRoleAssignmentAppScopeClient, err := rolemanagemententitlementmanagementroleassignmentappscope.NewRoleManagementEntitlementManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentAppScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentClient, err := rolemanagemententitlementmanagementroleassignment.NewRoleManagementEntitlementManagementRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentClient.Client)

	roleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient, err := rolemanagemententitlementmanagementroleassignmentdirectoryscope.NewRoleManagementEntitlementManagementRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentPrincipalClient, err := rolemanagemententitlementmanagementroleassignmentprincipal.NewRoleManagementEntitlementManagementRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentPrincipalClient.Client)

	roleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient, err := rolemanagemententitlementmanagementroleassignmentroledefinition.NewRoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient, err := rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing.NewRoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient, err := rolemanagemententitlementmanagementroleassignmentscheduleappscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleClient, err := rolemanagemententitlementmanagementroleassignmentschedule.NewRoleManagementEntitlementManagementRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient, err := rolemanagemententitlementmanagementroleassignmentscheduledirectoryscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstanceactivatedusing.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstanceappscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstance.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstancedirectoryscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstanceprincipal.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient, err := rolemanagemententitlementmanagementroleassignmentscheduleinstanceroledefinition.NewRoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient, err := rolemanagemententitlementmanagementroleassignmentscheduleprincipal.NewRoleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequestappscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequest.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequestdirectoryscope.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequestprincipal.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequestroledefinition.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient, err := rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	roleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient, err := rolemanagemententitlementmanagementroleassignmentscheduleroledefinition.NewRoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleDefinitionClient, err := rolemanagemententitlementmanagementroledefinition.NewRoleManagementEntitlementManagementRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient, err := rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom.NewRoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleappscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleClient, err := rolemanagemententitlementmanagementroleeligibilityschedule.NewRoleManagementEntitlementManagementRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient, err := rolemanagemententitlementmanagementroleeligibilityscheduledirectoryscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleinstanceappscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleinstance.NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleinstanceprincipal.NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition.NewRoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleprincipal.NewRoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequestappscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequest.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequestdirectoryscope.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequestroledefinition.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient, err := rolemanagemententitlementmanagementroleeligibilityschedulerequesttargetschedule.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	roleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient, err := rolemanagemententitlementmanagementroleeligibilityscheduleroledefinition.NewRoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient.Client)

	return &Client{
		RoleManagement:                                                                   roleManagementClient,
		RoleManagementDirectory:                                                          roleManagementDirectoryClient,
		RoleManagementDirectoryResourceNamespace:                                         roleManagementDirectoryResourceNamespaceClient,
		RoleManagementDirectoryResourceNamespaceResourceAction:                           roleManagementDirectoryResourceNamespaceResourceActionClient,
		RoleManagementDirectoryRoleAssignment:                                            roleManagementDirectoryRoleAssignmentClient,
		RoleManagementDirectoryRoleAssignmentAppScope:                                    roleManagementDirectoryRoleAssignmentAppScopeClient,
		RoleManagementDirectoryRoleAssignmentDirectoryScope:                              roleManagementDirectoryRoleAssignmentDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentPrincipal:                                   roleManagementDirectoryRoleAssignmentPrincipalClient,
		RoleManagementDirectoryRoleAssignmentRoleDefinition:                              roleManagementDirectoryRoleAssignmentRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentSchedule:                                    roleManagementDirectoryRoleAssignmentScheduleClient,
		RoleManagementDirectoryRoleAssignmentScheduleActivatedUsing:                      roleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleAppScope:                            roleManagementDirectoryRoleAssignmentScheduleAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleDirectoryScope:                      roleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstance:                            roleManagementDirectoryRoleAssignmentScheduleInstanceClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsing:              roleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScope:                    roleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScope:              roleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipal:                   roleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinition:              roleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentSchedulePrincipal:                           roleManagementDirectoryRoleAssignmentSchedulePrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequest:                             roleManagementDirectoryRoleAssignmentScheduleRequestClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsing:               roleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestAppScope:                     roleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScope:               roleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipal:                    roleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinition:               roleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestTargetSchedule:               roleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient,
		RoleManagementDirectoryRoleAssignmentScheduleRoleDefinition:                      roleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient,
		RoleManagementDirectoryRoleDefinition:                                            roleManagementDirectoryRoleDefinitionClient,
		RoleManagementDirectoryRoleDefinitionInheritsPermissionsFrom:                     roleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementDirectoryRoleEligibilitySchedule:                                   roleManagementDirectoryRoleEligibilityScheduleClient,
		RoleManagementDirectoryRoleEligibilityScheduleAppScope:                           roleManagementDirectoryRoleEligibilityScheduleAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleDirectoryScope:                     roleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstance:                           roleManagementDirectoryRoleEligibilityScheduleInstanceClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScope:                   roleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScope:             roleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipal:                  roleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinition:             roleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient,
		RoleManagementDirectoryRoleEligibilitySchedulePrincipal:                          roleManagementDirectoryRoleEligibilitySchedulePrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequest:                            roleManagementDirectoryRoleEligibilityScheduleRequestClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestAppScope:                    roleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScope:              roleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipal:                   roleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinition:              roleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestTargetSchedule:              roleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient,
		RoleManagementDirectoryRoleEligibilityScheduleRoleDefinition:                     roleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient,
		RoleManagementEntitlementManagement:                                              roleManagementEntitlementManagementClient,
		RoleManagementEntitlementManagementResourceNamespace:                             roleManagementEntitlementManagementResourceNamespaceClient,
		RoleManagementEntitlementManagementResourceNamespaceResourceAction:               roleManagementEntitlementManagementResourceNamespaceResourceActionClient,
		RoleManagementEntitlementManagementRoleAssignment:                                roleManagementEntitlementManagementRoleAssignmentClient,
		RoleManagementEntitlementManagementRoleAssignmentAppScope:                        roleManagementEntitlementManagementRoleAssignmentAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentDirectoryScope:                  roleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentPrincipal:                       roleManagementEntitlementManagementRoleAssignmentPrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentRoleDefinition:                  roleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentSchedule:                        roleManagementEntitlementManagementRoleAssignmentScheduleClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsing:          roleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleAppScope:                roleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScope:          roleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstance:                roleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing:  roleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScope:        roleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope:  roleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipal:       roleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition:  roleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipal:               roleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequest:                 roleManagementEntitlementManagementRoleAssignmentScheduleRequestClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsing:   roleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScope:         roleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScope:   roleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipal:        roleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinition:   roleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule:   roleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinition:          roleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleDefinition:                                roleManagementEntitlementManagementRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFrom:         roleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementEntitlementManagementRoleEligibilitySchedule:                       roleManagementEntitlementManagementRoleEligibilityScheduleClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleAppScope:               roleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScope:         roleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstance:               roleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScope:       roleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope: roleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipal:      roleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition: roleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipal:              roleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequest:                roleManagementEntitlementManagementRoleEligibilityScheduleRequestClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScope:        roleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope:  roleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipal:       roleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinition:  roleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetSchedule:  roleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinition:         roleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient,
	}, nil
}
