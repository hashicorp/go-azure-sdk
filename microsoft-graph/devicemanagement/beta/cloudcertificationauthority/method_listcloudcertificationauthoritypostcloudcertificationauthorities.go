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

type ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudCertificationAuthority
}

type ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudCertificationAuthority
}

type ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions() ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions {
	return ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions{}
}

func (o ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudCertificationAuthorityPostCloudCertificationAuthorities - Invoke action postCloudCertificationAuthority
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthorityPostCloudCertificationAuthorities(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions) (result ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCustomPager{},
		Path:          fmt.Sprintf("%s/postCloudCertificationAuthority", id.ID()),
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

// ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesComplete retrieves all the results into a single object
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesComplete(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions) (ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteResult, error) {
	return c.ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx, id, options, CloudCertificationAuthorityOperationPredicate{})
}

// ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudCertificationAuthorityClient) ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesOperationOptions, predicate CloudCertificationAuthorityOperationPredicate) (result ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteResult, err error) {
	items := make([]beta.CloudCertificationAuthority, 0)

	resp, err := c.ListCloudCertificationAuthorityPostCloudCertificationAuthorities(ctx, id, options)
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

	result = ListCloudCertificationAuthorityPostCloudCertificationAuthoritiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
