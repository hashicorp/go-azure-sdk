package profileeducationalactivity

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

type ListProfileEducationalActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.EducationalActivity
}

type ListProfileEducationalActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.EducationalActivity
}

type ListProfileEducationalActivitiesOperationOptions struct {
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

func DefaultListProfileEducationalActivitiesOperationOptions() ListProfileEducationalActivitiesOperationOptions {
	return ListProfileEducationalActivitiesOperationOptions{}
}

func (o ListProfileEducationalActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListProfileEducationalActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListProfileEducationalActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListProfileEducationalActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileEducationalActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileEducationalActivities - Get educationalActivities from users. Represents data that a user has supplied
// related to undergraduate, graduate, postgraduate or other educational activities.
func (c ProfileEducationalActivityClient) ListProfileEducationalActivities(ctx context.Context, id beta.UserId, options ListProfileEducationalActivitiesOperationOptions) (result ListProfileEducationalActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListProfileEducationalActivitiesCustomPager{},
		Path:          fmt.Sprintf("%s/profile/educationalActivities", id.ID()),
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
		Values *[]beta.EducationalActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileEducationalActivitiesComplete retrieves all the results into a single object
func (c ProfileEducationalActivityClient) ListProfileEducationalActivitiesComplete(ctx context.Context, id beta.UserId, options ListProfileEducationalActivitiesOperationOptions) (ListProfileEducationalActivitiesCompleteResult, error) {
	return c.ListProfileEducationalActivitiesCompleteMatchingPredicate(ctx, id, options, EducationalActivityOperationPredicate{})
}

// ListProfileEducationalActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileEducationalActivityClient) ListProfileEducationalActivitiesCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListProfileEducationalActivitiesOperationOptions, predicate EducationalActivityOperationPredicate) (result ListProfileEducationalActivitiesCompleteResult, err error) {
	items := make([]beta.EducationalActivity, 0)

	resp, err := c.ListProfileEducationalActivities(ctx, id, options)
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

	result = ListProfileEducationalActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
