package termsofuseagreement

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

type ListTermsOfUseAgreementsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Agreement
}

type ListTermsOfUseAgreementsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Agreement
}

type ListTermsOfUseAgreementsOperationOptions struct {
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

func DefaultListTermsOfUseAgreementsOperationOptions() ListTermsOfUseAgreementsOperationOptions {
	return ListTermsOfUseAgreementsOperationOptions{}
}

func (o ListTermsOfUseAgreementsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTermsOfUseAgreementsOperationOptions) ToOData() *odata.Query {
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

func (o ListTermsOfUseAgreementsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTermsOfUseAgreementsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreements - List agreements. Retrieve a list of agreement objects.
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreements(ctx context.Context, options ListTermsOfUseAgreementsOperationOptions) (result ListTermsOfUseAgreementsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTermsOfUseAgreementsCustomPager{},
		Path:          "/identityGovernance/termsOfUse/agreements",
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
		Values *[]beta.Agreement `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreementsComplete(ctx context.Context, options ListTermsOfUseAgreementsOperationOptions) (ListTermsOfUseAgreementsCompleteResult, error) {
	return c.ListTermsOfUseAgreementsCompleteMatchingPredicate(ctx, options, AgreementOperationPredicate{})
}

// ListTermsOfUseAgreementsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementClient) ListTermsOfUseAgreementsCompleteMatchingPredicate(ctx context.Context, options ListTermsOfUseAgreementsOperationOptions, predicate AgreementOperationPredicate) (result ListTermsOfUseAgreementsCompleteResult, err error) {
	items := make([]beta.Agreement, 0)

	resp, err := c.ListTermsOfUseAgreements(ctx, options)
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

	result = ListTermsOfUseAgreementsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
