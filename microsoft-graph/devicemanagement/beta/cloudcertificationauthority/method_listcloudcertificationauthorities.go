package cloudcertificationauthority

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

type ListCloudCertificationAuthoritiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudCertificationAuthority
}

type ListCloudCertificationAuthoritiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudCertificationAuthority
}

type ListCloudCertificationAuthoritiesOperationOptions struct {
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

func DefaultListCloudCertificationAuthoritiesOperationOptions() ListCloudCertificationAuthoritiesOperationOptions {
	return ListCloudCertificationAuthoritiesOperationOptions{}
}

func (o ListCloudCertificationAuthoritiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudCertificationAuthoritiesOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudCertificationAuthoritiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudCertificationAuthoritiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudCertificationAuthoritiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudCertificationAuthorities - Get cloudCertificationAuthority from deviceManagement. Collection of
// CloudCertificationAuthority records associated with account.
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthorities(ctx context.Context, options ListCloudCertificationAuthoritiesOperationOptions) (result ListCloudCertificationAuthoritiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudCertificationAuthoritiesCustomPager{},
		Path:          "/deviceManagement/cloudCertificationAuthority",
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
		Values *[]beta.CloudCertificationAuthority `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudCertificationAuthoritiesComplete retrieves all the results into a single object
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthoritiesComplete(ctx context.Context, options ListCloudCertificationAuthoritiesOperationOptions) (ListCloudCertificationAuthoritiesCompleteResult, error) {
	return c.ListCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx, options, CloudCertificationAuthorityOperationPredicate{})
}

// ListCloudCertificationAuthoritiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx context.Context, options ListCloudCertificationAuthoritiesOperationOptions, predicate CloudCertificationAuthorityOperationPredicate) (result ListCloudCertificationAuthoritiesCompleteResult, err error) {
	items := make([]beta.CloudCertificationAuthority, 0)

	resp, err := c.ListCloudCertificationAuthorities(ctx, options)
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

	result = ListCloudCertificationAuthoritiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
