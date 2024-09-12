package appconsentrequestsforapproval

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

type ListAppConsentRequestsForApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppConsentRequest
}

type ListAppConsentRequestsForApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppConsentRequest
}

type ListAppConsentRequestsForApprovalsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListAppConsentRequestsForApprovalsOperationOptions() ListAppConsentRequestsForApprovalsOperationOptions {
	return ListAppConsentRequestsForApprovalsOperationOptions{}
}

func (o ListAppConsentRequestsForApprovalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppConsentRequestsForApprovalsOperationOptions) ToOData() *odata.Query {
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

func (o ListAppConsentRequestsForApprovalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppConsentRequestsForApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestsForApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestsForApprovals - Get appConsentRequestsForApproval from users
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovals(ctx context.Context, id beta.UserId, options ListAppConsentRequestsForApprovalsOperationOptions) (result ListAppConsentRequestsForApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppConsentRequestsForApprovalsCustomPager{},
		Path:          fmt.Sprintf("%s/appConsentRequestsForApproval", id.ID()),
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
		Values *[]beta.AppConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentRequestsForApprovalsComplete retrieves all the results into a single object
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovalsComplete(ctx context.Context, id beta.UserId, options ListAppConsentRequestsForApprovalsOperationOptions) (ListAppConsentRequestsForApprovalsCompleteResult, error) {
	return c.ListAppConsentRequestsForApprovalsCompleteMatchingPredicate(ctx, id, options, AppConsentRequestOperationPredicate{})
}

// ListAppConsentRequestsForApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovalsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options ListAppConsentRequestsForApprovalsOperationOptions, predicate AppConsentRequestOperationPredicate) (result ListAppConsentRequestsForApprovalsCompleteResult, err error) {
	items := make([]beta.AppConsentRequest, 0)

	resp, err := c.ListAppConsentRequestsForApprovals(ctx, id, options)
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

	result = ListAppConsentRequestsForApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
