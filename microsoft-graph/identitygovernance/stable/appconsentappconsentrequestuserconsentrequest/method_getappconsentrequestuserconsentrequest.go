package appconsentappconsentrequestuserconsentrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppConsentRequestUserConsentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserConsentRequest
}

type GetAppConsentRequestUserConsentRequestOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAppConsentRequestUserConsentRequestOperationOptions() GetAppConsentRequestUserConsentRequestOperationOptions {
	return GetAppConsentRequestUserConsentRequestOperationOptions{}
}

func (o GetAppConsentRequestUserConsentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAppConsentRequestUserConsentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAppConsentRequestUserConsentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAppConsentRequestUserConsentRequest - Get userConsentRequest. Read the properties and relationships of a
// userConsentRequest object.
func (c AppConsentAppConsentRequestUserConsentRequestClient) GetAppConsentRequestUserConsentRequest(ctx context.Context, id stable.IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, options GetAppConsentRequestUserConsentRequestOperationOptions) (result GetAppConsentRequestUserConsentRequestOperationResponse, err error) {
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

	var model stable.UserConsentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
