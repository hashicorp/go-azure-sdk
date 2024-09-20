package certificateauthoritycertificatebasedapplicationconfiguration

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

type ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CertificateBasedApplicationConfiguration
}

type ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CertificateBasedApplicationConfiguration
}

type ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions struct {
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

func DefaultListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions() ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions {
	return ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions{}
}

func (o ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCertificateAuthorityCertificateBasedApplicationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCertificateAuthorityCertificateBasedApplicationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCertificateAuthorityCertificateBasedApplicationConfigurations - List certificateBasedApplicationConfigurations.
// Get a list of certificateBasedApplicationConfiguration objects.
func (c CertificateAuthorityCertificateBasedApplicationConfigurationClient) ListCertificateAuthorityCertificateBasedApplicationConfigurations(ctx context.Context, options ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions) (result ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCertificateAuthorityCertificateBasedApplicationConfigurationsCustomPager{},
		Path:          "/directory/certificateAuthorities/certificateBasedApplicationConfigurations",
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
		Values *[]beta.CertificateBasedApplicationConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCertificateAuthorityCertificateBasedApplicationConfigurationsComplete retrieves all the results into a single object
func (c CertificateAuthorityCertificateBasedApplicationConfigurationClient) ListCertificateAuthorityCertificateBasedApplicationConfigurationsComplete(ctx context.Context, options ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions) (ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult, error) {
	return c.ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate(ctx, options, CertificateBasedApplicationConfigurationOperationPredicate{})
}

// ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CertificateAuthorityCertificateBasedApplicationConfigurationClient) ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteMatchingPredicate(ctx context.Context, options ListCertificateAuthorityCertificateBasedApplicationConfigurationsOperationOptions, predicate CertificateBasedApplicationConfigurationOperationPredicate) (result ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult, err error) {
	items := make([]beta.CertificateBasedApplicationConfiguration, 0)

	resp, err := c.ListCertificateAuthorityCertificateBasedApplicationConfigurations(ctx, options)
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

	result = ListCertificateAuthorityCertificateBasedApplicationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
