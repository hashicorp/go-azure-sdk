package termsofuseagreementfile

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

type ListTermsOfUseAgreementFilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AgreementFileLocalization
}

type ListTermsOfUseAgreementFilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AgreementFileLocalization
}

type ListTermsOfUseAgreementFilesOperationOptions struct {
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

func DefaultListTermsOfUseAgreementFilesOperationOptions() ListTermsOfUseAgreementFilesOperationOptions {
	return ListTermsOfUseAgreementFilesOperationOptions{}
}

func (o ListTermsOfUseAgreementFilesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTermsOfUseAgreementFilesOperationOptions) ToOData() *odata.Query {
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

func (o ListTermsOfUseAgreementFilesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTermsOfUseAgreementFilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTermsOfUseAgreementFilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTermsOfUseAgreementFiles - Get files from identityGovernance. PDFs linked to this agreement. This property is in
// the process of being deprecated. Use the file property instead. Supports $expand.
func (c TermsOfUseAgreementFileClient) ListTermsOfUseAgreementFiles(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFilesOperationOptions) (result ListTermsOfUseAgreementFilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTermsOfUseAgreementFilesCustomPager{},
		Path:          fmt.Sprintf("%s/files", id.ID()),
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
		Values *[]stable.AgreementFileLocalization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTermsOfUseAgreementFilesComplete retrieves all the results into a single object
func (c TermsOfUseAgreementFileClient) ListTermsOfUseAgreementFilesComplete(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFilesOperationOptions) (ListTermsOfUseAgreementFilesCompleteResult, error) {
	return c.ListTermsOfUseAgreementFilesCompleteMatchingPredicate(ctx, id, options, AgreementFileLocalizationOperationPredicate{})
}

// ListTermsOfUseAgreementFilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TermsOfUseAgreementFileClient) ListTermsOfUseAgreementFilesCompleteMatchingPredicate(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementId, options ListTermsOfUseAgreementFilesOperationOptions, predicate AgreementFileLocalizationOperationPredicate) (result ListTermsOfUseAgreementFilesCompleteResult, err error) {
	items := make([]stable.AgreementFileLocalization, 0)

	resp, err := c.ListTermsOfUseAgreementFiles(ctx, id, options)
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

	result = ListTermsOfUseAgreementFilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
