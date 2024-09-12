package signin

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSignInOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.SignIn
}

type GetSignInOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetSignInOperationOptions() GetSignInOperationOptions {
	return GetSignInOperationOptions{}
}

func (o GetSignInOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSignInOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSignInOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSignIn - Get signIn. Retrieve a specific Microsoft Entra user sign-in event for your tenant. Sign-ins that are
// interactive in nature (where a username/password is passed as part of auth token) and successful federated sign-ins
// are currently included in the sign-in logs.
func (c SignInClient) GetSignIn(ctx context.Context, id stable.AuditLogSignInId, options GetSignInOperationOptions) (result GetSignInOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.SignIn
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
