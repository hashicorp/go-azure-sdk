package agreementacceptance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAgreementAcceptancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementAcceptance
}

type ListAgreementAcceptancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementAcceptance
}

type ListAgreementAcceptancesOperationOptions struct {
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

func DefaultListAgreementAcceptancesOperationOptions() ListAgreementAcceptancesOperationOptions {
	return ListAgreementAcceptancesOperationOptions{}
}

func (o ListAgreementAcceptancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAgreementAcceptancesOperationOptions) ToOData() *odata.Query {
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

func (o ListAgreementAcceptancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAgreementAcceptancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAgreementAcceptancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAgreementAcceptances - Get agreementAcceptances from users. The user's terms of use acceptance statuses.
// Read-only. Nullable.
func (c AgreementAcceptanceClient) ListAgreementAcceptances(ctx context.Context, id stable.UserId, options ListAgreementAcceptancesOperationOptions) (result ListAgreementAcceptancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAgreementAcceptancesCustomPager{},
		Path:          fmt.Sprintf("%s/agreementAcceptances", id.ID()),
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
		Values *[]stable.AgreementAcceptance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAgreementAcceptancesComplete retrieves all the results into a single object
func (c AgreementAcceptanceClient) ListAgreementAcceptancesComplete(ctx context.Context, id stable.UserId, options ListAgreementAcceptancesOperationOptions) (ListAgreementAcceptancesCompleteResult, error) {
	return c.ListAgreementAcceptancesCompleteMatchingPredicate(ctx, id, options, AgreementAcceptanceOperationPredicate{})
}

// ListAgreementAcceptancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AgreementAcceptanceClient) ListAgreementAcceptancesCompleteMatchingPredicate(ctx context.Context, id stable.UserId, options ListAgreementAcceptancesOperationOptions, predicate AgreementAcceptanceOperationPredicate) (result ListAgreementAcceptancesCompleteResult, err error) {
	items := make([]stable.AgreementAcceptance, 0)

	resp, err := c.ListAgreementAcceptances(ctx, id, options)
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

	result = ListAgreementAcceptancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
