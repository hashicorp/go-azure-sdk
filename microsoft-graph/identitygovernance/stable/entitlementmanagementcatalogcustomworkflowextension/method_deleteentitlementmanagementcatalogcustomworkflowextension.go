package entitlementmanagementcatalogcustomworkflowextension

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions() DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions {
	return DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions{}
}

func (o DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementCatalogCustomWorkflowExtension - Delete accessPackageAssignmentRequestWorkflowExtension.
// Delete an accessPackageAssignmentRequestWorkflowExtension object. The custom workflow extension must first be removed
// from any associated policies before it can be deleted. Follow these steps to remove the custom workflow extension
// from any associated policies
func (c EntitlementManagementCatalogCustomWorkflowExtensionClient) DeleteEntitlementManagementCatalogCustomWorkflowExtension(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdCustomWorkflowExtensionId, options DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationOptions) (result DeleteEntitlementManagementCatalogCustomWorkflowExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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
