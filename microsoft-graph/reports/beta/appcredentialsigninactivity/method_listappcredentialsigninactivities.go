package appcredentialsigninactivity

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

type ListAppCredentialSignInActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppCredentialSignInActivity
}

type ListAppCredentialSignInActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppCredentialSignInActivity
}

type ListAppCredentialSignInActivitiesOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListAppCredentialSignInActivitiesOperationOptions() ListAppCredentialSignInActivitiesOperationOptions {
	return ListAppCredentialSignInActivitiesOperationOptions{}
}

func (o ListAppCredentialSignInActivitiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppCredentialSignInActivitiesOperationOptions) ToOData() *odata.Query {
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

func (o ListAppCredentialSignInActivitiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppCredentialSignInActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppCredentialSignInActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppCredentialSignInActivities - List appCredentialSignInActivities. Get a list of appCredentialSignInActivity
// objects that contains recent activity of application credentials.
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivities(ctx context.Context, options ListAppCredentialSignInActivitiesOperationOptions) (result ListAppCredentialSignInActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppCredentialSignInActivitiesCustomPager{},
		Path:          "/reports/appCredentialSignInActivities",
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
		Values *[]beta.AppCredentialSignInActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppCredentialSignInActivitiesComplete retrieves all the results into a single object
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivitiesComplete(ctx context.Context, options ListAppCredentialSignInActivitiesOperationOptions) (ListAppCredentialSignInActivitiesCompleteResult, error) {
	return c.ListAppCredentialSignInActivitiesCompleteMatchingPredicate(ctx, options, AppCredentialSignInActivityOperationPredicate{})
}

// ListAppCredentialSignInActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivitiesCompleteMatchingPredicate(ctx context.Context, options ListAppCredentialSignInActivitiesOperationOptions, predicate AppCredentialSignInActivityOperationPredicate) (result ListAppCredentialSignInActivitiesCompleteResult, err error) {
	items := make([]beta.AppCredentialSignInActivity, 0)

	resp, err := c.ListAppCredentialSignInActivities(ctx, options)
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

	result = ListAppCredentialSignInActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
