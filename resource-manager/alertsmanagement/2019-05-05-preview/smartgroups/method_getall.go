package smartgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SmartGroup
}

type GetAllCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SmartGroup
}

type GetAllOperationOptions struct {
	MonitorCondition    *MonitorCondition
	MonitorService      *MonitorService
	PageCount           *int64
	Severity            *Severity
	SmartGroupState     *AlertState
	SortBy              *SmartGroupsSortByFields
	SortOrder           *SortOrder
	TargetResource      *string
	TargetResourceGroup *string
	TargetResourceType  *string
	TimeRange           *TimeRange
}

func DefaultGetAllOperationOptions() GetAllOperationOptions {
	return GetAllOperationOptions{}
}

func (o GetAllOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAllOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetAllOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.MonitorCondition != nil {
		out.Append("monitorCondition", fmt.Sprintf("%v", *o.MonitorCondition))
	}
	if o.MonitorService != nil {
		out.Append("monitorService", fmt.Sprintf("%v", *o.MonitorService))
	}
	if o.PageCount != nil {
		out.Append("pageCount", fmt.Sprintf("%v", *o.PageCount))
	}
	if o.Severity != nil {
		out.Append("severity", fmt.Sprintf("%v", *o.Severity))
	}
	if o.SmartGroupState != nil {
		out.Append("smartGroupState", fmt.Sprintf("%v", *o.SmartGroupState))
	}
	if o.SortBy != nil {
		out.Append("sortBy", fmt.Sprintf("%v", *o.SortBy))
	}
	if o.SortOrder != nil {
		out.Append("sortOrder", fmt.Sprintf("%v", *o.SortOrder))
	}
	if o.TargetResource != nil {
		out.Append("targetResource", fmt.Sprintf("%v", *o.TargetResource))
	}
	if o.TargetResourceGroup != nil {
		out.Append("targetResourceGroup", fmt.Sprintf("%v", *o.TargetResourceGroup))
	}
	if o.TargetResourceType != nil {
		out.Append("targetResourceType", fmt.Sprintf("%v", *o.TargetResourceType))
	}
	if o.TimeRange != nil {
		out.Append("timeRange", fmt.Sprintf("%v", *o.TimeRange))
	}
	return &out
}

type GetAllCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetAllCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAll ...
func (c SmartGroupsClient) GetAll(ctx context.Context, id commonids.SubscriptionId, options GetAllOperationOptions) (result GetAllOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetAllCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.AlertsManagement/smartGroups", id.ID()),
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
		Values *[]SmartGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllComplete retrieves all the results into a single object
func (c SmartGroupsClient) GetAllComplete(ctx context.Context, id commonids.SubscriptionId, options GetAllOperationOptions) (GetAllCompleteResult, error) {
	return c.GetAllCompleteMatchingPredicate(ctx, id, options, SmartGroupOperationPredicate{})
}

// GetAllCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SmartGroupsClient) GetAllCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, options GetAllOperationOptions, predicate SmartGroupOperationPredicate) (result GetAllCompleteResult, err error) {
	items := make([]SmartGroup, 0)

	resp, err := c.GetAll(ctx, id, options)
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

	result = GetAllCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
