package datasecurityandgovernanceactivity

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

type DeleteDataSecurityAndGovernanceActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDataSecurityAndGovernanceActivityOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteDataSecurityAndGovernanceActivityOperationOptions() DeleteDataSecurityAndGovernanceActivityOperationOptions {
	return DeleteDataSecurityAndGovernanceActivityOperationOptions{}
}

func (o DeleteDataSecurityAndGovernanceActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDataSecurityAndGovernanceActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteDataSecurityAndGovernanceActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDataSecurityAndGovernanceActivity - Delete navigation property activities for users
func (c DataSecurityAndGovernanceActivityClient) DeleteDataSecurityAndGovernanceActivity(ctx context.Context, id beta.UserId, options DeleteDataSecurityAndGovernanceActivityOperationOptions) (result DeleteDataSecurityAndGovernanceActivityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dataSecurityAndGovernance/activities", id.ID()),
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
