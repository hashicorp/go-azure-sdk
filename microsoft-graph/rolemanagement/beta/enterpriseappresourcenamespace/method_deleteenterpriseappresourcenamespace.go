package enterpriseappresourcenamespace

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteEnterpriseAppResourceNamespaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEnterpriseAppResourceNamespaceOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEnterpriseAppResourceNamespaceOperationOptions() DeleteEnterpriseAppResourceNamespaceOperationOptions {
	return DeleteEnterpriseAppResourceNamespaceOperationOptions{}
}

func (o DeleteEnterpriseAppResourceNamespaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEnterpriseAppResourceNamespaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEnterpriseAppResourceNamespaceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEnterpriseAppResourceNamespace - Delete navigation property resourceNamespaces for roleManagement
func (c EnterpriseAppResourceNamespaceClient) DeleteEnterpriseAppResourceNamespace(ctx context.Context, id beta.RoleManagementEnterpriseAppIdResourceNamespaceId, options DeleteEnterpriseAppResourceNamespaceOperationOptions) (result DeleteEnterpriseAppResourceNamespaceOperationResponse, err error) {
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
