package termsofuseagreementfileversion

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

type ListTermsOfUseAgreementFileVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AgreementFileVersion
}

type ListTermsOfUseAgreementFileVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AgreementFileVersion
}

type ListTermsOfUseAgreementFileVersionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListTermsOfUseAgreementFileVersionsOperationOptions() ListTermsOfUseAgreementFileVersionsOperationOptions {
	return ListTermsOfUseAgreementFileVersionsOperationOptions{}
}

func (o ListTermsOfUseAgreementFileVersionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTermsOfUseAgreementFileVersionsOperationOptions) ToOData() *odata.Query {
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

func (o ListTermsOfUseAgreementFileVersionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTermsOfUseAgreementFileVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFileVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFileVersions - Get versions from identityGovernance. Read-only. Customized versions of the
// terms of use agreement in the Microsoft Entra tenant.
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersions(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementIdFileId, options ListTermsOfUseAgreementFileVersionsOperationOptions) (result ListTermsOfUseAgreementFileVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTermsOfUseAgreementFileVersionsCustomPager{},
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]beta.AgreementFileVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFileVersionsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersionsComplete(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementIdFileId, options ListTermsOfUseAgreementFileVersionsOperationOptions) (ListTermsOfUseAgreementFileVersionsCompleteResult, error) {
	return c.ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate(ctx, id, options, AgreementFileVersionOperationPredicate{})
}

// ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileVersionClient) ListTermsOfUseAgreementFileVersionsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementIdFileId, options ListTermsOfUseAgreementFileVersionsOperationOptions, predicate AgreementFileVersionOperationPredicate) (result ListTermsOfUseAgreementFileVersionsCompleteResult, err error) {
	items := make([]beta.AgreementFileVersion, 0)

	resp, err := c.ListTermsOfUseAgreementFileVersions(ctx, id, options)
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

	result = ListTermsOfUseAgreementFileVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
