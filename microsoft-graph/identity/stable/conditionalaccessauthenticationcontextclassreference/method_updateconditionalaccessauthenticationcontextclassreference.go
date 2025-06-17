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

type UpdateConditionalAccessAuthenticationContextClassReferenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions() UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions {
	return UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions{}
}

func (o UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateConditionalAccessAuthenticationContextClassReference - Create or Update authenticationContextClassReference.
// Create an authenticationContextClassReference object, if the ID has not been used. If ID has been used, this call
// updates the authenticationContextClassReference object.
func (c ConditionalAccessAuthenticationContextClassReferenceClient) UpdateConditionalAccessAuthenticationContextClassReference(ctx context.Context, id stable.IdentityConditionalAccessAuthenticationContextClassReferenceId, input stable.AuthenticationContextClassReference, options UpdateConditionalAccessAuthenticationContextClassReferenceOperationOptions) (result UpdateConditionalAccessAuthenticationContextClassReferenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
