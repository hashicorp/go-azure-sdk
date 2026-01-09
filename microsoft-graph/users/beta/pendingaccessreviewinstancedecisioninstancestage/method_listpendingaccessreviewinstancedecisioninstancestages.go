package pendingaccessreviewinstancedecisioninstancestage

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

type ListPendingAccessReviewInstanceDecisionInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions struct {
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

func DefaultListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions() ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions {
	return ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions{}
}

func (o ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceDecisionInstanceStages - Get stages from users. If the instance has multiple stages,
// this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence,
// number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent
// accessReviewScheduleDefinition.
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStages(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions) (result ListPendingAccessReviewInstanceDecisionInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager{},
		Path:          fmt.Sprintf("%s/instance/stages", id.ID()),
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
		Values *[]beta.AccessReviewStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceDecisionInstanceStagesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStagesComplete(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions) (ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate(ctx, id, options, AccessReviewStageOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdDecisionId, options ListPendingAccessReviewInstanceDecisionInstanceStagesOperationOptions, predicate AccessReviewStageOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult, err error) {
	items := make([]beta.AccessReviewStage, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInstanceStages(ctx, id, options)
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

	result = ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
