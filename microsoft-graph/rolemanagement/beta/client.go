package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpc"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementcloudpcroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectory"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectoryroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectorytransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectorytransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectorytransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectorytransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementdirectorytransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapp"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseappresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseappresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseappresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapproleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapptransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapptransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapptransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapptransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententerpriseapptransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentapprovalstep"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleassignmentscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduledirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequestappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequestdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequestroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementroleeligibilityscheduleroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementtransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementtransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementtransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementtransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagemententitlementmanagementtransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchange"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangecustomappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeresourcenamespace"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeresourcenamespaceresourceaction"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeresourcenamespaceresourceactionauthenticationcontext"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeresourcenamespaceresourceactionresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroledefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangeroledefinitioninheritspermissionsfrom"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangetransitiveroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangetransitiveroleassignmentappscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangetransitiveroleassignmentdirectoryscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangetransitiveroleassignmentprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/rolemanagementexchangetransitiveroleassignmentroledefinition"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	RoleManagement                                                                          *rolemanagement.RoleManagementClient
	RoleManagementCloudPC                                                                   *rolemanagementcloudpc.RoleManagementCloudPCClient
	RoleManagementCloudPCResourceNamespace                                                  *rolemanagementcloudpcresourcenamespace.RoleManagementCloudPCResourceNamespaceClient
	RoleManagementCloudPCResourceNamespaceResourceAction                                    *rolemanagementcloudpcresourcenamespaceresourceaction.RoleManagementCloudPCResourceNamespaceResourceActionClient
	RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContext               *rolemanagementcloudpcresourcenamespaceresourceactionauthenticationcontext.RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementCloudPCResourceNamespaceResourceActionResourceScope                       *rolemanagementcloudpcresourcenamespaceresourceactionresourcescope.RoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient
	RoleManagementCloudPCRoleAssignment                                                     *rolemanagementcloudpcroleassignment.RoleManagementCloudPCRoleAssignmentClient
	RoleManagementCloudPCRoleAssignmentAppScope                                             *rolemanagementcloudpcroleassignmentappscope.RoleManagementCloudPCRoleAssignmentAppScopeClient
	RoleManagementCloudPCRoleAssignmentDirectoryScope                                       *rolemanagementcloudpcroleassignmentdirectoryscope.RoleManagementCloudPCRoleAssignmentDirectoryScopeClient
	RoleManagementCloudPCRoleAssignmentPrincipal                                            *rolemanagementcloudpcroleassignmentprincipal.RoleManagementCloudPCRoleAssignmentPrincipalClient
	RoleManagementCloudPCRoleAssignmentRoleDefinition                                       *rolemanagementcloudpcroleassignmentroledefinition.RoleManagementCloudPCRoleAssignmentRoleDefinitionClient
	RoleManagementCloudPCRoleDefinition                                                     *rolemanagementcloudpcroledefinition.RoleManagementCloudPCRoleDefinitionClient
	RoleManagementCloudPCRoleDefinitionInheritsPermissionsFrom                              *rolemanagementcloudpcroledefinitioninheritspermissionsfrom.RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient
	RoleManagementDeviceManagement                                                          *rolemanagementdevicemanagement.RoleManagementDeviceManagementClient
	RoleManagementDeviceManagementResourceNamespace                                         *rolemanagementdevicemanagementresourcenamespace.RoleManagementDeviceManagementResourceNamespaceClient
	RoleManagementDeviceManagementResourceNamespaceResourceAction                           *rolemanagementdevicemanagementresourcenamespaceresourceaction.RoleManagementDeviceManagementResourceNamespaceResourceActionClient
	RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContext      *rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext.RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementDeviceManagementResourceNamespaceResourceActionResourceScope              *rolemanagementdevicemanagementresourcenamespaceresourceactionresourcescope.RoleManagementDeviceManagementResourceNamespaceResourceActionResourceScopeClient
	RoleManagementDeviceManagementRoleAssignment                                            *rolemanagementdevicemanagementroleassignment.RoleManagementDeviceManagementRoleAssignmentClient
	RoleManagementDeviceManagementRoleAssignmentAppScope                                    *rolemanagementdevicemanagementroleassignmentappscope.RoleManagementDeviceManagementRoleAssignmentAppScopeClient
	RoleManagementDeviceManagementRoleAssignmentDirectoryScope                              *rolemanagementdevicemanagementroleassignmentdirectoryscope.RoleManagementDeviceManagementRoleAssignmentDirectoryScopeClient
	RoleManagementDeviceManagementRoleAssignmentPrincipal                                   *rolemanagementdevicemanagementroleassignmentprincipal.RoleManagementDeviceManagementRoleAssignmentPrincipalClient
	RoleManagementDeviceManagementRoleAssignmentRoleDefinition                              *rolemanagementdevicemanagementroleassignmentroledefinition.RoleManagementDeviceManagementRoleAssignmentRoleDefinitionClient
	RoleManagementDeviceManagementRoleDefinition                                            *rolemanagementdevicemanagementroledefinition.RoleManagementDeviceManagementRoleDefinitionClient
	RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFrom                     *rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom.RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient
	RoleManagementDirectory                                                                 *rolemanagementdirectory.RoleManagementDirectoryClient
	RoleManagementDirectoryResourceNamespace                                                *rolemanagementdirectoryresourcenamespace.RoleManagementDirectoryResourceNamespaceClient
	RoleManagementDirectoryResourceNamespaceResourceAction                                  *rolemanagementdirectoryresourcenamespaceresourceaction.RoleManagementDirectoryResourceNamespaceResourceActionClient
	RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContext             *rolemanagementdirectoryresourcenamespaceresourceactionauthenticationcontext.RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementDirectoryResourceNamespaceResourceActionResourceScope                     *rolemanagementdirectoryresourcenamespaceresourceactionresourcescope.RoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient
	RoleManagementDirectoryRoleAssignment                                                   *rolemanagementdirectoryroleassignment.RoleManagementDirectoryRoleAssignmentClient
	RoleManagementDirectoryRoleAssignmentAppScope                                           *rolemanagementdirectoryroleassignmentappscope.RoleManagementDirectoryRoleAssignmentAppScopeClient
	RoleManagementDirectoryRoleAssignmentApproval                                           *rolemanagementdirectoryroleassignmentapproval.RoleManagementDirectoryRoleAssignmentApprovalClient
	RoleManagementDirectoryRoleAssignmentApprovalStep                                       *rolemanagementdirectoryroleassignmentapprovalstep.RoleManagementDirectoryRoleAssignmentApprovalStepClient
	RoleManagementDirectoryRoleAssignmentDirectoryScope                                     *rolemanagementdirectoryroleassignmentdirectoryscope.RoleManagementDirectoryRoleAssignmentDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentPrincipal                                          *rolemanagementdirectoryroleassignmentprincipal.RoleManagementDirectoryRoleAssignmentPrincipalClient
	RoleManagementDirectoryRoleAssignmentRoleDefinition                                     *rolemanagementdirectoryroleassignmentroledefinition.RoleManagementDirectoryRoleAssignmentRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentSchedule                                           *rolemanagementdirectoryroleassignmentschedule.RoleManagementDirectoryRoleAssignmentScheduleClient
	RoleManagementDirectoryRoleAssignmentScheduleActivatedUsing                             *rolemanagementdirectoryroleassignmentscheduleactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleAppScope                                   *rolemanagementdirectoryroleassignmentscheduleappscope.RoleManagementDirectoryRoleAssignmentScheduleAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleDirectoryScope                             *rolemanagementdirectoryroleassignmentscheduledirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstance                                   *rolemanagementdirectoryroleassignmentscheduleinstance.RoleManagementDirectoryRoleAssignmentScheduleInstanceClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsing                     *rolemanagementdirectoryroleassignmentscheduleinstanceactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScope                           *rolemanagementdirectoryroleassignmentscheduleinstanceappscope.RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScope                     *rolemanagementdirectoryroleassignmentscheduleinstancedirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipal                          *rolemanagementdirectoryroleassignmentscheduleinstanceprincipal.RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinition                     *rolemanagementdirectoryroleassignmentscheduleinstanceroledefinition.RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentSchedulePrincipal                                  *rolemanagementdirectoryroleassignmentscheduleprincipal.RoleManagementDirectoryRoleAssignmentSchedulePrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleRequest                                    *rolemanagementdirectoryroleassignmentschedulerequest.RoleManagementDirectoryRoleAssignmentScheduleRequestClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsing                      *rolemanagementdirectoryroleassignmentschedulerequestactivatedusing.RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestAppScope                            *rolemanagementdirectoryroleassignmentschedulerequestappscope.RoleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScope                      *rolemanagementdirectoryroleassignmentschedulerequestdirectoryscope.RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipal                           *rolemanagementdirectoryroleassignmentschedulerequestprincipal.RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinition                      *rolemanagementdirectoryroleassignmentschedulerequestroledefinition.RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient
	RoleManagementDirectoryRoleAssignmentScheduleRequestTargetSchedule                      *rolemanagementdirectoryroleassignmentschedulerequesttargetschedule.RoleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient
	RoleManagementDirectoryRoleAssignmentScheduleRoleDefinition                             *rolemanagementdirectoryroleassignmentscheduleroledefinition.RoleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient
	RoleManagementDirectoryRoleDefinition                                                   *rolemanagementdirectoryroledefinition.RoleManagementDirectoryRoleDefinitionClient
	RoleManagementDirectoryRoleDefinitionInheritsPermissionsFrom                            *rolemanagementdirectoryroledefinitioninheritspermissionsfrom.RoleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient
	RoleManagementDirectoryRoleEligibilitySchedule                                          *rolemanagementdirectoryroleeligibilityschedule.RoleManagementDirectoryRoleEligibilityScheduleClient
	RoleManagementDirectoryRoleEligibilityScheduleAppScope                                  *rolemanagementdirectoryroleeligibilityscheduleappscope.RoleManagementDirectoryRoleEligibilityScheduleAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleDirectoryScope                            *rolemanagementdirectoryroleeligibilityscheduledirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstance                                  *rolemanagementdirectoryroleeligibilityscheduleinstance.RoleManagementDirectoryRoleEligibilityScheduleInstanceClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScope                          *rolemanagementdirectoryroleeligibilityscheduleinstanceappscope.RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScope                    *rolemanagementdirectoryroleeligibilityscheduleinstancedirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipal                         *rolemanagementdirectoryroleeligibilityscheduleinstanceprincipal.RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinition                    *rolemanagementdirectoryroleeligibilityscheduleinstanceroledefinition.RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient
	RoleManagementDirectoryRoleEligibilitySchedulePrincipal                                 *rolemanagementdirectoryroleeligibilityscheduleprincipal.RoleManagementDirectoryRoleEligibilitySchedulePrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleRequest                                   *rolemanagementdirectoryroleeligibilityschedulerequest.RoleManagementDirectoryRoleEligibilityScheduleRequestClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestAppScope                           *rolemanagementdirectoryroleeligibilityschedulerequestappscope.RoleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScope                     *rolemanagementdirectoryroleeligibilityschedulerequestdirectoryscope.RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipal                          *rolemanagementdirectoryroleeligibilityschedulerequestprincipal.RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinition                     *rolemanagementdirectoryroleeligibilityschedulerequestroledefinition.RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient
	RoleManagementDirectoryRoleEligibilityScheduleRequestTargetSchedule                     *rolemanagementdirectoryroleeligibilityschedulerequesttargetschedule.RoleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient
	RoleManagementDirectoryRoleEligibilityScheduleRoleDefinition                            *rolemanagementdirectoryroleeligibilityscheduleroledefinition.RoleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient
	RoleManagementDirectoryTransitiveRoleAssignment                                         *rolemanagementdirectorytransitiveroleassignment.RoleManagementDirectoryTransitiveRoleAssignmentClient
	RoleManagementDirectoryTransitiveRoleAssignmentAppScope                                 *rolemanagementdirectorytransitiveroleassignmentappscope.RoleManagementDirectoryTransitiveRoleAssignmentAppScopeClient
	RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScope                           *rolemanagementdirectorytransitiveroleassignmentdirectoryscope.RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient
	RoleManagementDirectoryTransitiveRoleAssignmentPrincipal                                *rolemanagementdirectorytransitiveroleassignmentprincipal.RoleManagementDirectoryTransitiveRoleAssignmentPrincipalClient
	RoleManagementDirectoryTransitiveRoleAssignmentRoleDefinition                           *rolemanagementdirectorytransitiveroleassignmentroledefinition.RoleManagementDirectoryTransitiveRoleAssignmentRoleDefinitionClient
	RoleManagementEnterpriseApp                                                             *rolemanagemententerpriseapp.RoleManagementEnterpriseAppClient
	RoleManagementEnterpriseAppResourceNamespace                                            *rolemanagemententerpriseappresourcenamespace.RoleManagementEnterpriseAppResourceNamespaceClient
	RoleManagementEnterpriseAppResourceNamespaceResourceAction                              *rolemanagemententerpriseappresourcenamespaceresourceaction.RoleManagementEnterpriseAppResourceNamespaceResourceActionClient
	RoleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContext         *rolemanagemententerpriseappresourcenamespaceresourceactionauthenticationcontext.RoleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScope                 *rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope.RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient
	RoleManagementEnterpriseAppRoleAssignment                                               *rolemanagemententerpriseapproleassignment.RoleManagementEnterpriseAppRoleAssignmentClient
	RoleManagementEnterpriseAppRoleAssignmentAppScope                                       *rolemanagemententerpriseapproleassignmentappscope.RoleManagementEnterpriseAppRoleAssignmentAppScopeClient
	RoleManagementEnterpriseAppRoleAssignmentApproval                                       *rolemanagemententerpriseapproleassignmentapproval.RoleManagementEnterpriseAppRoleAssignmentApprovalClient
	RoleManagementEnterpriseAppRoleAssignmentApprovalStep                                   *rolemanagemententerpriseapproleassignmentapprovalstep.RoleManagementEnterpriseAppRoleAssignmentApprovalStepClient
	RoleManagementEnterpriseAppRoleAssignmentDirectoryScope                                 *rolemanagemententerpriseapproleassignmentdirectoryscope.RoleManagementEnterpriseAppRoleAssignmentDirectoryScopeClient
	RoleManagementEnterpriseAppRoleAssignmentPrincipal                                      *rolemanagemententerpriseapproleassignmentprincipal.RoleManagementEnterpriseAppRoleAssignmentPrincipalClient
	RoleManagementEnterpriseAppRoleAssignmentRoleDefinition                                 *rolemanagemententerpriseapproleassignmentroledefinition.RoleManagementEnterpriseAppRoleAssignmentRoleDefinitionClient
	RoleManagementEnterpriseAppRoleAssignmentSchedule                                       *rolemanagemententerpriseapproleassignmentschedule.RoleManagementEnterpriseAppRoleAssignmentScheduleClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsing                         *rolemanagemententerpriseapproleassignmentscheduleactivatedusing.RoleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsingClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleAppScope                               *rolemanagemententerpriseapproleassignmentscheduleappscope.RoleManagementEnterpriseAppRoleAssignmentScheduleAppScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScope                         *rolemanagemententerpriseapproleassignmentscheduledirectoryscope.RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstance                               *rolemanagemententerpriseapproleassignmentscheduleinstance.RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing                 *rolemanagemententerpriseapproleassignmentscheduleinstanceactivatedusing.RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScope                       *rolemanagemententerpriseapproleassignmentscheduleinstanceappscope.RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope                 *rolemanagemententerpriseapproleassignmentscheduleinstancedirectoryscope.RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipal                      *rolemanagemententerpriseapproleassignmentscheduleinstanceprincipal.RoleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipalClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition                 *rolemanagemententerpriseapproleassignmentscheduleinstanceroledefinition.RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient
	RoleManagementEnterpriseAppRoleAssignmentSchedulePrincipal                              *rolemanagemententerpriseapproleassignmentscheduleprincipal.RoleManagementEnterpriseAppRoleAssignmentSchedulePrincipalClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequest                                *rolemanagemententerpriseapproleassignmentschedulerequest.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsing                  *rolemanagemententerpriseapproleassignmentschedulerequestactivatedusing.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScope                        *rolemanagemententerpriseapproleassignmentschedulerequestappscope.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScope                  *rolemanagemententerpriseapproleassignmentschedulerequestdirectoryscope.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipal                       *rolemanagemententerpriseapproleassignmentschedulerequestprincipal.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipalClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinition                  *rolemanagemententerpriseapproleassignmentschedulerequestroledefinition.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetSchedule                  *rolemanagemententerpriseapproleassignmentschedulerequesttargetschedule.RoleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient
	RoleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinition                         *rolemanagemententerpriseapproleassignmentscheduleroledefinition.RoleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinitionClient
	RoleManagementEnterpriseAppRoleDefinition                                               *rolemanagemententerpriseapproledefinition.RoleManagementEnterpriseAppRoleDefinitionClient
	RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFrom                        *rolemanagemententerpriseapproledefinitioninheritspermissionsfrom.RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient
	RoleManagementEnterpriseAppRoleEligibilitySchedule                                      *rolemanagemententerpriseapproleeligibilityschedule.RoleManagementEnterpriseAppRoleEligibilityScheduleClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleAppScope                              *rolemanagemententerpriseapproleeligibilityscheduleappscope.RoleManagementEnterpriseAppRoleEligibilityScheduleAppScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScope                        *rolemanagemententerpriseapproleeligibilityscheduledirectoryscope.RoleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleInstance                              *rolemanagemententerpriseapproleeligibilityscheduleinstance.RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScope                      *rolemanagemententerpriseapproleeligibilityscheduleinstanceappscope.RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope                *rolemanagemententerpriseapproleeligibilityscheduleinstancedirectoryscope.RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipal                     *rolemanagemententerpriseapproleeligibilityscheduleinstanceprincipal.RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition                *rolemanagemententerpriseapproleeligibilityscheduleinstanceroledefinition.RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient
	RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipal                             *rolemanagemententerpriseapproleeligibilityscheduleprincipal.RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequest                               *rolemanagemententerpriseapproleeligibilityschedulerequest.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScope                       *rolemanagemententerpriseapproleeligibilityschedulerequestappscope.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScope                 *rolemanagemententerpriseapproleeligibilityschedulerequestdirectoryscope.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipal                      *rolemanagemententerpriseapproleeligibilityschedulerequestprincipal.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipalClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinition                 *rolemanagemententerpriseapproleeligibilityschedulerequestroledefinition.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetSchedule                 *rolemanagemententerpriseapproleeligibilityschedulerequesttargetschedule.RoleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient
	RoleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinition                        *rolemanagemententerpriseapproleeligibilityscheduleroledefinition.RoleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinitionClient
	RoleManagementEnterpriseAppTransitiveRoleAssignment                                     *rolemanagemententerpriseapptransitiveroleassignment.RoleManagementEnterpriseAppTransitiveRoleAssignmentClient
	RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScope                             *rolemanagemententerpriseapptransitiveroleassignmentappscope.RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient
	RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScope                       *rolemanagemententerpriseapptransitiveroleassignmentdirectoryscope.RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient
	RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipal                            *rolemanagemententerpriseapptransitiveroleassignmentprincipal.RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient
	RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinition                       *rolemanagemententerpriseapptransitiveroleassignmentroledefinition.RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient
	RoleManagementEntitlementManagement                                                     *rolemanagemententitlementmanagement.RoleManagementEntitlementManagementClient
	RoleManagementEntitlementManagementResourceNamespace                                    *rolemanagemententitlementmanagementresourcenamespace.RoleManagementEntitlementManagementResourceNamespaceClient
	RoleManagementEntitlementManagementResourceNamespaceResourceAction                      *rolemanagemententitlementmanagementresourcenamespaceresourceaction.RoleManagementEntitlementManagementResourceNamespaceResourceActionClient
	RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContext *rolemanagemententitlementmanagementresourcenamespaceresourceactionauthenticationcontext.RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScope         *rolemanagemententitlementmanagementresourcenamespaceresourceactionresourcescope.RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient
	RoleManagementEntitlementManagementRoleAssignment                                       *rolemanagemententitlementmanagementroleassignment.RoleManagementEntitlementManagementRoleAssignmentClient
	RoleManagementEntitlementManagementRoleAssignmentAppScope                               *rolemanagemententitlementmanagementroleassignmentappscope.RoleManagementEntitlementManagementRoleAssignmentAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentApproval                               *rolemanagemententitlementmanagementroleassignmentapproval.RoleManagementEntitlementManagementRoleAssignmentApprovalClient
	RoleManagementEntitlementManagementRoleAssignmentApprovalStep                           *rolemanagemententitlementmanagementroleassignmentapprovalstep.RoleManagementEntitlementManagementRoleAssignmentApprovalStepClient
	RoleManagementEntitlementManagementRoleAssignmentDirectoryScope                         *rolemanagemententitlementmanagementroleassignmentdirectoryscope.RoleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentPrincipal                              *rolemanagemententitlementmanagementroleassignmentprincipal.RoleManagementEntitlementManagementRoleAssignmentPrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentRoleDefinition                         *rolemanagemententitlementmanagementroleassignmentroledefinition.RoleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentSchedule                               *rolemanagemententitlementmanagementroleassignmentschedule.RoleManagementEntitlementManagementRoleAssignmentScheduleClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsing                 *rolemanagemententitlementmanagementroleassignmentscheduleactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleAppScope                       *rolemanagemententitlementmanagementroleassignmentscheduleappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScope                 *rolemanagemententitlementmanagementroleassignmentscheduledirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstance                       *rolemanagemententitlementmanagementroleassignmentscheduleinstance.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing         *rolemanagemententitlementmanagementroleassignmentscheduleinstanceactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScope               *rolemanagemententitlementmanagementroleassignmentscheduleinstanceappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope         *rolemanagemententitlementmanagementroleassignmentscheduleinstancedirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipal              *rolemanagemententitlementmanagementroleassignmentscheduleinstanceprincipal.RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition         *rolemanagemententitlementmanagementroleassignmentscheduleinstanceroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipal                      *rolemanagemententitlementmanagementroleassignmentscheduleprincipal.RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequest                        *rolemanagemententitlementmanagementroleassignmentschedulerequest.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsing          *rolemanagemententitlementmanagementroleassignmentschedulerequestactivatedusing.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScope                *rolemanagemententitlementmanagementroleassignmentschedulerequestappscope.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScope          *rolemanagemententitlementmanagementroleassignmentschedulerequestdirectoryscope.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipal               *rolemanagemententitlementmanagementroleassignmentschedulerequestprincipal.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinition          *rolemanagemententitlementmanagementroleassignmentschedulerequestroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule          *rolemanagemententitlementmanagementroleassignmentschedulerequesttargetschedule.RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient
	RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinition                 *rolemanagemententitlementmanagementroleassignmentscheduleroledefinition.RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient
	RoleManagementEntitlementManagementRoleDefinition                                       *rolemanagemententitlementmanagementroledefinition.RoleManagementEntitlementManagementRoleDefinitionClient
	RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFrom                *rolemanagemententitlementmanagementroledefinitioninheritspermissionsfrom.RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient
	RoleManagementEntitlementManagementRoleEligibilitySchedule                              *rolemanagemententitlementmanagementroleeligibilityschedule.RoleManagementEntitlementManagementRoleEligibilityScheduleClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleAppScope                      *rolemanagemententitlementmanagementroleeligibilityscheduleappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScope                *rolemanagemententitlementmanagementroleeligibilityscheduledirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstance                      *rolemanagemententitlementmanagementroleeligibilityscheduleinstance.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScope              *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope        *rolemanagemententitlementmanagementroleeligibilityscheduleinstancedirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipal             *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceprincipal.RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition        *rolemanagemententitlementmanagementroleeligibilityscheduleinstanceroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient
	RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipal                     *rolemanagemententitlementmanagementroleeligibilityscheduleprincipal.RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequest                       *rolemanagemententitlementmanagementroleeligibilityschedulerequest.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScope               *rolemanagemententitlementmanagementroleeligibilityschedulerequestappscope.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope         *rolemanagemententitlementmanagementroleeligibilityschedulerequestdirectoryscope.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipal              *rolemanagemententitlementmanagementroleeligibilityschedulerequestprincipal.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinition         *rolemanagemententitlementmanagementroleeligibilityschedulerequestroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetSchedule         *rolemanagemententitlementmanagementroleeligibilityschedulerequesttargetschedule.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient
	RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinition                *rolemanagemententitlementmanagementroleeligibilityscheduleroledefinition.RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient
	RoleManagementEntitlementManagementTransitiveRoleAssignment                             *rolemanagemententitlementmanagementtransitiveroleassignment.RoleManagementEntitlementManagementTransitiveRoleAssignmentClient
	RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScope                     *rolemanagemententitlementmanagementtransitiveroleassignmentappscope.RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient
	RoleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScope               *rolemanagemententitlementmanagementtransitiveroleassignmentdirectoryscope.RoleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient
	RoleManagementEntitlementManagementTransitiveRoleAssignmentPrincipal                    *rolemanagemententitlementmanagementtransitiveroleassignmentprincipal.RoleManagementEntitlementManagementTransitiveRoleAssignmentPrincipalClient
	RoleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinition               *rolemanagemententitlementmanagementtransitiveroleassignmentroledefinition.RoleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClient
	RoleManagementExchange                                                                  *rolemanagementexchange.RoleManagementExchangeClient
	RoleManagementExchangeCustomAppScope                                                    *rolemanagementexchangecustomappscope.RoleManagementExchangeCustomAppScopeClient
	RoleManagementExchangeResourceNamespace                                                 *rolemanagementexchangeresourcenamespace.RoleManagementExchangeResourceNamespaceClient
	RoleManagementExchangeResourceNamespaceResourceAction                                   *rolemanagementexchangeresourcenamespaceresourceaction.RoleManagementExchangeResourceNamespaceResourceActionClient
	RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContext              *rolemanagementexchangeresourcenamespaceresourceactionauthenticationcontext.RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient
	RoleManagementExchangeResourceNamespaceResourceActionResourceScope                      *rolemanagementexchangeresourcenamespaceresourceactionresourcescope.RoleManagementExchangeResourceNamespaceResourceActionResourceScopeClient
	RoleManagementExchangeRoleAssignment                                                    *rolemanagementexchangeroleassignment.RoleManagementExchangeRoleAssignmentClient
	RoleManagementExchangeRoleAssignmentAppScope                                            *rolemanagementexchangeroleassignmentappscope.RoleManagementExchangeRoleAssignmentAppScopeClient
	RoleManagementExchangeRoleAssignmentDirectoryScope                                      *rolemanagementexchangeroleassignmentdirectoryscope.RoleManagementExchangeRoleAssignmentDirectoryScopeClient
	RoleManagementExchangeRoleAssignmentPrincipal                                           *rolemanagementexchangeroleassignmentprincipal.RoleManagementExchangeRoleAssignmentPrincipalClient
	RoleManagementExchangeRoleAssignmentRoleDefinition                                      *rolemanagementexchangeroleassignmentroledefinition.RoleManagementExchangeRoleAssignmentRoleDefinitionClient
	RoleManagementExchangeRoleDefinition                                                    *rolemanagementexchangeroledefinition.RoleManagementExchangeRoleDefinitionClient
	RoleManagementExchangeRoleDefinitionInheritsPermissionsFrom                             *rolemanagementexchangeroledefinitioninheritspermissionsfrom.RoleManagementExchangeRoleDefinitionInheritsPermissionsFromClient
	RoleManagementExchangeTransitiveRoleAssignment                                          *rolemanagementexchangetransitiveroleassignment.RoleManagementExchangeTransitiveRoleAssignmentClient
	RoleManagementExchangeTransitiveRoleAssignmentAppScope                                  *rolemanagementexchangetransitiveroleassignmentappscope.RoleManagementExchangeTransitiveRoleAssignmentAppScopeClient
	RoleManagementExchangeTransitiveRoleAssignmentDirectoryScope                            *rolemanagementexchangetransitiveroleassignmentdirectoryscope.RoleManagementExchangeTransitiveRoleAssignmentDirectoryScopeClient
	RoleManagementExchangeTransitiveRoleAssignmentPrincipal                                 *rolemanagementexchangetransitiveroleassignmentprincipal.RoleManagementExchangeTransitiveRoleAssignmentPrincipalClient
	RoleManagementExchangeTransitiveRoleAssignmentRoleDefinition                            *rolemanagementexchangetransitiveroleassignmentroledefinition.RoleManagementExchangeTransitiveRoleAssignmentRoleDefinitionClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	roleManagementClient, err := rolemanagement.NewRoleManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagement client: %+v", err)
	}
	configureFunc(roleManagementClient.Client)

	roleManagementCloudPCClient, err := rolemanagementcloudpc.NewRoleManagementCloudPCClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPC client: %+v", err)
	}
	configureFunc(roleManagementCloudPCClient.Client)

	roleManagementCloudPCResourceNamespaceClient, err := rolemanagementcloudpcresourcenamespace.NewRoleManagementCloudPCResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementCloudPCResourceNamespaceClient.Client)

	roleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagementcloudpcresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementCloudPCResourceNamespaceResourceActionClient, err := rolemanagementcloudpcresourcenamespaceresourceaction.NewRoleManagementCloudPCResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementCloudPCResourceNamespaceResourceActionClient.Client)

	roleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient, err := rolemanagementcloudpcresourcenamespaceresourceactionresourcescope.NewRoleManagementCloudPCResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementCloudPCRoleAssignmentAppScopeClient, err := rolemanagementcloudpcroleassignmentappscope.NewRoleManagementCloudPCRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleAssignmentAppScopeClient.Client)

	roleManagementCloudPCRoleAssignmentClient, err := rolemanagementcloudpcroleassignment.NewRoleManagementCloudPCRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleAssignmentClient.Client)

	roleManagementCloudPCRoleAssignmentDirectoryScopeClient, err := rolemanagementcloudpcroleassignmentdirectoryscope.NewRoleManagementCloudPCRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleAssignmentDirectoryScopeClient.Client)

	roleManagementCloudPCRoleAssignmentPrincipalClient, err := rolemanagementcloudpcroleassignmentprincipal.NewRoleManagementCloudPCRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleAssignmentPrincipalClient.Client)

	roleManagementCloudPCRoleAssignmentRoleDefinitionClient, err := rolemanagementcloudpcroleassignmentroledefinition.NewRoleManagementCloudPCRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleAssignmentRoleDefinitionClient.Client)

	roleManagementCloudPCRoleDefinitionClient, err := rolemanagementcloudpcroledefinition.NewRoleManagementCloudPCRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleDefinitionClient.Client)

	roleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient, err := rolemanagementcloudpcroledefinitioninheritspermissionsfrom.NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementCloudPCRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient.Client)

	roleManagementDeviceManagementClient, err := rolemanagementdevicemanagement.NewRoleManagementDeviceManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagement client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementClient.Client)

	roleManagementDeviceManagementResourceNamespaceClient, err := rolemanagementdevicemanagementresourcenamespace.NewRoleManagementDeviceManagementResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementResourceNamespaceClient.Client)

	roleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementDeviceManagementResourceNamespaceResourceActionClient, err := rolemanagementdevicemanagementresourcenamespaceresourceaction.NewRoleManagementDeviceManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementResourceNamespaceResourceActionClient.Client)

	roleManagementDeviceManagementResourceNamespaceResourceActionResourceScopeClient, err := rolemanagementdevicemanagementresourcenamespaceresourceactionresourcescope.NewRoleManagementDeviceManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementDeviceManagementRoleAssignmentAppScopeClient, err := rolemanagementdevicemanagementroleassignmentappscope.NewRoleManagementDeviceManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleAssignmentAppScopeClient.Client)

	roleManagementDeviceManagementRoleAssignmentClient, err := rolemanagementdevicemanagementroleassignment.NewRoleManagementDeviceManagementRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleAssignmentClient.Client)

	roleManagementDeviceManagementRoleAssignmentDirectoryScopeClient, err := rolemanagementdevicemanagementroleassignmentdirectoryscope.NewRoleManagementDeviceManagementRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleAssignmentDirectoryScopeClient.Client)

	roleManagementDeviceManagementRoleAssignmentPrincipalClient, err := rolemanagementdevicemanagementroleassignmentprincipal.NewRoleManagementDeviceManagementRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleAssignmentPrincipalClient.Client)

	roleManagementDeviceManagementRoleAssignmentRoleDefinitionClient, err := rolemanagementdevicemanagementroleassignmentroledefinition.NewRoleManagementDeviceManagementRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleAssignmentRoleDefinitionClient.Client)

	roleManagementDeviceManagementRoleDefinitionClient, err := rolemanagementdevicemanagementroledefinition.NewRoleManagementDeviceManagementRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleDefinitionClient.Client)

	roleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient, err := rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom.NewRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient.Client)

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

	roleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagementdirectoryresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementDirectoryResourceNamespaceResourceActionClient, err := rolemanagementdirectoryresourcenamespaceresourceaction.NewRoleManagementDirectoryResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementDirectoryResourceNamespaceResourceActionClient.Client)

	roleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient, err := rolemanagementdirectoryresourcenamespaceresourceactionresourcescope.NewRoleManagementDirectoryResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementDirectoryRoleAssignmentAppScopeClient, err := rolemanagementdirectoryroleassignmentappscope.NewRoleManagementDirectoryRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentAppScopeClient.Client)

	roleManagementDirectoryRoleAssignmentApprovalClient, err := rolemanagementdirectoryroleassignmentapproval.NewRoleManagementDirectoryRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentApprovalClient.Client)

	roleManagementDirectoryRoleAssignmentApprovalStepClient, err := rolemanagementdirectoryroleassignmentapprovalstep.NewRoleManagementDirectoryRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(roleManagementDirectoryRoleAssignmentApprovalStepClient.Client)

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

	roleManagementDirectoryTransitiveRoleAssignmentAppScopeClient, err := rolemanagementdirectorytransitiveroleassignmentappscope.NewRoleManagementDirectoryTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryTransitiveRoleAssignmentAppScopeClient.Client)

	roleManagementDirectoryTransitiveRoleAssignmentClient, err := rolemanagementdirectorytransitiveroleassignment.NewRoleManagementDirectoryTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementDirectoryTransitiveRoleAssignmentClient.Client)

	roleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient, err := rolemanagementdirectorytransitiveroleassignmentdirectoryscope.NewRoleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient.Client)

	roleManagementDirectoryTransitiveRoleAssignmentPrincipalClient, err := rolemanagementdirectorytransitiveroleassignmentprincipal.NewRoleManagementDirectoryTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementDirectoryTransitiveRoleAssignmentPrincipalClient.Client)

	roleManagementDirectoryTransitiveRoleAssignmentRoleDefinitionClient, err := rolemanagementdirectorytransitiveroleassignmentroledefinition.NewRoleManagementDirectoryTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementDirectoryTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementDirectoryTransitiveRoleAssignmentRoleDefinitionClient.Client)

	roleManagementEnterpriseAppClient, err := rolemanagemententerpriseapp.NewRoleManagementEnterpriseAppClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseApp client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppClient.Client)

	roleManagementEnterpriseAppResourceNamespaceClient, err := rolemanagemententerpriseappresourcenamespace.NewRoleManagementEnterpriseAppResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppResourceNamespaceClient.Client)

	roleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagemententerpriseappresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementEnterpriseAppResourceNamespaceResourceActionClient, err := rolemanagemententerpriseappresourcenamespaceresourceaction.NewRoleManagementEnterpriseAppResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppResourceNamespaceResourceActionClient.Client)

	roleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient, err := rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope.NewRoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentAppScopeClient, err := rolemanagemententerpriseapproleassignmentappscope.NewRoleManagementEnterpriseAppRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentAppScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentApprovalClient, err := rolemanagemententerpriseapproleassignmentapproval.NewRoleManagementEnterpriseAppRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentApprovalClient.Client)

	roleManagementEnterpriseAppRoleAssignmentApprovalStepClient, err := rolemanagemententerpriseapproleassignmentapprovalstep.NewRoleManagementEnterpriseAppRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentApprovalStepClient.Client)

	roleManagementEnterpriseAppRoleAssignmentClient, err := rolemanagemententerpriseapproleassignment.NewRoleManagementEnterpriseAppRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentClient.Client)

	roleManagementEnterpriseAppRoleAssignmentDirectoryScopeClient, err := rolemanagemententerpriseapproleassignmentdirectoryscope.NewRoleManagementEnterpriseAppRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentPrincipalClient, err := rolemanagemententerpriseapproleassignmentprincipal.NewRoleManagementEnterpriseAppRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentPrincipalClient.Client)

	roleManagementEnterpriseAppRoleAssignmentRoleDefinitionClient, err := rolemanagemententerpriseapproleassignmentroledefinition.NewRoleManagementEnterpriseAppRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsingClient, err := rolemanagemententerpriseapproleassignmentscheduleactivatedusing.NewRoleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsingClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleAppScopeClient, err := rolemanagemententerpriseapproleassignmentscheduleappscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleAppScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleClient, err := rolemanagemententerpriseapproleassignmentschedule.NewRoleManagementEnterpriseAppRoleAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentSchedule client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient, err := rolemanagemententerpriseapproleassignmentscheduledirectoryscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient, err := rolemanagemententerpriseapproleassignmentscheduleinstanceactivatedusing.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient, err := rolemanagemententerpriseapproleassignmentscheduleinstanceappscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient, err := rolemanagemententerpriseapproleassignmentscheduleinstance.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient, err := rolemanagemententerpriseapproleassignmentscheduleinstancedirectoryscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipalClient, err := rolemanagemententerpriseapproleassignmentscheduleinstanceprincipal.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipalClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient, err := rolemanagemententerpriseapproleassignmentscheduleinstanceroledefinition.NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleAssignmentSchedulePrincipalClient, err := rolemanagemententerpriseapproleassignmentscheduleprincipal.NewRoleManagementEnterpriseAppRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentSchedulePrincipalClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient, err := rolemanagemententerpriseapproleassignmentschedulerequestactivatedusing.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScopeClient, err := rolemanagemententerpriseapproleassignmentschedulerequestappscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestClient, err := rolemanagemententerpriseapproleassignmentschedulerequest.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient, err := rolemanagemententerpriseapproleassignmentschedulerequestdirectoryscope.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipalClient, err := rolemanagemententerpriseapproleassignmentschedulerequestprincipal.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipalClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient, err := rolemanagemententerpriseapproleassignmentschedulerequestroledefinition.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient, err := rolemanagemententerpriseapproleassignmentschedulerequesttargetschedule.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient.Client)

	roleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinitionClient, err := rolemanagemententerpriseapproleassignmentscheduleroledefinition.NewRoleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleDefinitionClient, err := rolemanagemententerpriseapproledefinition.NewRoleManagementEnterpriseAppRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient, err := rolemanagemententerpriseapproledefinitioninheritspermissionsfrom.NewRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleAppScopeClient, err := rolemanagemententerpriseapproleeligibilityscheduleappscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleAppScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleClient, err := rolemanagemententerpriseapproleeligibilityschedule.NewRoleManagementEnterpriseAppRoleEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilitySchedule client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScopeClient, err := rolemanagemententerpriseapproleeligibilityscheduledirectoryscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient, err := rolemanagemententerpriseapproleeligibilityscheduleinstanceappscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient, err := rolemanagemententerpriseapproleeligibilityscheduleinstance.NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient, err := rolemanagemententerpriseapproleeligibilityscheduleinstancedirectoryscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient, err := rolemanagemententerpriseapproleeligibilityscheduleinstanceprincipal.NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient, err := rolemanagemententerpriseapproleeligibilityscheduleinstanceroledefinition.NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient, err := rolemanagemententerpriseapproleeligibilityscheduleprincipal.NewRoleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScopeClient, err := rolemanagemententerpriseapproleeligibilityschedulerequestappscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestClient, err := rolemanagemententerpriseapproleeligibilityschedulerequest.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient, err := rolemanagemententerpriseapproleeligibilityschedulerequestdirectoryscope.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipalClient, err := rolemanagemententerpriseapproleeligibilityschedulerequestprincipal.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipalClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient, err := rolemanagemententerpriseapproleeligibilityschedulerequestroledefinition.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient, err := rolemanagemententerpriseapproleeligibilityschedulerequesttargetschedule.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient.Client)

	roleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinitionClient, err := rolemanagemententerpriseapproleeligibilityscheduleroledefinition.NewRoleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinitionClient.Client)

	roleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient, err := rolemanagemententerpriseapptransitiveroleassignmentappscope.NewRoleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient.Client)

	roleManagementEnterpriseAppTransitiveRoleAssignmentClient, err := rolemanagemententerpriseapptransitiveroleassignment.NewRoleManagementEnterpriseAppTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppTransitiveRoleAssignmentClient.Client)

	roleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient, err := rolemanagemententerpriseapptransitiveroleassignmentdirectoryscope.NewRoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient.Client)

	roleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient, err := rolemanagemententerpriseapptransitiveroleassignmentprincipal.NewRoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient.Client)

	roleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient, err := rolemanagemententerpriseapptransitiveroleassignmentroledefinition.NewRoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient.Client)

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

	roleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagemententitlementmanagementresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementEntitlementManagementResourceNamespaceResourceActionClient, err := rolemanagemententitlementmanagementresourcenamespaceresourceaction.NewRoleManagementEntitlementManagementResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementResourceNamespaceResourceActionClient.Client)

	roleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient, err := rolemanagemententitlementmanagementresourcenamespaceresourceactionresourcescope.NewRoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentAppScopeClient, err := rolemanagemententitlementmanagementroleassignmentappscope.NewRoleManagementEntitlementManagementRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentAppScopeClient.Client)

	roleManagementEntitlementManagementRoleAssignmentApprovalClient, err := rolemanagemententitlementmanagementroleassignmentapproval.NewRoleManagementEntitlementManagementRoleAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentApproval client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentApprovalClient.Client)

	roleManagementEntitlementManagementRoleAssignmentApprovalStepClient, err := rolemanagemententitlementmanagementroleassignmentapprovalstep.NewRoleManagementEntitlementManagementRoleAssignmentApprovalStepClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementRoleAssignmentApprovalStep client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementRoleAssignmentApprovalStepClient.Client)

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

	roleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient, err := rolemanagemententitlementmanagementtransitiveroleassignmentappscope.NewRoleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient.Client)

	roleManagementEntitlementManagementTransitiveRoleAssignmentClient, err := rolemanagemententitlementmanagementtransitiveroleassignment.NewRoleManagementEntitlementManagementTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementTransitiveRoleAssignmentClient.Client)

	roleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient, err := rolemanagemententitlementmanagementtransitiveroleassignmentdirectoryscope.NewRoleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient.Client)

	roleManagementEntitlementManagementTransitiveRoleAssignmentPrincipalClient, err := rolemanagemententitlementmanagementtransitiveroleassignmentprincipal.NewRoleManagementEntitlementManagementTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementTransitiveRoleAssignmentPrincipalClient.Client)

	roleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClient, err := rolemanagemententitlementmanagementtransitiveroleassignmentroledefinition.NewRoleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClient.Client)

	roleManagementExchangeClient, err := rolemanagementexchange.NewRoleManagementExchangeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchange client: %+v", err)
	}
	configureFunc(roleManagementExchangeClient.Client)

	roleManagementExchangeCustomAppScopeClient, err := rolemanagementexchangecustomappscope.NewRoleManagementExchangeCustomAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeCustomAppScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeCustomAppScopeClient.Client)

	roleManagementExchangeResourceNamespaceClient, err := rolemanagementexchangeresourcenamespace.NewRoleManagementExchangeResourceNamespaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeResourceNamespace client: %+v", err)
	}
	configureFunc(roleManagementExchangeResourceNamespaceClient.Client)

	roleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient, err := rolemanagementexchangeresourcenamespaceresourceactionauthenticationcontext.NewRoleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContext client: %+v", err)
	}
	configureFunc(roleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient.Client)

	roleManagementExchangeResourceNamespaceResourceActionClient, err := rolemanagementexchangeresourcenamespaceresourceaction.NewRoleManagementExchangeResourceNamespaceResourceActionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeResourceNamespaceResourceAction client: %+v", err)
	}
	configureFunc(roleManagementExchangeResourceNamespaceResourceActionClient.Client)

	roleManagementExchangeResourceNamespaceResourceActionResourceScopeClient, err := rolemanagementexchangeresourcenamespaceresourceactionresourcescope.NewRoleManagementExchangeResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeResourceNamespaceResourceActionResourceScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeResourceNamespaceResourceActionResourceScopeClient.Client)

	roleManagementExchangeRoleAssignmentAppScopeClient, err := rolemanagementexchangeroleassignmentappscope.NewRoleManagementExchangeRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleAssignmentAppScopeClient.Client)

	roleManagementExchangeRoleAssignmentClient, err := rolemanagementexchangeroleassignment.NewRoleManagementExchangeRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleAssignmentClient.Client)

	roleManagementExchangeRoleAssignmentDirectoryScopeClient, err := rolemanagementexchangeroleassignmentdirectoryscope.NewRoleManagementExchangeRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleAssignmentDirectoryScopeClient.Client)

	roleManagementExchangeRoleAssignmentPrincipalClient, err := rolemanagementexchangeroleassignmentprincipal.NewRoleManagementExchangeRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleAssignmentPrincipalClient.Client)

	roleManagementExchangeRoleAssignmentRoleDefinitionClient, err := rolemanagementexchangeroleassignmentroledefinition.NewRoleManagementExchangeRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleAssignmentRoleDefinitionClient.Client)

	roleManagementExchangeRoleDefinitionClient, err := rolemanagementexchangeroledefinition.NewRoleManagementExchangeRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleDefinitionClient.Client)

	roleManagementExchangeRoleDefinitionInheritsPermissionsFromClient, err := rolemanagementexchangeroledefinitioninheritspermissionsfrom.NewRoleManagementExchangeRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeRoleDefinitionInheritsPermissionsFrom client: %+v", err)
	}
	configureFunc(roleManagementExchangeRoleDefinitionInheritsPermissionsFromClient.Client)

	roleManagementExchangeTransitiveRoleAssignmentAppScopeClient, err := rolemanagementexchangetransitiveroleassignmentappscope.NewRoleManagementExchangeTransitiveRoleAssignmentAppScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeTransitiveRoleAssignmentAppScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeTransitiveRoleAssignmentAppScopeClient.Client)

	roleManagementExchangeTransitiveRoleAssignmentClient, err := rolemanagementexchangetransitiveroleassignment.NewRoleManagementExchangeTransitiveRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeTransitiveRoleAssignment client: %+v", err)
	}
	configureFunc(roleManagementExchangeTransitiveRoleAssignmentClient.Client)

	roleManagementExchangeTransitiveRoleAssignmentDirectoryScopeClient, err := rolemanagementexchangetransitiveroleassignmentdirectoryscope.NewRoleManagementExchangeTransitiveRoleAssignmentDirectoryScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeTransitiveRoleAssignmentDirectoryScope client: %+v", err)
	}
	configureFunc(roleManagementExchangeTransitiveRoleAssignmentDirectoryScopeClient.Client)

	roleManagementExchangeTransitiveRoleAssignmentPrincipalClient, err := rolemanagementexchangetransitiveroleassignmentprincipal.NewRoleManagementExchangeTransitiveRoleAssignmentPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeTransitiveRoleAssignmentPrincipal client: %+v", err)
	}
	configureFunc(roleManagementExchangeTransitiveRoleAssignmentPrincipalClient.Client)

	roleManagementExchangeTransitiveRoleAssignmentRoleDefinitionClient, err := rolemanagementexchangetransitiveroleassignmentroledefinition.NewRoleManagementExchangeTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RoleManagementExchangeTransitiveRoleAssignmentRoleDefinition client: %+v", err)
	}
	configureFunc(roleManagementExchangeTransitiveRoleAssignmentRoleDefinitionClient.Client)

	return &Client{
		RoleManagement:                                       roleManagementClient,
		RoleManagementCloudPC:                                roleManagementCloudPCClient,
		RoleManagementCloudPCResourceNamespace:               roleManagementCloudPCResourceNamespaceClient,
		RoleManagementCloudPCResourceNamespaceResourceAction: roleManagementCloudPCResourceNamespaceResourceActionClient,
		RoleManagementCloudPCResourceNamespaceResourceActionAuthenticationContext:               roleManagementCloudPCResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementCloudPCResourceNamespaceResourceActionResourceScope:                       roleManagementCloudPCResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementCloudPCRoleAssignment:                                                     roleManagementCloudPCRoleAssignmentClient,
		RoleManagementCloudPCRoleAssignmentAppScope:                                             roleManagementCloudPCRoleAssignmentAppScopeClient,
		RoleManagementCloudPCRoleAssignmentDirectoryScope:                                       roleManagementCloudPCRoleAssignmentDirectoryScopeClient,
		RoleManagementCloudPCRoleAssignmentPrincipal:                                            roleManagementCloudPCRoleAssignmentPrincipalClient,
		RoleManagementCloudPCRoleAssignmentRoleDefinition:                                       roleManagementCloudPCRoleAssignmentRoleDefinitionClient,
		RoleManagementCloudPCRoleDefinition:                                                     roleManagementCloudPCRoleDefinitionClient,
		RoleManagementCloudPCRoleDefinitionInheritsPermissionsFrom:                              roleManagementCloudPCRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementDeviceManagement:                                                          roleManagementDeviceManagementClient,
		RoleManagementDeviceManagementResourceNamespace:                                         roleManagementDeviceManagementResourceNamespaceClient,
		RoleManagementDeviceManagementResourceNamespaceResourceAction:                           roleManagementDeviceManagementResourceNamespaceResourceActionClient,
		RoleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContext:      roleManagementDeviceManagementResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementDeviceManagementResourceNamespaceResourceActionResourceScope:              roleManagementDeviceManagementResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementDeviceManagementRoleAssignment:                                            roleManagementDeviceManagementRoleAssignmentClient,
		RoleManagementDeviceManagementRoleAssignmentAppScope:                                    roleManagementDeviceManagementRoleAssignmentAppScopeClient,
		RoleManagementDeviceManagementRoleAssignmentDirectoryScope:                              roleManagementDeviceManagementRoleAssignmentDirectoryScopeClient,
		RoleManagementDeviceManagementRoleAssignmentPrincipal:                                   roleManagementDeviceManagementRoleAssignmentPrincipalClient,
		RoleManagementDeviceManagementRoleAssignmentRoleDefinition:                              roleManagementDeviceManagementRoleAssignmentRoleDefinitionClient,
		RoleManagementDeviceManagementRoleDefinition:                                            roleManagementDeviceManagementRoleDefinitionClient,
		RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFrom:                     roleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementDirectory:                                                                 roleManagementDirectoryClient,
		RoleManagementDirectoryResourceNamespace:                                                roleManagementDirectoryResourceNamespaceClient,
		RoleManagementDirectoryResourceNamespaceResourceAction:                                  roleManagementDirectoryResourceNamespaceResourceActionClient,
		RoleManagementDirectoryResourceNamespaceResourceActionAuthenticationContext:             roleManagementDirectoryResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementDirectoryResourceNamespaceResourceActionResourceScope:                     roleManagementDirectoryResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementDirectoryRoleAssignment:                                                   roleManagementDirectoryRoleAssignmentClient,
		RoleManagementDirectoryRoleAssignmentAppScope:                                           roleManagementDirectoryRoleAssignmentAppScopeClient,
		RoleManagementDirectoryRoleAssignmentApproval:                                           roleManagementDirectoryRoleAssignmentApprovalClient,
		RoleManagementDirectoryRoleAssignmentApprovalStep:                                       roleManagementDirectoryRoleAssignmentApprovalStepClient,
		RoleManagementDirectoryRoleAssignmentDirectoryScope:                                     roleManagementDirectoryRoleAssignmentDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentPrincipal:                                          roleManagementDirectoryRoleAssignmentPrincipalClient,
		RoleManagementDirectoryRoleAssignmentRoleDefinition:                                     roleManagementDirectoryRoleAssignmentRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentSchedule:                                           roleManagementDirectoryRoleAssignmentScheduleClient,
		RoleManagementDirectoryRoleAssignmentScheduleActivatedUsing:                             roleManagementDirectoryRoleAssignmentScheduleActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleAppScope:                                   roleManagementDirectoryRoleAssignmentScheduleAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleDirectoryScope:                             roleManagementDirectoryRoleAssignmentScheduleDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstance:                                   roleManagementDirectoryRoleAssignmentScheduleInstanceClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsing:                     roleManagementDirectoryRoleAssignmentScheduleInstanceActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceAppScope:                           roleManagementDirectoryRoleAssignmentScheduleInstanceAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScope:                     roleManagementDirectoryRoleAssignmentScheduleInstanceDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstancePrincipal:                          roleManagementDirectoryRoleAssignmentScheduleInstancePrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinition:                     roleManagementDirectoryRoleAssignmentScheduleInstanceRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentSchedulePrincipal:                                  roleManagementDirectoryRoleAssignmentSchedulePrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequest:                                    roleManagementDirectoryRoleAssignmentScheduleRequestClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsing:                      roleManagementDirectoryRoleAssignmentScheduleRequestActivatedUsingClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestAppScope:                            roleManagementDirectoryRoleAssignmentScheduleRequestAppScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScope:                      roleManagementDirectoryRoleAssignmentScheduleRequestDirectoryScopeClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestPrincipal:                           roleManagementDirectoryRoleAssignmentScheduleRequestPrincipalClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinition:                      roleManagementDirectoryRoleAssignmentScheduleRequestRoleDefinitionClient,
		RoleManagementDirectoryRoleAssignmentScheduleRequestTargetSchedule:                      roleManagementDirectoryRoleAssignmentScheduleRequestTargetScheduleClient,
		RoleManagementDirectoryRoleAssignmentScheduleRoleDefinition:                             roleManagementDirectoryRoleAssignmentScheduleRoleDefinitionClient,
		RoleManagementDirectoryRoleDefinition:                                                   roleManagementDirectoryRoleDefinitionClient,
		RoleManagementDirectoryRoleDefinitionInheritsPermissionsFrom:                            roleManagementDirectoryRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementDirectoryRoleEligibilitySchedule:                                          roleManagementDirectoryRoleEligibilityScheduleClient,
		RoleManagementDirectoryRoleEligibilityScheduleAppScope:                                  roleManagementDirectoryRoleEligibilityScheduleAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleDirectoryScope:                            roleManagementDirectoryRoleEligibilityScheduleDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstance:                                  roleManagementDirectoryRoleEligibilityScheduleInstanceClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceAppScope:                          roleManagementDirectoryRoleEligibilityScheduleInstanceAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScope:                    roleManagementDirectoryRoleEligibilityScheduleInstanceDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstancePrincipal:                         roleManagementDirectoryRoleEligibilityScheduleInstancePrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinition:                    roleManagementDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient,
		RoleManagementDirectoryRoleEligibilitySchedulePrincipal:                                 roleManagementDirectoryRoleEligibilitySchedulePrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequest:                                   roleManagementDirectoryRoleEligibilityScheduleRequestClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestAppScope:                           roleManagementDirectoryRoleEligibilityScheduleRequestAppScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScope:                     roleManagementDirectoryRoleEligibilityScheduleRequestDirectoryScopeClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestPrincipal:                          roleManagementDirectoryRoleEligibilityScheduleRequestPrincipalClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinition:                     roleManagementDirectoryRoleEligibilityScheduleRequestRoleDefinitionClient,
		RoleManagementDirectoryRoleEligibilityScheduleRequestTargetSchedule:                     roleManagementDirectoryRoleEligibilityScheduleRequestTargetScheduleClient,
		RoleManagementDirectoryRoleEligibilityScheduleRoleDefinition:                            roleManagementDirectoryRoleEligibilityScheduleRoleDefinitionClient,
		RoleManagementDirectoryTransitiveRoleAssignment:                                         roleManagementDirectoryTransitiveRoleAssignmentClient,
		RoleManagementDirectoryTransitiveRoleAssignmentAppScope:                                 roleManagementDirectoryTransitiveRoleAssignmentAppScopeClient,
		RoleManagementDirectoryTransitiveRoleAssignmentDirectoryScope:                           roleManagementDirectoryTransitiveRoleAssignmentDirectoryScopeClient,
		RoleManagementDirectoryTransitiveRoleAssignmentPrincipal:                                roleManagementDirectoryTransitiveRoleAssignmentPrincipalClient,
		RoleManagementDirectoryTransitiveRoleAssignmentRoleDefinition:                           roleManagementDirectoryTransitiveRoleAssignmentRoleDefinitionClient,
		RoleManagementEnterpriseApp:                                                             roleManagementEnterpriseAppClient,
		RoleManagementEnterpriseAppResourceNamespace:                                            roleManagementEnterpriseAppResourceNamespaceClient,
		RoleManagementEnterpriseAppResourceNamespaceResourceAction:                              roleManagementEnterpriseAppResourceNamespaceResourceActionClient,
		RoleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContext:         roleManagementEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementEnterpriseAppResourceNamespaceResourceActionResourceScope:                 roleManagementEnterpriseAppResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementEnterpriseAppRoleAssignment:                                               roleManagementEnterpriseAppRoleAssignmentClient,
		RoleManagementEnterpriseAppRoleAssignmentAppScope:                                       roleManagementEnterpriseAppRoleAssignmentAppScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentApproval:                                       roleManagementEnterpriseAppRoleAssignmentApprovalClient,
		RoleManagementEnterpriseAppRoleAssignmentApprovalStep:                                   roleManagementEnterpriseAppRoleAssignmentApprovalStepClient,
		RoleManagementEnterpriseAppRoleAssignmentDirectoryScope:                                 roleManagementEnterpriseAppRoleAssignmentDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentPrincipal:                                      roleManagementEnterpriseAppRoleAssignmentPrincipalClient,
		RoleManagementEnterpriseAppRoleAssignmentRoleDefinition:                                 roleManagementEnterpriseAppRoleAssignmentRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleAssignmentSchedule:                                       roleManagementEnterpriseAppRoleAssignmentScheduleClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsing:                         roleManagementEnterpriseAppRoleAssignmentScheduleActivatedUsingClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleAppScope:                               roleManagementEnterpriseAppRoleAssignmentScheduleAppScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScope:                         roleManagementEnterpriseAppRoleAssignmentScheduleDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstance:                               roleManagementEnterpriseAppRoleAssignmentScheduleInstanceClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsing:                 roleManagementEnterpriseAppRoleAssignmentScheduleInstanceActivatedUsingClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScope:                       roleManagementEnterpriseAppRoleAssignmentScheduleInstanceAppScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScope:                 roleManagementEnterpriseAppRoleAssignmentScheduleInstanceDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipal:                      roleManagementEnterpriseAppRoleAssignmentScheduleInstancePrincipalClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinition:                 roleManagementEnterpriseAppRoleAssignmentScheduleInstanceRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleAssignmentSchedulePrincipal:                              roleManagementEnterpriseAppRoleAssignmentSchedulePrincipalClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequest:                                roleManagementEnterpriseAppRoleAssignmentScheduleRequestClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsing:                  roleManagementEnterpriseAppRoleAssignmentScheduleRequestActivatedUsingClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScope:                        roleManagementEnterpriseAppRoleAssignmentScheduleRequestAppScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScope:                  roleManagementEnterpriseAppRoleAssignmentScheduleRequestDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipal:                       roleManagementEnterpriseAppRoleAssignmentScheduleRequestPrincipalClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinition:                  roleManagementEnterpriseAppRoleAssignmentScheduleRequestRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetSchedule:                  roleManagementEnterpriseAppRoleAssignmentScheduleRequestTargetScheduleClient,
		RoleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinition:                         roleManagementEnterpriseAppRoleAssignmentScheduleRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleDefinition:                                               roleManagementEnterpriseAppRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFrom:                        roleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementEnterpriseAppRoleEligibilitySchedule:                                      roleManagementEnterpriseAppRoleEligibilityScheduleClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleAppScope:                              roleManagementEnterpriseAppRoleEligibilityScheduleAppScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScope:                        roleManagementEnterpriseAppRoleEligibilityScheduleDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleInstance:                              roleManagementEnterpriseAppRoleEligibilityScheduleInstanceClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScope:                      roleManagementEnterpriseAppRoleEligibilityScheduleInstanceAppScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScope:                roleManagementEnterpriseAppRoleEligibilityScheduleInstanceDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipal:                     roleManagementEnterpriseAppRoleEligibilityScheduleInstancePrincipalClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinition:                roleManagementEnterpriseAppRoleEligibilityScheduleInstanceRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleEligibilitySchedulePrincipal:                             roleManagementEnterpriseAppRoleEligibilitySchedulePrincipalClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequest:                               roleManagementEnterpriseAppRoleEligibilityScheduleRequestClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScope:                       roleManagementEnterpriseAppRoleEligibilityScheduleRequestAppScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScope:                 roleManagementEnterpriseAppRoleEligibilityScheduleRequestDirectoryScopeClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipal:                      roleManagementEnterpriseAppRoleEligibilityScheduleRequestPrincipalClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinition:                 roleManagementEnterpriseAppRoleEligibilityScheduleRequestRoleDefinitionClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetSchedule:                 roleManagementEnterpriseAppRoleEligibilityScheduleRequestTargetScheduleClient,
		RoleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinition:                        roleManagementEnterpriseAppRoleEligibilityScheduleRoleDefinitionClient,
		RoleManagementEnterpriseAppTransitiveRoleAssignment:                                     roleManagementEnterpriseAppTransitiveRoleAssignmentClient,
		RoleManagementEnterpriseAppTransitiveRoleAssignmentAppScope:                             roleManagementEnterpriseAppTransitiveRoleAssignmentAppScopeClient,
		RoleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScope:                       roleManagementEnterpriseAppTransitiveRoleAssignmentDirectoryScopeClient,
		RoleManagementEnterpriseAppTransitiveRoleAssignmentPrincipal:                            roleManagementEnterpriseAppTransitiveRoleAssignmentPrincipalClient,
		RoleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinition:                       roleManagementEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient,
		RoleManagementEntitlementManagement:                                                     roleManagementEntitlementManagementClient,
		RoleManagementEntitlementManagementResourceNamespace:                                    roleManagementEntitlementManagementResourceNamespaceClient,
		RoleManagementEntitlementManagementResourceNamespaceResourceAction:                      roleManagementEntitlementManagementResourceNamespaceResourceActionClient,
		RoleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContext: roleManagementEntitlementManagementResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementEntitlementManagementResourceNamespaceResourceActionResourceScope:         roleManagementEntitlementManagementResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementEntitlementManagementRoleAssignment:                                       roleManagementEntitlementManagementRoleAssignmentClient,
		RoleManagementEntitlementManagementRoleAssignmentAppScope:                               roleManagementEntitlementManagementRoleAssignmentAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentApproval:                               roleManagementEntitlementManagementRoleAssignmentApprovalClient,
		RoleManagementEntitlementManagementRoleAssignmentApprovalStep:                           roleManagementEntitlementManagementRoleAssignmentApprovalStepClient,
		RoleManagementEntitlementManagementRoleAssignmentDirectoryScope:                         roleManagementEntitlementManagementRoleAssignmentDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentPrincipal:                              roleManagementEntitlementManagementRoleAssignmentPrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentRoleDefinition:                         roleManagementEntitlementManagementRoleAssignmentRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentSchedule:                               roleManagementEntitlementManagementRoleAssignmentScheduleClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsing:                 roleManagementEntitlementManagementRoleAssignmentScheduleActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleAppScope:                       roleManagementEntitlementManagementRoleAssignmentScheduleAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScope:                 roleManagementEntitlementManagementRoleAssignmentScheduleDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstance:                       roleManagementEntitlementManagementRoleAssignmentScheduleInstanceClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsing:         roleManagementEntitlementManagementRoleAssignmentScheduleInstanceActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScope:               roleManagementEntitlementManagementRoleAssignmentScheduleInstanceAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScope:         roleManagementEntitlementManagementRoleAssignmentScheduleInstanceDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipal:              roleManagementEntitlementManagementRoleAssignmentScheduleInstancePrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinition:         roleManagementEntitlementManagementRoleAssignmentScheduleInstanceRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentSchedulePrincipal:                      roleManagementEntitlementManagementRoleAssignmentSchedulePrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequest:                        roleManagementEntitlementManagementRoleAssignmentScheduleRequestClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsing:          roleManagementEntitlementManagementRoleAssignmentScheduleRequestActivatedUsingClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScope:                roleManagementEntitlementManagementRoleAssignmentScheduleRequestAppScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScope:          roleManagementEntitlementManagementRoleAssignmentScheduleRequestDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipal:               roleManagementEntitlementManagementRoleAssignmentScheduleRequestPrincipalClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinition:          roleManagementEntitlementManagementRoleAssignmentScheduleRequestRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetSchedule:          roleManagementEntitlementManagementRoleAssignmentScheduleRequestTargetScheduleClient,
		RoleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinition:                 roleManagementEntitlementManagementRoleAssignmentScheduleRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleDefinition:                                       roleManagementEntitlementManagementRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFrom:                roleManagementEntitlementManagementRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementEntitlementManagementRoleEligibilitySchedule:                              roleManagementEntitlementManagementRoleEligibilityScheduleClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleAppScope:                      roleManagementEntitlementManagementRoleEligibilityScheduleAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScope:                roleManagementEntitlementManagementRoleEligibilityScheduleDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstance:                      roleManagementEntitlementManagementRoleEligibilityScheduleInstanceClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScope:              roleManagementEntitlementManagementRoleEligibilityScheduleInstanceAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScope:        roleManagementEntitlementManagementRoleEligibilityScheduleInstanceDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipal:             roleManagementEntitlementManagementRoleEligibilityScheduleInstancePrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinition:        roleManagementEntitlementManagementRoleEligibilityScheduleInstanceRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleEligibilitySchedulePrincipal:                     roleManagementEntitlementManagementRoleEligibilitySchedulePrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequest:                       roleManagementEntitlementManagementRoleEligibilityScheduleRequestClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScope:               roleManagementEntitlementManagementRoleEligibilityScheduleRequestAppScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope:         roleManagementEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipal:              roleManagementEntitlementManagementRoleEligibilityScheduleRequestPrincipalClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinition:         roleManagementEntitlementManagementRoleEligibilityScheduleRequestRoleDefinitionClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetSchedule:         roleManagementEntitlementManagementRoleEligibilityScheduleRequestTargetScheduleClient,
		RoleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinition:                roleManagementEntitlementManagementRoleEligibilityScheduleRoleDefinitionClient,
		RoleManagementEntitlementManagementTransitiveRoleAssignment:                             roleManagementEntitlementManagementTransitiveRoleAssignmentClient,
		RoleManagementEntitlementManagementTransitiveRoleAssignmentAppScope:                     roleManagementEntitlementManagementTransitiveRoleAssignmentAppScopeClient,
		RoleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScope:               roleManagementEntitlementManagementTransitiveRoleAssignmentDirectoryScopeClient,
		RoleManagementEntitlementManagementTransitiveRoleAssignmentPrincipal:                    roleManagementEntitlementManagementTransitiveRoleAssignmentPrincipalClient,
		RoleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinition:               roleManagementEntitlementManagementTransitiveRoleAssignmentRoleDefinitionClient,
		RoleManagementExchange:                                                     roleManagementExchangeClient,
		RoleManagementExchangeCustomAppScope:                                       roleManagementExchangeCustomAppScopeClient,
		RoleManagementExchangeResourceNamespace:                                    roleManagementExchangeResourceNamespaceClient,
		RoleManagementExchangeResourceNamespaceResourceAction:                      roleManagementExchangeResourceNamespaceResourceActionClient,
		RoleManagementExchangeResourceNamespaceResourceActionAuthenticationContext: roleManagementExchangeResourceNamespaceResourceActionAuthenticationContextClient,
		RoleManagementExchangeResourceNamespaceResourceActionResourceScope:         roleManagementExchangeResourceNamespaceResourceActionResourceScopeClient,
		RoleManagementExchangeRoleAssignment:                                       roleManagementExchangeRoleAssignmentClient,
		RoleManagementExchangeRoleAssignmentAppScope:                               roleManagementExchangeRoleAssignmentAppScopeClient,
		RoleManagementExchangeRoleAssignmentDirectoryScope:                         roleManagementExchangeRoleAssignmentDirectoryScopeClient,
		RoleManagementExchangeRoleAssignmentPrincipal:                              roleManagementExchangeRoleAssignmentPrincipalClient,
		RoleManagementExchangeRoleAssignmentRoleDefinition:                         roleManagementExchangeRoleAssignmentRoleDefinitionClient,
		RoleManagementExchangeRoleDefinition:                                       roleManagementExchangeRoleDefinitionClient,
		RoleManagementExchangeRoleDefinitionInheritsPermissionsFrom:                roleManagementExchangeRoleDefinitionInheritsPermissionsFromClient,
		RoleManagementExchangeTransitiveRoleAssignment:                             roleManagementExchangeTransitiveRoleAssignmentClient,
		RoleManagementExchangeTransitiveRoleAssignmentAppScope:                     roleManagementExchangeTransitiveRoleAssignmentAppScopeClient,
		RoleManagementExchangeTransitiveRoleAssignmentDirectoryScope:               roleManagementExchangeTransitiveRoleAssignmentDirectoryScopeClient,
		RoleManagementExchangeTransitiveRoleAssignmentPrincipal:                    roleManagementExchangeTransitiveRoleAssignmentPrincipalClient,
		RoleManagementExchangeTransitiveRoleAssignmentRoleDefinition:               roleManagementExchangeTransitiveRoleAssignmentRoleDefinitionClient,
	}, nil
}
