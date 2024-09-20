package appconsentappconsentrequestuserconsentrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAppConsentRequestUserConsentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserConsentRequest
}

type CreateAppConsentRequestUserConsentRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAppConsentRequestUserConsentRequestOperationOptions() CreateAppConsentRequestUserConsentRequestOperationOptions {
	return CreateAppConsentRequestUserConsentRequestOperationOptions{}
}

func (o CreateAppConsentRequestUserConsentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAppConsentRequestUserConsentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAppConsentRequestUserConsentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAppConsentRequestUserConsentRequest - Create new navigation property to userConsentRequests for
// identityGovernance
func (c AppConsentAppConsentRequestUserConsentRequestClient) CreateAppConsentRequestUserConsentRequest(ctx context.Context, id stable.IdentityGovernanceAppConsentAppConsentRequestId, input stable.UserConsentRequest, options CreateAppConsentRequestUserConsentRequestOperationOptions) (result CreateAppConsentRequestUserConsentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userConsentRequests", id.ID()),
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

	var model stable.UserConsentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
