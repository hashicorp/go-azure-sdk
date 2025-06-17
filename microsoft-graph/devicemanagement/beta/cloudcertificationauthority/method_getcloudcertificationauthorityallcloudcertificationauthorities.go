package cloudcertificationauthority

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

type GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudCertificationAuthority
}

type GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudCertificationAuthority
}

type GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions() GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions {
	return GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions{}
}

func (o GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions) ToOData() *odata.Query {
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

func (o GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetCloudCertificationAuthorityAllCloudCertificationAuthorities - Invoke action getAllCloudCertificationAuthority
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllCloudCertificationAuthorities(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions) (result GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCustomPager{},
		Path:          fmt.Sprintf("%s/getAllCloudCertificationAuthority", id.ID()),
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

// GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesComplete retrieves all the results into a single object
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesComplete(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions) (GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteResult, error) {
	return c.GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx, id, options, CloudCertificationAuthorityOperationPredicate{})
}

// GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudCertificationAuthorityClient) GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesOperationOptions, predicate CloudCertificationAuthorityOperationPredicate) (result GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteResult, err error) {
	items := make([]beta.CloudCertificationAuthority, 0)

	resp, err := c.GetCloudCertificationAuthorityAllCloudCertificationAuthorities(ctx, id, options)
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

	result = GetCloudCertificationAuthorityAllCloudCertificationAuthoritiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
