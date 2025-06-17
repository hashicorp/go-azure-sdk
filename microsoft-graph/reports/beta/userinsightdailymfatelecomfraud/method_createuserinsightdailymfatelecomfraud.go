package userinsightdailymfatelecomfraud

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserInsightDailyMfaTelecomFraudOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MfaTelecomFraudMetric
}

type CreateUserInsightDailyMfaTelecomFraudOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserInsightDailyMfaTelecomFraudOperationOptions() CreateUserInsightDailyMfaTelecomFraudOperationOptions {
	return CreateUserInsightDailyMfaTelecomFraudOperationOptions{}
}

func (o CreateUserInsightDailyMfaTelecomFraudOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserInsightDailyMfaTelecomFraudOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserInsightDailyMfaTelecomFraudOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserInsightDailyMfaTelecomFraud - Create new navigation property to mfaTelecomFraud for reports
func (c UserInsightDailyMfaTelecomFraudClient) CreateUserInsightDailyMfaTelecomFraud(ctx context.Context, input beta.MfaTelecomFraudMetric, options CreateUserInsightDailyMfaTelecomFraudOperationOptions) (result CreateUserInsightDailyMfaTelecomFraudOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/userInsights/daily/mfaTelecomFraud",
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

	var model beta.MfaTelecomFraudMetric
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
