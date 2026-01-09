package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/cloudpcroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defender"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/defenderroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/devicemanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directoryroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directorytransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directorytransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directorytransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directorytransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/directorytransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseappresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapproleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapptransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapptransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapptransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapptransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/enterpriseapptransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementtransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementtransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementtransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementtransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementtransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchange"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangecustomappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangeroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangetransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangetransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangetransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangetransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/exchangetransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagement"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CloudPC                                                                   *cloudpc.CloudPCClient
	CloudPCResourceNamespace                                                  *cloudpcresourcenamespace.CloudPCResourceNamespaceClient
	CloudPCResourceNamespaceResourceAction                                    *cloudpcresourcenamespaceresourceaction.CloudPCResourceNamespaceResourceActionClient
	CloudPCResourceNamespaceResourceActionAuthenticationContext               *cloudpcresourcenamespaceresourceactionauthenticationcontext.CloudPCResourceNamespaceResourceActionAuthenticationContextClient
	CloudPCResourceNamespaceResourceActionResourceScope                       *cloudpcresourcenamespaceresourceactionresourcescope.CloudPCResourceNamespaceResourceActionResourceScopeClient
	CloudPCRoleAssignment                                                     *cloudpcroleassignment.CloudPCRoleAssignmentClient
	CloudPCRoleAssignmentAppScope                                             *cloudpcroleassignmentappscope.CloudPCRoleAssignmentAppScopeClient
	CloudPCRoleAssignmentDirectoryScope                                       *cloudpcroleassignmentdirectoryscope.CloudPCRoleAssignmentDirectoryScopeClient
	CloudPCRoleAssignmentPrincipal                                            *cloudpcroleassignmentprincipal.CloudPCRoleAssignmentPrincipalClient
	CloudPCRoleAssignmentRoleDefinition                                       *cloudpcroleassignmentroledefinition.CloudPCRoleAssignmentRoleDefinitionClient
	CloudPCRoleDefinition                                                     *cloudpcroledefinition.CloudPCRoleDefinitionClient
	CloudPCRoleDefinitionInheritsPermissionsFrom                              *cloudpcroledefinitioninheritspermissionsfrom.CloudPCRoleDefinitionInheritsPermissionsFromClient
	Defender                                                                  *defender.DefenderClient
	DefenderResourceNamespace                                                 *defenderresourcenamespace.DefenderResourceNamespaceClient
	DefenderResourceNamespaceResourceAction                                   *defenderresourcenamespaceresourceaction.DefenderResourceNamespaceResourceActionClient
	DefenderResourceNamespaceResourceActionAuthenticationContext              *defenderresourcenamespaceresourceactionauthenticationcontext.DefenderResourceNamespaceResourceActionAuthenticationContextClient
	DefenderResourceNamespaceResourceActionResourceScope                      *defenderresourcenamespaceresourceactionresourcescope.DefenderResourceNamespaceResourceActionResourceScopeClient
	DefenderRoleAssignment                                                    *defenderroleassignment.DefenderRoleAssignmentClient
	DefenderRoleAssignmentAppScope                                            *defenderroleassignmentappscope.DefenderRoleAssignmentAppScopeClient
	DefenderRoleAssignmentDirectoryScope                                      *defenderroleassignmentdirectoryscope.DefenderRoleAssignmentDirectoryScopeClient
	DefenderRoleAssignmentPrincipal                                           *defenderroleassignmentprincipal.DefenderRoleAssignmentPrincipalClient
	DefenderRoleAssignmentRoleDefinition                                      *defenderroleassignmentroledefinition.DefenderRoleAssignmentRoleDefinitionClient
	DefenderRoleDefinition                                                    *defenderroledefinition.DefenderRoleDefinitionClient
	DefenderRoleDefinitionInheritsPermissionsFrom                             *defenderroledefinitioninheritspermissionsfrom.DefenderRoleDefinitionInheritsPermissionsFromClient
	DeviceManagement                                                          *devicemanagement.DeviceManagementClient
	DeviceManagementResourceNamespace                                         *devicemanagementresourcenamespace.DeviceManagementResourceNamespaceClient
	DeviceManagementResourceNamespaceResourceAction                           *devicemanagementresourcenamespaceresourceaction.DeviceManagementResourceNamespaceResourceActionClient
	DeviceManagementResourceNamespaceResourceActionAuthenticationContext      *devicemanagementresourcenamespaceresourceactionauthenticationcontext.DeviceManagementResourceNamespaceResourceActionAuthenticationContextClient
	DeviceManagementResourceNamespaceResourceActionResourceScope              *devicemanagementresourcenamespaceresourceactionresourcescope.DeviceManagementResourceNamespaceResourceActionResourceScopeClient
	DeviceManagementRoleAssignment                                            *devicemanagementroleassignment.DeviceManagementRoleAssignmentClient
	DeviceManagementRoleAssignmentAppScope                                    *devicemanagementroleassignmentappscope.DeviceManagementRoleAssignmentAppScopeClient
	DeviceManagementRoleAssignmentDirectoryScope                              *devicemanagementroleassignmentdirectoryscope.DeviceManagementRoleAssignmentDirectoryScopeClient
	DeviceManagementRoleAssignmentPrincipal                                   *devicemanagementroleassignmentprincipal.DeviceManagementRoleAssignmentPrincipalClient
	DeviceManagementRoleAssignmentRoleDefinition                              *devicemanagementroleassignmentroledefinition.DeviceManagementRoleAssignmentRoleDefinitionClient
	DeviceManagementRoleDefinition                                            *devicemanagementroledefinition.DeviceManagementRoleDefinitionClient
	DeviceManagementRoleDefinitionInheritsPermissionsFrom                     *devicemanagementroledefinitioninheritspermissionsfrom.DeviceManagementRoleDefinitionInheritsPermissionsFromClient
	Directory                                                                 *directory.DirectoryClient
	DirectoryResourceNamespace                                                *directoryresourcenamespace.DirectoryResourceNamespaceClient
	DirectoryResourceNamespaceResourceAction                                  *directoryresourcenamespaceresourceaction.DirectoryResourceNamespaceResourceActionClient
	DirectoryResourceNamespaceResourceActionAuthenticationContext             *directoryresourcenamespaceresourceactionauthenticationcontext.DirectoryResourceNamespaceResourceActionAuthenticationContextClient
	DirectoryResourceNamespaceResourceActionResourceScope                     *directoryresourcenamespaceresourceactionresourcescope.DirectoryResourceNamespaceResourceActionResourceScopeClient
	DirectoryRoleAssignment                                                   *directoryroleassignment.DirectoryRoleAssignmentClient
	DirectoryRoleAssignmentAppScope                                           *directoryroleassignmentappscope.DirectoryRoleAssignmentAppScopeClient
	DirectoryRoleAssignmentApproval                                           *directoryroleassignmentapproval.DirectoryRoleAssignmentApprovalClient
	DirectoryRoleAssignmentApprovalStep                                       *directoryroleassignmentapprovalstep.DirectoryRoleAssignmentApprovalStepClient
	DirectoryRoleAssignmentDirectoryScope                                     *directoryroleassignmentdirectoryscope.DirectoryRoleAssignmentDirectoryScopeClient
	DirectoryRoleAssignmentPrincipal                                          *directoryroleassignmentprincipal.DirectoryRoleAssignmentPrincipalClient
	DirectoryRoleAssignmentRoleDefinition                                     *directoryroleassignmentroledefinition.DirectoryRoleAssignmentRoleDefinitionClient
	DirectoryRoleAssignmentSchedule                                           *directoryroleassignmentschedule.DirectoryRoleAssignmentScheduleClient
	DirectoryRoleAssignmentScheduleActivatedUsing                             *directoryroleassignmentscheduleactivatedusing.DirectoryRoleAssignmentScheduleActivatedUsingClient
	DirectoryRoleAssignmentScheduleAppScope                                   *directoryroleassignmentscheduleappscope.DirectoryRoleAssignmentScheduleAppScopeClient
	DirectoryRoleAssignmentScheduleDirectoryScope                             *directoryroleassignmentscheduledirectoryscope.DirectoryRoleAssignmentScheduleDirectoryScopeClient
	DirectoryRoleAssignmentScheduleInstance                                   *directoryroleassignmentscheduleinstance.DirectoryRoleAssignmentScheduleInstanceClient
	DirectoryRoleAssignmentScheduleInstanceActivatedUsing                     *directoryroleassignmentscheduleinstanceactivatedusing.DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient
	DirectoryRoleAssignmentScheduleInstanceAppScope                           *directoryroleassignmentscheduleinstanceappscope.DirectoryRoleAssignmentScheduleInstanceAppScopeClient
	DirectoryRoleAssignmentScheduleInstanceDirectoryScope                     *directoryroleassignmentscheduleinstancedirectoryscope.DirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient
	DirectoryRoleAssignmentScheduleInstancePrincipal                          *directoryroleassignmentscheduleinstanceprincipal.DirectoryRoleAssignmentScheduleInstancePrincipalClient
	DirectoryRoleAssignmentScheduleInstanceRoleDefinition                     *directoryroleassignmentscheduleinstanceroledefinition.DirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient
	DirectoryRoleAssignmentSchedulePrincipal                                  *directoryroleassignmentscheduleprincipal.DirectoryRoleAssignmentSchedulePrincipalClient
	DirectoryRoleAssignmentScheduleRequest                                    *directoryroleassignmentschedulerequest.DirectoryRoleAssignmentScheduleRequestClient
	DirectoryRoleAssignmentScheduleRequestActivatedUsing                      *directoryroleassignmentschedulerequestactivatedusing.DirectoryRoleAssignmentScheduleRequestActivatedUsingClient
	DirectoryRoleAssignmentScheduleRequestAppScope                            *directoryroleassignmentschedulerequestappscope.DirectoryRoleAssignmentScheduleRequestAppScopeClient
	DirectoryRoleAssignmentScheduleRequestDirectoryScope                      *directoryroleassignmentschedulerequestdirectoryscope.DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient
	DirectoryRoleAssignmentScheduleRequestPrincipal                           *directoryroleassignmentschedulerequestprincipal.DirectoryRoleAssignmentScheduleRequestPrincipalClient
	DirectoryRoleAssignmentScheduleRequestRoleDefinition                      *directoryroleassignmentschedulerequestroledefinition.DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient
	DirectoryRoleAssignmentScheduleRequestTargetSchedule                      *directoryroleassignmentschedulerequesttargetschedule.DirectoryRoleAssignmentScheduleRequestTargetScheduleClient
	DirectoryRoleAssignmentScheduleRoleDefinition                             *directoryroleassignmentscheduleroledefinition.DirectoryRoleAssignmentScheduleRoleDefinitionClient
	DirectoryRoleDefinition                                                   *directoryroledefinition.DirectoryRoleDefinitionClient
	DirectoryRoleDefinitionInheritsPermissionsFrom                            *directoryroledefinitioninheritspermissionsfrom.DirectoryRoleDefinitionInheritsPermissionsFromClient
	DirectoryRoleEligibilitySchedule                                          *directoryroleeligibilityschedule.DirectoryRoleEligibilityScheduleClient
	DirectoryRoleEligibilityScheduleAppScope                                  *directoryroleeligibilityscheduleappscope.DirectoryRoleEligibilityScheduleAppScopeClient
	DirectoryRoleEligibilityScheduleDirectoryScope                            *directoryroleeligibilityscheduledirectoryscope.DirectoryRoleEligibilityScheduleDirectoryScopeClient
	DirectoryRoleEligibilityScheduleInstance                                  *directoryroleeligibilityscheduleinstance.DirectoryRoleEligibilityScheduleInstanceClient
	DirectoryRoleEligibilityScheduleInstanceAppScope                          *directoryroleeligibilityscheduleinstanceappscope.DirectoryRoleEligibilityScheduleInstanceAppScopeClient
	DirectoryRoleEligibilityScheduleInstanceDirectoryScope                    *directoryroleeligibilityscheduleinstancedirectoryscope.DirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient
	DirectoryRoleEligibilityScheduleInstancePrincipal                         *directoryroleeligibilityscheduleinstanceprincipal.DirectoryRoleEligibilityScheduleInstancePrincipalClient
	DirectoryRoleEligibilityScheduleInstanceRoleDefinition                    *directoryroleeligibilityscheduleinstanceroledefinition.DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient
	DirectoryRoleEligibilitySchedulePrincipal                                 *directoryroleeligibilityscheduleprincipal.DirectoryRoleEligibilitySchedulePrincipalClient
	DirectoryRoleEligibilityScheduleRequest                                   *directoryroleeligibilityschedulerequest.DirectoryRoleEligibilityScheduleRequestClient
	DirectoryRoleEligibilityScheduleRequestAppScope                           *directoryroleeligibilityschedulerequestappscope.DirectoryRoleEligibilityScheduleRequestAppScopeClient
	DirectoryRoleEligibilityScheduleRequestDirectoryScope                     *directoryroleeligibilityschedulerequestdirectoryscope.DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient
	DirectoryRoleEligibilityScheduleRequestPrincipal                          *directoryroleeligibilityschedulerequestprincipal.DirectoryRoleEligibilityScheduleRequestPrincipalClient
	DirectoryRoleEligibilityScheduleRequestRoleDefinition                     *directoryroleeligibilityschedulerequestroledefinition.DirectoryRoleEligibilityScheduleRequestRoleDefinitionClient
	DirectoryRoleEligibilityScheduleRequestTargetSchedule                     *directoryroleeligibilityschedulerequesttargetschedule.DirectoryRoleEligibilityScheduleRequestTargetScheduleClient
	DirectoryRoleEligibilityScheduleRoleDefinition                            *directoryroleeligibilityscheduleroledefinition.DirectoryRoleEligibilityScheduleRoleDefinitionClient
	DirectoryTransitiveRoleAssignment                                         *directorytransitiveroleassignment.DirectoryTransitiveRoleAssignmentClient
	DirectoryTransitiveRoleAssignmentAppScope                                 *directorytransitiveroleassignmentappscope.DirectoryTransitiveRoleAssignmentAppScopeClient
	DirectoryTransitiveRoleAssignmentDirectoryScope                           *directorytransitiveroleassignmentdirectoryscope.DirectoryTransitiveRoleAssignmentDirectoryScopeClient
	DirectoryTransitiveRoleAssignmentPrincipal                                *directorytransitiveroleassignmentprincipal.DirectoryTransitiveRoleAssignmentPrincipalClient
	DirectoryTransitiveRoleAssignmentRoleDefinition                           *directorytransitiveroleassignmentroledefinition.DirectoryTransitiveRoleAssignmentRoleDefinitionClient
	EnterpriseApp                                                             *enterpriseapp.EnterpriseAppClient
	EnterpriseAppResourceNamespace                                            *enterpriseappresourcenamespace.EnterpriseAppResourceNamespaceClient
	EnterpriseAppResourceNamespaceResourceAction                              *enterpriseappresourcenamespaceresourceaction.EnterpriseAppResourceNamespaceResourceActionClient
	EnterpriseAppResourceNamespaceResourceActionAuthenticationContext         *enterpriseappresourcenamespaceresourceactionauthenticationcontext.EnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient
	EnterpriseAppResourceNamespaceResourceActionResourceScope                 *enterpriseappresourcenamespaceresourceactionresourcescope.EnterpriseAppResourceNamespaceResourceActionResourceScopeClient
	EnterpriseAppRoleAssignment                                               *enterpriseapproleassignment.EnterpriseAppRoleAssignmentClient
	EnterpriseAppRoleAssignmentAppScope                                       *enterpriseapproleassignmentappscope.EnterpriseAppRoleAssignmentAppScopeClient
	EnterpriseAppRoleAssignmentApproval                                       *enterpriseapproleassignmentapproval.EnterpriseAppRoleAssignmentApprovalClient
	EnterpriseAppRoleAssignmentApprovalStep                                   *enterpriseapproleassignmentapprovalstep.EnterpriseAppRoleAssignmentApprovalStepClient
	EnterpriseAppRoleAssignmentDirectoryScope                                 *enterpriseapproleassignmentdirectoryscope.EnterpriseAppRoleAssignmentDirectoryScopeClient
	EnterpriseAppRoleAssignmentPrincipal                                      *enterpriseapproleassignmentprincipal.EnterpriseAppRoleAssignmentPrincipalClient
	EnterpriseAppRoleAssignmentRoleDefinition                                 *enterpriseapproleassignmentroledefinition.EnterpriseAppRoleAssignmentRoleDefinitionClient
	EnterpriseAppRoleAssignmentSchedule                                       *enterpriseapproleassignmentschedule.EnterpriseAppRoleAssignmentScheduleClient
	EnterpriseAppRoleAssignmentScheduleActivatedUsing                         *enterpriseapproleassignmentscheduleactivatedusing.EnterpriseAppRoleAssignmentScheduleActivatedUsingClient
	EnterpriseAppRoleAssignmentScheduleAppScope                               *enterpriseapproleassignmentscheduleappscope.EnterpriseAppRoleAssignmentScheduleAppScopeClient
	EnterpriseAppRoleAssignmentScheduleDirectoryScope                         *enterpriseapproleassignmentscheduledirectoryscope.EnterpriseAppRoleAssignmentScheduleDirectoryScopeClient
	EnterpriseAppRoleAssignmentScheduleInstance                               *enterpriseapproleassignmentscheduleinstance.EnterpriseAppRoleAssignmentScheduleInstanceClient
	EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing                 *enterpriseapproleassignmentscheduleinstanceactivatedusing.EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient
	EnterpriseAppRoleAssignmentScheduleInstanceAppScope                       *enterpriseapproleassignmentscheduleinstanceappscope.EnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient
	EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope                 *enterpriseapproleassignmentscheduleinstancedirectoryscope.EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient
	EnterpriseAppRoleAssignmentScheduleInstancePrincipal                      *enterpriseapproleassignmentscheduleinstanceprincipal.EnterpriseAppRoleAssignmentScheduleInstancePrincipalClient
	EnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition                 *enterpriseapproleassignmentscheduleinstanceroledefinition.EnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient
	EnterpriseAppRoleAssignmentSchedulePrincipal                              *enterpriseapproleassignmentscheduleprincipal.EnterpriseAppRoleAssignmentSchedulePrincipalClient
	EnterpriseAppRoleAssignmentScheduleRequest                                *enterpriseapproleassignmentschedulerequest.EnterpriseAppRoleAssignmentScheduleRequestClient
	EnterpriseAppRoleAssignmentScheduleRequestActivatedUsing                  *enterpriseapproleassignmentschedulerequestactivatedusing.EnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient
	EnterpriseAppRoleAssignmentScheduleRequestAppScope                        *enterpriseapproleassignmentschedulerequestappscope.EnterpriseAppRoleAssignmentScheduleRequestAppScopeClient
	EnterpriseAppRoleAssignmentScheduleRequestDirectoryScope                  *enterpriseapproleassignmentschedulerequestdirectoryscope.EnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient
	EnterpriseAppRoleAssignmentScheduleRequestPrincipal                       *enterpriseapproleassignmentschedulerequestprincipal.EnterpriseAppRoleAssignmentScheduleRequestPrincipalClient
	EnterpriseAppRoleAssignmentScheduleRequestRoleDefinition                  *enterpriseapproleassignmentschedulerequestroledefinition.EnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient
	EnterpriseAppRoleAssignmentScheduleRequestTargetSchedule                  *enterpriseapproleassignmentschedulerequesttargetschedule.EnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient
	EnterpriseAppRoleAssignmentScheduleRoleDefinition                         *enterpriseapproleassignmentscheduleroledefinition.EnterpriseAppRoleAssignmentScheduleRoleDefinitionClient
	EnterpriseAppRoleDefinition                                               *enterpriseapproledefinition.EnterpriseAppRoleDefinitionClient
	EnterpriseAppRoleDefinitionInheritsPermissionsFrom                        *enterpriseapproledefinitioninheritspermissionsfrom.EnterpriseAppRoleDefinitionInheritsPermissionsFromClient
	EnterpriseAppRoleEligibilitySchedule                                      *enterpriseapproleeligibilityschedule.EnterpriseAppRoleEligibilityScheduleClient
	EnterpriseAppRoleEligibilityScheduleAppScope                              *enterpriseapproleeligibilityscheduleappscope.EnterpriseAppRoleEligibilityScheduleAppScopeClient
	EnterpriseAppRoleEligibilityScheduleDirectoryScope                        *enterpriseapproleeligibilityscheduledirectoryscope.EnterpriseAppRoleEligibilityScheduleDirectoryScopeClient
	EnterpriseAppRoleEligibilityScheduleInstance                              *enterpriseapproleeligibilityscheduleinstance.EnterpriseAppRoleEligibilityScheduleInstanceClient
	EnterpriseAppRoleEligibilityScheduleInstanceAppScope                      *enterpriseapproleeligibilityscheduleinstanceappscope.EnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient
	EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope                *enterpriseapproleeligibilityscheduleinstancedirectoryscope.EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient
	EnterpriseAppRoleEligibilityScheduleInstancePrincipal                     *enterpriseapproleeligibilityscheduleinstanceprincipal.EnterpriseAppRoleEligibilityScheduleInstancePrincipalClient
	EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition                *enterpriseapproleeligibilityscheduleinstanceroledefinition.EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient
	EnterpriseAppRoleEligibilitySchedulePrincipal                             *enterpriseapproleeligibilityscheduleprincipal.EnterpriseAppRoleEligibilitySchedulePrincipalClient
	EnterpriseAppRoleEligibilityScheduleRequest                               *enterpriseapproleeligibilityschedulerequest.EnterpriseAppRoleEligibilityScheduleRequestClient
	EnterpriseAppRoleEligibilityScheduleRequestAppScope                       *enterpriseapproleeligibilityschedulerequestappscope.EnterpriseAppRoleEligibilityScheduleRequestAppScopeClient
	EnterpriseAppRoleEligibilityScheduleRequestDirectoryScope                 *enterpriseapproleeligibilityschedulerequestdirectoryscope.EnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient
	EnterpriseAppRoleEligibilityScheduleRequestPrincipal                      *enterpriseapproleeligibilityschedulerequestprincipal.EnterpriseAppRoleEligibilityScheduleRequestPrincipalClient
	EnterpriseAppRoleEligibilityScheduleRequestRoleDefinition                 *enterpriseapproleeligibilityschedulerequestroledefinition.EnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient
	EnterpriseAppRoleEligibilityScheduleRequestTargetSchedule                 *enterpriseapproleeligibilityschedulerequesttargetschedule.EnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient
	EnterpriseAppRoleEligibilityScheduleRoleDefinition                        *enterpriseapproleeligibilityscheduleroledefinition.EnterpriseAppRoleEligibilityScheduleRoleDefinitionClient
	EnterpriseAppTransitiveRoleAssignment                                     *enterpriseapptransitiveroleassignment.EnterpriseAppTransitiveRoleAssignmentClient
	EnterpriseAppTransitiveRoleAssignmentAppScope                             *enterpriseapptransitiveroleassignmentappscope.EnterpriseAppTransitiveRoleAssignmentAppScopeClient
	EnterpriseAppTransitiveRoleAssignmentDirectoryScope                       *enterpriseapptransitiveroleassignmentdirectoryscope.EnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient
	EnterpriseAppTransitiveRoleAssignmentPrincipal                            *enterpriseapptransitiveroleassignmentprincipal.EnterpriseAppTransitiveRoleAssignmentPrincipalClient
	EnterpriseAppTransitiveRoleAssignmentRoleDefinition                       *enterpriseapptransitiveroleassignmentroledefinition.EnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient
	EntitlementManagement                                                     *entitlementmanagement.EntitlementManagementClient
	EntitlementManagementResourceNamespace                                    *entitlementmanagementresourcenamespace.EntitlementManagementResourceNamespaceClient
	EntitlementManagementResourceNamespaceResourceAction                      *entitlementmanagementresourcenamespaceresourceaction.EntitlementManagementResourceNamespaceResourceActionClient
	EntitlementManagementResourceNamespaceResourceActionAuthenticationContext *entitlementmanagementresourcenamespaceresourceactionauthenticationcontext.EntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient
	EntitlementManagementResourceNamespaceResourceActionResourceScope         *entitlementmanagementresourcenamespaceresourceactionresourcescope.EntitlementManagementResourceNamespaceResourceActionResourceScopeClient
	EntitlementManagementRoleAssignment                                       *entitlementmanagementroleassignment.EntitlementManagementRoleAssignmentClient
	EntitlementManagementRoleAssignmentAppScope                               *entitlementmanagementroleassignmentappscope.EntitlementManagementRoleAssignmentAppScopeClient
	EntitlementManagementRoleAssignmentApproval                               *entitlementmanagementroleassignmentapproval.EntitlementManagementRoleAssignmentApprovalClient
	EntitlementManagementRoleAssignmentApprovalStep                           *entitlementmanagementroleassignmentapprovalstep.EntitlementManagementRoleAssignmentApprovalStepClient
	EntitlementManagementRoleAssignmentDirectoryScope                         *entitlementmanagementroleassignmentdirectoryscope.EntitlementManagementRoleAssignmentDirectoryScopeClient
	EntitlementManagementRoleAssignmentPrincipal                              *entitlementmanagementroleassignmentprincipal.EntitlementManagementRoleAssignmentPrincipalClient
	EntitlementManagementRoleAssignmentRoleDefinition                         *entitlementmanagementroleassignmentroledefinition.EntitlementManagementRoleAssignmentRoleDefinitionClient
	EntitlementManagementRoleAssignmentSchedule                               *entitlementmanagementroleassignmentschedule.EntitlementManagementRoleAssignmentScheduleClient
	EntitlementManagementRoleAssignmentScheduleActivatedUsing                 *entitlementmanagementroleassignmentscheduleactivatedusing.EntitlementManagementRoleAssignmentScheduleActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleAppScope                       *entitlementmanagementroleassignmentscheduleappscope.EntitlementManagementRoleAssignmentScheduleAppScopeClient
	EntitlementManagementRoleAssignmentScheduleDirectoryScope                 *entitlementmanagementroleassignmentscheduledirectoryscope.EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleInstance                       *entitlementmanagementroleassignmentscheduleinstance.EntitlementManagementRoleAssignmentScheduleInstanceClient
	EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing         *entitlementmanagementroleassignmentscheduleinstanceactivatedusing.EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleInstanceAppScope               *entitlementmanagementroleassignmentscheduleinstanceappscope.EntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient
	EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope         *entitlementmanagementroleassignmentscheduleinstancedirectoryscope.EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleInstancePrincipal              *entitlementmanagementroleassignmentscheduleinstanceprincipal.EntitlementManagementRoleAssignmentScheduleInstancePrincipalClient
	EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition         *entitlementmanagementroleassignmentscheduleinstanceroledefinition.EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient
	EntitlementManagementRoleAssignmentSchedulePrincipal                      *entitlementmanagementroleassignmentscheduleprincipal.EntitlementManagementRoleAssignmentSchedulePrincipalClient
	EntitlementManagementRoleAssignmentScheduleRequest                        *entitlementmanagementroleassignmentschedulerequest.EntitlementManagementRoleAssignmentScheduleRequestClient
	EntitlementManagementRoleAssignmentScheduleRequestActivatedUsing          *entitlementmanagementroleassignmentschedulerequestactivatedusing.EntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient
	EntitlementManagementRoleAssignmentScheduleRequestAppScope                *entitlementmanagementroleassignmentschedulerequestappscope.EntitlementManagementRoleAssignmentScheduleRequestAppScopeClient
	EntitlementManagementRoleAssignmentScheduleRequestDirectoryScope          *entitlementmanagementroleassignmentschedulerequestdirectoryscope.EntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient
	EntitlementManagementRoleAssignmentScheduleRequestPrincipal               *entitlementmanagementroleassignmentschedulerequestprincipal.EntitlementManagementRoleAssignmentScheduleRequestPrincipalClient
	EntitlementManagementRoleAssignmentScheduleRequestRoleDefinition          *entitlementmanagementroleassignmentschedulerequestroledefinition.EntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient
	EntitlementManagementRoleAssignmentScheduleRequestTargetSchedule          *entitlementmanagementroleassignmentschedulerequesttargetschedule.EntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient
	EntitlementManagementRoleAssignmentScheduleRoleDefinition                 *entitlementmanagementroleassignmentscheduleroledefinition.EntitlementManagementRoleAssignmentScheduleRoleDefinitionClient
	EntitlementManagementRoleDefinition                                       *entitlementmanagementroledefinition.EntitlementManagementRoleDefinitionClient
	EntitlementManagementRoleDefinitionInheritsPermissionsFrom                *entitlementmanagementroledefinitioninheritspermissionsfrom.EntitlementManagementRoleDefinitionInheritsPermissionsFromClient
	EntitlementManagementRoleEligibilitySchedule                              *entitlementmanagementroleeligibilityschedule.EntitlementManagementRoleEligibilityScheduleClient
	EntitlementManagementRoleEligibilityScheduleAppScope                      *entitlementmanagementroleeligibilityscheduleappscope.EntitlementManagementRoleEligibilityScheduleAppScopeClient
	EntitlementManagementRoleEligibilityScheduleDirectoryScope                *entitlementmanagementroleeligibilityscheduledirectoryscope.EntitlementManagementRoleEligibilityScheduleDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleInstance                      *entitlementmanagementroleeligibilityscheduleinstance.EntitlementManagementRoleEligibilityScheduleInstanceClient
	EntitlementManagementRoleEligibilityScheduleInstanceAppScope              *entitlementmanagementroleeligibilityscheduleinstanceappscope.EntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient
	EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope        *entitlementmanagementroleeligibilityscheduleinstancedirectoryscope.EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleInstancePrincipal             *entitlementmanagementroleeligibilityscheduleinstanceprincipal.EntitlementManagementRoleEligibilityScheduleInstancePrincipalClient
	EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition        *entitlementmanagementroleeligibilityscheduleinstanceroledefinition.EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient
	EntitlementManagementRoleEligibilitySchedulePrincipal                     *entitlementmanagementroleeligibilityscheduleprincipal.EntitlementManagementRoleEligibilitySchedulePrincipalClient
	EntitlementManagementRoleEligibilityScheduleRequest                       *entitlementmanagementroleeligibilityschedulerequest.EntitlementManagementRoleEligibilityScheduleRequestClient
	EntitlementManagementRoleEligibilityScheduleRequestAppScope               *entitlementmanagementroleeligibilityschedulerequestappscope.EntitlementManagementRoleEligibilityScheduleRequestAppScopeClient
	EntitlementManagementRoleEligibilityScheduleRequestDirectoryScope         *entitlementmanagementroleeligibilityschedulerequestdirectoryscope.EntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient
	EntitlementManagementRoleEligibilityScheduleRequestPrincipal              *entitlementmanagementroleeligibilityschedulerequestprincipal.EntitlementManagementRoleEligibilityScheduleRequestPrincipalClient
	EntitlementManagementRoleEligibilityScheduleRequestRoleDefinition         *entitlementmanagementroleeligibilityschedulerequestroledefinition.EntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient
	EntitlementManagementRoleEligibilityScheduleRequestTargetSchedule         *entitlementmanagementroleeligibilityschedulerequesttargetschedule.EntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient
	EntitlementManagementRoleEligibilityScheduleRoleDefinition                *entitlementmanagementroleeligibilityscheduleroledefinition.EntitlementManagementRoleEligibilityScheduleRoleDefinitionClient
	EntitlementManagementTransitiveRoleAssignment                             *entitlementmanagementtransitiveroleassignment.EntitlementManagementTransitiveRoleAssignmentClient
	EntitlementManagementTransitiveRoleAssignmentAppScope                     *entitlementmanagementtransitiveroleassignmentappscope.EntitlementManagementTransitiveRoleAssignmentAppScopeClient
	EntitlementManagementTransitiveRoleAssignmentDirectoryScope               *entitlementmanagementtransitiveroleassignmentdirectoryscope.EntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient
	EntitlementManagementTransitiveRoleAssignmentPrincipal                    *entitlementmanagementtransitiveroleassignmentprincipal.EntitlementManagementTransitiveRoleAssignmentPrincipalClient
	EntitlementManagementTransitiveRoleAssignmentRoleDefinition               *entitlementmanagementtransitiveroleassignmentroledefinition.EntitlementManagementTransitiveRoleAssignmentRoleDefinitionClient
	Exchange                                                                  *exchange.ExchangeClient
	ExchangeCustomAppScope                                                    *exchangecustomappscope.ExchangeCustomAppScopeClient
	ExchangeResourceNamespace                                                 *exchangeresourcenamespace.ExchangeResourceNamespaceClient
	ExchangeResourceNamespaceResourceAction                                   *exchangeresourcenamespaceresourceaction.ExchangeResourceNamespaceResourceActionClient
	ExchangeResourceNamespaceResourceActionAuthenticationContext              *exchangeresourcenamespaceresourceactionauthenticationcontext.ExchangeResourceNamespaceResourceActionAuthenticationContextClient
	ExchangeResourceNamespaceResourceActionResourceScope                      *exchangeresourcenamespaceresourceactionresourcescope.ExchangeResourceNamespaceResourceActionResourceScopeClient
	ExchangeRoleAssignment                                                    *exchangeroleassignment.ExchangeRoleAssignmentClient
	ExchangeRoleAssignmentAppScope                                            *exchangeroleassignmentappscope.ExchangeRoleAssignmentAppScopeClient
	ExchangeRoleAssignmentDirectoryScope                                      *exchangeroleassignmentdirectoryscope.ExchangeRoleAssignmentDirectoryScopeClient
	ExchangeRoleAssignmentPrincipal                                           *exchangeroleassignmentprincipal.ExchangeRoleAssignmentPrincipalClient
	ExchangeRoleAssignmentRoleDefinition                                      *exchangeroleassignmentroledefinition.ExchangeRoleAssignmentRoleDefinitionClient
	ExchangeRoleDefinition                                                    *exchangeroledefinition.ExchangeRoleDefinitionClient
	ExchangeRoleDefinitionInheritsPermissionsFrom                             *exchangeroledefinitioninheritspermissionsfrom.ExchangeRoleDefinitionInheritsPermissionsFromClient
	ExchangeTransitiveRoleAssignment                                          *exchangetransitiveroleassignment.ExchangeTransitiveRoleAssignmentClient
	ExchangeTransitiveRoleAssignmentAppScope                                  *exchangetransitiveroleassignmentappscope.ExchangeTransitiveRoleAssignmentAppScopeClient
	ExchangeTransitiveRoleAssignmentDirectoryScope                            *exchangetransitiveroleassignmentdirectoryscope.ExchangeTransitiveRoleAssignmentDirectoryScopeClient
	ExchangeTransitiveRoleAssignmentPrincipal                                 *exchangetransitiveroleassignmentprincipal.ExchangeTransitiveRoleAssignmentPrincipalClient
	ExchangeTransitiveRoleAssignmentRoleDefinition                            *exchangetransitiveroleassignmentroledefinition.ExchangeTransitiveRoleAssignmentRoleDefinitionClient
	RoleManagement                                                            *rolemanagement.RoleManagementClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	cloudPCClient, err := cloudpc.NewCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPC client: %+v", err)
	}
	configureFunc(cloudPCClient.Client)

	cloudPCResourceNamespaceClient, err := cloudpcresourcenamespace.NewCloudPCResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCResourceNamespace client: %+v", err)
	}
	configureFunc(cloudPCResourceNamespaceClient.Client)

	cloudPCResourceNamespaceResourceActionAuthenticationContextClient, err := cloudpcresourcenamespaceresourceactionauthenticationcontext.NewCloudPCResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(cloudPCResourceNamespaceResourceActionAuthenticationContextClient.Client)

	cloudPCResourceNamespaceResourceActionClient, err := cloudpcresourcenamespaceresourceaction.NewCloudPCResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(cloudPCResourceNamespaceResourceActionClient.Client)

	cloudPCResourceNamespaceResourceActionResourceScopeClient, err := cloudpcresourcenamespaceresourceactionresourcescope.NewCloudPCResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(cloudPCResourceNamespaceResourceActionResourceScopeClient.Client)

	cloudPCRoleAssignmentAppScopeClient, err := cloudpcroleassignmentappscope.NewCloudPCRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(cloudPCRoleAssignmentAppScopeClient.Client)

	cloudPCRoleAssignmentClient, err := cloudpcroleassignment.NewCloudPCRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleAssignment client: %+v", err)
	}
	configureFunc(cloudPCRoleAssignmentClient.Client)

	cloudPCRoleAssignmentDirectoryScopeClient, err := cloudpcroleassignmentdirectoryscope.NewCloudPCRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(cloudPCRoleAssignmentDirectoryScopeClient.Client)

	cloudPCRoleAssignmentPrincipalClient, err := cloudpcroleassignmentprincipal.NewCloudPCRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(cloudPCRoleAssignmentPrincipalClient.Client)

	cloudPCRoleAssignmentRoleDefinitionClient, err := cloudpcroleassignmentroledefinition.NewCloudPCRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(cloudPCRoleAssignmentRoleDefinitionClient.Client)

	cloudPCRoleDefinitionClient, err := cloudpcroledefinition.NewCloudPCRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleDefinition client: %+v", err)
	}
	configureFunc(cloudPCRoleDefinitionClient.Client)

	cloudPCRoleDefinitionInheritsPermissionsFromClient, err := cloudpcroledefinitioninheritspermissionsfrom.NewCloudPCRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CloudPCRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(cloudPCRoleDefinitionInheritsPermissionsFromClient.Client)

	defenderClient, err := defender.NewDefenderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Defender client: %+v", err)
	}
	configureFunc(defenderClient.Client)

	defenderResourceNamespaceClient, err := defenderresourcenamespace.NewDefenderResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderResourceNamespace client: %+v", err)
	}
	configureFunc(defenderResourceNamespaceClient.Client)

	defenderResourceNamespaceResourceActionAuthenticationContextClient, err := defenderresourcenamespaceresourceactionauthenticationcontext.NewDefenderResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(defenderResourceNamespaceResourceActionAuthenticationContextClient.Client)

	defenderResourceNamespaceResourceActionClient, err := defenderresourcenamespaceresourceaction.NewDefenderResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(defenderResourceNamespaceResourceActionClient.Client)

	defenderResourceNamespaceResourceActionResourceScopeClient, err := defenderresourcenamespaceresourceactionresourcescope.NewDefenderResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(defenderResourceNamespaceResourceActionResourceScopeClient.Client)

	defenderRoleAssignmentAppScopeClient, err := defenderroleassignmentappscope.NewDefenderRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(defenderRoleAssignmentAppScopeClient.Client)

	defenderRoleAssignmentClient, err := defenderroleassignment.NewDefenderRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleAssignment client: %+v", err)
	}
	configureFunc(defenderRoleAssignmentClient.Client)

	defenderRoleAssignmentDirectoryScopeClient, err := defenderroleassignmentdirectoryscope.NewDefenderRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(defenderRoleAssignmentDirectoryScopeClient.Client)

	defenderRoleAssignmentPrincipalClient, err := defenderroleassignmentprincipal.NewDefenderRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(defenderRoleAssignmentPrincipalClient.Client)

	defenderRoleAssignmentRoleDefinitionClient, err := defenderroleassignmentroledefinition.NewDefenderRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(defenderRoleAssignmentRoleDefinitionClient.Client)

	defenderRoleDefinitionClient, err := defenderroledefinition.NewDefenderRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleDefinition client: %+v", err)
	}
	configureFunc(defenderRoleDefinitionClient.Client)

	defenderRoleDefinitionInheritsPermissionsFromClient, err := defenderroledefinitioninheritspermissionsfrom.NewDefenderRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(defenderRoleDefinitionInheritsPermissionsFromClient.Client)

	deviceManagementClient, err := devicemanagement.NewDeviceManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagement client: %+v", err)
	}
	configureFunc(deviceManagementClient.Client)

	deviceManagementResourceNamespaceClient, err := devicemanagementresourcenamespace.NewDeviceManagementResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementResourceNamespace client: %+v", err)
	}
	configureFunc(deviceManagementResourceNamespaceClient.Client)

	deviceManagementResourceNamespaceResourceActionAuthenticationContextClient, err := devicemanagementresourcenamespaceresourceactionauthenticationcontext.NewDeviceManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(deviceManagementResourceNamespaceResourceActionAuthenticationContextClient.Client)

	deviceManagementResourceNamespaceResourceActionClient, err := devicemanagementresourcenamespaceresourceaction.NewDeviceManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(deviceManagementResourceNamespaceResourceActionClient.Client)

	deviceManagementResourceNamespaceResourceActionResourceScopeClient, err := devicemanagementresourcenamespaceresourceactionresourcescope.NewDeviceManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(deviceManagementResourceNamespaceResourceActionResourceScopeClient.Client)

	deviceManagementRoleAssignmentAppScopeClient, err := devicemanagementroleassignmentappscope.NewDeviceManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(deviceManagementRoleAssignmentAppScopeClient.Client)

	deviceManagementRoleAssignmentClient, err := devicemanagementroleassignment.NewDeviceManagementRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleAssignment client: %+v", err)
	}
	configureFunc(deviceManagementRoleAssignmentClient.Client)

	deviceManagementRoleAssignmentDirectoryScopeClient, err := devicemanagementroleassignmentdirectoryscope.NewDeviceManagementRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(deviceManagementRoleAssignmentDirectoryScopeClient.Client)

	deviceManagementRoleAssignmentPrincipalClient, err := devicemanagementroleassignmentprincipal.NewDeviceManagementRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(deviceManagementRoleAssignmentPrincipalClient.Client)

	deviceManagementRoleAssignmentRoleDefinitionClient, err := devicemanagementroleassignmentroledefinition.NewDeviceManagementRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(deviceManagementRoleAssignmentRoleDefinitionClient.Client)

	deviceManagementRoleDefinitionClient, err := devicemanagementroledefinition.NewDeviceManagementRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleDefinition client: %+v", err)
	}
	configureFunc(deviceManagementRoleDefinitionClient.Client)

	deviceManagementRoleDefinitionInheritsPermissionsFromClient, err := devicemanagementroledefinitioninheritspermissionsfrom.NewDeviceManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceManagementRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(deviceManagementRoleDefinitionInheritsPermissionsFromClient.Client)

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

	directoryResourceNamespaceResourceActionAuthenticationContextClient, err := directoryresourcenamespaceresourceactionauthenticationcontext.NewDirectoryResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(directoryResourceNamespaceResourceActionAuthenticationContextClient.Client)

	directoryResourceNamespaceResourceActionClient, err := directoryresourcenamespaceresourceaction.NewDirectoryResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(directoryResourceNamespaceResourceActionClient.Client)

	directoryResourceNamespaceResourceActionResourceScopeClient, err := directoryresourcenamespaceresourceactionresourcescope.NewDirectoryResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(directoryResourceNamespaceResourceActionResourceScopeClient.Client)

	directoryRoleAssignmentAppScopeClient, err := directoryroleassignmentappscope.NewDirectoryRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentAppScopeClient.Client)

	directoryRoleAssignmentApprovalClient, err := directoryroleassignmentapproval.NewDirectoryRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentApprovalClient.Client)

	directoryRoleAssignmentApprovalStepClient, err := directoryroleassignmentapprovalstep.NewDirectoryRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(directoryRoleAssignmentApprovalStepClient.Client)

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

	directoryTransitiveRoleAssignmentAppScopeClient, err := directorytransitiveroleassignmentappscope.NewDirectoryTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(directoryTransitiveRoleAssignmentAppScopeClient.Client)

	directoryTransitiveRoleAssignmentClient, err := directorytransitiveroleassignment.NewDirectoryTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(directoryTransitiveRoleAssignmentClient.Client)

	directoryTransitiveRoleAssignmentDirectoryScopeClient, err := directorytransitiveroleassignmentdirectoryscope.NewDirectoryTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(directoryTransitiveRoleAssignmentDirectoryScopeClient.Client)

	directoryTransitiveRoleAssignmentPrincipalClient, err := directorytransitiveroleassignmentprincipal.NewDirectoryTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(directoryTransitiveRoleAssignmentPrincipalClient.Client)

	directoryTransitiveRoleAssignmentRoleDefinitionClient, err := directorytransitiveroleassignmentroledefinition.NewDirectoryTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(directoryTransitiveRoleAssignmentRoleDefinitionClient.Client)

	enterpriseAppClient, err := enterpriseapp.NewEnterpriseAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseApp client: %+v", err)
	}
	configureFunc(enterpriseAppClient.Client)

	enterpriseAppResourceNamespaceClient, err := enterpriseappresourcenamespace.NewEnterpriseAppResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppResourceNamespace client: %+v", err)
	}
	configureFunc(enterpriseAppResourceNamespaceClient.Client)

	enterpriseAppResourceNamespaceResourceActionAuthenticationContextClient, err := enterpriseappresourcenamespaceresourceactionauthenticationcontext.NewEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(enterpriseAppResourceNamespaceResourceActionAuthenticationContextClient.Client)

	enterpriseAppResourceNamespaceResourceActionClient, err := enterpriseappresourcenamespaceresourceaction.NewEnterpriseAppResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(enterpriseAppResourceNamespaceResourceActionClient.Client)

	enterpriseAppResourceNamespaceResourceActionResourceScopeClient, err := enterpriseappresourcenamespaceresourceactionresourcescope.NewEnterpriseAppResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(enterpriseAppResourceNamespaceResourceActionResourceScopeClient.Client)

	enterpriseAppRoleAssignmentAppScopeClient, err := enterpriseapproleassignmentappscope.NewEnterpriseAppRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentAppScopeClient.Client)

	enterpriseAppRoleAssignmentApprovalClient, err := enterpriseapproleassignmentapproval.NewEnterpriseAppRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentApprovalClient.Client)

	enterpriseAppRoleAssignmentApprovalStepClient, err := enterpriseapproleassignmentapprovalstep.NewEnterpriseAppRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentApprovalStepClient.Client)

	enterpriseAppRoleAssignmentClient, err := enterpriseapproleassignment.NewEnterpriseAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignment client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentClient.Client)

	enterpriseAppRoleAssignmentDirectoryScopeClient, err := enterpriseapproleassignmentdirectoryscope.NewEnterpriseAppRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentDirectoryScopeClient.Client)

	enterpriseAppRoleAssignmentPrincipalClient, err := enterpriseapproleassignmentprincipal.NewEnterpriseAppRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentPrincipalClient.Client)

	enterpriseAppRoleAssignmentRoleDefinitionClient, err := enterpriseapproleassignmentroledefinition.NewEnterpriseAppRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentRoleDefinitionClient.Client)

	enterpriseAppRoleAssignmentScheduleActivatedUsingClient, err := enterpriseapproleassignmentscheduleactivatedusing.NewEnterpriseAppRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleActivatedUsingClient.Client)

	enterpriseAppRoleAssignmentScheduleAppScopeClient, err := enterpriseapproleassignmentscheduleappscope.NewEnterpriseAppRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleAppScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleClient, err := enterpriseapproleassignmentschedule.NewEnterpriseAppRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleClient.Client)

	enterpriseAppRoleAssignmentScheduleDirectoryScopeClient, err := enterpriseapproleassignmentscheduledirectoryscope.NewEnterpriseAppRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleDirectoryScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient, err := enterpriseapproleassignmentscheduleinstanceactivatedusing.NewEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	enterpriseAppRoleAssignmentScheduleInstanceAppScopeClient, err := enterpriseapproleassignmentscheduleinstanceappscope.NewEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstanceAppScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleInstanceClient, err := enterpriseapproleassignmentscheduleinstance.NewEnterpriseAppRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstanceClient.Client)

	enterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient, err := enterpriseapproleassignmentscheduleinstancedirectoryscope.NewEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleInstancePrincipalClient, err := enterpriseapproleassignmentscheduleinstanceprincipal.NewEnterpriseAppRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstancePrincipalClient.Client)

	enterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient, err := enterpriseapproleassignmentscheduleinstanceroledefinition.NewEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	enterpriseAppRoleAssignmentSchedulePrincipalClient, err := enterpriseapproleassignmentscheduleprincipal.NewEnterpriseAppRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentSchedulePrincipalClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient, err := enterpriseapproleassignmentschedulerequestactivatedusing.NewEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestAppScopeClient, err := enterpriseapproleassignmentschedulerequestappscope.NewEnterpriseAppRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestAppScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestClient, err := enterpriseapproleassignmentschedulerequest.NewEnterpriseAppRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient, err := enterpriseapproleassignmentschedulerequestdirectoryscope.NewEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestPrincipalClient, err := enterpriseapproleassignmentschedulerequestprincipal.NewEnterpriseAppRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestPrincipalClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient, err := enterpriseapproleassignmentschedulerequestroledefinition.NewEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	enterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient, err := enterpriseapproleassignmentschedulerequesttargetschedule.NewEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	enterpriseAppRoleAssignmentScheduleRoleDefinitionClient, err := enterpriseapproleassignmentscheduleroledefinition.NewEnterpriseAppRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleAssignmentScheduleRoleDefinitionClient.Client)

	enterpriseAppRoleDefinitionClient, err := enterpriseapproledefinition.NewEnterpriseAppRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleDefinitionClient.Client)

	enterpriseAppRoleDefinitionInheritsPermissionsFromClient, err := enterpriseapproledefinitioninheritspermissionsfrom.NewEnterpriseAppRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(enterpriseAppRoleDefinitionInheritsPermissionsFromClient.Client)

	enterpriseAppRoleEligibilityScheduleAppScopeClient, err := enterpriseapproleeligibilityscheduleappscope.NewEnterpriseAppRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleAppScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleClient, err := enterpriseapproleeligibilityschedule.NewEnterpriseAppRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleClient.Client)

	enterpriseAppRoleEligibilityScheduleDirectoryScopeClient, err := enterpriseapproleeligibilityscheduledirectoryscope.NewEnterpriseAppRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleDirectoryScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleInstanceAppScopeClient, err := enterpriseapproleeligibilityscheduleinstanceappscope.NewEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleInstanceAppScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleInstanceClient, err := enterpriseapproleeligibilityscheduleinstance.NewEnterpriseAppRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleInstanceClient.Client)

	enterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient, err := enterpriseapproleeligibilityscheduleinstancedirectoryscope.NewEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleInstancePrincipalClient, err := enterpriseapproleeligibilityscheduleinstanceprincipal.NewEnterpriseAppRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleInstancePrincipalClient.Client)

	enterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient, err := enterpriseapproleeligibilityscheduleinstanceroledefinition.NewEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	enterpriseAppRoleEligibilitySchedulePrincipalClient, err := enterpriseapproleeligibilityscheduleprincipal.NewEnterpriseAppRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilitySchedulePrincipalClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestAppScopeClient, err := enterpriseapproleeligibilityschedulerequestappscope.NewEnterpriseAppRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestAppScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestClient, err := enterpriseapproleeligibilityschedulerequest.NewEnterpriseAppRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient, err := enterpriseapproleeligibilityschedulerequestdirectoryscope.NewEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestPrincipalClient, err := enterpriseapproleeligibilityschedulerequestprincipal.NewEnterpriseAppRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestPrincipalClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient, err := enterpriseapproleeligibilityschedulerequestroledefinition.NewEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	enterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient, err := enterpriseapproleeligibilityschedulerequesttargetschedule.NewEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	enterpriseAppRoleEligibilityScheduleRoleDefinitionClient, err := enterpriseapproleeligibilityscheduleroledefinition.NewEnterpriseAppRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppRoleEligibilityScheduleRoleDefinitionClient.Client)

	enterpriseAppTransitiveRoleAssignmentAppScopeClient, err := enterpriseapptransitiveroleassignmentappscope.NewEnterpriseAppTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(enterpriseAppTransitiveRoleAssignmentAppScopeClient.Client)

	enterpriseAppTransitiveRoleAssignmentClient, err := enterpriseapptransitiveroleassignment.NewEnterpriseAppTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(enterpriseAppTransitiveRoleAssignmentClient.Client)

	enterpriseAppTransitiveRoleAssignmentDirectoryScopeClient, err := enterpriseapptransitiveroleassignmentdirectoryscope.NewEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(enterpriseAppTransitiveRoleAssignmentDirectoryScopeClient.Client)

	enterpriseAppTransitiveRoleAssignmentPrincipalClient, err := enterpriseapptransitiveroleassignmentprincipal.NewEnterpriseAppTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(enterpriseAppTransitiveRoleAssignmentPrincipalClient.Client)

	enterpriseAppTransitiveRoleAssignmentRoleDefinitionClient, err := enterpriseapptransitiveroleassignmentroledefinition.NewEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnterpriseAppTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(enterpriseAppTransitiveRoleAssignmentRoleDefinitionClient.Client)

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

	entitlementManagementResourceNamespaceResourceActionAuthenticationContextClient, err := entitlementmanagementresourcenamespaceresourceactionauthenticationcontext.NewEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(entitlementManagementResourceNamespaceResourceActionAuthenticationContextClient.Client)

	entitlementManagementResourceNamespaceResourceActionClient, err := entitlementmanagementresourcenamespaceresourceaction.NewEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(entitlementManagementResourceNamespaceResourceActionClient.Client)

	entitlementManagementResourceNamespaceResourceActionResourceScopeClient, err := entitlementmanagementresourcenamespaceresourceactionresourcescope.NewEntitlementManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceNamespaceResourceActionResourceScopeClient.Client)

	entitlementManagementRoleAssignmentAppScopeClient, err := entitlementmanagementroleassignmentappscope.NewEntitlementManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentAppScopeClient.Client)

	entitlementManagementRoleAssignmentApprovalClient, err := entitlementmanagementroleassignmentapproval.NewEntitlementManagementRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentApprovalClient.Client)

	entitlementManagementRoleAssignmentApprovalStepClient, err := entitlementmanagementroleassignmentapprovalstep.NewEntitlementManagementRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(entitlementManagementRoleAssignmentApprovalStepClient.Client)

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

	entitlementManagementTransitiveRoleAssignmentAppScopeClient, err := entitlementmanagementtransitiveroleassignmentappscope.NewEntitlementManagementTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(entitlementManagementTransitiveRoleAssignmentAppScopeClient.Client)

	entitlementManagementTransitiveRoleAssignmentClient, err := entitlementmanagementtransitiveroleassignment.NewEntitlementManagementTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(entitlementManagementTransitiveRoleAssignmentClient.Client)

	entitlementManagementTransitiveRoleAssignmentDirectoryScopeClient, err := entitlementmanagementtransitiveroleassignmentdirectoryscope.NewEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(entitlementManagementTransitiveRoleAssignmentDirectoryScopeClient.Client)

	entitlementManagementTransitiveRoleAssignmentPrincipalClient, err := entitlementmanagementtransitiveroleassignmentprincipal.NewEntitlementManagementTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(entitlementManagementTransitiveRoleAssignmentPrincipalClient.Client)

	entitlementManagementTransitiveRoleAssignmentRoleDefinitionClient, err := entitlementmanagementtransitiveroleassignmentroledefinition.NewEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(entitlementManagementTransitiveRoleAssignmentRoleDefinitionClient.Client)

	exchangeClient, err := exchange.NewExchangeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Exchange client: %+v", err)
	}
	configureFunc(exchangeClient.Client)

	exchangeCustomAppScopeClient, err := exchangecustomappscope.NewExchangeCustomAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeCustomAppScope client: %+v", err)
	}
	configureFunc(exchangeCustomAppScopeClient.Client)

	exchangeResourceNamespaceClient, err := exchangeresourcenamespace.NewExchangeResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeResourceNamespace client: %+v", err)
	}
	configureFunc(exchangeResourceNamespaceClient.Client)

	exchangeResourceNamespaceResourceActionAuthenticationContextClient, err := exchangeresourcenamespaceresourceactionauthenticationcontext.NewExchangeResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(exchangeResourceNamespaceResourceActionAuthenticationContextClient.Client)

	exchangeResourceNamespaceResourceActionClient, err := exchangeresourcenamespaceresourceaction.NewExchangeResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(exchangeResourceNamespaceResourceActionClient.Client)

	exchangeResourceNamespaceResourceActionResourceScopeClient, err := exchangeresourcenamespaceresourceactionresourcescope.NewExchangeResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(exchangeResourceNamespaceResourceActionResourceScopeClient.Client)

	exchangeRoleAssignmentAppScopeClient, err := exchangeroleassignmentappscope.NewExchangeRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(exchangeRoleAssignmentAppScopeClient.Client)

	exchangeRoleAssignmentClient, err := exchangeroleassignment.NewExchangeRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleAssignment client: %+v", err)
	}
	configureFunc(exchangeRoleAssignmentClient.Client)

	exchangeRoleAssignmentDirectoryScopeClient, err := exchangeroleassignmentdirectoryscope.NewExchangeRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(exchangeRoleAssignmentDirectoryScopeClient.Client)

	exchangeRoleAssignmentPrincipalClient, err := exchangeroleassignmentprincipal.NewExchangeRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(exchangeRoleAssignmentPrincipalClient.Client)

	exchangeRoleAssignmentRoleDefinitionClient, err := exchangeroleassignmentroledefinition.NewExchangeRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(exchangeRoleAssignmentRoleDefinitionClient.Client)

	exchangeRoleDefinitionClient, err := exchangeroledefinition.NewExchangeRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleDefinition client: %+v", err)
	}
	configureFunc(exchangeRoleDefinitionClient.Client)

	exchangeRoleDefinitionInheritsPermissionsFromClient, err := exchangeroledefinitioninheritspermissionsfrom.NewExchangeRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(exchangeRoleDefinitionInheritsPermissionsFromClient.Client)

	exchangeTransitiveRoleAssignmentAppScopeClient, err := exchangetransitiveroleassignmentappscope.NewExchangeTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(exchangeTransitiveRoleAssignmentAppScopeClient.Client)

	exchangeTransitiveRoleAssignmentClient, err := exchangetransitiveroleassignment.NewExchangeTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(exchangeTransitiveRoleAssignmentClient.Client)

	exchangeTransitiveRoleAssignmentDirectoryScopeClient, err := exchangetransitiveroleassignmentdirectoryscope.NewExchangeTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(exchangeTransitiveRoleAssignmentDirectoryScopeClient.Client)

	exchangeTransitiveRoleAssignmentPrincipalClient, err := exchangetransitiveroleassignmentprincipal.NewExchangeTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(exchangeTransitiveRoleAssignmentPrincipalClient.Client)

	exchangeTransitiveRoleAssignmentRoleDefinitionClient, err := exchangetransitiveroleassignmentroledefinition.NewExchangeTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExchangeTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(exchangeTransitiveRoleAssignmentRoleDefinitionClient.Client)

	roleManagementClient, err := rolemanagement.NewRoleManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagement client: %+v", err)
	}
	configureFunc(roleManagementClient.Client)

	return &Client{
		CloudPC:                                cloudPCClient,
		CloudPCResourceNamespace:               cloudPCResourceNamespaceClient,
		CloudPCResourceNamespaceResourceAction: cloudPCResourceNamespaceResourceActionClient,
		CloudPCResourceNamespaceResourceActionAuthenticationContext: cloudPCResourceNamespaceResourceActionAuthenticationContextClient,
		CloudPCResourceNamespaceResourceActionResourceScope:         cloudPCResourceNamespaceResourceActionResourceScopeClient,
		CloudPCRoleAssignment:                        cloudPCRoleAssignmentClient,
		CloudPCRoleAssignmentAppScope:                cloudPCRoleAssignmentAppScopeClient,
		CloudPCRoleAssignmentDirectoryScope:          cloudPCRoleAssignmentDirectoryScopeClient,
		CloudPCRoleAssignmentPrincipal:               cloudPCRoleAssignmentPrincipalClient,
		CloudPCRoleAssignmentRoleDefinition:          cloudPCRoleAssignmentRoleDefinitionClient,
		CloudPCRoleDefinition:                        cloudPCRoleDefinitionClient,
		CloudPCRoleDefinitionInheritsPermissionsFrom: cloudPCRoleDefinitionInheritsPermissionsFromClient,
		Defender:                                defenderClient,
		DefenderResourceNamespace:               defenderResourceNamespaceClient,
		DefenderResourceNamespaceResourceAction: defenderResourceNamespaceResourceActionClient,
		DefenderResourceNamespaceResourceActionAuthenticationContext: defenderResourceNamespaceResourceActionAuthenticationContextClient,
		DefenderResourceNamespaceResourceActionResourceScope:         defenderResourceNamespaceResourceActionResourceScopeClient,
		DefenderRoleAssignment:                                               defenderRoleAssignmentClient,
		DefenderRoleAssignmentAppScope:                                       defenderRoleAssignmentAppScopeClient,
		DefenderRoleAssignmentDirectoryScope:                                 defenderRoleAssignmentDirectoryScopeClient,
		DefenderRoleAssignmentPrincipal:                                      defenderRoleAssignmentPrincipalClient,
		DefenderRoleAssignmentRoleDefinition:                                 defenderRoleAssignmentRoleDefinitionClient,
		DefenderRoleDefinition:                                               defenderRoleDefinitionClient,
		DefenderRoleDefinitionInheritsPermissionsFrom:                        defenderRoleDefinitionInheritsPermissionsFromClient,
		DeviceManagement:                                                     deviceManagementClient,
		DeviceManagementResourceNamespace:                                    deviceManagementResourceNamespaceClient,
		DeviceManagementResourceNamespaceResourceAction:                      deviceManagementResourceNamespaceResourceActionClient,
		DeviceManagementResourceNamespaceResourceActionAuthenticationContext: deviceManagementResourceNamespaceResourceActionAuthenticationContextClient,
		DeviceManagementResourceNamespaceResourceActionResourceScope:         deviceManagementResourceNamespaceResourceActionResourceScopeClient,
		DeviceManagementRoleAssignment:                                       deviceManagementRoleAssignmentClient,
		DeviceManagementRoleAssignmentAppScope:                               deviceManagementRoleAssignmentAppScopeClient,
		DeviceManagementRoleAssignmentDirectoryScope:                         deviceManagementRoleAssignmentDirectoryScopeClient,
		DeviceManagementRoleAssignmentPrincipal:                              deviceManagementRoleAssignmentPrincipalClient,
		DeviceManagementRoleAssignmentRoleDefinition:                         deviceManagementRoleAssignmentRoleDefinitionClient,
		DeviceManagementRoleDefinition:                                       deviceManagementRoleDefinitionClient,
		DeviceManagementRoleDefinitionInheritsPermissionsFrom:                deviceManagementRoleDefinitionInheritsPermissionsFromClient,
		Directory:                                directoryClient,
		DirectoryResourceNamespace:               directoryResourceNamespaceClient,
		DirectoryResourceNamespaceResourceAction: directoryResourceNamespaceResourceActionClient,
		DirectoryResourceNamespaceResourceActionAuthenticationContext: directoryResourceNamespaceResourceActionAuthenticationContextClient,
		DirectoryResourceNamespaceResourceActionResourceScope:         directoryResourceNamespaceResourceActionResourceScopeClient,
		DirectoryRoleAssignment:                                       directoryRoleAssignmentClient,
		DirectoryRoleAssignmentAppScope:                               directoryRoleAssignmentAppScopeClient,
		DirectoryRoleAssignmentApproval:                               directoryRoleAssignmentApprovalClient,
		DirectoryRoleAssignmentApprovalStep:                           directoryRoleAssignmentApprovalStepClient,
		DirectoryRoleAssignmentDirectoryScope:                         directoryRoleAssignmentDirectoryScopeClient,
		DirectoryRoleAssignmentPrincipal:                              directoryRoleAssignmentPrincipalClient,
		DirectoryRoleAssignmentRoleDefinition:                         directoryRoleAssignmentRoleDefinitionClient,
		DirectoryRoleAssignmentSchedule:                               directoryRoleAssignmentScheduleClient,
		DirectoryRoleAssignmentScheduleActivatedUsing:                 directoryRoleAssignmentScheduleActivatedUsingClient,
		DirectoryRoleAssignmentScheduleAppScope:                       directoryRoleAssignmentScheduleAppScopeClient,
		DirectoryRoleAssignmentScheduleDirectoryScope:                 directoryRoleAssignmentScheduleDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleInstance:                       directoryRoleAssignmentScheduleInstanceClient,
		DirectoryRoleAssignmentScheduleInstanceActivatedUsing:         directoryRoleAssignmentScheduleInstanceActivatedUsingClient,
		DirectoryRoleAssignmentScheduleInstanceAppScope:               directoryRoleAssignmentScheduleInstanceAppScopeClient,
		DirectoryRoleAssignmentScheduleInstanceDirectoryScope:         directoryRoleAssignmentScheduleInstanceDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleInstancePrincipal:              directoryRoleAssignmentScheduleInstancePrincipalClient,
		DirectoryRoleAssignmentScheduleInstanceRoleDefinition:         directoryRoleAssignmentScheduleInstanceRoleDefinitionClient,
		DirectoryRoleAssignmentSchedulePrincipal:                      directoryRoleAssignmentSchedulePrincipalClient,
		DirectoryRoleAssignmentScheduleRequest:                        directoryRoleAssignmentScheduleRequestClient,
		DirectoryRoleAssignmentScheduleRequestActivatedUsing:          directoryRoleAssignmentScheduleRequestActivatedUsingClient,
		DirectoryRoleAssignmentScheduleRequestAppScope:                directoryRoleAssignmentScheduleRequestAppScopeClient,
		DirectoryRoleAssignmentScheduleRequestDirectoryScope:          directoryRoleAssignmentScheduleRequestDirectoryScopeClient,
		DirectoryRoleAssignmentScheduleRequestPrincipal:               directoryRoleAssignmentScheduleRequestPrincipalClient,
		DirectoryRoleAssignmentScheduleRequestRoleDefinition:          directoryRoleAssignmentScheduleRequestRoleDefinitionClient,
		DirectoryRoleAssignmentScheduleRequestTargetSchedule:          directoryRoleAssignmentScheduleRequestTargetScheduleClient,
		DirectoryRoleAssignmentScheduleRoleDefinition:                 directoryRoleAssignmentScheduleRoleDefinitionClient,
		DirectoryRoleDefinition:                                       directoryRoleDefinitionClient,
		DirectoryRoleDefinitionInheritsPermissionsFrom:                directoryRoleDefinitionInheritsPermissionsFromClient,
		DirectoryRoleEligibilitySchedule:                              directoryRoleEligibilityScheduleClient,
		DirectoryRoleEligibilityScheduleAppScope:                      directoryRoleEligibilityScheduleAppScopeClient,
		DirectoryRoleEligibilityScheduleDirectoryScope:                directoryRoleEligibilityScheduleDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleInstance:                      directoryRoleEligibilityScheduleInstanceClient,
		DirectoryRoleEligibilityScheduleInstanceAppScope:              directoryRoleEligibilityScheduleInstanceAppScopeClient,
		DirectoryRoleEligibilityScheduleInstanceDirectoryScope:        directoryRoleEligibilityScheduleInstanceDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleInstancePrincipal:             directoryRoleEligibilityScheduleInstancePrincipalClient,
		DirectoryRoleEligibilityScheduleInstanceRoleDefinition:        directoryRoleEligibilityScheduleInstanceRoleDefinitionClient,
		DirectoryRoleEligibilitySchedulePrincipal:                     directoryRoleEligibilitySchedulePrincipalClient,
		DirectoryRoleEligibilityScheduleRequest:                       directoryRoleEligibilityScheduleRequestClient,
		DirectoryRoleEligibilityScheduleRequestAppScope:               directoryRoleEligibilityScheduleRequestAppScopeClient,
		DirectoryRoleEligibilityScheduleRequestDirectoryScope:         directoryRoleEligibilityScheduleRequestDirectoryScopeClient,
		DirectoryRoleEligibilityScheduleRequestPrincipal:              directoryRoleEligibilityScheduleRequestPrincipalClient,
		DirectoryRoleEligibilityScheduleRequestRoleDefinition:         directoryRoleEligibilityScheduleRequestRoleDefinitionClient,
		DirectoryRoleEligibilityScheduleRequestTargetSchedule:         directoryRoleEligibilityScheduleRequestTargetScheduleClient,
		DirectoryRoleEligibilityScheduleRoleDefinition:                directoryRoleEligibilityScheduleRoleDefinitionClient,
		DirectoryTransitiveRoleAssignment:                             directoryTransitiveRoleAssignmentClient,
		DirectoryTransitiveRoleAssignmentAppScope:                     directoryTransitiveRoleAssignmentAppScopeClient,
		DirectoryTransitiveRoleAssignmentDirectoryScope:               directoryTransitiveRoleAssignmentDirectoryScopeClient,
		DirectoryTransitiveRoleAssignmentPrincipal:                    directoryTransitiveRoleAssignmentPrincipalClient,
		DirectoryTransitiveRoleAssignmentRoleDefinition:               directoryTransitiveRoleAssignmentRoleDefinitionClient,
		EnterpriseApp:                                enterpriseAppClient,
		EnterpriseAppResourceNamespace:               enterpriseAppResourceNamespaceClient,
		EnterpriseAppResourceNamespaceResourceAction: enterpriseAppResourceNamespaceResourceActionClient,
		EnterpriseAppResourceNamespaceResourceActionAuthenticationContext:         enterpriseAppResourceNamespaceResourceActionAuthenticationContextClient,
		EnterpriseAppResourceNamespaceResourceActionResourceScope:                 enterpriseAppResourceNamespaceResourceActionResourceScopeClient,
		EnterpriseAppRoleAssignment:                                               enterpriseAppRoleAssignmentClient,
		EnterpriseAppRoleAssignmentAppScope:                                       enterpriseAppRoleAssignmentAppScopeClient,
		EnterpriseAppRoleAssignmentApproval:                                       enterpriseAppRoleAssignmentApprovalClient,
		EnterpriseAppRoleAssignmentApprovalStep:                                   enterpriseAppRoleAssignmentApprovalStepClient,
		EnterpriseAppRoleAssignmentDirectoryScope:                                 enterpriseAppRoleAssignmentDirectoryScopeClient,
		EnterpriseAppRoleAssignmentPrincipal:                                      enterpriseAppRoleAssignmentPrincipalClient,
		EnterpriseAppRoleAssignmentRoleDefinition:                                 enterpriseAppRoleAssignmentRoleDefinitionClient,
		EnterpriseAppRoleAssignmentSchedule:                                       enterpriseAppRoleAssignmentScheduleClient,
		EnterpriseAppRoleAssignmentScheduleActivatedUsing:                         enterpriseAppRoleAssignmentScheduleActivatedUsingClient,
		EnterpriseAppRoleAssignmentScheduleAppScope:                               enterpriseAppRoleAssignmentScheduleAppScopeClient,
		EnterpriseAppRoleAssignmentScheduleDirectoryScope:                         enterpriseAppRoleAssignmentScheduleDirectoryScopeClient,
		EnterpriseAppRoleAssignmentScheduleInstance:                               enterpriseAppRoleAssignmentScheduleInstanceClient,
		EnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing:                 enterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient,
		EnterpriseAppRoleAssignmentScheduleInstanceAppScope:                       enterpriseAppRoleAssignmentScheduleInstanceAppScopeClient,
		EnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope:                 enterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient,
		EnterpriseAppRoleAssignmentScheduleInstancePrincipal:                      enterpriseAppRoleAssignmentScheduleInstancePrincipalClient,
		EnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition:                 enterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient,
		EnterpriseAppRoleAssignmentSchedulePrincipal:                              enterpriseAppRoleAssignmentSchedulePrincipalClient,
		EnterpriseAppRoleAssignmentScheduleRequest:                                enterpriseAppRoleAssignmentScheduleRequestClient,
		EnterpriseAppRoleAssignmentScheduleRequestActivatedUsing:                  enterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient,
		EnterpriseAppRoleAssignmentScheduleRequestAppScope:                        enterpriseAppRoleAssignmentScheduleRequestAppScopeClient,
		EnterpriseAppRoleAssignmentScheduleRequestDirectoryScope:                  enterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient,
		EnterpriseAppRoleAssignmentScheduleRequestPrincipal:                       enterpriseAppRoleAssignmentScheduleRequestPrincipalClient,
		EnterpriseAppRoleAssignmentScheduleRequestRoleDefinition:                  enterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient,
		EnterpriseAppRoleAssignmentScheduleRequestTargetSchedule:                  enterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient,
		EnterpriseAppRoleAssignmentScheduleRoleDefinition:                         enterpriseAppRoleAssignmentScheduleRoleDefinitionClient,
		EnterpriseAppRoleDefinition:                                               enterpriseAppRoleDefinitionClient,
		EnterpriseAppRoleDefinitionInheritsPermissionsFrom:                        enterpriseAppRoleDefinitionInheritsPermissionsFromClient,
		EnterpriseAppRoleEligibilitySchedule:                                      enterpriseAppRoleEligibilityScheduleClient,
		EnterpriseAppRoleEligibilityScheduleAppScope:                              enterpriseAppRoleEligibilityScheduleAppScopeClient,
		EnterpriseAppRoleEligibilityScheduleDirectoryScope:                        enterpriseAppRoleEligibilityScheduleDirectoryScopeClient,
		EnterpriseAppRoleEligibilityScheduleInstance:                              enterpriseAppRoleEligibilityScheduleInstanceClient,
		EnterpriseAppRoleEligibilityScheduleInstanceAppScope:                      enterpriseAppRoleEligibilityScheduleInstanceAppScopeClient,
		EnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope:                enterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient,
		EnterpriseAppRoleEligibilityScheduleInstancePrincipal:                     enterpriseAppRoleEligibilityScheduleInstancePrincipalClient,
		EnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition:                enterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient,
		EnterpriseAppRoleEligibilitySchedulePrincipal:                             enterpriseAppRoleEligibilitySchedulePrincipalClient,
		EnterpriseAppRoleEligibilityScheduleRequest:                               enterpriseAppRoleEligibilityScheduleRequestClient,
		EnterpriseAppRoleEligibilityScheduleRequestAppScope:                       enterpriseAppRoleEligibilityScheduleRequestAppScopeClient,
		EnterpriseAppRoleEligibilityScheduleRequestDirectoryScope:                 enterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient,
		EnterpriseAppRoleEligibilityScheduleRequestPrincipal:                      enterpriseAppRoleEligibilityScheduleRequestPrincipalClient,
		EnterpriseAppRoleEligibilityScheduleRequestRoleDefinition:                 enterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient,
		EnterpriseAppRoleEligibilityScheduleRequestTargetSchedule:                 enterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient,
		EnterpriseAppRoleEligibilityScheduleRoleDefinition:                        enterpriseAppRoleEligibilityScheduleRoleDefinitionClient,
		EnterpriseAppTransitiveRoleAssignment:                                     enterpriseAppTransitiveRoleAssignmentClient,
		EnterpriseAppTransitiveRoleAssignmentAppScope:                             enterpriseAppTransitiveRoleAssignmentAppScopeClient,
		EnterpriseAppTransitiveRoleAssignmentDirectoryScope:                       enterpriseAppTransitiveRoleAssignmentDirectoryScopeClient,
		EnterpriseAppTransitiveRoleAssignmentPrincipal:                            enterpriseAppTransitiveRoleAssignmentPrincipalClient,
		EnterpriseAppTransitiveRoleAssignmentRoleDefinition:                       enterpriseAppTransitiveRoleAssignmentRoleDefinitionClient,
		EntitlementManagement:                                                     entitlementManagementClient,
		EntitlementManagementResourceNamespace:                                    entitlementManagementResourceNamespaceClient,
		EntitlementManagementResourceNamespaceResourceAction:                      entitlementManagementResourceNamespaceResourceActionClient,
		EntitlementManagementResourceNamespaceResourceActionAuthenticationContext: entitlementManagementResourceNamespaceResourceActionAuthenticationContextClient,
		EntitlementManagementResourceNamespaceResourceActionResourceScope:         entitlementManagementResourceNamespaceResourceActionResourceScopeClient,
		EntitlementManagementRoleAssignment:                                       entitlementManagementRoleAssignmentClient,
		EntitlementManagementRoleAssignmentAppScope:                               entitlementManagementRoleAssignmentAppScopeClient,
		EntitlementManagementRoleAssignmentApproval:                               entitlementManagementRoleAssignmentApprovalClient,
		EntitlementManagementRoleAssignmentApprovalStep:                           entitlementManagementRoleAssignmentApprovalStepClient,
		EntitlementManagementRoleAssignmentDirectoryScope:                         entitlementManagementRoleAssignmentDirectoryScopeClient,
		EntitlementManagementRoleAssignmentPrincipal:                              entitlementManagementRoleAssignmentPrincipalClient,
		EntitlementManagementRoleAssignmentRoleDefinition:                         entitlementManagementRoleAssignmentRoleDefinitionClient,
		EntitlementManagementRoleAssignmentSchedule:                               entitlementManagementRoleAssignmentScheduleClient,
		EntitlementManagementRoleAssignmentScheduleActivatedUsing:                 entitlementManagementRoleAssignmentScheduleActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleAppScope:                       entitlementManagementRoleAssignmentScheduleAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleDirectoryScope:                 entitlementManagementRoleAssignmentScheduleDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstance:                       entitlementManagementRoleAssignmentScheduleInstanceClient,
		EntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing:         entitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleInstanceAppScope:               entitlementManagementRoleAssignmentScheduleInstanceAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope:         entitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleInstancePrincipal:              entitlementManagementRoleAssignmentScheduleInstancePrincipalClient,
		EntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition:         entitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient,
		EntitlementManagementRoleAssignmentSchedulePrincipal:                      entitlementManagementRoleAssignmentSchedulePrincipalClient,
		EntitlementManagementRoleAssignmentScheduleRequest:                        entitlementManagementRoleAssignmentScheduleRequestClient,
		EntitlementManagementRoleAssignmentScheduleRequestActivatedUsing:          entitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient,
		EntitlementManagementRoleAssignmentScheduleRequestAppScope:                entitlementManagementRoleAssignmentScheduleRequestAppScopeClient,
		EntitlementManagementRoleAssignmentScheduleRequestDirectoryScope:          entitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient,
		EntitlementManagementRoleAssignmentScheduleRequestPrincipal:               entitlementManagementRoleAssignmentScheduleRequestPrincipalClient,
		EntitlementManagementRoleAssignmentScheduleRequestRoleDefinition:          entitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient,
		EntitlementManagementRoleAssignmentScheduleRequestTargetSchedule:          entitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient,
		EntitlementManagementRoleAssignmentScheduleRoleDefinition:                 entitlementManagementRoleAssignmentScheduleRoleDefinitionClient,
		EntitlementManagementRoleDefinition:                                       entitlementManagementRoleDefinitionClient,
		EntitlementManagementRoleDefinitionInheritsPermissionsFrom:                entitlementManagementRoleDefinitionInheritsPermissionsFromClient,
		EntitlementManagementRoleEligibilitySchedule:                              entitlementManagementRoleEligibilityScheduleClient,
		EntitlementManagementRoleEligibilityScheduleAppScope:                      entitlementManagementRoleEligibilityScheduleAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleDirectoryScope:                entitlementManagementRoleEligibilityScheduleDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstance:                      entitlementManagementRoleEligibilityScheduleInstanceClient,
		EntitlementManagementRoleEligibilityScheduleInstanceAppScope:              entitlementManagementRoleEligibilityScheduleInstanceAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope:        entitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleInstancePrincipal:             entitlementManagementRoleEligibilityScheduleInstancePrincipalClient,
		EntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition:        entitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient,
		EntitlementManagementRoleEligibilitySchedulePrincipal:                     entitlementManagementRoleEligibilitySchedulePrincipalClient,
		EntitlementManagementRoleEligibilityScheduleRequest:                       entitlementManagementRoleEligibilityScheduleRequestClient,
		EntitlementManagementRoleEligibilityScheduleRequestAppScope:               entitlementManagementRoleEligibilityScheduleRequestAppScopeClient,
		EntitlementManagementRoleEligibilityScheduleRequestDirectoryScope:         entitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient,
		EntitlementManagementRoleEligibilityScheduleRequestPrincipal:              entitlementManagementRoleEligibilityScheduleRequestPrincipalClient,
		EntitlementManagementRoleEligibilityScheduleRequestRoleDefinition:         entitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient,
		EntitlementManagementRoleEligibilityScheduleRequestTargetSchedule:         entitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient,
		EntitlementManagementRoleEligibilityScheduleRoleDefinition:                entitlementManagementRoleEligibilityScheduleRoleDefinitionClient,
		EntitlementManagementTransitiveRoleAssignment:                             entitlementManagementTransitiveRoleAssignmentClient,
		EntitlementManagementTransitiveRoleAssignmentAppScope:                     entitlementManagementTransitiveRoleAssignmentAppScopeClient,
		EntitlementManagementTransitiveRoleAssignmentDirectoryScope:               entitlementManagementTransitiveRoleAssignmentDirectoryScopeClient,
		EntitlementManagementTransitiveRoleAssignmentPrincipal:                    entitlementManagementTransitiveRoleAssignmentPrincipalClient,
		EntitlementManagementTransitiveRoleAssignmentRoleDefinition:               entitlementManagementTransitiveRoleAssignmentRoleDefinitionClient,
		Exchange:                                exchangeClient,
		ExchangeCustomAppScope:                  exchangeCustomAppScopeClient,
		ExchangeResourceNamespace:               exchangeResourceNamespaceClient,
		ExchangeResourceNamespaceResourceAction: exchangeResourceNamespaceResourceActionClient,
		ExchangeResourceNamespaceResourceActionAuthenticationContext: exchangeResourceNamespaceResourceActionAuthenticationContextClient,
		ExchangeResourceNamespaceResourceActionResourceScope:         exchangeResourceNamespaceResourceActionResourceScopeClient,
		ExchangeRoleAssignment:                         exchangeRoleAssignmentClient,
		ExchangeRoleAssignmentAppScope:                 exchangeRoleAssignmentAppScopeClient,
		ExchangeRoleAssignmentDirectoryScope:           exchangeRoleAssignmentDirectoryScopeClient,
		ExchangeRoleAssignmentPrincipal:                exchangeRoleAssignmentPrincipalClient,
		ExchangeRoleAssignmentRoleDefinition:           exchangeRoleAssignmentRoleDefinitionClient,
		ExchangeRoleDefinition:                         exchangeRoleDefinitionClient,
		ExchangeRoleDefinitionInheritsPermissionsFrom:  exchangeRoleDefinitionInheritsPermissionsFromClient,
		ExchangeTransitiveRoleAssignment:               exchangeTransitiveRoleAssignmentClient,
		ExchangeTransitiveRoleAssignmentAppScope:       exchangeTransitiveRoleAssignmentAppScopeClient,
		ExchangeTransitiveRoleAssignmentDirectoryScope: exchangeTransitiveRoleAssignmentDirectoryScopeClient,
		ExchangeTransitiveRoleAssignmentPrincipal:      exchangeTransitiveRoleAssignmentPrincipalClient,
		ExchangeTransitiveRoleAssignmentRoleDefinition: exchangeTransitiveRoleAssignmentRoleDefinitionClient,
		RoleManagement:                                 roleManagementClient,
	}, nil
}
