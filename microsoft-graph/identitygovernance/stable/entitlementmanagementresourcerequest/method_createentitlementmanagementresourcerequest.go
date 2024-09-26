package entitlementmanagementresourcerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementResourceRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceRequest
}

type CreateEntitlementManagementResourceRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceRequestOperationOptions() CreateEntitlementManagementResourceRequestOperationOptions {
	return CreateEntitlementManagementResourceRequestOperationOptions{}
}

func (o CreateEntitlementManagementResourceRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRequest - Create accessPackageResourceRequest. Create a new
// accessPackageResourceRequest object to request the addition of a resource to an access package catalog, update of a
// resource, or the removal of a resource from a catalog. A resource must be included in an access package catalog
// before a role of that resource can be added to an access package.
func (c EntitlementManagementResourceRequestClient) CreateEntitlementManagementResourceRequest(ctx context.Context, input stable.AccessPackageResourceRequest, options CreateEntitlementManagementResourceRequestOperationOptions) (result CreateEntitlementManagementResourceRequestOperationResponse, err error) {
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
		Path:          "/identityGovernance/entitlementManagement/resourceRequests",
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

	var model stable.AccessPackageResourceRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
