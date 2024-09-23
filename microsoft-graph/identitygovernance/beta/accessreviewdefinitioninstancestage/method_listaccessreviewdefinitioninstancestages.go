package accessreviewdefinitioninstancestage

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

type ListAccessReviewDefinitionInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewStage
}

type ListAccessReviewDefinitionInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewStage
}

type ListAccessReviewDefinitionInstanceStagesOperationOptions struct {
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

func DefaultListAccessReviewDefinitionInstanceStagesOperationOptions() ListAccessReviewDefinitionInstanceStagesOperationOptions {
	return ListAccessReviewDefinitionInstanceStagesOperationOptions{}
}

func (o ListAccessReviewDefinitionInstanceStagesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewDefinitionInstanceStagesOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewDefinitionInstanceStagesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewDefinitionInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewDefinitionInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewDefinitionInstanceStages - List stages. Retrieve the stages in a multi-stage access review instance.
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStages(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceStagesOperationOptions) (result ListAccessReviewDefinitionInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewDefinitionInstanceStagesCustomPager{},
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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

// ListAccessReviewDefinitionInstanceStagesComplete retrieves all the results into a single object
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStagesComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceStagesOperationOptions) (ListAccessReviewDefinitionInstanceStagesCompleteResult, error) {
	return c.ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate(ctx, id, options, AccessReviewStageOperationPredicate{})
}

// ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewDefinitionInstanceStageClient) ListAccessReviewDefinitionInstanceStagesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ListAccessReviewDefinitionInstanceStagesOperationOptions, predicate AccessReviewStageOperationPredicate) (result ListAccessReviewDefinitionInstanceStagesCompleteResult, err error) {
	items := make([]beta.AccessReviewStage, 0)

	resp, err := c.ListAccessReviewDefinitionInstanceStages(ctx, id, options)
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

	result = ListAccessReviewDefinitionInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
