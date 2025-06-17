package userinsightmonthlymfaregistereduser

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserInsightMonthlyMfaRegisteredUsersCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions() GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions {
	return GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions{}
}

func (o GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserInsightMonthlyMfaRegisteredUsersCount - Get the number of the resource
func (c UserInsightMonthlyMfaRegisteredUserClient) GetUserInsightMonthlyMfaRegisteredUsersCount(ctx context.Context, options GetUserInsightMonthlyMfaRegisteredUsersCountOperationOptions) (result GetUserInsightMonthlyMfaRegisteredUsersCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/userInsights/monthly/mfaRegisteredUsers/$count",
		RetryFunc:     options.RetryFunc,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
