package termsofuseagreementfilelocalization

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

type ListTermsOfUseAgreementFileLocalizationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AgreementFileLocalization
}

type ListTermsOfUseAgreementFileLocalizationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AgreementFileLocalization
}

type ListTermsOfUseAgreementFileLocalizationsOperationOptions struct {
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

func DefaultListTermsOfUseAgreementFileLocalizationsOperationOptions() ListTermsOfUseAgreementFileLocalizationsOperationOptions {
	return ListTermsOfUseAgreementFileLocalizationsOperationOptions{}
}

func (o ListTermsOfUseAgreementFileLocalizationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTermsOfUseAgreementFileLocalizationsOperationOptions) ToOData() *odata.Query {
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

func (o ListTermsOfUseAgreementFileLocalizationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTermsOfUseAgreementFileLocalizationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFileLocalizationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFileLocalizations - Get localizations from identityGovernance. The localized version of the
// terms of use agreement files attached to the agreement.
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizations(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFileLocalizationsOperationOptions) (result ListTermsOfUseAgreementFileLocalizationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTermsOfUseAgreementFileLocalizationsCustomPager{},
		Path:          fmt.Sprintf("%s/file/localizations", id.ID()),
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
		Values *[]beta.AgreementFileLocalization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFileLocalizationsComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizationsComplete(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFileLocalizationsOperationOptions) (ListTermsOfUseAgreementFileLocalizationsCompleteResult, error) {
	return c.ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate(ctx, id, options, AgreementFileLocalizationOperationPredicate{})
}

// ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileLocalizationClient) ListTermsOfUseAgreementFileLocalizationsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFileLocalizationsOperationOptions, predicate AgreementFileLocalizationOperationPredicate) (result ListTermsOfUseAgreementFileLocalizationsCompleteResult, err error) {
	items := make([]beta.AgreementFileLocalization, 0)

	resp, err := c.ListTermsOfUseAgreementFileLocalizations(ctx, id, options)
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

	result = ListTermsOfUseAgreementFileLocalizationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
