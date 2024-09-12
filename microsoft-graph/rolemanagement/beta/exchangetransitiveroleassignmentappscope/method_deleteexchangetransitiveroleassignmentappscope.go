package exchangetransitiveroleassignmentappscope

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

type DeleteExchangeTransitiveRoleAssignmentAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions() DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions {
	return DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions{}
}

func (o DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteExchangeTransitiveRoleAssignmentAppScope - Delete navigation property appScope for roleManagement
func (c ExchangeTransitiveRoleAssignmentAppScopeClient) DeleteExchangeTransitiveRoleAssignmentAppScope(ctx context.Context, id beta.RoleManagementExchangeTransitiveRoleAssignmentId, options DeleteExchangeTransitiveRoleAssignmentAppScopeOperationOptions) (result DeleteExchangeTransitiveRoleAssignmentAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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
