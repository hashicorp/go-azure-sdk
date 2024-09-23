package usercredentialusagedetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserCredentialUsageDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserCredentialUsageDetails
}

type CreateUserCredentialUsageDetailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserCredentialUsageDetailOperationOptions() CreateUserCredentialUsageDetailOperationOptions {
	return CreateUserCredentialUsageDetailOperationOptions{}
}

func (o CreateUserCredentialUsageDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserCredentialUsageDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserCredentialUsageDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserCredentialUsageDetail - Create new navigation property to userCredentialUsageDetails for reports
func (c UserCredentialUsageDetailClient) CreateUserCredentialUsageDetail(ctx context.Context, input beta.UserCredentialUsageDetails, options CreateUserCredentialUsageDetailOperationOptions) (result CreateUserCredentialUsageDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/userCredentialUsageDetails",
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

	var model beta.UserCredentialUsageDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
