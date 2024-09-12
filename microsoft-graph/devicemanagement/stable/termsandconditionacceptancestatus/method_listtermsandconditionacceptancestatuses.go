package termsandconditionacceptancestatus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTermsAndConditionAcceptanceStatusesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TermsAndConditionsAcceptanceStatus
}

type ListTermsAndConditionAcceptanceStatusesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TermsAndConditionsAcceptanceStatus
}

type ListTermsAndConditionAcceptanceStatusesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTermsAndConditionAcceptanceStatusesOperationOptions() ListTermsAndConditionAcceptanceStatusesOperationOptions {
	return ListTermsAndConditionAcceptanceStatusesOperationOptions{}
}

func (o ListTermsAndConditionAcceptanceStatusesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTermsAndConditionAcceptanceStatusesOperationOptions) ToOData() *odata.Query {
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

func (o ListTermsAndConditionAcceptanceStatusesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTermsAndConditionAcceptanceStatusesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsAndConditionAcceptanceStatusesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsAndConditionAcceptanceStatuses - List termsAndConditionsAcceptanceStatuses. List properties and
// relationships of the termsAndConditionsAcceptanceStatus objects.
func (c TermsAndConditionAcceptanceStatusClient) ListTermsAndConditionAcceptanceStatuses(ctx context.Context, id stable.DeviceManagementTermsAndConditionId, options ListTermsAndConditionAcceptanceStatusesOperationOptions) (result ListTermsAndConditionAcceptanceStatusesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTermsAndConditionAcceptanceStatusesCustomPager{},
		Path:          fmt.Sprintf("%s/acceptanceStatuses", id.ID()),
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
		Values *[]stable.TermsAndConditionsAcceptanceStatus `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsAndConditionAcceptanceStatusesComplete retrieves all the results into a single object
func (c TermsAndConditionAcceptanceStatusClient) ListTermsAndConditionAcceptanceStatusesComplete(ctx context.Context, id stable.DeviceManagementTermsAndConditionId, options ListTermsAndConditionAcceptanceStatusesOperationOptions) (ListTermsAndConditionAcceptanceStatusesCompleteResult, error) {
	return c.ListTermsAndConditionAcceptanceStatusesCompleteMatchingPredicate(ctx, id, options, TermsAndConditionsAcceptanceStatusOperationPredicate{})
}

// ListTermsAndConditionAcceptanceStatusesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsAndConditionAcceptanceStatusClient) ListTermsAndConditionAcceptanceStatusesCompleteMatchingPredicate(ctx context.Context, id stable.DeviceManagementTermsAndConditionId, options ListTermsAndConditionAcceptanceStatusesOperationOptions, predicate TermsAndConditionsAcceptanceStatusOperationPredicate) (result ListTermsAndConditionAcceptanceStatusesCompleteResult, err error) {
	items := make([]stable.TermsAndConditionsAcceptanceStatus, 0)

	resp, err := c.ListTermsAndConditionAcceptanceStatuses(ctx, id, options)
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

	result = ListTermsAndConditionAcceptanceStatusesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
