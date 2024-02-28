package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSubscriptionLevelAlertsByRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Alert
}

type ListSubscriptionLevelAlertsByRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Alert
}

type ListSubscriptionLevelAlertsByRegionOperationOptions struct {
	AutoDismissRuleName *string
	Expand              *string
	Filter              *string
	Select              *string
}

func DefaultListSubscriptionLevelAlertsByRegionOperationOptions() ListSubscriptionLevelAlertsByRegionOperationOptions {
	return ListSubscriptionLevelAlertsByRegionOperationOptions{}
}

func (o ListSubscriptionLevelAlertsByRegionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSubscriptionLevelAlertsByRegionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListSubscriptionLevelAlertsByRegionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.AutoDismissRuleName != nil {
		out.Append("autoDismissRuleName", fmt.Sprintf("%v", *o.AutoDismissRuleName))
	}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

// ListSubscriptionLevelAlertsByRegion ...
func (c AlertsClient) ListSubscriptionLevelAlertsByRegion(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions) (result ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/alerts", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Alert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSubscriptionLevelAlertsByRegionComplete retrieves all the results into a single object
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionComplete(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions) (ListSubscriptionLevelAlertsByRegionCompleteResult, error) {
	return c.ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx, id, options, AlertOperationPredicate{})
}

// ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, options ListSubscriptionLevelAlertsByRegionOperationOptions, predicate AlertOperationPredicate) (result ListSubscriptionLevelAlertsByRegionCompleteResult, err error) {
	items := make([]Alert, 0)

	resp, err := c.ListSubscriptionLevelAlertsByRegion(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListSubscriptionLevelAlertsByRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
