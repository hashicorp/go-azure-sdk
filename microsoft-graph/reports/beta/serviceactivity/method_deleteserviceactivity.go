package serviceactivity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteServiceActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteServiceActivityOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteServiceActivityOperationOptions() DeleteServiceActivityOperationOptions {
	return DeleteServiceActivityOperationOptions{}
}

func (o DeleteServiceActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteServiceActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteServiceActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteServiceActivity - Delete navigation property serviceActivity for reports
func (c ServiceActivityClient) DeleteServiceActivity(ctx context.Context, options DeleteServiceActivityOperationOptions) (result DeleteServiceActivityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/serviceActivity",
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
