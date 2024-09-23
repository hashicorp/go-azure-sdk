package appconsentrequestsforapprovaluserconsentrequest

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

type CreateAppConsentRequestsForApprovalUserConsentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserConsentRequest
}

type CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions() CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions {
	return CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions{}
}

func (o CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAppConsentRequestsForApprovalUserConsentRequest - Create new navigation property to userConsentRequests for
// users
func (c AppConsentRequestsForApprovalUserConsentRequestClient) CreateAppConsentRequestsForApprovalUserConsentRequest(ctx context.Context, id beta.UserIdAppConsentRequestsForApprovalId, input beta.UserConsentRequest, options CreateAppConsentRequestsForApprovalUserConsentRequestOperationOptions) (result CreateAppConsentRequestsForApprovalUserConsentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/userConsentRequests", id.ID()),
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

	var model beta.UserConsentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
