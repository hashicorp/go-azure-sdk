package userexperienceanalyticsmodelscore

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

type ListUserExperienceAnalyticsModelScoresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserExperienceAnalyticsModelScores
}

type ListUserExperienceAnalyticsModelScoresCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserExperienceAnalyticsModelScores
}

type ListUserExperienceAnalyticsModelScoresOperationOptions struct {
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

func DefaultListUserExperienceAnalyticsModelScoresOperationOptions() ListUserExperienceAnalyticsModelScoresOperationOptions {
	return ListUserExperienceAnalyticsModelScoresOperationOptions{}
}

func (o ListUserExperienceAnalyticsModelScoresOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserExperienceAnalyticsModelScoresOperationOptions) ToOData() *odata.Query {
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

func (o ListUserExperienceAnalyticsModelScoresOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserExperienceAnalyticsModelScoresCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserExperienceAnalyticsModelScoresCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserExperienceAnalyticsModelScores - Get userExperienceAnalyticsModelScores from deviceManagement. User
// experience analytics model scores
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScores(ctx context.Context, options ListUserExperienceAnalyticsModelScoresOperationOptions) (result ListUserExperienceAnalyticsModelScoresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserExperienceAnalyticsModelScoresCustomPager{},
		Path:          "/deviceManagement/userExperienceAnalyticsModelScores",
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
		Values *[]beta.UserExperienceAnalyticsModelScores `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserExperienceAnalyticsModelScoresComplete retrieves all the results into a single object
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScoresComplete(ctx context.Context, options ListUserExperienceAnalyticsModelScoresOperationOptions) (ListUserExperienceAnalyticsModelScoresCompleteResult, error) {
	return c.ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate(ctx, options, UserExperienceAnalyticsModelScoresOperationPredicate{})
}

// ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserExperienceAnalyticsModelScoreClient) ListUserExperienceAnalyticsModelScoresCompleteMatchingPredicate(ctx context.Context, options ListUserExperienceAnalyticsModelScoresOperationOptions, predicate UserExperienceAnalyticsModelScoresOperationPredicate) (result ListUserExperienceAnalyticsModelScoresCompleteResult, err error) {
	items := make([]beta.UserExperienceAnalyticsModelScores, 0)

	resp, err := c.ListUserExperienceAnalyticsModelScores(ctx, options)
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

	result = ListUserExperienceAnalyticsModelScoresCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
