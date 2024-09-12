package conditionalaccessauthenticationcontextclassreference

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConditionalAccessAuthenticationContextClassReferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationContextClassReference
}

type GetConditionalAccessAuthenticationContextClassReferenceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetConditionalAccessAuthenticationContextClassReferenceOperationOptions() GetConditionalAccessAuthenticationContextClassReferenceOperationOptions {
	return GetConditionalAccessAuthenticationContextClassReferenceOperationOptions{}
}

func (o GetConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConditionalAccessAuthenticationContextClassReference - Get authenticationContextClassReference. Retrieve the
// properties and relationships of a authenticationContextClassReference object.
func (c ConditionalAccessAuthenticationContextClassReferenceClient) GetConditionalAccessAuthenticationContextClassReference(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationContextClassReferenceId, options GetConditionalAccessAuthenticationContextClassReferenceOperationOptions) (result GetConditionalAccessAuthenticationContextClassReferenceOperationResponse, err error) {
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

	var model stable.AuthenticationContextClassReference
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
