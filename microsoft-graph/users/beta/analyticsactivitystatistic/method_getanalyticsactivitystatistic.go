package analyticsactivitystatistic

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAnalyticsActivityStatisticOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.ActivityStatistics
}

type GetAnalyticsActivityStatisticOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAnalyticsActivityStatisticOperationOptions() GetAnalyticsActivityStatisticOperationOptions {
	return GetAnalyticsActivityStatisticOperationOptions{}
}

func (o GetAnalyticsActivityStatisticOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAnalyticsActivityStatisticOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAnalyticsActivityStatisticOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAnalyticsActivityStatistic - Get activityStatistics from users. The collection of work activities that a user
// spent time on during and outside of working hours. Read-only. Nullable.
func (c AnalyticsActivityStatisticClient) GetAnalyticsActivityStatistic(ctx context.Context, id beta.UserIdAnalyticsActivityStatisticId, options GetAnalyticsActivityStatisticOperationOptions) (result GetAnalyticsActivityStatisticOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalActivityStatisticsImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
