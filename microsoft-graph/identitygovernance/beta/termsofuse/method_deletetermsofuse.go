package termsofuse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteTermsOfUseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteTermsOfUseOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteTermsOfUseOperationOptions() DeleteTermsOfUseOperationOptions {
	return DeleteTermsOfUseOperationOptions{}
}

func (o DeleteTermsOfUseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteTermsOfUseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteTermsOfUseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteTermsOfUse - Delete navigation property termsOfUse for identityGovernance
func (c TermsOfUseClient) DeleteTermsOfUse(ctx context.Context, options DeleteTermsOfUseOperationOptions) (result DeleteTermsOfUseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/identityGovernance/termsOfUse",
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
