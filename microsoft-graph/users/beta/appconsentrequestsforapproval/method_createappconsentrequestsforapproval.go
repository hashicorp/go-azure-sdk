package appconsentrequestsforapproval

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

type CreateAppConsentRequestsForApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AppConsentRequest
}

type CreateAppConsentRequestsForApprovalOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAppConsentRequestsForApprovalOperationOptions() CreateAppConsentRequestsForApprovalOperationOptions {
	return CreateAppConsentRequestsForApprovalOperationOptions{}
}

func (o CreateAppConsentRequestsForApprovalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAppConsentRequestsForApprovalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAppConsentRequestsForApprovalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAppConsentRequestsForApproval - Create new navigation property to appConsentRequestsForApproval for users
func (c AppConsentRequestsForApprovalClient) CreateAppConsentRequestsForApproval(ctx context.Context, id beta.UserId, input beta.AppConsentRequest, options CreateAppConsentRequestsForApprovalOperationOptions) (result CreateAppConsentRequestsForApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appConsentRequestsForApproval", id.ID()),
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

	var model beta.AppConsentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
