package entitlementmanagementresourcenamespaceresourceaction

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

type GetEntitlementManagementResourceNamespaceResourceActionsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions() GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions {
	return GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions{}
}

func (o GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementResourceNamespaceResourceActionsCount - Get the number of the resource
func (c EntitlementManagementResourceNamespaceResourceActionClient) GetEntitlementManagementResourceNamespaceResourceActionsCount(ctx context.Context, id stable.RoleManagementEntitlementManagementResourceNamespaceId, options GetEntitlementManagementResourceNamespaceResourceActionsCountOperationOptions) (result GetEntitlementManagementResourceNamespaceResourceActionsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resourceActions/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
