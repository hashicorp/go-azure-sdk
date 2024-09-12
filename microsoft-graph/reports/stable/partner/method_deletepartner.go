package partner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePartnerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePartnerOperationOptions struct {
	IfMatch *string
}

func DefaultDeletePartnerOperationOptions() DeletePartnerOperationOptions {
	return DeletePartnerOperationOptions{}
}

func (o DeletePartnerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePartnerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeletePartnerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePartner - Delete navigation property partners for reports
func (c PartnerClient) DeletePartner(ctx context.Context, options DeletePartnerOperationOptions) (result DeletePartnerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/partners",
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
