package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreview"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancecontactedreviewer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancestage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancestagedecision"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewdefinitioninstancestagedecisioninsight"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewhistorydefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/accessreviewhistorydefinitioninstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/appconsent"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/appconsentappconsentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/appconsentappconsentrequestuserconsentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/appconsentappconsentrequestuserconsentrequestapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/appconsentappconsentrequestuserconsentrequestapprovalstage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageaccesspackagesincompatiblewith"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentapprovalstage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicyaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicycatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesettingcustomextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageassignmentpolicyquestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageincompatibleaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageincompatiblegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageincompatiblegroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementaccesspackageresourcerolescopescoperesourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicyaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicycatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicycustomextensionstagesetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentpolicyquestion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequestaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequestassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmentrequestrequestor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementassignmenttarget"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogcustomworkflowextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourceroleresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementcatalogresourcescoperesourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementconnectedorganization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementconnectedorganizationexternalsponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementconnectedorganizationinternalsponsor"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceroleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogcustomworkflowextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestcatalogresourcescoperesourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerequestresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceroleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourceroleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescoperoleresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourceroleresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcerolescopescoperesourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescoperesource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescoperesourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescoperesourcerole"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescoperesourceroleresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementresourcescoperesourceroleresourceenvironment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/entitlementmanagementsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/identitygovernance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextension"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensioncreatedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensioncreatedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensionlastmodifiedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensionlastmodifiedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowcustomtaskextensionlastmodifiedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditem"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowcreatedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowcreatedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowcreatedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowexecutionscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowlastmodifiedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowlastmodifiedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowlastmodifiedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrun"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowruntaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowtasktaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversioncreatedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversioncreatedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversionlastmodifiedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversionlastmodifiedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversionlastmodifiedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowtaskdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflow"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowcreatedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowcreatedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowcreatedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowexecutionscope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowlastmodifiedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowlastmodifiedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowlastmodifiedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrun"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowruntaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreport"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskdefinition"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtaskreporttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtasktaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplate"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowtemplatetasktaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowuserprocessingresulttaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversioncreatedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversioncreatedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversioncreatedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversionlastmodifiedby"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversionlastmodifiedbymailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresult"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresultsubject"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectmailboxsetting"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/lifecycleworkflowworkflowversiontasktaskprocessingresulttask"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccess"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentapproval"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentapprovalstage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulegroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstanceactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstancegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstancegroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequestactivatedusing"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequestgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequestgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulegroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleinstancegroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleinstancegroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleinstanceprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequestgroup"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequestgroupserviceprovisioningerror"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequestprincipal"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequesttargetschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuse"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreement"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreementacceptance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreementfile"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreementfilelocalization"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreementfilelocalizationversion"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/termsofuseagreementfileversion"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessReview                                                                                                   *accessreview.AccessReviewClient
	AccessReviewDefinition                                                                                         *accessreviewdefinition.AccessReviewDefinitionClient
	AccessReviewDefinitionInstance                                                                                 *accessreviewdefinitioninstance.AccessReviewDefinitionInstanceClient
	AccessReviewDefinitionInstanceContactedReviewer                                                                *accessreviewdefinitioninstancecontactedreviewer.AccessReviewDefinitionInstanceContactedReviewerClient
	AccessReviewDefinitionInstanceDecision                                                                         *accessreviewdefinitioninstancedecision.AccessReviewDefinitionInstanceDecisionClient
	AccessReviewDefinitionInstanceDecisionInsight                                                                  *accessreviewdefinitioninstancedecisioninsight.AccessReviewDefinitionInstanceDecisionInsightClient
	AccessReviewDefinitionInstanceStage                                                                            *accessreviewdefinitioninstancestage.AccessReviewDefinitionInstanceStageClient
	AccessReviewDefinitionInstanceStageDecision                                                                    *accessreviewdefinitioninstancestagedecision.AccessReviewDefinitionInstanceStageDecisionClient
	AccessReviewDefinitionInstanceStageDecisionInsight                                                             *accessreviewdefinitioninstancestagedecisioninsight.AccessReviewDefinitionInstanceStageDecisionInsightClient
	AccessReviewHistoryDefinition                                                                                  *accessreviewhistorydefinition.AccessReviewHistoryDefinitionClient
	AccessReviewHistoryDefinitionInstance                                                                          *accessreviewhistorydefinitioninstance.AccessReviewHistoryDefinitionInstanceClient
	AppConsent                                                                                                     *appconsent.AppConsentClient
	AppConsentAppConsentRequest                                                                                    *appconsentappconsentrequest.AppConsentAppConsentRequestClient
	AppConsentAppConsentRequestUserConsentRequest                                                                  *appconsentappconsentrequestuserconsentrequest.AppConsentAppConsentRequestUserConsentRequestClient
	AppConsentAppConsentRequestUserConsentRequestApproval                                                          *appconsentappconsentrequestuserconsentrequestapproval.AppConsentAppConsentRequestUserConsentRequestApprovalClient
	AppConsentAppConsentRequestUserConsentRequestApprovalStage                                                     *appconsentappconsentrequestuserconsentrequestapprovalstage.AppConsentAppConsentRequestUserConsentRequestApprovalStageClient
	EntitlementManagement                                                                                          *entitlementmanagement.EntitlementManagementClient
	EntitlementManagementAccessPackage                                                                             *entitlementmanagementaccesspackage.EntitlementManagementAccessPackageClient
	EntitlementManagementAccessPackageAccessPackagesIncompatibleWith                                               *entitlementmanagementaccesspackageaccesspackagesincompatiblewith.EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient
	EntitlementManagementAccessPackageAssignmentApproval                                                           *entitlementmanagementaccesspackageassignmentapproval.EntitlementManagementAccessPackageAssignmentApprovalClient
	EntitlementManagementAccessPackageAssignmentApprovalStage                                                      *entitlementmanagementaccesspackageassignmentapprovalstage.EntitlementManagementAccessPackageAssignmentApprovalStageClient
	EntitlementManagementAccessPackageAssignmentPolicy                                                             *entitlementmanagementaccesspackageassignmentpolicy.EntitlementManagementAccessPackageAssignmentPolicyClient
	EntitlementManagementAccessPackageAssignmentPolicyAccessPackage                                                *entitlementmanagementaccesspackageassignmentpolicyaccesspackage.EntitlementManagementAccessPackageAssignmentPolicyAccessPackageClient
	EntitlementManagementAccessPackageAssignmentPolicyCatalog                                                      *entitlementmanagementaccesspackageassignmentpolicycatalog.EntitlementManagementAccessPackageAssignmentPolicyCatalogClient
	EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSetting                                  *entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting.EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient
	EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtension                   *entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesettingcustomextension.EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient
	EntitlementManagementAccessPackageAssignmentPolicyQuestion                                                     *entitlementmanagementaccesspackageassignmentpolicyquestion.EntitlementManagementAccessPackageAssignmentPolicyQuestionClient
	EntitlementManagementAccessPackageCatalog                                                                      *entitlementmanagementaccesspackagecatalog.EntitlementManagementAccessPackageCatalogClient
	EntitlementManagementAccessPackageIncompatibleAccessPackage                                                    *entitlementmanagementaccesspackageincompatibleaccesspackage.EntitlementManagementAccessPackageIncompatibleAccessPackageClient
	EntitlementManagementAccessPackageIncompatibleGroup                                                            *entitlementmanagementaccesspackageincompatiblegroup.EntitlementManagementAccessPackageIncompatibleGroupClient
	EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningError                                    *entitlementmanagementaccesspackageincompatiblegroupserviceprovisioningerror.EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient
	EntitlementManagementAccessPackageResourceRoleScope                                                            *entitlementmanagementaccesspackageresourcerolescope.EntitlementManagementAccessPackageResourceRoleScopeClient
	EntitlementManagementAccessPackageResourceRoleScopeRole                                                        *entitlementmanagementaccesspackageresourcerolescoperole.EntitlementManagementAccessPackageResourceRoleScopeRoleClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResource                                                *entitlementmanagementaccesspackageresourcerolescoperoleresource.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironment                                     *entitlementmanagementaccesspackageresourcerolescoperoleresourceenvironment.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRole                                            *entitlementmanagementaccesspackageresourcerolescoperoleresourcerole.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScope                                           *entitlementmanagementaccesspackageresourcerolescoperoleresourcescope.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResource                                   *entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesource.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironment                        *entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourceenvironment.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRole                               *entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourcerole.EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRoleClient
	EntitlementManagementAccessPackageResourceRoleScopeScope                                                       *entitlementmanagementaccesspackageresourcerolescopescope.EntitlementManagementAccessPackageResourceRoleScopeScopeClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResource                                               *entitlementmanagementaccesspackageresourcerolescopescoperesource.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironment                                    *entitlementmanagementaccesspackageresourcerolescopescoperesourceenvironment.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRole                                           *entitlementmanagementaccesspackageresourcerolescopescoperesourcerole.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResource                                   *entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresource.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironment                        *entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourceenvironment.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScope                              *entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient
	EntitlementManagementAccessPackageResourceRoleScopeScopeResourceScope                                          *entitlementmanagementaccesspackageresourcerolescopescoperesourcescope.EntitlementManagementAccessPackageResourceRoleScopeScopeResourceScopeClient
	EntitlementManagementAssignment                                                                                *entitlementmanagementassignment.EntitlementManagementAssignmentClient
	EntitlementManagementAssignmentAccessPackage                                                                   *entitlementmanagementassignmentaccesspackage.EntitlementManagementAssignmentAccessPackageClient
	EntitlementManagementAssignmentAssignmentPolicy                                                                *entitlementmanagementassignmentassignmentpolicy.EntitlementManagementAssignmentAssignmentPolicyClient
	EntitlementManagementAssignmentPolicy                                                                          *entitlementmanagementassignmentpolicy.EntitlementManagementAssignmentPolicyClient
	EntitlementManagementAssignmentPolicyAccessPackage                                                             *entitlementmanagementassignmentpolicyaccesspackage.EntitlementManagementAssignmentPolicyAccessPackageClient
	EntitlementManagementAssignmentPolicyCatalog                                                                   *entitlementmanagementassignmentpolicycatalog.EntitlementManagementAssignmentPolicyCatalogClient
	EntitlementManagementAssignmentPolicyCustomExtensionStageSetting                                               *entitlementmanagementassignmentpolicycustomextensionstagesetting.EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient
	EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtension                                *entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension.EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient
	EntitlementManagementAssignmentPolicyQuestion                                                                  *entitlementmanagementassignmentpolicyquestion.EntitlementManagementAssignmentPolicyQuestionClient
	EntitlementManagementAssignmentRequest                                                                         *entitlementmanagementassignmentrequest.EntitlementManagementAssignmentRequestClient
	EntitlementManagementAssignmentRequestAccessPackage                                                            *entitlementmanagementassignmentrequestaccesspackage.EntitlementManagementAssignmentRequestAccessPackageClient
	EntitlementManagementAssignmentRequestAssignment                                                               *entitlementmanagementassignmentrequestassignment.EntitlementManagementAssignmentRequestAssignmentClient
	EntitlementManagementAssignmentRequestRequestor                                                                *entitlementmanagementassignmentrequestrequestor.EntitlementManagementAssignmentRequestRequestorClient
	EntitlementManagementAssignmentTarget                                                                          *entitlementmanagementassignmenttarget.EntitlementManagementAssignmentTargetClient
	EntitlementManagementCatalog                                                                                   *entitlementmanagementcatalog.EntitlementManagementCatalogClient
	EntitlementManagementCatalogAccessPackage                                                                      *entitlementmanagementcatalogaccesspackage.EntitlementManagementCatalogAccessPackageClient
	EntitlementManagementCatalogCustomWorkflowExtension                                                            *entitlementmanagementcatalogcustomworkflowextension.EntitlementManagementCatalogCustomWorkflowExtensionClient
	EntitlementManagementCatalogResource                                                                           *entitlementmanagementcatalogresource.EntitlementManagementCatalogResourceClient
	EntitlementManagementCatalogResourceEnvironment                                                                *entitlementmanagementcatalogresourceenvironment.EntitlementManagementCatalogResourceEnvironmentClient
	EntitlementManagementCatalogResourceRole                                                                       *entitlementmanagementcatalogresourcerole.EntitlementManagementCatalogResourceRoleClient
	EntitlementManagementCatalogResourceRoleResource                                                               *entitlementmanagementcatalogresourceroleresource.EntitlementManagementCatalogResourceRoleResourceClient
	EntitlementManagementCatalogResourceRoleResourceEnvironment                                                    *entitlementmanagementcatalogresourceroleresourceenvironment.EntitlementManagementCatalogResourceRoleResourceEnvironmentClient
	EntitlementManagementCatalogResourceRoleResourceRole                                                           *entitlementmanagementcatalogresourceroleresourcerole.EntitlementManagementCatalogResourceRoleResourceRoleClient
	EntitlementManagementCatalogResourceRoleResourceScope                                                          *entitlementmanagementcatalogresourceroleresourcescope.EntitlementManagementCatalogResourceRoleResourceScopeClient
	EntitlementManagementCatalogResourceRoleResourceScopeResource                                                  *entitlementmanagementcatalogresourceroleresourcescoperesource.EntitlementManagementCatalogResourceRoleResourceScopeResourceClient
	EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironment                                       *entitlementmanagementcatalogresourceroleresourcescoperesourceenvironment.EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementCatalogResourceRoleResourceScopeResourceRole                                              *entitlementmanagementcatalogresourceroleresourcescoperesourcerole.EntitlementManagementCatalogResourceRoleResourceScopeResourceRoleClient
	EntitlementManagementCatalogResourceScope                                                                      *entitlementmanagementcatalogresourcescope.EntitlementManagementCatalogResourceScopeClient
	EntitlementManagementCatalogResourceScopeResource                                                              *entitlementmanagementcatalogresourcescoperesource.EntitlementManagementCatalogResourceScopeResourceClient
	EntitlementManagementCatalogResourceScopeResourceEnvironment                                                   *entitlementmanagementcatalogresourcescoperesourceenvironment.EntitlementManagementCatalogResourceScopeResourceEnvironmentClient
	EntitlementManagementCatalogResourceScopeResourceRole                                                          *entitlementmanagementcatalogresourcescoperesourcerole.EntitlementManagementCatalogResourceScopeResourceRoleClient
	EntitlementManagementCatalogResourceScopeResourceRoleResource                                                  *entitlementmanagementcatalogresourcescoperesourceroleresource.EntitlementManagementCatalogResourceScopeResourceRoleResourceClient
	EntitlementManagementCatalogResourceScopeResourceRoleResourceEnvironment                                       *entitlementmanagementcatalogresourcescoperesourceroleresourceenvironment.EntitlementManagementCatalogResourceScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementCatalogResourceScopeResourceRoleResourceScope                                             *entitlementmanagementcatalogresourcescoperesourceroleresourcescope.EntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient
	EntitlementManagementCatalogResourceScopeResourceScope                                                         *entitlementmanagementcatalogresourcescoperesourcescope.EntitlementManagementCatalogResourceScopeResourceScopeClient
	EntitlementManagementConnectedOrganization                                                                     *entitlementmanagementconnectedorganization.EntitlementManagementConnectedOrganizationClient
	EntitlementManagementConnectedOrganizationExternalSponsor                                                      *entitlementmanagementconnectedorganizationexternalsponsor.EntitlementManagementConnectedOrganizationExternalSponsorClient
	EntitlementManagementConnectedOrganizationInternalSponsor                                                      *entitlementmanagementconnectedorganizationinternalsponsor.EntitlementManagementConnectedOrganizationInternalSponsorClient
	EntitlementManagementResource                                                                                  *entitlementmanagementresource.EntitlementManagementResourceClient
	EntitlementManagementResourceEnvironment                                                                       *entitlementmanagementresourceenvironment.EntitlementManagementResourceEnvironmentClient
	EntitlementManagementResourceEnvironmentResource                                                               *entitlementmanagementresourceenvironmentresource.EntitlementManagementResourceEnvironmentResourceClient
	EntitlementManagementResourceEnvironmentResourceEnvironment                                                    *entitlementmanagementresourceenvironmentresourceenvironment.EntitlementManagementResourceEnvironmentResourceEnvironmentClient
	EntitlementManagementResourceEnvironmentResourceRole                                                           *entitlementmanagementresourceenvironmentresourcerole.EntitlementManagementResourceEnvironmentResourceRoleClient
	EntitlementManagementResourceEnvironmentResourceRoleResource                                                   *entitlementmanagementresourceenvironmentresourceroleresource.EntitlementManagementResourceEnvironmentResourceRoleResourceClient
	EntitlementManagementResourceEnvironmentResourceRoleResourceEnvironment                                        *entitlementmanagementresourceenvironmentresourceroleresourceenvironment.EntitlementManagementResourceEnvironmentResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceEnvironmentResourceRoleResourceScope                                              *entitlementmanagementresourceenvironmentresourceroleresourcescope.EntitlementManagementResourceEnvironmentResourceRoleResourceScopeClient
	EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResource                                      *entitlementmanagementresourceenvironmentresourceroleresourcescoperesource.EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient
	EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironment                           *entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment.EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceEnvironmentResourceScope                                                          *entitlementmanagementresourceenvironmentresourcescope.EntitlementManagementResourceEnvironmentResourceScopeClient
	EntitlementManagementResourceEnvironmentResourceScopeResource                                                  *entitlementmanagementresourceenvironmentresourcescoperesource.EntitlementManagementResourceEnvironmentResourceScopeResourceClient
	EntitlementManagementResourceEnvironmentResourceScopeResourceEnvironment                                       *entitlementmanagementresourceenvironmentresourcescoperesourceenvironment.EntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceEnvironmentResourceScopeResourceRole                                              *entitlementmanagementresourceenvironmentresourcescoperesourcerole.EntitlementManagementResourceEnvironmentResourceScopeResourceRoleClient
	EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResource                                      *entitlementmanagementresourceenvironmentresourcescoperesourceroleresource.EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient
	EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironment                           *entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment.EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRequest                                                                           *entitlementmanagementresourcerequest.EntitlementManagementResourceRequestClient
	EntitlementManagementResourceRequestCatalog                                                                    *entitlementmanagementresourcerequestcatalog.EntitlementManagementResourceRequestCatalogClient
	EntitlementManagementResourceRequestCatalogAccessPackage                                                       *entitlementmanagementresourcerequestcatalogaccesspackage.EntitlementManagementResourceRequestCatalogAccessPackageClient
	EntitlementManagementResourceRequestCatalogCustomWorkflowExtension                                             *entitlementmanagementresourcerequestcatalogcustomworkflowextension.EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient
	EntitlementManagementResourceRequestCatalogResource                                                            *entitlementmanagementresourcerequestcatalogresource.EntitlementManagementResourceRequestCatalogResourceClient
	EntitlementManagementResourceRequestCatalogResourceEnvironment                                                 *entitlementmanagementresourcerequestcatalogresourceenvironment.EntitlementManagementResourceRequestCatalogResourceEnvironmentClient
	EntitlementManagementResourceRequestCatalogResourceRole                                                        *entitlementmanagementresourcerequestcatalogresourcerole.EntitlementManagementResourceRequestCatalogResourceRoleClient
	EntitlementManagementResourceRequestCatalogResourceRoleResource                                                *entitlementmanagementresourcerequestcatalogresourceroleresource.EntitlementManagementResourceRequestCatalogResourceRoleResourceClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceEnvironment                                     *entitlementmanagementresourcerequestcatalogresourceroleresourceenvironment.EntitlementManagementResourceRequestCatalogResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceRole                                            *entitlementmanagementresourcerequestcatalogresourceroleresourcerole.EntitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceScope                                           *entitlementmanagementresourcerequestcatalogresourceroleresourcescope.EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResource                                   *entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesource.EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironment                        *entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourceenvironment.EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRole                               *entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole.EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient
	EntitlementManagementResourceRequestCatalogResourceScope                                                       *entitlementmanagementresourcerequestcatalogresourcescope.EntitlementManagementResourceRequestCatalogResourceScopeClient
	EntitlementManagementResourceRequestCatalogResourceScopeResource                                               *entitlementmanagementresourcerequestcatalogresourcescoperesource.EntitlementManagementResourceRequestCatalogResourceScopeResourceClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceEnvironment                                    *entitlementmanagementresourcerequestcatalogresourcescoperesourceenvironment.EntitlementManagementResourceRequestCatalogResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceRole                                           *entitlementmanagementresourcerequestcatalogresourcescoperesourcerole.EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResource                                   *entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresource.EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironment                        *entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourceenvironment.EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScope                              *entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourcescope.EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient
	EntitlementManagementResourceRequestCatalogResourceScopeResourceScope                                          *entitlementmanagementresourcerequestcatalogresourcescoperesourcescope.EntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient
	EntitlementManagementResourceRequestResource                                                                   *entitlementmanagementresourcerequestresource.EntitlementManagementResourceRequestResourceClient
	EntitlementManagementResourceRequestResourceEnvironment                                                        *entitlementmanagementresourcerequestresourceenvironment.EntitlementManagementResourceRequestResourceEnvironmentClient
	EntitlementManagementResourceRequestResourceRole                                                               *entitlementmanagementresourcerequestresourcerole.EntitlementManagementResourceRequestResourceRoleClient
	EntitlementManagementResourceRequestResourceRoleResource                                                       *entitlementmanagementresourcerequestresourceroleresource.EntitlementManagementResourceRequestResourceRoleResourceClient
	EntitlementManagementResourceRequestResourceRoleResourceEnvironment                                            *entitlementmanagementresourcerequestresourceroleresourceenvironment.EntitlementManagementResourceRequestResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRequestResourceRoleResourceScope                                                  *entitlementmanagementresourcerequestresourceroleresourcescope.EntitlementManagementResourceRequestResourceRoleResourceScopeClient
	EntitlementManagementResourceRequestResourceScope                                                              *entitlementmanagementresourcerequestresourcescope.EntitlementManagementResourceRequestResourceScopeClient
	EntitlementManagementResourceRequestResourceScopeResource                                                      *entitlementmanagementresourcerequestresourcescoperesource.EntitlementManagementResourceRequestResourceScopeResourceClient
	EntitlementManagementResourceRequestResourceScopeResourceEnvironment                                           *entitlementmanagementresourcerequestresourcescoperesourceenvironment.EntitlementManagementResourceRequestResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceRequestResourceScopeResourceRole                                                  *entitlementmanagementresourcerequestresourcescoperesourcerole.EntitlementManagementResourceRequestResourceScopeResourceRoleClient
	EntitlementManagementResourceRole                                                                              *entitlementmanagementresourcerole.EntitlementManagementResourceRoleClient
	EntitlementManagementResourceRoleResource                                                                      *entitlementmanagementresourceroleresource.EntitlementManagementResourceRoleResourceClient
	EntitlementManagementResourceRoleResourceEnvironment                                                           *entitlementmanagementresourceroleresourceenvironment.EntitlementManagementResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRoleResourceScope                                                                 *entitlementmanagementresourceroleresourcescope.EntitlementManagementResourceRoleResourceScopeClient
	EntitlementManagementResourceRoleResourceScopeResource                                                         *entitlementmanagementresourceroleresourcescoperesource.EntitlementManagementResourceRoleResourceScopeResourceClient
	EntitlementManagementResourceRoleResourceScopeResourceEnvironment                                              *entitlementmanagementresourceroleresourcescoperesourceenvironment.EntitlementManagementResourceRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceRoleScope                                                                         *entitlementmanagementresourcerolescope.EntitlementManagementResourceRoleScopeClient
	EntitlementManagementResourceRoleScopeRole                                                                     *entitlementmanagementresourcerolescoperole.EntitlementManagementResourceRoleScopeRoleClient
	EntitlementManagementResourceRoleScopeRoleResource                                                             *entitlementmanagementresourcerolescoperoleresource.EntitlementManagementResourceRoleScopeRoleResourceClient
	EntitlementManagementResourceRoleScopeRoleResourceEnvironment                                                  *entitlementmanagementresourcerolescoperoleresourceenvironment.EntitlementManagementResourceRoleScopeRoleResourceEnvironmentClient
	EntitlementManagementResourceRoleScopeRoleResourceRole                                                         *entitlementmanagementresourcerolescoperoleresourcerole.EntitlementManagementResourceRoleScopeRoleResourceRoleClient
	EntitlementManagementResourceRoleScopeRoleResourceScope                                                        *entitlementmanagementresourcerolescoperoleresourcescope.EntitlementManagementResourceRoleScopeRoleResourceScopeClient
	EntitlementManagementResourceRoleScopeRoleResourceScopeResource                                                *entitlementmanagementresourcerolescoperoleresourcescoperesource.EntitlementManagementResourceRoleScopeRoleResourceScopeResourceClient
	EntitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironment                                     *entitlementmanagementresourcerolescoperoleresourcescoperesourceenvironment.EntitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRole                                            *entitlementmanagementresourcerolescoperoleresourcescoperesourcerole.EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient
	EntitlementManagementResourceRoleScopeScope                                                                    *entitlementmanagementresourcerolescopescope.EntitlementManagementResourceRoleScopeScopeClient
	EntitlementManagementResourceRoleScopeScopeResource                                                            *entitlementmanagementresourcerolescopescoperesource.EntitlementManagementResourceRoleScopeScopeResourceClient
	EntitlementManagementResourceRoleScopeScopeResourceEnvironment                                                 *entitlementmanagementresourcerolescopescoperesourceenvironment.EntitlementManagementResourceRoleScopeScopeResourceEnvironmentClient
	EntitlementManagementResourceRoleScopeScopeResourceRole                                                        *entitlementmanagementresourcerolescopescoperesourcerole.EntitlementManagementResourceRoleScopeScopeResourceRoleClient
	EntitlementManagementResourceRoleScopeScopeResourceRoleResource                                                *entitlementmanagementresourcerolescopescoperesourceroleresource.EntitlementManagementResourceRoleScopeScopeResourceRoleResourceClient
	EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironment                                     *entitlementmanagementresourcerolescopescoperesourceroleresourceenvironment.EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScope                                           *entitlementmanagementresourcerolescopescoperesourceroleresourcescope.EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient
	EntitlementManagementResourceRoleScopeScopeResourceScope                                                       *entitlementmanagementresourcerolescopescoperesourcescope.EntitlementManagementResourceRoleScopeScopeResourceScopeClient
	EntitlementManagementResourceScope                                                                             *entitlementmanagementresourcescope.EntitlementManagementResourceScopeClient
	EntitlementManagementResourceScopeResource                                                                     *entitlementmanagementresourcescoperesource.EntitlementManagementResourceScopeResourceClient
	EntitlementManagementResourceScopeResourceEnvironment                                                          *entitlementmanagementresourcescoperesourceenvironment.EntitlementManagementResourceScopeResourceEnvironmentClient
	EntitlementManagementResourceScopeResourceRole                                                                 *entitlementmanagementresourcescoperesourcerole.EntitlementManagementResourceScopeResourceRoleClient
	EntitlementManagementResourceScopeResourceRoleResource                                                         *entitlementmanagementresourcescoperesourceroleresource.EntitlementManagementResourceScopeResourceRoleResourceClient
	EntitlementManagementResourceScopeResourceRoleResourceEnvironment                                              *entitlementmanagementresourcescoperesourceroleresourceenvironment.EntitlementManagementResourceScopeResourceRoleResourceEnvironmentClient
	EntitlementManagementSetting                                                                                   *entitlementmanagementsetting.EntitlementManagementSettingClient
	IdentityGovernance                                                                                             *identitygovernance.IdentityGovernanceClient
	LifecycleWorkflow                                                                                              *lifecycleworkflow.LifecycleWorkflowClient
	LifecycleWorkflowCustomTaskExtension                                                                           *lifecycleworkflowcustomtaskextension.LifecycleWorkflowCustomTaskExtensionClient
	LifecycleWorkflowCustomTaskExtensionCreatedBy                                                                  *lifecycleworkflowcustomtaskextensioncreatedby.LifecycleWorkflowCustomTaskExtensionCreatedByClient
	LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSetting                                                    *lifecycleworkflowcustomtaskextensioncreatedbymailboxsetting.LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient
	LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningError                                          *lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror.LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient
	LifecycleWorkflowCustomTaskExtensionLastModifiedBy                                                             *lifecycleworkflowcustomtaskextensionlastmodifiedby.LifecycleWorkflowCustomTaskExtensionLastModifiedByClient
	LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSetting                                               *lifecycleworkflowcustomtaskextensionlastmodifiedbymailboxsetting.LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient
	LifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningError                                     *lifecycleworkflowcustomtaskextensionlastmodifiedbyserviceprovisioningerror.LifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItem                                                                                   *lifecycleworkflowdeleteditem.LifecycleWorkflowDeletedItemClient
	LifecycleWorkflowDeletedItemWorkflow                                                                           *lifecycleworkflowdeleteditemworkflow.LifecycleWorkflowDeletedItemWorkflowClient
	LifecycleWorkflowDeletedItemWorkflowCreatedBy                                                                  *lifecycleworkflowdeleteditemworkflowcreatedby.LifecycleWorkflowDeletedItemWorkflowCreatedByClient
	LifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSetting                                                    *lifecycleworkflowdeleteditemworkflowcreatedbymailboxsetting.LifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningError                                          *lifecycleworkflowdeleteditemworkflowcreatedbyserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowExecutionScope                                                             *lifecycleworkflowdeleteditemworkflowexecutionscope.LifecycleWorkflowDeletedItemWorkflowExecutionScopeClient
	LifecycleWorkflowDeletedItemWorkflowLastModifiedBy                                                             *lifecycleworkflowdeleteditemworkflowlastmodifiedby.LifecycleWorkflowDeletedItemWorkflowLastModifiedByClient
	LifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSetting                                               *lifecycleworkflowdeleteditemworkflowlastmodifiedbymailboxsetting.LifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningError                                     *lifecycleworkflowdeleteditemworkflowlastmodifiedbyserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowRun                                                                        *lifecycleworkflowdeleteditemworkflowrun.LifecycleWorkflowDeletedItemWorkflowRunClient
	LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResult                                                    *lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubject                                             *lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSetting                               *lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningError                     *lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTask                                                *lifecycleworkflowdeleteditemworkflowruntaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResult                                                    *lifecycleworkflowdeleteditemworkflowrunuserprocessingresult.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubject                                             *lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSetting                               *lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningError                     *lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult                                *lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubject                         *lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting           *lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError *lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTask                            *lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTaskClient
	LifecycleWorkflowDeletedItemWorkflowTask                                                                       *lifecycleworkflowdeleteditemworkflowtask.LifecycleWorkflowDeletedItemWorkflowTaskClient
	LifecycleWorkflowDeletedItemWorkflowTaskReport                                                                 *lifecycleworkflowdeleteditemworkflowtaskreport.LifecycleWorkflowDeletedItemWorkflowTaskReportClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTask                                                             *lifecycleworkflowdeleteditemworkflowtaskreporttask.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinition                                                   *lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult                                             *lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubject                                      *lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting                        *lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError              *lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTask                                         *lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTaskClient
	LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResult                                                   *lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubject                                            *lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSetting                              *lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError                    *lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTask                                               *lifecycleworkflowdeleteditemworkflowtasktaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResult                                                       *lifecycleworkflowdeleteditemworkflowuserprocessingresult.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubject                                                *lifecycleworkflowdeleteditemworkflowuserprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSetting                                  *lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningError                        *lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult                                   *lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubject                            *lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting              *lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError    *lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTask                               *lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTaskClient
	LifecycleWorkflowDeletedItemWorkflowVersion                                                                    *lifecycleworkflowdeleteditemworkflowversion.LifecycleWorkflowDeletedItemWorkflowVersionClient
	LifecycleWorkflowDeletedItemWorkflowVersionCreatedBy                                                           *lifecycleworkflowdeleteditemworkflowversioncreatedby.LifecycleWorkflowDeletedItemWorkflowVersionCreatedByClient
	LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSetting                                             *lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting.LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningError                                   *lifecycleworkflowdeleteditemworkflowversioncreatedbyserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedBy                                                      *lifecycleworkflowdeleteditemworkflowversionlastmodifiedby.LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient
	LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSetting                                        *lifecycleworkflowdeleteditemworkflowversionlastmodifiedbymailboxsetting.LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningError                              *lifecycleworkflowdeleteditemworkflowversionlastmodifiedbyserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowVersionTask                                                                *lifecycleworkflowdeleteditemworkflowversiontask.LifecycleWorkflowDeletedItemWorkflowVersionTaskClient
	LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResult                                            *lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient
	LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubject                                     *lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubject.LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectClient
	LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting                       *lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError             *lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTask                                        *lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask.LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient
	LifecycleWorkflowSetting                                                                                       *lifecycleworkflowsetting.LifecycleWorkflowSettingClient
	LifecycleWorkflowTaskDefinition                                                                                *lifecycleworkflowtaskdefinition.LifecycleWorkflowTaskDefinitionClient
	LifecycleWorkflowWorkflow                                                                                      *lifecycleworkflowworkflow.LifecycleWorkflowWorkflowClient
	LifecycleWorkflowWorkflowCreatedBy                                                                             *lifecycleworkflowworkflowcreatedby.LifecycleWorkflowWorkflowCreatedByClient
	LifecycleWorkflowWorkflowCreatedByMailboxSetting                                                               *lifecycleworkflowworkflowcreatedbymailboxsetting.LifecycleWorkflowWorkflowCreatedByMailboxSettingClient
	LifecycleWorkflowWorkflowCreatedByServiceProvisioningError                                                     *lifecycleworkflowworkflowcreatedbyserviceprovisioningerror.LifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowExecutionScope                                                                        *lifecycleworkflowworkflowexecutionscope.LifecycleWorkflowWorkflowExecutionScopeClient
	LifecycleWorkflowWorkflowLastModifiedBy                                                                        *lifecycleworkflowworkflowlastmodifiedby.LifecycleWorkflowWorkflowLastModifiedByClient
	LifecycleWorkflowWorkflowLastModifiedByMailboxSetting                                                          *lifecycleworkflowworkflowlastmodifiedbymailboxsetting.LifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient
	LifecycleWorkflowWorkflowLastModifiedByServiceProvisioningError                                                *lifecycleworkflowworkflowlastmodifiedbyserviceprovisioningerror.LifecycleWorkflowWorkflowLastModifiedByServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowRun                                                                                   *lifecycleworkflowworkflowrun.LifecycleWorkflowWorkflowRunClient
	LifecycleWorkflowWorkflowRunTaskProcessingResult                                                               *lifecycleworkflowworkflowruntaskprocessingresult.LifecycleWorkflowWorkflowRunTaskProcessingResultClient
	LifecycleWorkflowWorkflowRunTaskProcessingResultSubject                                                        *lifecycleworkflowworkflowruntaskprocessingresultsubject.LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSetting                                          *lifecycleworkflowworkflowruntaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningError                                *lifecycleworkflowworkflowruntaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowRunTaskProcessingResultTask                                                           *lifecycleworkflowworkflowruntaskprocessingresulttask.LifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowRunUserProcessingResult                                                               *lifecycleworkflowworkflowrunuserprocessingresult.LifecycleWorkflowWorkflowRunUserProcessingResultClient
	LifecycleWorkflowWorkflowRunUserProcessingResultSubject                                                        *lifecycleworkflowworkflowrunuserprocessingresultsubject.LifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient
	LifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSetting                                          *lifecycleworkflowworkflowrunuserprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningError                                *lifecycleworkflowworkflowrunuserprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResult                                           *lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient
	LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubject                                    *lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubject.LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting                      *lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError            *lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTask                                       *lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresulttask.LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowTask                                                                                  *lifecycleworkflowworkflowtask.LifecycleWorkflowWorkflowTaskClient
	LifecycleWorkflowWorkflowTaskReport                                                                            *lifecycleworkflowworkflowtaskreport.LifecycleWorkflowWorkflowTaskReportClient
	LifecycleWorkflowWorkflowTaskReportTask                                                                        *lifecycleworkflowworkflowtaskreporttask.LifecycleWorkflowWorkflowTaskReportTaskClient
	LifecycleWorkflowWorkflowTaskReportTaskDefinition                                                              *lifecycleworkflowworkflowtaskreporttaskdefinition.LifecycleWorkflowWorkflowTaskReportTaskDefinitionClient
	LifecycleWorkflowWorkflowTaskReportTaskProcessingResult                                                        *lifecycleworkflowworkflowtaskreporttaskprocessingresult.LifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient
	LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubject                                                 *lifecycleworkflowworkflowtaskreporttaskprocessingresultsubject.LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting                                   *lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError                         *lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowTaskReportTaskProcessingResultTask                                                    *lifecycleworkflowworkflowtaskreporttaskprocessingresulttask.LifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowTaskTaskProcessingResult                                                              *lifecycleworkflowworkflowtasktaskprocessingresult.LifecycleWorkflowWorkflowTaskTaskProcessingResultClient
	LifecycleWorkflowWorkflowTaskTaskProcessingResultSubject                                                       *lifecycleworkflowworkflowtasktaskprocessingresultsubject.LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSetting                                         *lifecycleworkflowworkflowtasktaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError                               *lifecycleworkflowworkflowtasktaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowTaskTaskProcessingResultTask                                                          *lifecycleworkflowworkflowtasktaskprocessingresulttask.LifecycleWorkflowWorkflowTaskTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowTemplate                                                                              *lifecycleworkflowworkflowtemplate.LifecycleWorkflowWorkflowTemplateClient
	LifecycleWorkflowWorkflowTemplateTask                                                                          *lifecycleworkflowworkflowtemplatetask.LifecycleWorkflowWorkflowTemplateTaskClient
	LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResult                                                      *lifecycleworkflowworkflowtemplatetasktaskprocessingresult.LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient
	LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubject                                               *lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubject.LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSetting                                 *lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningError                       *lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTask                                                  *lifecycleworkflowworkflowtemplatetasktaskprocessingresulttask.LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowUserProcessingResult                                                                  *lifecycleworkflowworkflowuserprocessingresult.LifecycleWorkflowWorkflowUserProcessingResultClient
	LifecycleWorkflowWorkflowUserProcessingResultSubject                                                           *lifecycleworkflowworkflowuserprocessingresultsubject.LifecycleWorkflowWorkflowUserProcessingResultSubjectClient
	LifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSetting                                             *lifecycleworkflowworkflowuserprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningError                                   *lifecycleworkflowworkflowuserprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResult                                              *lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient
	LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubject                                       *lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubject.LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting                         *lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError               *lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTask                                          *lifecycleworkflowworkflowuserprocessingresulttaskprocessingresulttask.LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient
	LifecycleWorkflowWorkflowVersion                                                                               *lifecycleworkflowworkflowversion.LifecycleWorkflowWorkflowVersionClient
	LifecycleWorkflowWorkflowVersionCreatedBy                                                                      *lifecycleworkflowworkflowversioncreatedby.LifecycleWorkflowWorkflowVersionCreatedByClient
	LifecycleWorkflowWorkflowVersionCreatedByMailboxSetting                                                        *lifecycleworkflowworkflowversioncreatedbymailboxsetting.LifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient
	LifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningError                                              *lifecycleworkflowworkflowversioncreatedbyserviceprovisioningerror.LifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowVersionLastModifiedBy                                                                 *lifecycleworkflowworkflowversionlastmodifiedby.LifecycleWorkflowWorkflowVersionLastModifiedByClient
	LifecycleWorkflowWorkflowVersionLastModifiedByMailboxSetting                                                   *lifecycleworkflowworkflowversionlastmodifiedbymailboxsetting.LifecycleWorkflowWorkflowVersionLastModifiedByMailboxSettingClient
	LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningError                                         *lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror.LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowVersionTask                                                                           *lifecycleworkflowworkflowversiontask.LifecycleWorkflowWorkflowVersionTaskClient
	LifecycleWorkflowWorkflowVersionTaskTaskProcessingResult                                                       *lifecycleworkflowworkflowversiontasktaskprocessingresult.LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient
	LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubject                                                *lifecycleworkflowworkflowversiontasktaskprocessingresultsubject.LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectClient
	LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting                                  *lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectmailboxsetting.LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient
	LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError                        *lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror.LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient
	LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTask                                                   *lifecycleworkflowworkflowversiontasktaskprocessingresulttask.LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTaskClient
	PrivilegedAccess                                                                                               *privilegedaccess.PrivilegedAccessClient
	PrivilegedAccessGroup                                                                                          *privilegedaccessgroup.PrivilegedAccessGroupClient
	PrivilegedAccessGroupAssignmentApproval                                                                        *privilegedaccessgroupassignmentapproval.PrivilegedAccessGroupAssignmentApprovalClient
	PrivilegedAccessGroupAssignmentApprovalStage                                                                   *privilegedaccessgroupassignmentapprovalstage.PrivilegedAccessGroupAssignmentApprovalStageClient
	PrivilegedAccessGroupAssignmentSchedule                                                                        *privilegedaccessgroupassignmentschedule.PrivilegedAccessGroupAssignmentScheduleClient
	PrivilegedAccessGroupAssignmentScheduleActivatedUsing                                                          *privilegedaccessgroupassignmentscheduleactivatedusing.PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient
	PrivilegedAccessGroupAssignmentScheduleGroup                                                                   *privilegedaccessgroupassignmentschedulegroup.PrivilegedAccessGroupAssignmentScheduleGroupClient
	PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningError                                           *privilegedaccessgroupassignmentschedulegroupserviceprovisioningerror.PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupAssignmentScheduleInstance                                                                *privilegedaccessgroupassignmentscheduleinstance.PrivilegedAccessGroupAssignmentScheduleInstanceClient
	PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsing                                                  *privilegedaccessgroupassignmentscheduleinstanceactivatedusing.PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient
	PrivilegedAccessGroupAssignmentScheduleInstanceGroup                                                           *privilegedaccessgroupassignmentscheduleinstancegroup.PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient
	PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningError                                   *privilegedaccessgroupassignmentscheduleinstancegroupserviceprovisioningerror.PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupAssignmentScheduleInstancePrincipal                                                       *privilegedaccessgroupassignmentscheduleinstanceprincipal.PrivilegedAccessGroupAssignmentScheduleInstancePrincipalClient
	PrivilegedAccessGroupAssignmentSchedulePrincipal                                                               *privilegedaccessgroupassignmentscheduleprincipal.PrivilegedAccessGroupAssignmentSchedulePrincipalClient
	PrivilegedAccessGroupAssignmentScheduleRequest                                                                 *privilegedaccessgroupassignmentschedulerequest.PrivilegedAccessGroupAssignmentScheduleRequestClient
	PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsing                                                   *privilegedaccessgroupassignmentschedulerequestactivatedusing.PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient
	PrivilegedAccessGroupAssignmentScheduleRequestGroup                                                            *privilegedaccessgroupassignmentschedulerequestgroup.PrivilegedAccessGroupAssignmentScheduleRequestGroupClient
	PrivilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningError                                    *privilegedaccessgroupassignmentschedulerequestgroupserviceprovisioningerror.PrivilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupAssignmentScheduleRequestPrincipal                                                        *privilegedaccessgroupassignmentschedulerequestprincipal.PrivilegedAccessGroupAssignmentScheduleRequestPrincipalClient
	PrivilegedAccessGroupAssignmentScheduleRequestTargetSchedule                                                   *privilegedaccessgroupassignmentschedulerequesttargetschedule.PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient
	PrivilegedAccessGroupEligibilitySchedule                                                                       *privilegedaccessgroupeligibilityschedule.PrivilegedAccessGroupEligibilityScheduleClient
	PrivilegedAccessGroupEligibilityScheduleGroup                                                                  *privilegedaccessgroupeligibilityschedulegroup.PrivilegedAccessGroupEligibilityScheduleGroupClient
	PrivilegedAccessGroupEligibilityScheduleGroupServiceProvisioningError                                          *privilegedaccessgroupeligibilityschedulegroupserviceprovisioningerror.PrivilegedAccessGroupEligibilityScheduleGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupEligibilityScheduleInstance                                                               *privilegedaccessgroupeligibilityscheduleinstance.PrivilegedAccessGroupEligibilityScheduleInstanceClient
	PrivilegedAccessGroupEligibilityScheduleInstanceGroup                                                          *privilegedaccessgroupeligibilityscheduleinstancegroup.PrivilegedAccessGroupEligibilityScheduleInstanceGroupClient
	PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningError                                  *privilegedaccessgroupeligibilityscheduleinstancegroupserviceprovisioningerror.PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupEligibilityScheduleInstancePrincipal                                                      *privilegedaccessgroupeligibilityscheduleinstanceprincipal.PrivilegedAccessGroupEligibilityScheduleInstancePrincipalClient
	PrivilegedAccessGroupEligibilitySchedulePrincipal                                                              *privilegedaccessgroupeligibilityscheduleprincipal.PrivilegedAccessGroupEligibilitySchedulePrincipalClient
	PrivilegedAccessGroupEligibilityScheduleRequest                                                                *privilegedaccessgroupeligibilityschedulerequest.PrivilegedAccessGroupEligibilityScheduleRequestClient
	PrivilegedAccessGroupEligibilityScheduleRequestGroup                                                           *privilegedaccessgroupeligibilityschedulerequestgroup.PrivilegedAccessGroupEligibilityScheduleRequestGroupClient
	PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningError                                   *privilegedaccessgroupeligibilityschedulerequestgroupserviceprovisioningerror.PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient
	PrivilegedAccessGroupEligibilityScheduleRequestPrincipal                                                       *privilegedaccessgroupeligibilityschedulerequestprincipal.PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient
	PrivilegedAccessGroupEligibilityScheduleRequestTargetSchedule                                                  *privilegedaccessgroupeligibilityschedulerequesttargetschedule.PrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient
	TermsOfUse                                                                                                     *termsofuse.TermsOfUseClient
	TermsOfUseAgreement                                                                                            *termsofuseagreement.TermsOfUseAgreementClient
	TermsOfUseAgreementAcceptance                                                                                  *termsofuseagreementacceptance.TermsOfUseAgreementAcceptanceClient
	TermsOfUseAgreementFile                                                                                        *termsofuseagreementfile.TermsOfUseAgreementFileClient
	TermsOfUseAgreementFileLocalization                                                                            *termsofuseagreementfilelocalization.TermsOfUseAgreementFileLocalizationClient
	TermsOfUseAgreementFileLocalizationVersion                                                                     *termsofuseagreementfilelocalizationversion.TermsOfUseAgreementFileLocalizationVersionClient
	TermsOfUseAgreementFileVersion                                                                                 *termsofuseagreementfileversion.TermsOfUseAgreementFileVersionClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	accessReviewClient, err := accessreview.NewAccessReviewClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReview client: %+v", err)
	}
	configureFunc(accessReviewClient.Client)

	accessReviewDefinitionClient, err := accessreviewdefinition.NewAccessReviewDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinition client: %+v", err)
	}
	configureFunc(accessReviewDefinitionClient.Client)

	accessReviewDefinitionInstanceClient, err := accessreviewdefinitioninstance.NewAccessReviewDefinitionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstance client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceClient.Client)

	accessReviewDefinitionInstanceContactedReviewerClient, err := accessreviewdefinitioninstancecontactedreviewer.NewAccessReviewDefinitionInstanceContactedReviewerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceContactedReviewer client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceContactedReviewerClient.Client)

	accessReviewDefinitionInstanceDecisionClient, err := accessreviewdefinitioninstancedecision.NewAccessReviewDefinitionInstanceDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceDecision client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceDecisionClient.Client)

	accessReviewDefinitionInstanceDecisionInsightClient, err := accessreviewdefinitioninstancedecisioninsight.NewAccessReviewDefinitionInstanceDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceDecisionInsight client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceDecisionInsightClient.Client)

	accessReviewDefinitionInstanceStageClient, err := accessreviewdefinitioninstancestage.NewAccessReviewDefinitionInstanceStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceStage client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceStageClient.Client)

	accessReviewDefinitionInstanceStageDecisionClient, err := accessreviewdefinitioninstancestagedecision.NewAccessReviewDefinitionInstanceStageDecisionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceStageDecision client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceStageDecisionClient.Client)

	accessReviewDefinitionInstanceStageDecisionInsightClient, err := accessreviewdefinitioninstancestagedecisioninsight.NewAccessReviewDefinitionInstanceStageDecisionInsightClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewDefinitionInstanceStageDecisionInsight client: %+v", err)
	}
	configureFunc(accessReviewDefinitionInstanceStageDecisionInsightClient.Client)

	accessReviewHistoryDefinitionClient, err := accessreviewhistorydefinition.NewAccessReviewHistoryDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewHistoryDefinition client: %+v", err)
	}
	configureFunc(accessReviewHistoryDefinitionClient.Client)

	accessReviewHistoryDefinitionInstanceClient, err := accessreviewhistorydefinitioninstance.NewAccessReviewHistoryDefinitionInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessReviewHistoryDefinitionInstance client: %+v", err)
	}
	configureFunc(accessReviewHistoryDefinitionInstanceClient.Client)

	appConsentAppConsentRequestClient, err := appconsentappconsentrequest.NewAppConsentAppConsentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentAppConsentRequest client: %+v", err)
	}
	configureFunc(appConsentAppConsentRequestClient.Client)

	appConsentAppConsentRequestUserConsentRequestApprovalClient, err := appconsentappconsentrequestuserconsentrequestapproval.NewAppConsentAppConsentRequestUserConsentRequestApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentAppConsentRequestUserConsentRequestApproval client: %+v", err)
	}
	configureFunc(appConsentAppConsentRequestUserConsentRequestApprovalClient.Client)

	appConsentAppConsentRequestUserConsentRequestApprovalStageClient, err := appconsentappconsentrequestuserconsentrequestapprovalstage.NewAppConsentAppConsentRequestUserConsentRequestApprovalStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentAppConsentRequestUserConsentRequestApprovalStage client: %+v", err)
	}
	configureFunc(appConsentAppConsentRequestUserConsentRequestApprovalStageClient.Client)

	appConsentAppConsentRequestUserConsentRequestClient, err := appconsentappconsentrequestuserconsentrequest.NewAppConsentAppConsentRequestUserConsentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsentAppConsentRequestUserConsentRequest client: %+v", err)
	}
	configureFunc(appConsentAppConsentRequestUserConsentRequestClient.Client)

	appConsentClient, err := appconsent.NewAppConsentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppConsent client: %+v", err)
	}
	configureFunc(appConsentClient.Client)

	entitlementManagementAccessPackageAccessPackagesIncompatibleWithClient, err := entitlementmanagementaccesspackageaccesspackagesincompatiblewith.NewEntitlementManagementAccessPackageAccessPackagesIncompatibleWithClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAccessPackagesIncompatibleWith client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAccessPackagesIncompatibleWithClient.Client)

	entitlementManagementAccessPackageAssignmentApprovalClient, err := entitlementmanagementaccesspackageassignmentapproval.NewEntitlementManagementAccessPackageAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentApproval client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentApprovalClient.Client)

	entitlementManagementAccessPackageAssignmentApprovalStageClient, err := entitlementmanagementaccesspackageassignmentapprovalstage.NewEntitlementManagementAccessPackageAssignmentApprovalStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentApprovalStage client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentApprovalStageClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyAccessPackageClient, err := entitlementmanagementaccesspackageassignmentpolicyaccesspackage.NewEntitlementManagementAccessPackageAssignmentPolicyAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicyAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyAccessPackageClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyCatalogClient, err := entitlementmanagementaccesspackageassignmentpolicycatalog.NewEntitlementManagementAccessPackageAssignmentPolicyCatalogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicyCatalog client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyCatalogClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyClient, err := entitlementmanagementaccesspackageassignmentpolicy.NewEntitlementManagementAccessPackageAssignmentPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicy client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient, err := entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesetting.NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSetting client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient, err := entitlementmanagementaccesspackageassignmentpolicycustomextensionstagesettingcustomextension.NewEntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtension client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient.Client)

	entitlementManagementAccessPackageAssignmentPolicyQuestionClient, err := entitlementmanagementaccesspackageassignmentpolicyquestion.NewEntitlementManagementAccessPackageAssignmentPolicyQuestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageAssignmentPolicyQuestion client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageAssignmentPolicyQuestionClient.Client)

	entitlementManagementAccessPackageCatalogClient, err := entitlementmanagementaccesspackagecatalog.NewEntitlementManagementAccessPackageCatalogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageCatalog client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageCatalogClient.Client)

	entitlementManagementAccessPackageClient, err := entitlementmanagementaccesspackage.NewEntitlementManagementAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageClient.Client)

	entitlementManagementAccessPackageIncompatibleAccessPackageClient, err := entitlementmanagementaccesspackageincompatibleaccesspackage.NewEntitlementManagementAccessPackageIncompatibleAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageIncompatibleAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageIncompatibleAccessPackageClient.Client)

	entitlementManagementAccessPackageIncompatibleGroupClient, err := entitlementmanagementaccesspackageincompatiblegroup.NewEntitlementManagementAccessPackageIncompatibleGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageIncompatibleGroup client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageIncompatibleGroupClient.Client)

	entitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient, err := entitlementmanagementaccesspackageincompatiblegroupserviceprovisioningerror.NewEntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeClient, err := entitlementmanagementaccesspackageresourcerolescope.NewEntitlementManagementAccessPackageResourceRoleScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScope client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleClient, err := entitlementmanagementaccesspackageresourcerolescoperole.NewEntitlementManagementAccessPackageResourceRoleScopeRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRole client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresource.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourceenvironment.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourcerole.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourcescope.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesource.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourceenvironment.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRoleClient, err := entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesourcerole.NewEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRoleClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeClient, err := entitlementmanagementaccesspackageresourcerolescopescope.NewEntitlementManagementAccessPackageResourceRoleScopeScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScope client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesource.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourceenvironment.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourcerole.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresource.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourceenvironment.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourceroleresourcescope.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient.Client)

	entitlementManagementAccessPackageResourceRoleScopeScopeResourceScopeClient, err := entitlementmanagementaccesspackageresourcerolescopescoperesourcescope.NewEntitlementManagementAccessPackageResourceRoleScopeScopeResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAccessPackageResourceRoleScopeScopeResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementAccessPackageResourceRoleScopeScopeResourceScopeClient.Client)

	entitlementManagementAssignmentAccessPackageClient, err := entitlementmanagementassignmentaccesspackage.NewEntitlementManagementAssignmentAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentAccessPackageClient.Client)

	entitlementManagementAssignmentAssignmentPolicyClient, err := entitlementmanagementassignmentassignmentpolicy.NewEntitlementManagementAssignmentAssignmentPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentAssignmentPolicy client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentAssignmentPolicyClient.Client)

	entitlementManagementAssignmentClient, err := entitlementmanagementassignment.NewEntitlementManagementAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignment client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentClient.Client)

	entitlementManagementAssignmentPolicyAccessPackageClient, err := entitlementmanagementassignmentpolicyaccesspackage.NewEntitlementManagementAssignmentPolicyAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicyAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyAccessPackageClient.Client)

	entitlementManagementAssignmentPolicyCatalogClient, err := entitlementmanagementassignmentpolicycatalog.NewEntitlementManagementAssignmentPolicyCatalogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicyCatalog client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyCatalogClient.Client)

	entitlementManagementAssignmentPolicyClient, err := entitlementmanagementassignmentpolicy.NewEntitlementManagementAssignmentPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicy client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyClient.Client)

	entitlementManagementAssignmentPolicyCustomExtensionStageSettingClient, err := entitlementmanagementassignmentpolicycustomextensionstagesetting.NewEntitlementManagementAssignmentPolicyCustomExtensionStageSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicyCustomExtensionStageSetting client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyCustomExtensionStageSettingClient.Client)

	entitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient, err := entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension.NewEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtension client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient.Client)

	entitlementManagementAssignmentPolicyQuestionClient, err := entitlementmanagementassignmentpolicyquestion.NewEntitlementManagementAssignmentPolicyQuestionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentPolicyQuestion client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentPolicyQuestionClient.Client)

	entitlementManagementAssignmentRequestAccessPackageClient, err := entitlementmanagementassignmentrequestaccesspackage.NewEntitlementManagementAssignmentRequestAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentRequestAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentRequestAccessPackageClient.Client)

	entitlementManagementAssignmentRequestAssignmentClient, err := entitlementmanagementassignmentrequestassignment.NewEntitlementManagementAssignmentRequestAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentRequestAssignment client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentRequestAssignmentClient.Client)

	entitlementManagementAssignmentRequestClient, err := entitlementmanagementassignmentrequest.NewEntitlementManagementAssignmentRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentRequest client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentRequestClient.Client)

	entitlementManagementAssignmentRequestRequestorClient, err := entitlementmanagementassignmentrequestrequestor.NewEntitlementManagementAssignmentRequestRequestorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentRequestRequestor client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentRequestRequestorClient.Client)

	entitlementManagementAssignmentTargetClient, err := entitlementmanagementassignmenttarget.NewEntitlementManagementAssignmentTargetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementAssignmentTarget client: %+v", err)
	}
	configureFunc(entitlementManagementAssignmentTargetClient.Client)

	entitlementManagementCatalogAccessPackageClient, err := entitlementmanagementcatalogaccesspackage.NewEntitlementManagementCatalogAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogAccessPackageClient.Client)

	entitlementManagementCatalogClient, err := entitlementmanagementcatalog.NewEntitlementManagementCatalogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalog client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogClient.Client)

	entitlementManagementCatalogCustomWorkflowExtensionClient, err := entitlementmanagementcatalogcustomworkflowextension.NewEntitlementManagementCatalogCustomWorkflowExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogCustomWorkflowExtension client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogCustomWorkflowExtensionClient.Client)

	entitlementManagementCatalogResourceClient, err := entitlementmanagementcatalogresource.NewEntitlementManagementCatalogResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResource client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceClient.Client)

	entitlementManagementCatalogResourceEnvironmentClient, err := entitlementmanagementcatalogresourceenvironment.NewEntitlementManagementCatalogResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceEnvironmentClient.Client)

	entitlementManagementCatalogResourceRoleClient, err := entitlementmanagementcatalogresourcerole.NewEntitlementManagementCatalogResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleClient.Client)

	entitlementManagementCatalogResourceRoleResourceClient, err := entitlementmanagementcatalogresourceroleresource.NewEntitlementManagementCatalogResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceClient.Client)

	entitlementManagementCatalogResourceRoleResourceEnvironmentClient, err := entitlementmanagementcatalogresourceroleresourceenvironment.NewEntitlementManagementCatalogResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementCatalogResourceRoleResourceRoleClient, err := entitlementmanagementcatalogresourceroleresourcerole.NewEntitlementManagementCatalogResourceRoleResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceRoleClient.Client)

	entitlementManagementCatalogResourceRoleResourceScopeClient, err := entitlementmanagementcatalogresourceroleresourcescope.NewEntitlementManagementCatalogResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceScopeClient.Client)

	entitlementManagementCatalogResourceRoleResourceScopeResourceClient, err := entitlementmanagementcatalogresourceroleresourcescoperesource.NewEntitlementManagementCatalogResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceScopeResourceClient.Client)

	entitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementcatalogresourceroleresourcescoperesourceenvironment.NewEntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementCatalogResourceRoleResourceScopeResourceRoleClient, err := entitlementmanagementcatalogresourceroleresourcescoperesourcerole.NewEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceRoleResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceRoleResourceScopeResourceRoleClient.Client)

	entitlementManagementCatalogResourceScopeClient, err := entitlementmanagementcatalogresourcescope.NewEntitlementManagementCatalogResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeClient.Client)

	entitlementManagementCatalogResourceScopeResourceClient, err := entitlementmanagementcatalogresourcescoperesource.NewEntitlementManagementCatalogResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceClient.Client)

	entitlementManagementCatalogResourceScopeResourceEnvironmentClient, err := entitlementmanagementcatalogresourcescoperesourceenvironment.NewEntitlementManagementCatalogResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementCatalogResourceScopeResourceRoleClient, err := entitlementmanagementcatalogresourcescoperesourcerole.NewEntitlementManagementCatalogResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceRoleClient.Client)

	entitlementManagementCatalogResourceScopeResourceRoleResourceClient, err := entitlementmanagementcatalogresourcescoperesourceroleresource.NewEntitlementManagementCatalogResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceRoleResourceClient.Client)

	entitlementManagementCatalogResourceScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementcatalogresourcescoperesourceroleresourceenvironment.NewEntitlementManagementCatalogResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient, err := entitlementmanagementcatalogresourcescoperesourceroleresourcescope.NewEntitlementManagementCatalogResourceScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient.Client)

	entitlementManagementCatalogResourceScopeResourceScopeClient, err := entitlementmanagementcatalogresourcescoperesourcescope.NewEntitlementManagementCatalogResourceScopeResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementCatalogResourceScopeResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementCatalogResourceScopeResourceScopeClient.Client)

	entitlementManagementClient, err := entitlementmanagement.NewEntitlementManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagement client: %+v", err)
	}
	configureFunc(entitlementManagementClient.Client)

	entitlementManagementConnectedOrganizationClient, err := entitlementmanagementconnectedorganization.NewEntitlementManagementConnectedOrganizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementConnectedOrganization client: %+v", err)
	}
	configureFunc(entitlementManagementConnectedOrganizationClient.Client)

	entitlementManagementConnectedOrganizationExternalSponsorClient, err := entitlementmanagementconnectedorganizationexternalsponsor.NewEntitlementManagementConnectedOrganizationExternalSponsorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementConnectedOrganizationExternalSponsor client: %+v", err)
	}
	configureFunc(entitlementManagementConnectedOrganizationExternalSponsorClient.Client)

	entitlementManagementConnectedOrganizationInternalSponsorClient, err := entitlementmanagementconnectedorganizationinternalsponsor.NewEntitlementManagementConnectedOrganizationInternalSponsorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementConnectedOrganizationInternalSponsor client: %+v", err)
	}
	configureFunc(entitlementManagementConnectedOrganizationInternalSponsorClient.Client)

	entitlementManagementResourceClient, err := entitlementmanagementresource.NewEntitlementManagementResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceClient.Client)

	entitlementManagementResourceEnvironmentClient, err := entitlementmanagementresourceenvironment.NewEntitlementManagementResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentClient.Client)

	entitlementManagementResourceEnvironmentResourceClient, err := entitlementmanagementresourceenvironmentresource.NewEntitlementManagementResourceEnvironmentResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceClient.Client)

	entitlementManagementResourceEnvironmentResourceEnvironmentClient, err := entitlementmanagementresourceenvironmentresourceenvironment.NewEntitlementManagementResourceEnvironmentResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceEnvironmentClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleClient, err := entitlementmanagementresourceenvironmentresourcerole.NewEntitlementManagementResourceEnvironmentResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleResourceClient, err := entitlementmanagementresourceenvironmentresourceroleresource.NewEntitlementManagementResourceEnvironmentResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleResourceClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourceenvironmentresourceroleresourceenvironment.NewEntitlementManagementResourceEnvironmentResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleResourceScopeClient, err := entitlementmanagementresourceenvironmentresourceroleresourcescope.NewEntitlementManagementResourceEnvironmentResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient, err := entitlementmanagementresourceenvironmentresourceroleresourcescoperesource.NewEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient.Client)

	entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourceenvironmentresourceroleresourcescoperesourceenvironment.NewEntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeClient, err := entitlementmanagementresourceenvironmentresourcescope.NewEntitlementManagementResourceEnvironmentResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeResourceClient, err := entitlementmanagementresourceenvironmentresourcescoperesource.NewEntitlementManagementResourceEnvironmentResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeResourceClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourceenvironmentresourcescoperesourceenvironment.NewEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeResourceRoleClient, err := entitlementmanagementresourceenvironmentresourcescoperesourcerole.NewEntitlementManagementResourceEnvironmentResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient, err := entitlementmanagementresourceenvironmentresourcescoperesourceroleresource.NewEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient.Client)

	entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourceenvironmentresourcescoperesourceroleresourceenvironment.NewEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogAccessPackageClient, err := entitlementmanagementresourcerequestcatalogaccesspackage.NewEntitlementManagementResourceRequestCatalogAccessPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogAccessPackage client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogAccessPackageClient.Client)

	entitlementManagementResourceRequestCatalogClient, err := entitlementmanagementresourcerequestcatalog.NewEntitlementManagementResourceRequestCatalogClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalog client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogClient.Client)

	entitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient, err := entitlementmanagementresourcerequestcatalogcustomworkflowextension.NewEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogCustomWorkflowExtension client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient.Client)

	entitlementManagementResourceRequestCatalogResourceClient, err := entitlementmanagementresourcerequestcatalogresource.NewEntitlementManagementResourceRequestCatalogResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceClient.Client)

	entitlementManagementResourceRequestCatalogResourceEnvironmentClient, err := entitlementmanagementresourcerequestcatalogresourceenvironment.NewEntitlementManagementResourceRequestCatalogResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleClient, err := entitlementmanagementresourcerequestcatalogresourcerole.NewEntitlementManagementResourceRequestCatalogResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceClient, err := entitlementmanagementresourcerequestcatalogresourceroleresource.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourceenvironment.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourcerole.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourcescope.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesource.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourceenvironment.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient, err := entitlementmanagementresourcerequestcatalogresourceroleresourcescoperesourcerole.NewEntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeClient, err := entitlementmanagementresourcerequestcatalogresourcescope.NewEntitlementManagementResourceRequestCatalogResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesource.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourceenvironment.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourcerole.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresource.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourceenvironment.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourceroleresourcescope.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient, err := entitlementmanagementresourcerequestcatalogresourcescoperesourcescope.NewEntitlementManagementResourceRequestCatalogResourceScopeResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestCatalogResourceScopeResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient.Client)

	entitlementManagementResourceRequestClient, err := entitlementmanagementresourcerequest.NewEntitlementManagementResourceRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequest client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestClient.Client)

	entitlementManagementResourceRequestResourceClient, err := entitlementmanagementresourcerequestresource.NewEntitlementManagementResourceRequestResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceClient.Client)

	entitlementManagementResourceRequestResourceEnvironmentClient, err := entitlementmanagementresourcerequestresourceenvironment.NewEntitlementManagementResourceRequestResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestResourceRoleClient, err := entitlementmanagementresourcerequestresourcerole.NewEntitlementManagementResourceRequestResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceRoleClient.Client)

	entitlementManagementResourceRequestResourceRoleResourceClient, err := entitlementmanagementresourcerequestresourceroleresource.NewEntitlementManagementResourceRequestResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceRoleResourceClient.Client)

	entitlementManagementResourceRequestResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourcerequestresourceroleresourceenvironment.NewEntitlementManagementResourceRequestResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestResourceRoleResourceScopeClient, err := entitlementmanagementresourcerequestresourceroleresourcescope.NewEntitlementManagementResourceRequestResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceRequestResourceScopeClient, err := entitlementmanagementresourcerequestresourcescope.NewEntitlementManagementResourceRequestResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceScopeClient.Client)

	entitlementManagementResourceRequestResourceScopeResourceClient, err := entitlementmanagementresourcerequestresourcescoperesource.NewEntitlementManagementResourceRequestResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceScopeResourceClient.Client)

	entitlementManagementResourceRequestResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourcerequestresourcescoperesourceenvironment.NewEntitlementManagementResourceRequestResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRequestResourceScopeResourceRoleClient, err := entitlementmanagementresourcerequestresourcescoperesourcerole.NewEntitlementManagementResourceRequestResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRequestResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRequestResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceRoleClient, err := entitlementmanagementresourcerole.NewEntitlementManagementResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleClient.Client)

	entitlementManagementResourceRoleResourceClient, err := entitlementmanagementresourceroleresource.NewEntitlementManagementResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleResourceClient.Client)

	entitlementManagementResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourceroleresourceenvironment.NewEntitlementManagementResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleResourceScopeClient, err := entitlementmanagementresourceroleresourcescope.NewEntitlementManagementResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceRoleResourceScopeResourceClient, err := entitlementmanagementresourceroleresourcescoperesource.NewEntitlementManagementResourceRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleResourceScopeResourceClient.Client)

	entitlementManagementResourceRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourceroleresourcescoperesourceenvironment.NewEntitlementManagementResourceRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleScopeClient, err := entitlementmanagementresourcerolescope.NewEntitlementManagementResourceRoleScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeClient.Client)

	entitlementManagementResourceRoleScopeRoleClient, err := entitlementmanagementresourcerolescoperole.NewEntitlementManagementResourceRoleScopeRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceClient, err := entitlementmanagementresourcerolescoperoleresource.NewEntitlementManagementResourceRoleScopeRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceEnvironmentClient, err := entitlementmanagementresourcerolescoperoleresourceenvironment.NewEntitlementManagementResourceRoleScopeRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceRoleClient, err := entitlementmanagementresourcerolescoperoleresourcerole.NewEntitlementManagementResourceRoleScopeRoleResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceRoleClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceScopeClient, err := entitlementmanagementresourcerolescoperoleresourcescope.NewEntitlementManagementResourceRoleScopeRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceScopeClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceScopeResourceClient, err := entitlementmanagementresourcerolescoperoleresourcescoperesource.NewEntitlementManagementResourceRoleScopeRoleResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceScopeResourceClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourcerolescoperoleresourcescoperesourceenvironment.NewEntitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient, err := entitlementmanagementresourcerolescoperoleresourcescoperesourcerole.NewEntitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceRoleScopeScopeClient, err := entitlementmanagementresourcerolescopescope.NewEntitlementManagementResourceRoleScopeScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceClient, err := entitlementmanagementresourcerolescopescoperesource.NewEntitlementManagementResourceRoleScopeScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceEnvironmentClient, err := entitlementmanagementresourcerolescopescoperesourceenvironment.NewEntitlementManagementResourceRoleScopeScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceRoleClient, err := entitlementmanagementresourcerolescopescoperesourcerole.NewEntitlementManagementResourceRoleScopeScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceRoleClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceRoleResourceClient, err := entitlementmanagementresourcerolescopescoperesourceroleresource.NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceRoleResourceClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourcerolescopescoperesourceroleresourceenvironment.NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient, err := entitlementmanagementresourcerolescopescoperesourceroleresourcescope.NewEntitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient.Client)

	entitlementManagementResourceRoleScopeScopeResourceScopeClient, err := entitlementmanagementresourcerolescopescoperesourcescope.NewEntitlementManagementResourceRoleScopeScopeResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceRoleScopeScopeResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceRoleScopeScopeResourceScopeClient.Client)

	entitlementManagementResourceScopeClient, err := entitlementmanagementresourcescope.NewEntitlementManagementResourceScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScope client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeClient.Client)

	entitlementManagementResourceScopeResourceClient, err := entitlementmanagementresourcescoperesource.NewEntitlementManagementResourceScopeResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScopeResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeResourceClient.Client)

	entitlementManagementResourceScopeResourceEnvironmentClient, err := entitlementmanagementresourcescoperesourceenvironment.NewEntitlementManagementResourceScopeResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScopeResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeResourceEnvironmentClient.Client)

	entitlementManagementResourceScopeResourceRoleClient, err := entitlementmanagementresourcescoperesourcerole.NewEntitlementManagementResourceScopeResourceRoleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScopeResourceRole client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeResourceRoleClient.Client)

	entitlementManagementResourceScopeResourceRoleResourceClient, err := entitlementmanagementresourcescoperesourceroleresource.NewEntitlementManagementResourceScopeResourceRoleResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScopeResourceRoleResource client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeResourceRoleResourceClient.Client)

	entitlementManagementResourceScopeResourceRoleResourceEnvironmentClient, err := entitlementmanagementresourcescoperesourceroleresourceenvironment.NewEntitlementManagementResourceScopeResourceRoleResourceEnvironmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementResourceScopeResourceRoleResourceEnvironment client: %+v", err)
	}
	configureFunc(entitlementManagementResourceScopeResourceRoleResourceEnvironmentClient.Client)

	entitlementManagementSettingClient, err := entitlementmanagementsetting.NewEntitlementManagementSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EntitlementManagementSetting client: %+v", err)
	}
	configureFunc(entitlementManagementSettingClient.Client)

	identityGovernanceClient, err := identitygovernance.NewIdentityGovernanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IdentityGovernance client: %+v", err)
	}
	configureFunc(identityGovernanceClient.Client)

	lifecycleWorkflowClient, err := lifecycleworkflow.NewLifecycleWorkflowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflow client: %+v", err)
	}
	configureFunc(lifecycleWorkflowClient.Client)

	lifecycleWorkflowCustomTaskExtensionClient, err := lifecycleworkflowcustomtaskextension.NewLifecycleWorkflowCustomTaskExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtension client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionClient.Client)

	lifecycleWorkflowCustomTaskExtensionCreatedByClient, err := lifecycleworkflowcustomtaskextensioncreatedby.NewLifecycleWorkflowCustomTaskExtensionCreatedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionCreatedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionCreatedByClient.Client)

	lifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient, err := lifecycleworkflowcustomtaskextensioncreatedbymailboxsetting.NewLifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient.Client)

	lifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient, err := lifecycleworkflowcustomtaskextensioncreatedbyserviceprovisioningerror.NewLifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowCustomTaskExtensionLastModifiedByClient, err := lifecycleworkflowcustomtaskextensionlastmodifiedby.NewLifecycleWorkflowCustomTaskExtensionLastModifiedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionLastModifiedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionLastModifiedByClient.Client)

	lifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient, err := lifecycleworkflowcustomtaskextensionlastmodifiedbymailboxsetting.NewLifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient.Client)

	lifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningErrorClient, err := lifecycleworkflowcustomtaskextensionlastmodifiedbyserviceprovisioningerror.NewLifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemClient, err := lifecycleworkflowdeleteditem.NewLifecycleWorkflowDeletedItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItem client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemClient.Client)

	lifecycleWorkflowDeletedItemWorkflowClient, err := lifecycleworkflowdeleteditemworkflow.NewLifecycleWorkflowDeletedItemWorkflowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflow client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowClient.Client)

	lifecycleWorkflowDeletedItemWorkflowCreatedByClient, err := lifecycleworkflowdeleteditemworkflowcreatedby.NewLifecycleWorkflowDeletedItemWorkflowCreatedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowCreatedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowCreatedByClient.Client)

	lifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowcreatedbymailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowcreatedbyserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowExecutionScopeClient, err := lifecycleworkflowdeleteditemworkflowexecutionscope.NewLifecycleWorkflowDeletedItemWorkflowExecutionScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowExecutionScope client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowExecutionScopeClient.Client)

	lifecycleWorkflowDeletedItemWorkflowLastModifiedByClient, err := lifecycleworkflowdeleteditemworkflowlastmodifiedby.NewLifecycleWorkflowDeletedItemWorkflowLastModifiedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowLastModifiedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowLastModifiedByClient.Client)

	lifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowlastmodifiedbymailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowlastmodifiedbyserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunClient, err := lifecycleworkflowdeleteditemworkflowrun.NewLifecycleWorkflowDeletedItemWorkflowRunClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRun client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowruntaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowruntaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowruntaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowrunuserprocessingresulttaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskClient, err := lifecycleworkflowdeleteditemworkflowtask.NewLifecycleWorkflowDeletedItemWorkflowTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportClient, err := lifecycleworkflowdeleteditemworkflowtaskreport.NewLifecycleWorkflowDeletedItemWorkflowTaskReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReport client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttask.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinition client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowtasktaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionClient, err := lifecycleworkflowdeleteditemworkflowversion.NewLifecycleWorkflowDeletedItemWorkflowVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersion client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionCreatedByClient, err := lifecycleworkflowdeleteditemworkflowversioncreatedby.NewLifecycleWorkflowDeletedItemWorkflowVersionCreatedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionCreatedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionCreatedByClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowversioncreatedbymailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowversioncreatedbyserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient, err := lifecycleworkflowdeleteditemworkflowversionlastmodifiedby.NewLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowversionlastmodifiedbymailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowversionlastmodifiedbyserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskClient, err := lifecycleworkflowdeleteditemworkflowversiontask.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient, err := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectClient, err := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubject.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient, err := lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask.NewLifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowSettingClient, err := lifecycleworkflowsetting.NewLifecycleWorkflowSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowSettingClient.Client)

	lifecycleWorkflowTaskDefinitionClient, err := lifecycleworkflowtaskdefinition.NewLifecycleWorkflowTaskDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowTaskDefinition client: %+v", err)
	}
	configureFunc(lifecycleWorkflowTaskDefinitionClient.Client)

	lifecycleWorkflowWorkflowClient, err := lifecycleworkflowworkflow.NewLifecycleWorkflowWorkflowClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflow client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowClient.Client)

	lifecycleWorkflowWorkflowCreatedByClient, err := lifecycleworkflowworkflowcreatedby.NewLifecycleWorkflowWorkflowCreatedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowCreatedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowCreatedByClient.Client)

	lifecycleWorkflowWorkflowCreatedByMailboxSettingClient, err := lifecycleworkflowworkflowcreatedbymailboxsetting.NewLifecycleWorkflowWorkflowCreatedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowCreatedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowCreatedByMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient, err := lifecycleworkflowworkflowcreatedbyserviceprovisioningerror.NewLifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowCreatedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowExecutionScopeClient, err := lifecycleworkflowworkflowexecutionscope.NewLifecycleWorkflowWorkflowExecutionScopeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowExecutionScope client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowExecutionScopeClient.Client)

	lifecycleWorkflowWorkflowLastModifiedByClient, err := lifecycleworkflowworkflowlastmodifiedby.NewLifecycleWorkflowWorkflowLastModifiedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowLastModifiedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowLastModifiedByClient.Client)

	lifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient, err := lifecycleworkflowworkflowlastmodifiedbymailboxsetting.NewLifecycleWorkflowWorkflowLastModifiedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowLastModifiedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowLastModifiedByServiceProvisioningErrorClient, err := lifecycleworkflowworkflowlastmodifiedbyserviceprovisioningerror.NewLifecycleWorkflowWorkflowLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowLastModifiedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowLastModifiedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowRunClient, err := lifecycleworkflowworkflowrun.NewLifecycleWorkflowWorkflowRunClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRun client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunClient.Client)

	lifecycleWorkflowWorkflowRunTaskProcessingResultClient, err := lifecycleworkflowworkflowruntaskprocessingresult.NewLifecycleWorkflowWorkflowRunTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowruntaskprocessingresultsubject.NewLifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowruntaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowruntaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowruntaskprocessingresulttask.NewLifecycleWorkflowWorkflowRunTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultClient, err := lifecycleworkflowworkflowrunuserprocessingresult.NewLifecycleWorkflowWorkflowRunUserProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient, err := lifecycleworkflowworkflowrunuserprocessingresultsubject.NewLifecycleWorkflowWorkflowRunUserProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowrunuserprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowrunuserprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient, err := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubject.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowrunuserprocessingresulttaskprocessingresulttask.NewLifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowTaskClient, err := lifecycleworkflowworkflowtask.NewLifecycleWorkflowWorkflowTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskClient.Client)

	lifecycleWorkflowWorkflowTaskReportClient, err := lifecycleworkflowworkflowtaskreport.NewLifecycleWorkflowWorkflowTaskReportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReport client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskClient, err := lifecycleworkflowworkflowtaskreporttask.NewLifecycleWorkflowWorkflowTaskReportTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskDefinitionClient, err := lifecycleworkflowworkflowtaskreporttaskdefinition.NewLifecycleWorkflowWorkflowTaskReportTaskDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskDefinition client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskDefinitionClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient, err := lifecycleworkflowworkflowtaskreporttaskprocessingresult.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowtaskreporttaskprocessingresultsubject.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowtaskreporttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowtaskreporttaskprocessingresulttask.NewLifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskReportTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowTaskTaskProcessingResultClient, err := lifecycleworkflowworkflowtasktaskprocessingresult.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowtasktaskprocessingresultsubject.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowtasktaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowtasktaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowTaskTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowtasktaskprocessingresulttask.NewLifecycleWorkflowWorkflowTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTaskTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTaskTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowTemplateClient, err := lifecycleworkflowworkflowtemplate.NewLifecycleWorkflowWorkflowTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplate client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskClient, err := lifecycleworkflowworkflowtemplatetask.NewLifecycleWorkflowWorkflowTemplateTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient, err := lifecycleworkflowworkflowtemplatetasktaskprocessingresult.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubject.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowtemplatetasktaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowtemplatetasktaskprocessingresulttask.NewLifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultClient, err := lifecycleworkflowworkflowuserprocessingresult.NewLifecycleWorkflowWorkflowUserProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultSubjectClient, err := lifecycleworkflowworkflowuserprocessingresultsubject.NewLifecycleWorkflowWorkflowUserProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowuserprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowuserprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient, err := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubject.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowuserprocessingresulttaskprocessingresulttask.NewLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient.Client)

	lifecycleWorkflowWorkflowVersionClient, err := lifecycleworkflowworkflowversion.NewLifecycleWorkflowWorkflowVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersion client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionClient.Client)

	lifecycleWorkflowWorkflowVersionCreatedByClient, err := lifecycleworkflowworkflowversioncreatedby.NewLifecycleWorkflowWorkflowVersionCreatedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionCreatedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionCreatedByClient.Client)

	lifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient, err := lifecycleworkflowworkflowversioncreatedbymailboxsetting.NewLifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionCreatedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningErrorClient, err := lifecycleworkflowworkflowversioncreatedbyserviceprovisioningerror.NewLifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowVersionLastModifiedByClient, err := lifecycleworkflowworkflowversionlastmodifiedby.NewLifecycleWorkflowWorkflowVersionLastModifiedByClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionLastModifiedBy client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionLastModifiedByClient.Client)

	lifecycleWorkflowWorkflowVersionLastModifiedByMailboxSettingClient, err := lifecycleworkflowworkflowversionlastmodifiedbymailboxsetting.NewLifecycleWorkflowWorkflowVersionLastModifiedByMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionLastModifiedByMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionLastModifiedByMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient, err := lifecycleworkflowworkflowversionlastmodifiedbyserviceprovisioningerror.NewLifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowVersionTaskClient, err := lifecycleworkflowworkflowversiontask.NewLifecycleWorkflowWorkflowVersionTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskClient.Client)

	lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient, err := lifecycleworkflowworkflowversiontasktaskprocessingresult.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTaskTaskProcessingResult client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient.Client)

	lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectClient, err := lifecycleworkflowworkflowversiontasktaskprocessingresultsubject.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubject client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectClient.Client)

	lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient, err := lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectmailboxsetting.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient.Client)

	lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient, err := lifecycleworkflowworkflowversiontasktaskprocessingresultsubjectserviceprovisioningerror.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient.Client)

	lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTaskClient, err := lifecycleworkflowworkflowversiontasktaskprocessingresulttask.NewLifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTaskClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTask client: %+v", err)
	}
	configureFunc(lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTaskClient.Client)

	privilegedAccessClient, err := privilegedaccess.NewPrivilegedAccessClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccess client: %+v", err)
	}
	configureFunc(privilegedAccessClient.Client)

	privilegedAccessGroupAssignmentApprovalClient, err := privilegedaccessgroupassignmentapproval.NewPrivilegedAccessGroupAssignmentApprovalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentApproval client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentApprovalClient.Client)

	privilegedAccessGroupAssignmentApprovalStageClient, err := privilegedaccessgroupassignmentapprovalstage.NewPrivilegedAccessGroupAssignmentApprovalStageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentApprovalStage client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentApprovalStageClient.Client)

	privilegedAccessGroupAssignmentScheduleActivatedUsingClient, err := privilegedaccessgroupassignmentscheduleactivatedusing.NewPrivilegedAccessGroupAssignmentScheduleActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleActivatedUsing client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleActivatedUsingClient.Client)

	privilegedAccessGroupAssignmentScheduleClient, err := privilegedaccessgroupassignmentschedule.NewPrivilegedAccessGroupAssignmentScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentSchedule client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleClient.Client)

	privilegedAccessGroupAssignmentScheduleGroupClient, err := privilegedaccessgroupassignmentschedulegroup.NewPrivilegedAccessGroupAssignmentScheduleGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleGroupClient.Client)

	privilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient, err := privilegedaccessgroupassignmentschedulegroupserviceprovisioningerror.NewPrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient, err := privilegedaccessgroupassignmentscheduleinstanceactivatedusing.NewPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsing client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient.Client)

	privilegedAccessGroupAssignmentScheduleInstanceClient, err := privilegedaccessgroupassignmentscheduleinstance.NewPrivilegedAccessGroupAssignmentScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleInstance client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleInstanceClient.Client)

	privilegedAccessGroupAssignmentScheduleInstanceGroupClient, err := privilegedaccessgroupassignmentscheduleinstancegroup.NewPrivilegedAccessGroupAssignmentScheduleInstanceGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleInstanceGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleInstanceGroupClient.Client)

	privilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient, err := privilegedaccessgroupassignmentscheduleinstancegroupserviceprovisioningerror.NewPrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupAssignmentScheduleInstancePrincipalClient, err := privilegedaccessgroupassignmentscheduleinstanceprincipal.NewPrivilegedAccessGroupAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleInstancePrincipalClient.Client)

	privilegedAccessGroupAssignmentSchedulePrincipalClient, err := privilegedaccessgroupassignmentscheduleprincipal.NewPrivilegedAccessGroupAssignmentSchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentSchedulePrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentSchedulePrincipalClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient, err := privilegedaccessgroupassignmentschedulerequestactivatedusing.NewPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsing client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestClient, err := privilegedaccessgroupassignmentschedulerequest.NewPrivilegedAccessGroupAssignmentScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequest client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestGroupClient, err := privilegedaccessgroupassignmentschedulerequestgroup.NewPrivilegedAccessGroupAssignmentScheduleRequestGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequestGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestGroupClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningErrorClient, err := privilegedaccessgroupassignmentschedulerequestgroupserviceprovisioningerror.NewPrivilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestPrincipalClient, err := privilegedaccessgroupassignmentschedulerequestprincipal.NewPrivilegedAccessGroupAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestPrincipalClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient, err := privilegedaccessgroupassignmentschedulerequesttargetschedule.NewPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupAssignmentScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(privilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient.Client)

	privilegedAccessGroupClient, err := privilegedaccessgroup.NewPrivilegedAccessGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupClient.Client)

	privilegedAccessGroupEligibilityScheduleClient, err := privilegedaccessgroupeligibilityschedule.NewPrivilegedAccessGroupEligibilityScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilitySchedule client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleClient.Client)

	privilegedAccessGroupEligibilityScheduleGroupClient, err := privilegedaccessgroupeligibilityschedulegroup.NewPrivilegedAccessGroupEligibilityScheduleGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleGroupClient.Client)

	privilegedAccessGroupEligibilityScheduleGroupServiceProvisioningErrorClient, err := privilegedaccessgroupeligibilityschedulegroupserviceprovisioningerror.NewPrivilegedAccessGroupEligibilityScheduleGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupEligibilityScheduleInstanceClient, err := privilegedaccessgroupeligibilityscheduleinstance.NewPrivilegedAccessGroupEligibilityScheduleInstanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleInstance client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleInstanceClient.Client)

	privilegedAccessGroupEligibilityScheduleInstanceGroupClient, err := privilegedaccessgroupeligibilityscheduleinstancegroup.NewPrivilegedAccessGroupEligibilityScheduleInstanceGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleInstanceGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleInstanceGroupClient.Client)

	privilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient, err := privilegedaccessgroupeligibilityscheduleinstancegroupserviceprovisioningerror.NewPrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupEligibilityScheduleInstancePrincipalClient, err := privilegedaccessgroupeligibilityscheduleinstanceprincipal.NewPrivilegedAccessGroupEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleInstancePrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleInstancePrincipalClient.Client)

	privilegedAccessGroupEligibilitySchedulePrincipalClient, err := privilegedaccessgroupeligibilityscheduleprincipal.NewPrivilegedAccessGroupEligibilitySchedulePrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilitySchedulePrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilitySchedulePrincipalClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestClient, err := privilegedaccessgroupeligibilityschedulerequest.NewPrivilegedAccessGroupEligibilityScheduleRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleRequest client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleRequestClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestGroupClient, err := privilegedaccessgroupeligibilityschedulerequestgroup.NewPrivilegedAccessGroupEligibilityScheduleRequestGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleRequestGroup client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleRequestGroupClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient, err := privilegedaccessgroupeligibilityschedulerequestgroupserviceprovisioningerror.NewPrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningError client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestPrincipalClient, err := privilegedaccessgroupeligibilityschedulerequestprincipal.NewPrivilegedAccessGroupEligibilityScheduleRequestPrincipalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleRequestPrincipal client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleRequestPrincipalClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient, err := privilegedaccessgroupeligibilityschedulerequesttargetschedule.NewPrivilegedAccessGroupEligibilityScheduleRequestTargetScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivilegedAccessGroupEligibilityScheduleRequestTargetSchedule client: %+v", err)
	}
	configureFunc(privilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient.Client)

	termsOfUseAgreementAcceptanceClient, err := termsofuseagreementacceptance.NewTermsOfUseAgreementAcceptanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreementAcceptance client: %+v", err)
	}
	configureFunc(termsOfUseAgreementAcceptanceClient.Client)

	termsOfUseAgreementClient, err := termsofuseagreement.NewTermsOfUseAgreementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreement client: %+v", err)
	}
	configureFunc(termsOfUseAgreementClient.Client)

	termsOfUseAgreementFileClient, err := termsofuseagreementfile.NewTermsOfUseAgreementFileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreementFile client: %+v", err)
	}
	configureFunc(termsOfUseAgreementFileClient.Client)

	termsOfUseAgreementFileLocalizationClient, err := termsofuseagreementfilelocalization.NewTermsOfUseAgreementFileLocalizationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreementFileLocalization client: %+v", err)
	}
	configureFunc(termsOfUseAgreementFileLocalizationClient.Client)

	termsOfUseAgreementFileLocalizationVersionClient, err := termsofuseagreementfilelocalizationversion.NewTermsOfUseAgreementFileLocalizationVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreementFileLocalizationVersion client: %+v", err)
	}
	configureFunc(termsOfUseAgreementFileLocalizationVersionClient.Client)

	termsOfUseAgreementFileVersionClient, err := termsofuseagreementfileversion.NewTermsOfUseAgreementFileVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUseAgreementFileVersion client: %+v", err)
	}
	configureFunc(termsOfUseAgreementFileVersionClient.Client)

	termsOfUseClient, err := termsofuse.NewTermsOfUseClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TermsOfUse client: %+v", err)
	}
	configureFunc(termsOfUseClient.Client)

	return &Client{
		AccessReview:                                                                                                   accessReviewClient,
		AccessReviewDefinition:                                                                                         accessReviewDefinitionClient,
		AccessReviewDefinitionInstance:                                                                                 accessReviewDefinitionInstanceClient,
		AccessReviewDefinitionInstanceContactedReviewer:                                                                accessReviewDefinitionInstanceContactedReviewerClient,
		AccessReviewDefinitionInstanceDecision:                                                                         accessReviewDefinitionInstanceDecisionClient,
		AccessReviewDefinitionInstanceDecisionInsight:                                                                  accessReviewDefinitionInstanceDecisionInsightClient,
		AccessReviewDefinitionInstanceStage:                                                                            accessReviewDefinitionInstanceStageClient,
		AccessReviewDefinitionInstanceStageDecision:                                                                    accessReviewDefinitionInstanceStageDecisionClient,
		AccessReviewDefinitionInstanceStageDecisionInsight:                                                             accessReviewDefinitionInstanceStageDecisionInsightClient,
		AccessReviewHistoryDefinition:                                                                                  accessReviewHistoryDefinitionClient,
		AccessReviewHistoryDefinitionInstance:                                                                          accessReviewHistoryDefinitionInstanceClient,
		AppConsent:                                                                                                     appConsentClient,
		AppConsentAppConsentRequest:                                                                                    appConsentAppConsentRequestClient,
		AppConsentAppConsentRequestUserConsentRequest:                                                                  appConsentAppConsentRequestUserConsentRequestClient,
		AppConsentAppConsentRequestUserConsentRequestApproval:                                                          appConsentAppConsentRequestUserConsentRequestApprovalClient,
		AppConsentAppConsentRequestUserConsentRequestApprovalStage:                                                     appConsentAppConsentRequestUserConsentRequestApprovalStageClient,
		EntitlementManagement:                                                                                          entitlementManagementClient,
		EntitlementManagementAccessPackage:                                                                             entitlementManagementAccessPackageClient,
		EntitlementManagementAccessPackageAccessPackagesIncompatibleWith:                                               entitlementManagementAccessPackageAccessPackagesIncompatibleWithClient,
		EntitlementManagementAccessPackageAssignmentApproval:                                                           entitlementManagementAccessPackageAssignmentApprovalClient,
		EntitlementManagementAccessPackageAssignmentApprovalStage:                                                      entitlementManagementAccessPackageAssignmentApprovalStageClient,
		EntitlementManagementAccessPackageAssignmentPolicy:                                                             entitlementManagementAccessPackageAssignmentPolicyClient,
		EntitlementManagementAccessPackageAssignmentPolicyAccessPackage:                                                entitlementManagementAccessPackageAssignmentPolicyAccessPackageClient,
		EntitlementManagementAccessPackageAssignmentPolicyCatalog:                                                      entitlementManagementAccessPackageAssignmentPolicyCatalogClient,
		EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSetting:                                  entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingClient,
		EntitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtension:                   entitlementManagementAccessPackageAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient,
		EntitlementManagementAccessPackageAssignmentPolicyQuestion:                                                     entitlementManagementAccessPackageAssignmentPolicyQuestionClient,
		EntitlementManagementAccessPackageCatalog:                                                                      entitlementManagementAccessPackageCatalogClient,
		EntitlementManagementAccessPackageIncompatibleAccessPackage:                                                    entitlementManagementAccessPackageIncompatibleAccessPackageClient,
		EntitlementManagementAccessPackageIncompatibleGroup:                                                            entitlementManagementAccessPackageIncompatibleGroupClient,
		EntitlementManagementAccessPackageIncompatibleGroupServiceProvisioningError:                                    entitlementManagementAccessPackageIncompatibleGroupServiceProvisioningErrorClient,
		EntitlementManagementAccessPackageResourceRoleScope:                                                            entitlementManagementAccessPackageResourceRoleScopeClient,
		EntitlementManagementAccessPackageResourceRoleScopeRole:                                                        entitlementManagementAccessPackageResourceRoleScopeRoleClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResource:                                                entitlementManagementAccessPackageResourceRoleScopeRoleResourceClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironment:                                     entitlementManagementAccessPackageResourceRoleScopeRoleResourceEnvironmentClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceRole:                                            entitlementManagementAccessPackageResourceRoleScopeRoleResourceRoleClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScope:                                           entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResource:                                   entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironment:                        entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRole:                               entitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRoleClient,
		EntitlementManagementAccessPackageResourceRoleScopeScope:                                                       entitlementManagementAccessPackageResourceRoleScopeScopeClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResource:                                               entitlementManagementAccessPackageResourceRoleScopeScopeResourceClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironment:                                    entitlementManagementAccessPackageResourceRoleScopeScopeResourceEnvironmentClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRole:                                           entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResource:                                   entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironment:                        entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScope:                              entitlementManagementAccessPackageResourceRoleScopeScopeResourceRoleResourceScopeClient,
		EntitlementManagementAccessPackageResourceRoleScopeScopeResourceScope:                                          entitlementManagementAccessPackageResourceRoleScopeScopeResourceScopeClient,
		EntitlementManagementAssignment:                                                                                entitlementManagementAssignmentClient,
		EntitlementManagementAssignmentAccessPackage:                                                                   entitlementManagementAssignmentAccessPackageClient,
		EntitlementManagementAssignmentAssignmentPolicy:                                                                entitlementManagementAssignmentAssignmentPolicyClient,
		EntitlementManagementAssignmentPolicy:                                                                          entitlementManagementAssignmentPolicyClient,
		EntitlementManagementAssignmentPolicyAccessPackage:                                                             entitlementManagementAssignmentPolicyAccessPackageClient,
		EntitlementManagementAssignmentPolicyCatalog:                                                                   entitlementManagementAssignmentPolicyCatalogClient,
		EntitlementManagementAssignmentPolicyCustomExtensionStageSetting:                                               entitlementManagementAssignmentPolicyCustomExtensionStageSettingClient,
		EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtension:                                entitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient,
		EntitlementManagementAssignmentPolicyQuestion:                                                                  entitlementManagementAssignmentPolicyQuestionClient,
		EntitlementManagementAssignmentRequest:                                                                         entitlementManagementAssignmentRequestClient,
		EntitlementManagementAssignmentRequestAccessPackage:                                                            entitlementManagementAssignmentRequestAccessPackageClient,
		EntitlementManagementAssignmentRequestAssignment:                                                               entitlementManagementAssignmentRequestAssignmentClient,
		EntitlementManagementAssignmentRequestRequestor:                                                                entitlementManagementAssignmentRequestRequestorClient,
		EntitlementManagementAssignmentTarget:                                                                          entitlementManagementAssignmentTargetClient,
		EntitlementManagementCatalog:                                                                                   entitlementManagementCatalogClient,
		EntitlementManagementCatalogAccessPackage:                                                                      entitlementManagementCatalogAccessPackageClient,
		EntitlementManagementCatalogCustomWorkflowExtension:                                                            entitlementManagementCatalogCustomWorkflowExtensionClient,
		EntitlementManagementCatalogResource:                                                                           entitlementManagementCatalogResourceClient,
		EntitlementManagementCatalogResourceEnvironment:                                                                entitlementManagementCatalogResourceEnvironmentClient,
		EntitlementManagementCatalogResourceRole:                                                                       entitlementManagementCatalogResourceRoleClient,
		EntitlementManagementCatalogResourceRoleResource:                                                               entitlementManagementCatalogResourceRoleResourceClient,
		EntitlementManagementCatalogResourceRoleResourceEnvironment:                                                    entitlementManagementCatalogResourceRoleResourceEnvironmentClient,
		EntitlementManagementCatalogResourceRoleResourceRole:                                                           entitlementManagementCatalogResourceRoleResourceRoleClient,
		EntitlementManagementCatalogResourceRoleResourceScope:                                                          entitlementManagementCatalogResourceRoleResourceScopeClient,
		EntitlementManagementCatalogResourceRoleResourceScopeResource:                                                  entitlementManagementCatalogResourceRoleResourceScopeResourceClient,
		EntitlementManagementCatalogResourceRoleResourceScopeResourceEnvironment:                                       entitlementManagementCatalogResourceRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementCatalogResourceRoleResourceScopeResourceRole:                                              entitlementManagementCatalogResourceRoleResourceScopeResourceRoleClient,
		EntitlementManagementCatalogResourceScope:                                                                      entitlementManagementCatalogResourceScopeClient,
		EntitlementManagementCatalogResourceScopeResource:                                                              entitlementManagementCatalogResourceScopeResourceClient,
		EntitlementManagementCatalogResourceScopeResourceEnvironment:                                                   entitlementManagementCatalogResourceScopeResourceEnvironmentClient,
		EntitlementManagementCatalogResourceScopeResourceRole:                                                          entitlementManagementCatalogResourceScopeResourceRoleClient,
		EntitlementManagementCatalogResourceScopeResourceRoleResource:                                                  entitlementManagementCatalogResourceScopeResourceRoleResourceClient,
		EntitlementManagementCatalogResourceScopeResourceRoleResourceEnvironment:                                       entitlementManagementCatalogResourceScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementCatalogResourceScopeResourceRoleResourceScope:                                             entitlementManagementCatalogResourceScopeResourceRoleResourceScopeClient,
		EntitlementManagementCatalogResourceScopeResourceScope:                                                         entitlementManagementCatalogResourceScopeResourceScopeClient,
		EntitlementManagementConnectedOrganization:                                                                     entitlementManagementConnectedOrganizationClient,
		EntitlementManagementConnectedOrganizationExternalSponsor:                                                      entitlementManagementConnectedOrganizationExternalSponsorClient,
		EntitlementManagementConnectedOrganizationInternalSponsor:                                                      entitlementManagementConnectedOrganizationInternalSponsorClient,
		EntitlementManagementResource:                                                                                  entitlementManagementResourceClient,
		EntitlementManagementResourceEnvironment:                                                                       entitlementManagementResourceEnvironmentClient,
		EntitlementManagementResourceEnvironmentResource:                                                               entitlementManagementResourceEnvironmentResourceClient,
		EntitlementManagementResourceEnvironmentResourceEnvironment:                                                    entitlementManagementResourceEnvironmentResourceEnvironmentClient,
		EntitlementManagementResourceEnvironmentResourceRole:                                                           entitlementManagementResourceEnvironmentResourceRoleClient,
		EntitlementManagementResourceEnvironmentResourceRoleResource:                                                   entitlementManagementResourceEnvironmentResourceRoleResourceClient,
		EntitlementManagementResourceEnvironmentResourceRoleResourceEnvironment:                                        entitlementManagementResourceEnvironmentResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceEnvironmentResourceRoleResourceScope:                                              entitlementManagementResourceEnvironmentResourceRoleResourceScopeClient,
		EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResource:                                      entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceClient,
		EntitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironment:                           entitlementManagementResourceEnvironmentResourceRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceEnvironmentResourceScope:                                                          entitlementManagementResourceEnvironmentResourceScopeClient,
		EntitlementManagementResourceEnvironmentResourceScopeResource:                                                  entitlementManagementResourceEnvironmentResourceScopeResourceClient,
		EntitlementManagementResourceEnvironmentResourceScopeResourceEnvironment:                                       entitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceEnvironmentResourceScopeResourceRole:                                              entitlementManagementResourceEnvironmentResourceScopeResourceRoleClient,
		EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResource:                                      entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient,
		EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironment:                           entitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRequest:                                                                           entitlementManagementResourceRequestClient,
		EntitlementManagementResourceRequestCatalog:                                                                    entitlementManagementResourceRequestCatalogClient,
		EntitlementManagementResourceRequestCatalogAccessPackage:                                                       entitlementManagementResourceRequestCatalogAccessPackageClient,
		EntitlementManagementResourceRequestCatalogCustomWorkflowExtension:                                             entitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient,
		EntitlementManagementResourceRequestCatalogResource:                                                            entitlementManagementResourceRequestCatalogResourceClient,
		EntitlementManagementResourceRequestCatalogResourceEnvironment:                                                 entitlementManagementResourceRequestCatalogResourceEnvironmentClient,
		EntitlementManagementResourceRequestCatalogResourceRole:                                                        entitlementManagementResourceRequestCatalogResourceRoleClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResource:                                                entitlementManagementResourceRequestCatalogResourceRoleResourceClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceEnvironment:                                     entitlementManagementResourceRequestCatalogResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceRole:                                            entitlementManagementResourceRequestCatalogResourceRoleResourceRoleClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceScope:                                           entitlementManagementResourceRequestCatalogResourceRoleResourceScopeClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResource:                                   entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironment:                        entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRole:                               entitlementManagementResourceRequestCatalogResourceRoleResourceScopeResourceRoleClient,
		EntitlementManagementResourceRequestCatalogResourceScope:                                                       entitlementManagementResourceRequestCatalogResourceScopeClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResource:                                               entitlementManagementResourceRequestCatalogResourceScopeResourceClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceEnvironment:                                    entitlementManagementResourceRequestCatalogResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceRole:                                           entitlementManagementResourceRequestCatalogResourceScopeResourceRoleClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResource:                                   entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironment:                        entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScope:                              entitlementManagementResourceRequestCatalogResourceScopeResourceRoleResourceScopeClient,
		EntitlementManagementResourceRequestCatalogResourceScopeResourceScope:                                          entitlementManagementResourceRequestCatalogResourceScopeResourceScopeClient,
		EntitlementManagementResourceRequestResource:                                                                   entitlementManagementResourceRequestResourceClient,
		EntitlementManagementResourceRequestResourceEnvironment:                                                        entitlementManagementResourceRequestResourceEnvironmentClient,
		EntitlementManagementResourceRequestResourceRole:                                                               entitlementManagementResourceRequestResourceRoleClient,
		EntitlementManagementResourceRequestResourceRoleResource:                                                       entitlementManagementResourceRequestResourceRoleResourceClient,
		EntitlementManagementResourceRequestResourceRoleResourceEnvironment:                                            entitlementManagementResourceRequestResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRequestResourceRoleResourceScope:                                                  entitlementManagementResourceRequestResourceRoleResourceScopeClient,
		EntitlementManagementResourceRequestResourceScope:                                                              entitlementManagementResourceRequestResourceScopeClient,
		EntitlementManagementResourceRequestResourceScopeResource:                                                      entitlementManagementResourceRequestResourceScopeResourceClient,
		EntitlementManagementResourceRequestResourceScopeResourceEnvironment:                                           entitlementManagementResourceRequestResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceRequestResourceScopeResourceRole:                                                  entitlementManagementResourceRequestResourceScopeResourceRoleClient,
		EntitlementManagementResourceRole:                                                                              entitlementManagementResourceRoleClient,
		EntitlementManagementResourceRoleResource:                                                                      entitlementManagementResourceRoleResourceClient,
		EntitlementManagementResourceRoleResourceEnvironment:                                                           entitlementManagementResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRoleResourceScope:                                                                 entitlementManagementResourceRoleResourceScopeClient,
		EntitlementManagementResourceRoleResourceScopeResource:                                                         entitlementManagementResourceRoleResourceScopeResourceClient,
		EntitlementManagementResourceRoleResourceScopeResourceEnvironment:                                              entitlementManagementResourceRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceRoleScope:                                                                         entitlementManagementResourceRoleScopeClient,
		EntitlementManagementResourceRoleScopeRole:                                                                     entitlementManagementResourceRoleScopeRoleClient,
		EntitlementManagementResourceRoleScopeRoleResource:                                                             entitlementManagementResourceRoleScopeRoleResourceClient,
		EntitlementManagementResourceRoleScopeRoleResourceEnvironment:                                                  entitlementManagementResourceRoleScopeRoleResourceEnvironmentClient,
		EntitlementManagementResourceRoleScopeRoleResourceRole:                                                         entitlementManagementResourceRoleScopeRoleResourceRoleClient,
		EntitlementManagementResourceRoleScopeRoleResourceScope:                                                        entitlementManagementResourceRoleScopeRoleResourceScopeClient,
		EntitlementManagementResourceRoleScopeRoleResourceScopeResource:                                                entitlementManagementResourceRoleScopeRoleResourceScopeResourceClient,
		EntitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironment:                                     entitlementManagementResourceRoleScopeRoleResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceRoleScopeRoleResourceScopeResourceRole:                                            entitlementManagementResourceRoleScopeRoleResourceScopeResourceRoleClient,
		EntitlementManagementResourceRoleScopeScope:                                                                    entitlementManagementResourceRoleScopeScopeClient,
		EntitlementManagementResourceRoleScopeScopeResource:                                                            entitlementManagementResourceRoleScopeScopeResourceClient,
		EntitlementManagementResourceRoleScopeScopeResourceEnvironment:                                                 entitlementManagementResourceRoleScopeScopeResourceEnvironmentClient,
		EntitlementManagementResourceRoleScopeScopeResourceRole:                                                        entitlementManagementResourceRoleScopeScopeResourceRoleClient,
		EntitlementManagementResourceRoleScopeScopeResourceRoleResource:                                                entitlementManagementResourceRoleScopeScopeResourceRoleResourceClient,
		EntitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironment:                                     entitlementManagementResourceRoleScopeScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementResourceRoleScopeScopeResourceRoleResourceScope:                                           entitlementManagementResourceRoleScopeScopeResourceRoleResourceScopeClient,
		EntitlementManagementResourceRoleScopeScopeResourceScope:                                                       entitlementManagementResourceRoleScopeScopeResourceScopeClient,
		EntitlementManagementResourceScope:                                                                             entitlementManagementResourceScopeClient,
		EntitlementManagementResourceScopeResource:                                                                     entitlementManagementResourceScopeResourceClient,
		EntitlementManagementResourceScopeResourceEnvironment:                                                          entitlementManagementResourceScopeResourceEnvironmentClient,
		EntitlementManagementResourceScopeResourceRole:                                                                 entitlementManagementResourceScopeResourceRoleClient,
		EntitlementManagementResourceScopeResourceRoleResource:                                                         entitlementManagementResourceScopeResourceRoleResourceClient,
		EntitlementManagementResourceScopeResourceRoleResourceEnvironment:                                              entitlementManagementResourceScopeResourceRoleResourceEnvironmentClient,
		EntitlementManagementSetting:                                                                                   entitlementManagementSettingClient,
		IdentityGovernance:                                                                                             identityGovernanceClient,
		LifecycleWorkflow:                                                                                              lifecycleWorkflowClient,
		LifecycleWorkflowCustomTaskExtension:                                                                           lifecycleWorkflowCustomTaskExtensionClient,
		LifecycleWorkflowCustomTaskExtensionCreatedBy:                                                                  lifecycleWorkflowCustomTaskExtensionCreatedByClient,
		LifecycleWorkflowCustomTaskExtensionCreatedByMailboxSetting:                                                    lifecycleWorkflowCustomTaskExtensionCreatedByMailboxSettingClient,
		LifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningError:                                          lifecycleWorkflowCustomTaskExtensionCreatedByServiceProvisioningErrorClient,
		LifecycleWorkflowCustomTaskExtensionLastModifiedBy:                                                             lifecycleWorkflowCustomTaskExtensionLastModifiedByClient,
		LifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSetting:                                               lifecycleWorkflowCustomTaskExtensionLastModifiedByMailboxSettingClient,
		LifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningError:                                     lifecycleWorkflowCustomTaskExtensionLastModifiedByServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItem:                                                                                   lifecycleWorkflowDeletedItemClient,
		LifecycleWorkflowDeletedItemWorkflow:                                                                           lifecycleWorkflowDeletedItemWorkflowClient,
		LifecycleWorkflowDeletedItemWorkflowCreatedBy:                                                                  lifecycleWorkflowDeletedItemWorkflowCreatedByClient,
		LifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSetting:                                                    lifecycleWorkflowDeletedItemWorkflowCreatedByMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningError:                                          lifecycleWorkflowDeletedItemWorkflowCreatedByServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowExecutionScope:                                                             lifecycleWorkflowDeletedItemWorkflowExecutionScopeClient,
		LifecycleWorkflowDeletedItemWorkflowLastModifiedBy:                                                             lifecycleWorkflowDeletedItemWorkflowLastModifiedByClient,
		LifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSetting:                                               lifecycleWorkflowDeletedItemWorkflowLastModifiedByMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningError:                                     lifecycleWorkflowDeletedItemWorkflowLastModifiedByServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowRun:                                                                        lifecycleWorkflowDeletedItemWorkflowRunClient,
		LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResult:                                                    lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubject:                                             lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSetting:                               lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningError:                     lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTask:                                                lifecycleWorkflowDeletedItemWorkflowRunTaskProcessingResultTaskClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResult:                                                    lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubject:                                             lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSetting:                               lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningError:                     lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResult:                                lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubject:                         lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting:           lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError: lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTask:                            lifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultTaskProcessingResultTaskClient,
		LifecycleWorkflowDeletedItemWorkflowTask:                                                                       lifecycleWorkflowDeletedItemWorkflowTaskClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReport:                                                                 lifecycleWorkflowDeletedItemWorkflowTaskReportClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTask:                                                             lifecycleWorkflowDeletedItemWorkflowTaskReportTaskClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinition:                                                   lifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResult:                                             lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubject:                                      lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting:                        lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError:              lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTask:                                         lifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultTaskClient,
		LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResult:                                                   lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubject:                                            lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSetting:                              lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError:                    lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTask:                                               lifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultTaskClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResult:                                                       lifecycleWorkflowDeletedItemWorkflowUserProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubject:                                                lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSetting:                                  lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningError:                        lifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult:                                   lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubject:                            lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting:              lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError:    lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTask:                               lifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultTaskClient,
		LifecycleWorkflowDeletedItemWorkflowVersion:                                                                    lifecycleWorkflowDeletedItemWorkflowVersionClient,
		LifecycleWorkflowDeletedItemWorkflowVersionCreatedBy:                                                           lifecycleWorkflowDeletedItemWorkflowVersionCreatedByClient,
		LifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSetting:                                             lifecycleWorkflowDeletedItemWorkflowVersionCreatedByMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningError:                                   lifecycleWorkflowDeletedItemWorkflowVersionCreatedByServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedBy:                                                      lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByClient,
		LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSetting:                                        lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningError:                              lifecycleWorkflowDeletedItemWorkflowVersionLastModifiedByServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTask:                                                                lifecycleWorkflowDeletedItemWorkflowVersionTaskClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResult:                                            lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubject:                                     lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting:                       lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError:             lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTask:                                        lifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient,
		LifecycleWorkflowSetting:                                                                                       lifecycleWorkflowSettingClient,
		LifecycleWorkflowTaskDefinition:                                                                                lifecycleWorkflowTaskDefinitionClient,
		LifecycleWorkflowWorkflow:                                                                                      lifecycleWorkflowWorkflowClient,
		LifecycleWorkflowWorkflowCreatedBy:                                                                             lifecycleWorkflowWorkflowCreatedByClient,
		LifecycleWorkflowWorkflowCreatedByMailboxSetting:                                                               lifecycleWorkflowWorkflowCreatedByMailboxSettingClient,
		LifecycleWorkflowWorkflowCreatedByServiceProvisioningError:                                                     lifecycleWorkflowWorkflowCreatedByServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowExecutionScope:                                                                        lifecycleWorkflowWorkflowExecutionScopeClient,
		LifecycleWorkflowWorkflowLastModifiedBy:                                                                        lifecycleWorkflowWorkflowLastModifiedByClient,
		LifecycleWorkflowWorkflowLastModifiedByMailboxSetting:                                                          lifecycleWorkflowWorkflowLastModifiedByMailboxSettingClient,
		LifecycleWorkflowWorkflowLastModifiedByServiceProvisioningError:                                                lifecycleWorkflowWorkflowLastModifiedByServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowRun:                                                                                   lifecycleWorkflowWorkflowRunClient,
		LifecycleWorkflowWorkflowRunTaskProcessingResult:                                                               lifecycleWorkflowWorkflowRunTaskProcessingResultClient,
		LifecycleWorkflowWorkflowRunTaskProcessingResultSubject:                                                        lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSetting:                                          lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningError:                                lifecycleWorkflowWorkflowRunTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowRunTaskProcessingResultTask:                                                           lifecycleWorkflowWorkflowRunTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowRunUserProcessingResult:                                                               lifecycleWorkflowWorkflowRunUserProcessingResultClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultSubject:                                                        lifecycleWorkflowWorkflowRunUserProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSetting:                                          lifecycleWorkflowWorkflowRunUserProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningError:                                lifecycleWorkflowWorkflowRunUserProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResult:                                           lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubject:                                    lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSetting:                      lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError:            lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTask:                                       lifecycleWorkflowWorkflowRunUserProcessingResultTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowTask:                                                                                  lifecycleWorkflowWorkflowTaskClient,
		LifecycleWorkflowWorkflowTaskReport:                                                                            lifecycleWorkflowWorkflowTaskReportClient,
		LifecycleWorkflowWorkflowTaskReportTask:                                                                        lifecycleWorkflowWorkflowTaskReportTaskClient,
		LifecycleWorkflowWorkflowTaskReportTaskDefinition:                                                              lifecycleWorkflowWorkflowTaskReportTaskDefinitionClient,
		LifecycleWorkflowWorkflowTaskReportTaskProcessingResult:                                                        lifecycleWorkflowWorkflowTaskReportTaskProcessingResultClient,
		LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubject:                                                 lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSetting:                                   lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningError:                         lifecycleWorkflowWorkflowTaskReportTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowTaskReportTaskProcessingResultTask:                                                    lifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowTaskTaskProcessingResult:                                                              lifecycleWorkflowWorkflowTaskTaskProcessingResultClient,
		LifecycleWorkflowWorkflowTaskTaskProcessingResultSubject:                                                       lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSetting:                                         lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningError:                               lifecycleWorkflowWorkflowTaskTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowTaskTaskProcessingResultTask:                                                          lifecycleWorkflowWorkflowTaskTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowTemplate:                                                                              lifecycleWorkflowWorkflowTemplateClient,
		LifecycleWorkflowWorkflowTemplateTask:                                                                          lifecycleWorkflowWorkflowTemplateTaskClient,
		LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResult:                                                      lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultClient,
		LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubject:                                               lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSetting:                                 lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningError:                       lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTask:                                                  lifecycleWorkflowWorkflowTemplateTaskTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowUserProcessingResult:                                                                  lifecycleWorkflowWorkflowUserProcessingResultClient,
		LifecycleWorkflowWorkflowUserProcessingResultSubject:                                                           lifecycleWorkflowWorkflowUserProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSetting:                                             lifecycleWorkflowWorkflowUserProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningError:                                   lifecycleWorkflowWorkflowUserProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResult:                                              lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultClient,
		LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubject:                                       lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSetting:                         lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningError:               lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTask:                                          lifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultTaskClient,
		LifecycleWorkflowWorkflowVersion:                                                                               lifecycleWorkflowWorkflowVersionClient,
		LifecycleWorkflowWorkflowVersionCreatedBy:                                                                      lifecycleWorkflowWorkflowVersionCreatedByClient,
		LifecycleWorkflowWorkflowVersionCreatedByMailboxSetting:                                                        lifecycleWorkflowWorkflowVersionCreatedByMailboxSettingClient,
		LifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningError:                                              lifecycleWorkflowWorkflowVersionCreatedByServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowVersionLastModifiedBy:                                                                 lifecycleWorkflowWorkflowVersionLastModifiedByClient,
		LifecycleWorkflowWorkflowVersionLastModifiedByMailboxSetting:                                                   lifecycleWorkflowWorkflowVersionLastModifiedByMailboxSettingClient,
		LifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningError:                                         lifecycleWorkflowWorkflowVersionLastModifiedByServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowVersionTask:                                                                           lifecycleWorkflowWorkflowVersionTaskClient,
		LifecycleWorkflowWorkflowVersionTaskTaskProcessingResult:                                                       lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultClient,
		LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubject:                                                lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectClient,
		LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSetting:                                  lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectMailboxSettingClient,
		LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningError:                        lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultSubjectServiceProvisioningErrorClient,
		LifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTask:                                                   lifecycleWorkflowWorkflowVersionTaskTaskProcessingResultTaskClient,
		PrivilegedAccess:                                                              privilegedAccessClient,
		PrivilegedAccessGroup:                                                         privilegedAccessGroupClient,
		PrivilegedAccessGroupAssignmentApproval:                                       privilegedAccessGroupAssignmentApprovalClient,
		PrivilegedAccessGroupAssignmentApprovalStage:                                  privilegedAccessGroupAssignmentApprovalStageClient,
		PrivilegedAccessGroupAssignmentSchedule:                                       privilegedAccessGroupAssignmentScheduleClient,
		PrivilegedAccessGroupAssignmentScheduleActivatedUsing:                         privilegedAccessGroupAssignmentScheduleActivatedUsingClient,
		PrivilegedAccessGroupAssignmentScheduleGroup:                                  privilegedAccessGroupAssignmentScheduleGroupClient,
		PrivilegedAccessGroupAssignmentScheduleGroupServiceProvisioningError:          privilegedAccessGroupAssignmentScheduleGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupAssignmentScheduleInstance:                               privilegedAccessGroupAssignmentScheduleInstanceClient,
		PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsing:                 privilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient,
		PrivilegedAccessGroupAssignmentScheduleInstanceGroup:                          privilegedAccessGroupAssignmentScheduleInstanceGroupClient,
		PrivilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningError:  privilegedAccessGroupAssignmentScheduleInstanceGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupAssignmentScheduleInstancePrincipal:                      privilegedAccessGroupAssignmentScheduleInstancePrincipalClient,
		PrivilegedAccessGroupAssignmentSchedulePrincipal:                              privilegedAccessGroupAssignmentSchedulePrincipalClient,
		PrivilegedAccessGroupAssignmentScheduleRequest:                                privilegedAccessGroupAssignmentScheduleRequestClient,
		PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsing:                  privilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient,
		PrivilegedAccessGroupAssignmentScheduleRequestGroup:                           privilegedAccessGroupAssignmentScheduleRequestGroupClient,
		PrivilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningError:   privilegedAccessGroupAssignmentScheduleRequestGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupAssignmentScheduleRequestPrincipal:                       privilegedAccessGroupAssignmentScheduleRequestPrincipalClient,
		PrivilegedAccessGroupAssignmentScheduleRequestTargetSchedule:                  privilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient,
		PrivilegedAccessGroupEligibilitySchedule:                                      privilegedAccessGroupEligibilityScheduleClient,
		PrivilegedAccessGroupEligibilityScheduleGroup:                                 privilegedAccessGroupEligibilityScheduleGroupClient,
		PrivilegedAccessGroupEligibilityScheduleGroupServiceProvisioningError:         privilegedAccessGroupEligibilityScheduleGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupEligibilityScheduleInstance:                              privilegedAccessGroupEligibilityScheduleInstanceClient,
		PrivilegedAccessGroupEligibilityScheduleInstanceGroup:                         privilegedAccessGroupEligibilityScheduleInstanceGroupClient,
		PrivilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningError: privilegedAccessGroupEligibilityScheduleInstanceGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupEligibilityScheduleInstancePrincipal:                     privilegedAccessGroupEligibilityScheduleInstancePrincipalClient,
		PrivilegedAccessGroupEligibilitySchedulePrincipal:                             privilegedAccessGroupEligibilitySchedulePrincipalClient,
		PrivilegedAccessGroupEligibilityScheduleRequest:                               privilegedAccessGroupEligibilityScheduleRequestClient,
		PrivilegedAccessGroupEligibilityScheduleRequestGroup:                          privilegedAccessGroupEligibilityScheduleRequestGroupClient,
		PrivilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningError:  privilegedAccessGroupEligibilityScheduleRequestGroupServiceProvisioningErrorClient,
		PrivilegedAccessGroupEligibilityScheduleRequestPrincipal:                      privilegedAccessGroupEligibilityScheduleRequestPrincipalClient,
		PrivilegedAccessGroupEligibilityScheduleRequestTargetSchedule:                 privilegedAccessGroupEligibilityScheduleRequestTargetScheduleClient,
		TermsOfUse:                                 termsOfUseClient,
		TermsOfUseAgreement:                        termsOfUseAgreementClient,
		TermsOfUseAgreementAcceptance:              termsOfUseAgreementAcceptanceClient,
		TermsOfUseAgreementFile:                    termsOfUseAgreementFileClient,
		TermsOfUseAgreementFileLocalization:        termsOfUseAgreementFileLocalizationClient,
		TermsOfUseAgreementFileLocalizationVersion: termsOfUseAgreementFileLocalizationVersionClient,
		TermsOfUseAgreementFileVersion:             termsOfUseAgreementFileVersionClient,
	}, nil
}
