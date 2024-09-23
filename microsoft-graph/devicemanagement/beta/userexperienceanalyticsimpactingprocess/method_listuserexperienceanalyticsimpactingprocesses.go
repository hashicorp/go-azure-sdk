package userexperienceanalyticsimpactingprocess

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

type ListUserExperienceAnalyticsImpactingProcessesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsImpactingProcess
}

type ListUserExperienceAnalyticsImpactingProcessesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsImpactingProcess
}

type ListUserExperienceAnalyticsImpactingProcessesOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsImpactingProcessesOperationOptions() ListUserExperienceAnalyticsImpactingProcessesOperationOptions {
	return ListUserExperienceAnalyticsImpactingProcessesOperationOptions{}
}

func (o ListUserExperienceAnalyticsImpactingProcessesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsImpactingProcessesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsImpactingProcessesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsImpactingProcessesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsImpactingProcessesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsImpactingProcesses - Get userExperienceAnalyticsImpactingProcess from deviceManagement.
// User experience analytics impacting process
func (c UserExperienceAnalyticsImpactingProcessClient) ListUserExperienceAnalyticsImpactingProcesses(ctx context.Context, options ListUserExperienceAnalyticsImpactingProcessesOperationOptions) (result ListUserExperienceAnalyticsImpactingProcessesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsImpactingProcessesCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsImpactingProcess",
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
		Values *[]beta.UserExperienceAnalyticsImpactingProcess `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsImpactingProcessesComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsImpactingProcessClient) ListUserExperienceAnalyticsImpactingProcessesComplete(ctx context.Context, options ListUserExperienceAnalyticsImpactingProcessesOperationOptions) (ListUserExperienceAnalyticsImpactingProcessesCompleteResult, error) {
	return c.ListUserExperienceAnalyticsImpactingProcessesCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsImpactingProcessOperationPredicate{})
}

// ListUserExperienceAnalyticsImpactingProcessesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsImpactingProcessClient) ListUserExperienceAnalyticsImpactingProcessesCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsImpactingProcessesOperationOptions, predicate UserExperienceAnalyticsImpactingProcessOperationPredicate) (result ListUserExperienceAnalyticsImpactingProcessesCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsImpactingProcess, 0)

	resp, err := c.ListUserExperienceAnalyticsImpactingProcesses(ctx, options)
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

	result = ListUserExperienceAnalyticsImpactingProcessesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
