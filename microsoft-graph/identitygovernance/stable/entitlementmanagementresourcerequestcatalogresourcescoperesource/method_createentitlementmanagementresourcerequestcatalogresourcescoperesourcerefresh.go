package entitlementmanagementresourcerequestcatalogresourcescoperesource

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

type CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions() CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions {
	return CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefresh - Invoke action refresh
func (c EntitlementManagementResourceRequestCatalogResourceScopeResourceClient) CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefresh(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRequestIdCatalogResourceIdScopeId, options CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationOptions) (result CreateEntitlementManagementResourceRequestCatalogResourceScopeResourceRefreshOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/refresh", id.ID()),
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
