package certificateauthoritymutualtlsoauthconfiguration

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

type ListCertificateAuthorityMutualTlsOauthConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MutualTlsOauthConfiguration
}

type ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MutualTlsOauthConfiguration
}

type ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions struct {
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

func DefaultListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions() ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions {
	return ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions{}
}

func (o ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCertificateAuthorityMutualTlsOauthConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCertificateAuthorityMutualTlsOauthConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCertificateAuthorityMutualTlsOauthConfigurations - List mutualTlsOauthConfigurations. Get a list of the available
// mutualTlsOauthConfiguration resources.
func (c CertificateAuthorityMutualTlsOauthConfigurationClient) ListCertificateAuthorityMutualTlsOauthConfigurations(ctx context.Context, options ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions) (result ListCertificateAuthorityMutualTlsOauthConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCertificateAuthorityMutualTlsOauthConfigurationsCustomPager{},
		Path:          "/directory/certificateAuthorities/mutualTlsOauthConfigurations",
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
		Values *[]beta.MutualTlsOauthConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCertificateAuthorityMutualTlsOauthConfigurationsComplete retrieves all the results into a single object
func (c CertificateAuthorityMutualTlsOauthConfigurationClient) ListCertificateAuthorityMutualTlsOauthConfigurationsComplete(ctx context.Context, options ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions) (ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteResult, error) {
	return c.ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteMatchingPredicate(ctx, options, MutualTlsOauthConfigurationOperationPredicate{})
}

// ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateAuthorityMutualTlsOauthConfigurationClient) ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListCertificateAuthorityMutualTlsOauthConfigurationsOperationOptions, predicate MutualTlsOauthConfigurationOperationPredicate) (result ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteResult, err error) {
	items := make([]beta.MutualTlsOauthConfiguration, 0)

	resp, err := c.ListCertificateAuthorityMutualTlsOauthConfigurations(ctx, options)
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

	result = ListCertificateAuthorityMutualTlsOauthConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
