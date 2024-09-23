package accessreviewhistorydefinitioninstance

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

type ListAccessReviewHistoryDefinitionInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewHistoryInstance
}

type ListAccessReviewHistoryDefinitionInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewHistoryInstance
}

type ListAccessReviewHistoryDefinitionInstancesOperationOptions struct {
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

func DefaultListAccessReviewHistoryDefinitionInstancesOperationOptions() ListAccessReviewHistoryDefinitionInstancesOperationOptions {
	return ListAccessReviewHistoryDefinitionInstancesOperationOptions{}
}

func (o ListAccessReviewHistoryDefinitionInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAccessReviewHistoryDefinitionInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListAccessReviewHistoryDefinitionInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAccessReviewHistoryDefinitionInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAccessReviewHistoryDefinitionInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAccessReviewHistoryDefinitionInstances - List instances (of an accessReviewHistoryDefinition). Retrieve the
// instances of an access review history definition created in the last 30 days.
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstances(ctx context.Context, id beta.IdentityGovernanceAccessReviewHistoryDefinitionId, options ListAccessReviewHistoryDefinitionInstancesOperationOptions) (result ListAccessReviewHistoryDefinitionInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAccessReviewHistoryDefinitionInstancesCustomPager{},
		Path:          fmt.Sprintf("%s/instances", id.ID()),
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
		Values *[]beta.AccessReviewHistoryInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAccessReviewHistoryDefinitionInstancesComplete retrieves all the results into a single object
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstancesComplete(ctx context.Context, id beta.IdentityGovernanceAccessReviewHistoryDefinitionId, options ListAccessReviewHistoryDefinitionInstancesOperationOptions) (ListAccessReviewHistoryDefinitionInstancesCompleteResult, error) {
	return c.ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate(ctx, id, options, AccessReviewHistoryInstanceOperationPredicate{})
}

// ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AccessReviewHistoryDefinitionInstanceClient) ListAccessReviewHistoryDefinitionInstancesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceAccessReviewHistoryDefinitionId, options ListAccessReviewHistoryDefinitionInstancesOperationOptions, predicate AccessReviewHistoryInstanceOperationPredicate) (result ListAccessReviewHistoryDefinitionInstancesCompleteResult, err error) {
	items := make([]beta.AccessReviewHistoryInstance, 0)

	resp, err := c.ListAccessReviewHistoryDefinitionInstances(ctx, id, options)
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

	result = ListAccessReviewHistoryDefinitionInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
