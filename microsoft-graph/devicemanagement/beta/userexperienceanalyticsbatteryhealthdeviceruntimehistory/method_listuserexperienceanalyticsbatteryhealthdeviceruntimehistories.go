package userexperienceanalyticsbatteryhealthdeviceruntimehistory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
}

type ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
}

type ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions() ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions {
	return ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions{}
}

func (o ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistories - Get
// userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory from deviceManagement. User Experience Analytics Battery
// Health Device Runtime History
func (c UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient) ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistories(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions) (result ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient) ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesComplete(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions) (ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationPredicate{})
}

// ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryClient) ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesOperationOptions, predicate UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryOperationPredicate) (result ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory, 0)

	resp, err := c.ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistories(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
