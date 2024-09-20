package federationconfiguration

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

type ListFederationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InternalDomainFederation
}

type ListFederationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InternalDomainFederation
}

type ListFederationConfigurationsOperationOptions struct {
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

func DefaultListFederationConfigurationsOperationOptions() ListFederationConfigurationsOperationOptions {
	return ListFederationConfigurationsOperationOptions{}
}

func (o ListFederationConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListFederationConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListFederationConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListFederationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFederationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFederationConfigurations - List internalDomainFederations. Read the properties of the internalDomainFederation
// objects for the domain. This API returns only one object in the collection.
func (c FederationConfigurationClient) ListFederationConfigurations(ctx context.Context, id beta.DomainId, options ListFederationConfigurationsOperationOptions) (result ListFederationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListFederationConfigurationsCustomPager{},
		Path:          fmt.Sprintf("%s/federationConfiguration", id.ID()),
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
		Values *[]beta.InternalDomainFederation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListFederationConfigurationsComplete retrieves all the results into a single object
func (c FederationConfigurationClient) ListFederationConfigurationsComplete(ctx context.Context, id beta.DomainId, options ListFederationConfigurationsOperationOptions) (ListFederationConfigurationsCompleteResult, error) {
	return c.ListFederationConfigurationsCompleteMatchingPredicate(ctx, id, options, InternalDomainFederationOperationPredicate{})
}

// ListFederationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FederationConfigurationClient) ListFederationConfigurationsCompleteMatchingPredicate(ctx context.Context, id beta.DomainId, options ListFederationConfigurationsOperationOptions, predicate InternalDomainFederationOperationPredicate) (result ListFederationConfigurationsCompleteResult, err error) {
	items := make([]beta.InternalDomainFederation, 0)

	resp, err := c.ListFederationConfigurations(ctx, id, options)
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

	result = ListFederationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
