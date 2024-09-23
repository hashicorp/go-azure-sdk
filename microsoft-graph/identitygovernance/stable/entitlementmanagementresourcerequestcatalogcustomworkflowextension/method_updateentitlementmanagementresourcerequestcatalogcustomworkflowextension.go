package entitlementmanagementresourcerequestcatalogcustomworkflowextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions() UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions {
	return UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions{}
}

func (o UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtension - Update the navigation property
// customWorkflowExtensions in identityGovernance
func (c EntitlementManagementResourceRequestCatalogCustomWorkflowExtensionClient) UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtension(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogCustomWorkflowExtensionId, input stable.CustomCalloutExtension, options UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationOptions) (result UpdateEntitlementManagementResourceRequestCatalogCustomWorkflowExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
