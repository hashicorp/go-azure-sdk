package pendingaccessreviewinstancestage

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

type ListPendingAccessReviewInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceStagesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListPendingAccessReviewInstanceStagesOperationOptions() ListPendingAccessReviewInstanceStagesOperationOptions {
	return ListPendingAccessReviewInstanceStagesOperationOptions{}
}

func (o ListPendingAccessReviewInstanceStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstanceStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceStages - Get stages from me. If the instance has multiple stages, this returns the
// collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and
// settings of stages on a review instance are created based on the accessReviewStageSettings on the parent
// accessReviewScheduleDefinition.
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStages(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceStagesOperationOptions) (result ListPendingAccessReviewInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceStagesCustomPager{},
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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
		Values *[]beta.AccessReviewStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceStagesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStagesComplete(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceStagesOperationOptions) (ListPendingAccessReviewInstanceStagesCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate(ctx, id, options, AccessReviewStageOperationPredicate{})
}

// ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options ListPendingAccessReviewInstanceStagesOperationOptions, predicate AccessReviewStageOperationPredicate) (result ListPendingAccessReviewInstanceStagesCompleteResult, err error) {
	items := make([]beta.AccessReviewStage, 0)

	resp, err := c.ListPendingAccessReviewInstanceStages(ctx, id, options)
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

	result = ListPendingAccessReviewInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
