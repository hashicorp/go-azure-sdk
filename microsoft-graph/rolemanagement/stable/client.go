package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/directoryroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/entitlementmanagementroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/stable/rolemanagement"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	Directory                                                          *directory.DirectoryClient
	DirectoryResourceNamespace                                         *directoryresourcenamespace.DirectoryResourceNamespaceClient
	DirectoryResourceNamespaceResourceAction                           *directoryresourcenamespaceresourceaction.DirectoryResourceNamespaceResourceActionClient
	DirectoryRoleAssignment                                            *directoryroleassignment.DirectoryRoleAssignmentClient
	DirectoryRoleAssignmentAppScope                                    *directoryroleassignmentappscope.DirectoryRoleAssignmentAppScopeClient
	DirectoryRoleAssignmentDirectoryScope                              *directoryroleassignmentdirectoryscope.DirectoryRoleAssignmentDirectoryScopeClient
	DirectoryRoleAssignmentPrincipal                                   *directoryroleassignmentprincipal.DirectoryRoleAssignmentPrincipalClient
	DirectoryRoleAssignmentRoleDefinition                              *directoryroleassignmentroledefinition.DirectoryRoleAssignmentRoleDefinitionClient
	DirectoryRoleAssignmentSchedule                                    *directoryroleassignmentschedule.DirectoryRoleAssignmentScheduleClient
	DirectoryRoleAssignmentScheduleActivatedUsing                      *directoryroleassignmentscheduleactivatedusing.DirectoryRoleAssignmentScheduleActivatedUsingClient
	DirectoryRoleAssignmentScheduleAppScope                            *directoryroleassignmentscheduleappscope.DirectoryRoleAssignmentScheduleAppScopeClient
	DirectoryRoleAssignmentScheduleDirectoryScope                      *directoryroleassignmentscheduledirectoryscope.DirectoryRoleAssignmentScheduleDirectoryScopeClient
	DirectoryRoleAssignmentScheduleInstance                            *directoryroleassignmentscheduleinstance.DirectoryRoleAssignmentScheduleInstanceClient
	DirectoryRoleAssignmentScheduleInstanceActivatedUsing              *directoryroleassignmentscheduleinstanceactivatedusing.DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient
	DirectoryRoleAssignmentScheduleInstanceAppScope                    *directoryroleassignmentscheduleinstanceappscope.DirectoryRoleAssignmentScheduleInstanceAppScopeClient
	DirectoryRoleAssignmentScheduleInstanceDirectoryScope              *directoryroleassignmentscheduleinstancedirectoryscope.DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient
	DirectoryRoleAssignmentScheduleInstancePrincipal                   *directoryroleassignmentscheduleinstanceprincipal.DirectoryRoleAssignmentScheduleInstancePrincipalClient
	DirectoryRoleAssignmentScheduleInstanceRoleDefinition              *directoryroleassignmentscheduleinstanceroledefinition.DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient
	DirectoryRoleAssignmentSchedulePrincipal                           *directoryroleassignmentscheduleprincipal.DirectoryRoleAssignmentSchedulePrincipalClient
	DirectoryRoleAssignmentScheduleRequest                             *directoryroleassignmentschedulerequest.DirectoryRoleAssignmentScheduleRequestClient
	DirectoryRoleAssignmentScheduleRequestActivatedUsing               *directoryroleassignmentschedulerequestactivatedusing.DirectoryRoleAssignmentScheduleRequestActivatedUsingClient
	DirectoryRoleAssignmentScheduleRequestAppScope                     *directoryroleassignmentschedulerequestappscope.DirectoryRoleAssignmentScheduleRequestAppScopeClient
	DirectoryRoleAssignmentScheduleRequestDirectoryScope               *directoryroleassignmentschedulerequestdirectoryscope.DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient
	DirectoryRoleAssignmentScheduleRequestPrincipal                    *directoryroleassignmentschedulerequestprincipal.DirectoryRoleAssignmentScheduleRequestPrincipalClient
	DirectoryRoleAssignmentScheduleRequestRoleDefinition               *directoryroleassignmentschedulerequestroledefinition.DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient
	DirectoryRoleAssignmentScheduleRequestTargetSchedule               *directoryroleassignmentschedulerequesttargetschedule.DirectoryRoleAssignmentScheduleRequestTargetScheduleClient
	DirectoryRoleAssignmentScheduleRoleDefinition                      *directoryroleassignmentscheduleroledefinition.DirectoryRoleAssignmentScheduleRoleDefinitionClient
	DirectoryRoleDefinition                                            *directoryroledefinition.DirectoryRoleDefinitionClient
	DirectoryRoleDefinitionInheritsPermissionsFrom                     *directoryroledefinitioninheritspermissionsfrom.DirectoryRoleDefinitionInheritsPermissionsFromClient
	DirectoryRoleEligibilitySchedule                                   *directoryroleeligibilityschedule.DirectoryRoleEligibilityScheduleClient
	DirectoryRoleEligibilityScheduleAppScope                           *directoryroleeligibilityscheduleappscope.DirectoryRoleEligibilityScheduleAppScopeClient
	DirectoryRoleEligibilityScheduleDirectoryScope                     *directoryroleeligibilityscheduledirectoryscope.DirectoryRoleEligibilityScheduleDirectoryScopeClient
	DirectoryRoleEligibilityScheduleInstance                           *directoryroleeligibilityscheduleinstance.DirectoryRoleEligibilityScheduleInstanceClient
	DirectoryRoleEligibilityScheduleInstanceAppScope                   *directoryroleeligibilityscheduleinstanceappscope.DirectoryRoleEligibilityScheduleInstanceAppScopeClient
	DirectoryRoleEligibilityScheduleInstanceDirectoryScope             *directoryroleeligibilityscheduleinstancedirectoryscope.DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient
	DirectoryRoleEligibilityScheduleInstancePrincipal                  *directoryroleeligibilityscheduleinstanceprincipal.DirectoryRoleEligibilityScheduleInstancePrincipalClient
	DirectoryRoleEligibilityScheduleInstanceRoleDefinition             *directoryroleeligibilityscheduleinstanceroledefinition.DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient
	DirectoryRoleEligibilitySchedulePrincipal                          *directoryroleeligibilityscheduleprincipal.DirectoryRoleEligibilitySchedulePrincipalClient
	DirectoryRoleEligibilityScheduleRequest                            *directoryroleeligibilityschedulerequest.DirectoryRoleEligibilityScheduleRequestClient
	DirectoryRoleEligibilityScheduleRequestAppScope                    *directoryroleeligibilityschedulerequestappscope.DirectoryRoleEligibilityScheduleRequestAppScopeClient
	DirectoryRoleEligibilityScheduleRequestDirectoryScope              *directoryroleeligibilityschedulerequestdirectoryscope.DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient
	DirectoryRoleEligibilityScheduleRequestPrincipal                   *directoryroleeligibilityschedulerequestprincipal.DirectoryRoleEligibilityScheduleRequestPrincipalClient
	DirectoryRoleEligibilityScheduleRequestRoleDefinition              *directoryroleeligibilityschedulerequestroledefinition.DirectoryRoleEligibilityScheduleRequestRoleDefinitionClient
	DirectoryRoleEligibilityScheduleRequestTargetSchedule              *directoryroleeligibilityschedulerequesttargetschedule.DirectoryRoleEligibilityScheduleRequestTargetScheduleClient
	DirectoryRoleEligibilityScheduleRoleDefinition                     *directoryroleeligibilityscheduleroledefinition.DirectoryRoleEligibilityScheduleRoleDefinitionClient
	EntitlementManagement                                              *entitlementmanagement.EntitlementManagementClient
	EntitlementManagementResourceNamespace                             *entitlementmanagementresourcenamespace.EntitlementManagementResourceNamespaceClient
	EntitlementManagementResourceNamespaceResourceAction               *entitlementmanagementresourcenamespaceresourceaction.EntitlementManagementResourceNamespaceResourceActionClient
	EntitlementManagementRoleAssignment                                *entitlementmanagementroleassignment.EntitlementManagementRoleAssignmentClient
	EntitlementManagementRoleAssignmentAppScope                        *entitlementmanagementroleassignmentappscope.EntitlementManagementRoleAssignmentAppScopeClient
	EntitlementManagementRoleAssignmentDirectoryScope                  *entitlementmanagementroleassignmentdirectoryscope.EntitlementManagementRoleAssignmentDirectoryScopeClient
	EntitlementManagementRoleAssignmentPrincipal                       *entitlementmanagementroleassignmentprincipal.EntitlementManagementRoleAssignmentPrincipalClient
	EntitlementManagementRoleAssignmentRoleDefinition                  *entitlementmanagementroleassignmentroledefinition.EntitlementManagementRoleAssignmentRoleDefinitionClient
	EntitlementManagementRoleAssignmentSchedule                        *entitlementmanagementroleassignmentschedule.EntitlementManagementRoleAssignmentScheduleClient
	EntitlementManagementRoleAssignmentScheduleActivatedUsing          *entitlementmanagementroleassignmentscheduleactivatedusing.EntitlementManagementRoleAssignmentScheduleActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleAppScope                *entitlementmanagementroleassignmentscheduleappscope.EntitlementManagementRoleAssignmentScheduleAppScopeClient
	EntitlementManagementRoleAssignmentScheduleDirectoryScope          *entitlementmanagementroleassignmentscheduledirectoryscope.EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleInstance                *entitlementmanagementroleassignmentscheduleinstance.EntitlementManagementRoleAssignmentScheduleInstanceClient
	EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing  *entitlementmanagementroleassignmentscheduleinstanceactivatedusing.EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleInstanceAppScope        *entitlementmanagementroleassignmentscheduleinstanceappscope.EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient
	EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope  *entitlementmanagementroleassignmentscheduleinstancedirectoryscope.EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleInstancePrincipal       *entitlementmanagementroleassignmentscheduleinstanceprincipal.EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient
	EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition  *entitlementmanagementroleassignmentscheduleinstanceroledefinition.EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient
	EntitlementManagementRoleAssignmentSchedulePrincipal               *entitlementmanagementroleassignmentscheduleprincipal.EntitlementManagementRoleAssignmentSchedulePrincipalClient
	EntitlementManagementRoleAssignmentScheduleRequest                 *entitlementmanagementroleassignmentschedulerequest.EntitlementManagementRoleAssignmentScheduleRequestClient
	EntitlementManagementRoleAssignmentScheduleRequestActivatedUsing   *entitlementmanagementroleassignmentschedulerequestactivatedusing.EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleRequestAppScope         *entitlementmanagementroleassignmentschedulerequestappscope.EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient
	EntitlementManagementRoleAssignmentScheduleRequestDirectoryScope   *entitlementmanagementroleassignmentschedulerequestdirectoryscope.EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleRequestPrincipal        *entitlementmanagementroleassignmentschedulerequestprincipal.EntitlementManagementRoleAssignmentScheduleRequestPrincipalClient
	EntitlementManagementRoleAssignmentScheduleRequestRoleDefinition   *entitlementmanagementroleassignmentschedulerequestroledefinition.EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient
	EntitlementManagementRoleAssignmentScheduleRequestTargetSchedule   *entitlementmanagementroleassignmentschedulerequesttargetschedule.EntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient
	EntitlementManagementRoleAssignmentScheduleRoleDefinition          *entitlementmanagementroleassignmentscheduleroledefinition.EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient
	EntitlementManagementRoleDefinition                                *entitlementmanagementroledefinition.EntitlementManagementRoleDefinitionClient
	EntitlementManagementRoleDefinitionInheritsPermissionsFrom         *entitlementmanagementroledefinitioninheritspermissionsfrom.EntitlementManagementRoleDefinitionInheritsPermissionsFromClient
	EntitlementManagementRoleEligibilitySchedule                       *entitlementmanagementroleeligibilityschedule.EntitlementManagementRoleEligibilityScheduleClient
	EntitlementManagementRoleEligibilityScheduleAppScope               *entitlementmanagementroleeligibilityscheduleappscope.EntitlementManagementRoleEligibilityScheduleAppScopeClient
	EntitlementManagementRoleEligibilityScheduleDirectoryScope         *entitlementmanagementroleeligibilityscheduledirectoryscope.EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleInstance               *entitlementmanagementroleeligibilityscheduleinstance.EntitlementManagementRoleEligibilityScheduleInstanceClient
	EntitlementManagementRoleEligibilityScheduleInstanceAppScope       *entitlementmanagementroleeligibilityscheduleinstanceappscope.EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient
	EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope *entitlementmanagementroleeligibilityscheduleinstancedirectoryscope.EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleInstancePrincipal      *entitlementmanagementroleeligibilityscheduleinstanceprincipal.EntitlementManagementRoleEligibilityScheduleInstancePrincipalClient
	EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition *entitlementmanagementroleeligibilityscheduleinstanceroledefinition.EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient
	EntitlementManagementRoleEligibilitySchedulePrincipal              *entitlementmanagementroleeligibilityscheduleprincipal.EntitlementManagementRoleEligibilitySchedulePrincipalClient
	EntitlementManagementRoleEligibilityScheduleRequest                *entitlementmanagementroleeligibilityschedulerequest.EntitlementManagementRoleEligibilityScheduleRequestClient
	EntitlementManagementRoleEligibilityScheduleRequestAppScope        *entitlementmanagementroleeligibilityschedulerequestappscope.EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient
	EntitlementManagementRoleEligibilityScheduleRequestDirectoryScope  *entitlementmanagementroleeligibilityschedulerequestdirectoryscope.EntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleRequestPrincipal       *entitlementmanagementroleeligibilityschedulerequestprincipal.EntitlementManagementRoleEligibilityScheduleRequestPrincipalClient
	EntitlementManagementRoleEligibilityScheduleRequestRoleDefinition  *entitlementmanagementroleeligibilityschedulerequestroledefinition.EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient
	EntitlementManagementRoleEligibilityScheduleRequestTargetSchedule  *entitlementmanagementroleeligibilityschedulerequesttargetschedule.EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient
	EntitlementManagementRoleEligibilityScheduleRoleDefinition         *entitlementmanagementroleeligibilityscheduleroledefinition.EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient
	RoleManagement                                                     *rolemanagement.RoleManagementClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	directoryClient, err := directory.NewDirectoryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Directory client: %+v", err)
	}
	configureFunc(directoryClient.Client)

	directoryResourceNamespaceClient, err := directoryresourcenamespace.NewDirectoryResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryResourceNamespace client: %+v", err)
	}
	configureFunc(directoryResourceNamespaceClient.Client)

	directoryResourceNamespaceResourceActionClient, err := directoryresourcenamespaceresourceaction.NewDirectoryResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(directoryResourceNamespaceResourceActionClient.Client)

	directoryRoleAssignmentAppScopeClient, err := directoryroleassignmentappscope.NewDirectoryRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentAppScopeClient.Client)

	directoryRoleAssignmentClient, err := directoryroleassignment.NewDirectoryRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignment client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentClient.Client)

	directoryRoleAssignmentDirectoryScopeClient, err := directoryroleassignmentdirectoryscope.NewDirectoryRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentDirectoryScopeClient.Client)

	directoryRoleAssignmentPrincipalClient, err := directoryroleassignmentprincipal.NewDirectoryRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentPrincipalClient.Client)

	directoryRoleAssignmentRoleDefinitionClient, err := directoryroleassignmentroledefinition.NewDirectoryRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentRoleDefinitionClient.Client)

	directoryRoleAssignmentScheduleActivatedUsingClient, err := directoryroleassignmentscheduleactivatedusing.NewDirectoryRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleActivatedUsingClient.Client)

	directoryRoleAssignmentScheduleAppScopeClient, err := directoryroleassignmentscheduleappscope.NewDirectoryRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleAppScopeClient.Client)

	directoryRoleAssignmentScheduleClient, err := directoryroleassignmentschedule.NewDirectoryRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleClient.Client)

	directoryRoleAssignmentScheduleDirectoryScopeClient, err := directoryroleassignmentscheduledirectoryscope.NewDirectoryRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleDirectoryScopeClient.Client)

	directoryRoleAssignmentScheduleInstanceActivatedUsingClient, err := directoryroleassignmentscheduleinstanceactivatedusing.NewDirectoryRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	directoryRoleAssignmentScheduleInstanceAppScopeClient, err := directoryroleassignmentscheduleinstanceappscope.NewDirectoryRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstanceAppScopeClient.Client)

	directoryRoleAssignmentScheduleInstanceClient, err := directoryroleassignmentscheduleinstance.NewDirectoryRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstanceClient.Client)

	directoryRoleAssignmentScheduleInstanceDirectoryScopeClient, err := directoryroleassignmentscheduleinstancedirectoryscope.NewDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	directoryRoleAssignmentScheduleInstancePrincipalClient, err := directoryroleassignmentscheduleinstanceprincipal.NewDirectoryRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstancePrincipalClient.Client)

	directoryRoleAssignmentScheduleInstanceRoleDefinitionClient, err := directoryroleassignmentscheduleinstanceroledefinition.NewDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	directoryRoleAssignmentSchedulePrincipalClient, err := directoryroleassignmentscheduleprincipal.NewDirectoryRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentSchedulePrincipalClient.Client)

	directoryRoleAssignmentScheduleRequestActivatedUsingClient, err := directoryroleassignmentschedulerequestactivatedusing.NewDirectoryRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	directoryRoleAssignmentScheduleRequestAppScopeClient, err := directoryroleassignmentschedulerequestappscope.NewDirectoryRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestAppScopeClient.Client)

	directoryRoleAssignmentScheduleRequestClient, err := directoryroleassignmentschedulerequest.NewDirectoryRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestClient.Client)

	directoryRoleAssignmentScheduleRequestDirectoryScopeClient, err := directoryroleassignmentschedulerequestdirectoryscope.NewDirectoryRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	directoryRoleAssignmentScheduleRequestPrincipalClient, err := directoryroleassignmentschedulerequestprincipal.NewDirectoryRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestPrincipalClient.Client)

	directoryRoleAssignmentScheduleRequestRoleDefinitionClient, err := directoryroleassignmentschedulerequestroledefinition.NewDirectoryRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	directoryRoleAssignmentScheduleRequestTargetScheduleClient, err := directoryroleassignmentschedulerequesttargetschedule.NewDirectoryRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	directoryRoleAssignmentScheduleRoleDefinitionClient, err := directoryroleassignmentscheduleroledefinition.NewDirectoryRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentScheduleRoleDefinitionClient.Client)

	directoryRoleDefinitionClient, err := directoryroledefinition.NewDirectoryRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleDefinitionClient.Client)

	directoryRoleDefinitionInheritsPermissionsFromClient, err := directoryroledefinitioninheritspermissionsfrom.NewDirectoryRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(directoryRoleDefinitionInheritsPermissionsFromClient.Client)

	directoryRoleEligibilityScheduleAppScopeClient, err := directoryroleeligibilityscheduleappscope.NewDirectoryRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleAppScopeClient.Client)

	directoryRoleEligibilityScheduleClient, err := directoryroleeligibilityschedule.NewDirectoryRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleClient.Client)

	directoryRoleEligibilityScheduleDirectoryScopeClient, err := directoryroleeligibilityscheduledirectoryscope.NewDirectoryRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleDirectoryScopeClient.Client)

	directoryRoleEligibilityScheduleInstanceAppScopeClient, err := directoryroleeligibilityscheduleinstanceappscope.NewDirectoryRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleInstanceAppScopeClient.Client)

	directoryRoleEligibilityScheduleInstanceClient, err := directoryroleeligibilityscheduleinstance.NewDirectoryRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleInstanceClient.Client)

	directoryRoleEligibilityScheduleInstanceDirectoryScopeClient, err := directoryroleeligibilityscheduleinstancedirectoryscope.NewDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	directoryRoleEligibilityScheduleInstancePrincipalClient, err := directoryroleeligibilityscheduleinstanceprincipal.NewDirectoryRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleInstancePrincipalClient.Client)

	directoryRoleEligibilityScheduleInstanceRoleDefinitionClient, err := directoryroleeligibilityscheduleinstanceroledefinition.NewDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	directoryRoleEligibilitySchedulePrincipalClient, err := directoryroleeligibilityscheduleprincipal.NewDirectoryRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(directoryRoleEligibilitySchedulePrincipalClient.Client)

	directoryRoleEligibilityScheduleRequestAppScopeClient, err := directoryroleeligibilityschedulerequestappscope.NewDirectoryRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestAppScopeClient.Client)

	directoryRoleEligibilityScheduleRequestClient, err := directoryroleeligibilityschedulerequest.NewDirectoryRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestClient.Client)

	directoryRoleEligibilityScheduleRequestDirectoryScopeClient, err := directoryroleeligibilityschedulerequestdirectoryscope.NewDirectoryRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	directoryRoleEligibilityScheduleRequestPrincipalClient, err := directoryroleeligibilityschedulerequestprincipal.NewDirectoryRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestPrincipalClient.Client)

	directoryRoleEligibilityScheduleRequestRoleDefinitionClient, err := directoryroleeligibilityschedulerequestroledefinition.NewDirectoryRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	directoryRoleEligibilityScheduleRequestTargetScheduleClient, err := directoryroleeligibilityschedulerequesttargetschedule.NewDirectoryRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	directoryRoleEligibilityScheduleRoleDefinitionClient, err := directoryroleeligibilityscheduleroledefinition.NewDirectoryRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(directoryRoleEligibilityScheduleRoleDefinitionClient.Client)

	entitlementManagementClient, err := entitlementmanagement.NewEntitlementManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagement client: %+v", err)
	}
	configureFunc(entitlementManagementClient.Client)

	entitlementManagementResourceNamespaceClient, err := entitlementmanagementresourcenamespace.NewEntitlementManagementResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceNamespace client: %+v", err)
	}
	configureFunc(entitlementManagementResourceNamespaceClient.Client)

	entitlementManagementResourceNamespaceResourceActionClient, err := entitlementmanagementresourcenamespaceresourceaction.NewEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(entitlementManagementResourceNamespaceResourceActionClient.Client)

	entitlementManagementRoleAssignmentAppScopeClient, err := entitlementmanagementroleassignmentappscope.NewEntitlementManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentAppScopeClient.Client)

	entitlementManagementRoleAssignmentClient, err := entitlementmanagementroleassignment.NewEntitlementManagementRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignment client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentClient.Client)

	entitlementManagementRoleAssignmentDirectoryScopeClient, err := entitlementmanagementroleassignmentdirectoryscope.NewEntitlementManagementRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentDirectoryScopeClient.Client)

	entitlementManagementRoleAssignmentPrincipalClient, err := entitlementmanagementroleassignmentprincipal.NewEntitlementManagementRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentPrincipalClient.Client)

	entitlementManagementRoleAssignmentRoleDefinitionClient, err := entitlementmanagementroleassignmentroledefinition.NewEntitlementManagementRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentRoleDefinitionClient.Client)

	entitlementManagementRoleAssignmentScheduleActivatedUsingClient, err := entitlementmanagementroleassignmentscheduleactivatedusing.NewEntitlementManagementRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleActivatedUsingClient.Client)

	entitlementManagementRoleAssignmentScheduleAppScopeClient, err := entitlementmanagementroleassignmentscheduleappscope.NewEntitlementManagementRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleAppScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleClient, err := entitlementmanagementroleassignmentschedule.NewEntitlementManagementRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleClient.Client)

	entitlementManagementRoleAssignmentScheduleDirectoryScopeClient, err := entitlementmanagementroleassignmentscheduledirectoryscope.NewEntitlementManagementRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleDirectoryScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient, err := entitlementmanagementroleassignmentscheduleinstanceactivatedusing.NewEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	entitlementManagementRoleAssignmentScheduleInstanceAppScopeClient, err := entitlementmanagementroleassignmentscheduleinstanceappscope.NewEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstanceAppScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleInstanceClient, err := entitlementmanagementroleassignmentscheduleinstance.NewEntitlementManagementRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstanceClient.Client)

	entitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient, err := entitlementmanagementroleassignmentscheduleinstancedirectoryscope.NewEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleInstancePrincipalClient, err := entitlementmanagementroleassignmentscheduleinstanceprincipal.NewEntitlementManagementRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstancePrincipalClient.Client)

	entitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient, err := entitlementmanagementroleassignmentscheduleinstanceroledefinition.NewEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	entitlementManagementRoleAssignmentSchedulePrincipalClient, err := entitlementmanagementroleassignmentscheduleprincipal.NewEntitlementManagementRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentSchedulePrincipalClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient, err := entitlementmanagementroleassignmentschedulerequestactivatedusing.NewEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestAppScopeClient, err := entitlementmanagementroleassignmentschedulerequestappscope.NewEntitlementManagementRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestAppScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestClient, err := entitlementmanagementroleassignmentschedulerequest.NewEntitlementManagementRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient, err := entitlementmanagementroleassignmentschedulerequestdirectoryscope.NewEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestPrincipalClient, err := entitlementmanagementroleassignmentschedulerequestprincipal.NewEntitlementManagementRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestPrincipalClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient, err := entitlementmanagementroleassignmentschedulerequestroledefinition.NewEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	entitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient, err := entitlementmanagementroleassignmentschedulerequesttargetschedule.NewEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	entitlementManagementRoleAssignmentScheduleRoleDefinitionClient, err := entitlementmanagementroleassignmentscheduleroledefinition.NewEntitlementManagementRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentScheduleRoleDefinitionClient.Client)

	entitlementManagementRoleDefinitionClient, err := entitlementmanagementroledefinition.NewEntitlementManagementRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleDefinitionClient.Client)

	entitlementManagementRoleDefinitionInheritsPermissionsFromClient, err := entitlementmanagementroledefinitioninheritspermissionsfrom.NewEntitlementManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(entitlementManagementRoleDefinitionInheritsPermissionsFromClient.Client)

	entitlementManagementRoleEligibilityScheduleAppScopeClient, err := entitlementmanagementroleeligibilityscheduleappscope.NewEntitlementManagementRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleAppScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleClient, err := entitlementmanagementroleeligibilityschedule.NewEntitlementManagementRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleClient.Client)

	entitlementManagementRoleEligibilityScheduleDirectoryScopeClient, err := entitlementmanagementroleeligibilityscheduledirectoryscope.NewEntitlementManagementRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleDirectoryScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleInstanceAppScopeClient, err := entitlementmanagementroleeligibilityscheduleinstanceappscope.NewEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleInstanceAppScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleInstanceClient, err := entitlementmanagementroleeligibilityscheduleinstance.NewEntitlementManagementRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleInstanceClient.Client)

	entitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient, err := entitlementmanagementroleeligibilityscheduleinstancedirectoryscope.NewEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleInstancePrincipalClient, err := entitlementmanagementroleeligibilityscheduleinstanceprincipal.NewEntitlementManagementRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleInstancePrincipalClient.Client)

	entitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient, err := entitlementmanagementroleeligibilityscheduleinstanceroledefinition.NewEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	entitlementManagementRoleEligibilitySchedulePrincipalClient, err := entitlementmanagementroleeligibilityscheduleprincipal.NewEntitlementManagementRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilitySchedulePrincipalClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestAppScopeClient, err := entitlementmanagementroleeligibilityschedulerequestappscope.NewEntitlementManagementRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestAppScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestClient, err := entitlementmanagementroleeligibilityschedulerequest.NewEntitlementManagementRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient, err := entitlementmanagementroleeligibilityschedulerequestdirectoryscope.NewEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestPrincipalClient, err := entitlementmanagementroleeligibilityschedulerequestprincipal.NewEntitlementManagementRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestPrincipalClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient, err := entitlementmanagementroleeligibilityschedulerequestroledefinition.NewEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	entitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient, err := entitlementmanagementroleeligibilityschedulerequesttargetschedule.NewEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	entitlementManagementRoleEligibilityScheduleRoleDefinitionClient, err := entitlementmanagementroleeligibilityscheduleroledefinition.NewEntitlementManagementRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementRoleEligibilityScheduleRoleDefinitionClient.Client)

	roleManagementClient, err := rolemanagement.NewRoleManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagement client: %+v", err)
	}
	configureFunc(roleManagementClient.Client)

	return &Client{
		Directory:                                                          directoryClient,
		DirectoryResourceNamespace:                                         directoryResourceNamespaceClient,
		DirectoryResourceNamespaceResourceAction:                           directoryResourceNamespaceResourceActionClient,
		DirectoryRoleAssignment:                                            directoryRoleAssignmentClient,
		DirectoryRoleAssignmentAppScope:                                    directoryRoleAssignmentAppScopeClient,
		DirectoryRoleAssignmentDirectoryScope:                              directoryRoleAssignmentDirectoryScopeClient,
		DirectoryRoleAssignmentPrincipal:                                   directoryRoleAssignmentPrincipalClient,
		DirectoryRoleAssignmentRoleDefinition:                              directoryRoleAssignmentRoleDefinitionClient,
		DirectoryRoleAssignmentSchedule:                                    directoryRoleAssignmentScheduleClient,
		DirectoryRoleAssignmentScheduleActivatedUsing:                      directoryRoleAssignmentScheduleActivatedUsingClient,
		DirectoryRoleAssignmentScheduleAppScope:                            directoryRoleAssignmentScheduleAppScopeClient,
		DirectoryRoleAssignmentScheduleDirectoryScope:                      directoryRoleAssignmentScheduleDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleInstance:                            directoryRoleAssignmentScheduleInstanceClient,
		DirectoryRoleAssignmentScheduleInstanceActivatedUsing:              directoryRoleAssignmentScheduleInstanceActivatedUsingClient,
		DirectoryRoleAssignmentScheduleInstanceAppScope:                    directoryRoleAssignmentScheduleInstanceAppScopeClient,
		DirectoryRoleAssignmentScheduleInstanceDirectoryScope:              directoryRoleAssignmentScheduleInstanceDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleInstancePrincipal:                   directoryRoleAssignmentScheduleInstancePrincipalClient,
		DirectoryRoleAssignmentScheduleInstanceRoleDefinition:              directoryRoleAssignmentScheduleInstanceRoleDefinitionClient,
		DirectoryRoleAssignmentSchedulePrincipal:                           directoryRoleAssignmentSchedulePrincipalClient,
		DirectoryRoleAssignmentScheduleRequest:                             directoryRoleAssignmentScheduleRequestClient,
		DirectoryRoleAssignmentScheduleRequestActivatedUsing:               directoryRoleAssignmentScheduleRequestActivatedUsingClient,
		DirectoryRoleAssignmentScheduleRequestAppScope:                     directoryRoleAssignmentScheduleRequestAppScopeClient,
		DirectoryRoleAssignmentScheduleRequestDirectoryScope:               directoryRoleAssignmentScheduleRequestDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleRequestPrincipal:                    directoryRoleAssignmentScheduleRequestPrincipalClient,
		DirectoryRoleAssignmentScheduleRequestRoleDefinition:               directoryRoleAssignmentScheduleRequestRoleDefinitionClient,
		DirectoryRoleAssignmentScheduleRequestTargetSchedule:               directoryRoleAssignmentScheduleRequestTargetScheduleClient,
		DirectoryRoleAssignmentScheduleRoleDefinition:                      directoryRoleAssignmentScheduleRoleDefinitionClient,
		DirectoryRoleDefinition:                                            directoryRoleDefinitionClient,
		DirectoryRoleDefinitionInheritsPermissionsFrom:                     directoryRoleDefinitionInheritsPermissionsFromClient,
		DirectoryRoleEligibilitySchedule:                                   directoryRoleEligibilityScheduleClient,
		DirectoryRoleEligibilityScheduleAppScope:                           directoryRoleEligibilityScheduleAppScopeClient,
		DirectoryRoleEligibilityScheduleDirectoryScope:                     directoryRoleEligibilityScheduleDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleInstance:                           directoryRoleEligibilityScheduleInstanceClient,
		DirectoryRoleEligibilityScheduleInstanceAppScope:                   directoryRoleEligibilityScheduleInstanceAppScopeClient,
		DirectoryRoleEligibilityScheduleInstanceDirectoryScope:             directoryRoleEligibilityScheduleInstanceDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleInstancePrincipal:                  directoryRoleEligibilityScheduleInstancePrincipalClient,
		DirectoryRoleEligibilityScheduleInstanceRoleDefinition:             directoryRoleEligibilityScheduleInstanceRoleDefinitionClient,
		DirectoryRoleEligibilitySchedulePrincipal:                          directoryRoleEligibilitySchedulePrincipalClient,
		DirectoryRoleEligibilityScheduleRequest:                            directoryRoleEligibilityScheduleRequestClient,
		DirectoryRoleEligibilityScheduleRequestAppScope:                    directoryRoleEligibilityScheduleRequestAppScopeClient,
		DirectoryRoleEligibilityScheduleRequestDirectoryScope:              directoryRoleEligibilityScheduleRequestDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleRequestPrincipal:                   directoryRoleEligibilityScheduleRequestPrincipalClient,
		DirectoryRoleEligibilityScheduleRequestRoleDefinition:              directoryRoleEligibilityScheduleRequestRoleDefinitionClient,
		DirectoryRoleEligibilityScheduleRequestTargetSchedule:              directoryRoleEligibilityScheduleRequestTargetScheduleClient,
		DirectoryRoleEligibilityScheduleRoleDefinition:                     directoryRoleEligibilityScheduleRoleDefinitionClient,
		EntitlementManagement:                                              entitlementManagementClient,
		EntitlementManagementResourceNamespace:                             entitlementManagementResourceNamespaceClient,
		EntitlementManagementResourceNamespaceResourceAction:               entitlementManagementResourceNamespaceResourceActionClient,
		EntitlementManagementRoleAssignment:                                entitlementManagementRoleAssignmentClient,
		EntitlementManagementRoleAssignmentAppScope:                        entitlementManagementRoleAssignmentAppScopeClient,
		EntitlementManagementRoleAssignmentDirectoryScope:                  entitlementManagementRoleAssignmentDirectoryScopeClient,
		EntitlementManagementRoleAssignmentPrincipal:                       entitlementManagementRoleAssignmentPrincipalClient,
		EntitlementManagementRoleAssignmentRoleDefinition:                  entitlementManagementRoleAssignmentRoleDefinitionClient,
		EntitlementManagementRoleAssignmentSchedule:                        entitlementManagementRoleAssignmentScheduleClient,
		EntitlementManagementRoleAssignmentScheduleActivatedUsing:          entitlementManagementRoleAssignmentScheduleActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleAppScope:                entitlementManagementRoleAssignmentScheduleAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleDirectoryScope:          entitlementManagementRoleAssignmentScheduleDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstance:                entitlementManagementRoleAssignmentScheduleInstanceClient,
		EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing:  entitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleInstanceAppScope:        entitlementManagementRoleAssignmentScheduleInstanceAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope:  entitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstancePrincipal:       entitlementManagementRoleAssignmentScheduleInstancePrincipalClient,
		EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition:  entitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient,
		EntitlementManagementRoleAssignmentSchedulePrincipal:               entitlementManagementRoleAssignmentSchedulePrincipalClient,
		EntitlementManagementRoleAssignmentScheduleRequest:                 entitlementManagementRoleAssignmentScheduleRequestClient,
		EntitlementManagementRoleAssignmentScheduleRequestActivatedUsing:   entitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleRequestAppScope:         entitlementManagementRoleAssignmentScheduleRequestAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleRequestDirectoryScope:   entitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleRequestPrincipal:        entitlementManagementRoleAssignmentScheduleRequestPrincipalClient,
		EntitlementManagementRoleAssignmentScheduleRequestRoleDefinition:   entitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient,
		EntitlementManagementRoleAssignmentScheduleRequestTargetSchedule:   entitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient,
		EntitlementManagementRoleAssignmentScheduleRoleDefinition:          entitlementManagementRoleAssignmentScheduleRoleDefinitionClient,
		EntitlementManagementRoleDefinition:                                entitlementManagementRoleDefinitionClient,
		EntitlementManagementRoleDefinitionInheritsPermissionsFrom:         entitlementManagementRoleDefinitionInheritsPermissionsFromClient,
		EntitlementManagementRoleEligibilitySchedule:                       entitlementManagementRoleEligibilityScheduleClient,
		EntitlementManagementRoleEligibilityScheduleAppScope:               entitlementManagementRoleEligibilityScheduleAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleDirectoryScope:         entitlementManagementRoleEligibilityScheduleDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstance:               entitlementManagementRoleEligibilityScheduleInstanceClient,
		EntitlementManagementRoleEligibilityScheduleInstanceAppScope:       entitlementManagementRoleEligibilityScheduleInstanceAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope: entitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstancePrincipal:      entitlementManagementRoleEligibilityScheduleInstancePrincipalClient,
		EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition: entitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient,
		EntitlementManagementRoleEligibilitySchedulePrincipal:              entitlementManagementRoleEligibilitySchedulePrincipalClient,
		EntitlementManagementRoleEligibilityScheduleRequest:                entitlementManagementRoleEligibilityScheduleRequestClient,
		EntitlementManagementRoleEligibilityScheduleRequestAppScope:        entitlementManagementRoleEligibilityScheduleRequestAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleRequestDirectoryScope:  entitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleRequestPrincipal:       entitlementManagementRoleEligibilityScheduleRequestPrincipalClient,
		EntitlementManagementRoleEligibilityScheduleRequestRoleDefinition:  entitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient,
		EntitlementManagementRoleEligibilityScheduleRequestTargetSchedule:  entitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient,
		EntitlementManagementRoleEligibilityScheduleRoleDefinition:         entitlementManagementRoleEligibilityScheduleRoleDefinitionClient,
		RoleManagement: roleManagementClient,
	}, nil
}
