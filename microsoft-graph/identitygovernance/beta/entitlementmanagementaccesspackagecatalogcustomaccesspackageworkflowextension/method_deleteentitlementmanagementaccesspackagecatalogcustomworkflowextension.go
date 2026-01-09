package entitlementmanagementaccesspackagecatalogcustomaccesspackageworkflowextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions() DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions {
	return DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtension - Delete
// accessPackageAssignmentRequestWorkflowExtension. Delete an accessPackageAssignmentRequestWorkflowExtension object.
// The custom workflow extension must first be removed from any associated policies before it can be deleted. Follow
// these steps to remove the custom workflow extension from any associated policies: 1. First retrieve the
// accessPackageCatalogId by calling the Get accessPackageAssignmentPolicies operation and appending
// ?$expand=accessPackage($expand=accessPackageCatalog) to the query. For example,
// https://graph.microsoft.com/beta/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies?$expand=accessPackage($expand=accessPackageCatalog).
// 2. Use the access package catalog ID and retrieve the ID of the accessPackageCustomWorkflowExtension object that you
// want to delete by running the List accessPackageCustomWorkflowExtensions operation. 3. Call the Update
// accessPackageAssignmentPolicy operation to remove the custom workflow extension object from the policy. For an
// example, see Example 3: Remove the customExtensionStageSettings from a policy.
func (c EntitlementManagementAccessPackageCatalogCustomAccessPackageWorkflowExtensionClient) DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtension(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogIdCustomAccessPackageWorkflowExtensionId, options DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationOptions) (result DeleteEntitlementManagementAccessPackageCatalogCustomWorkflowExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
