package onlinemeetingaiinsight

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

type CreateOnlineMeetingAiInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CallAiInsight
}

type CreateOnlineMeetingAiInsightOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOnlineMeetingAiInsightOperationOptions() CreateOnlineMeetingAiInsightOperationOptions {
	return CreateOnlineMeetingAiInsightOperationOptions{}
}

func (o CreateOnlineMeetingAiInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingAiInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingAiInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingAiInsight - Create new navigation property to aiInsights for users
func (c OnlineMeetingAiInsightClient) CreateOnlineMeetingAiInsight(ctx context.Context, id beta.UserIdOnlineMeetingId, input beta.CallAiInsight, options CreateOnlineMeetingAiInsightOperationOptions) (result CreateOnlineMeetingAiInsightOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/aiInsights", id.ID()),
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

	var model beta.CallAiInsight
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
