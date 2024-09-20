package profileinterest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateProfileInterestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateProfileInterestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateProfileInterestOperationOptions() UpdateProfileInterestOperationOptions {
	return UpdateProfileInterestOperationOptions{}
}

func (o UpdateProfileInterestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateProfileInterestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateProfileInterestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateProfileInterest - Update personinterest. Update the properties of a personInterest object in a user's profile.
func (c ProfileInterestClient) UpdateProfileInterest(ctx context.Context, id beta.MeProfileInterestId, input beta.PersonInterest, options UpdateProfileInterestOperationOptions) (result UpdateProfileInterestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
