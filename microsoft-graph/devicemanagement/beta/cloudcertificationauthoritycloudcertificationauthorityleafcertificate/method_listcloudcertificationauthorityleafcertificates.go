package cloudcertificationauthoritycloudcertificationauthorityleafcertificate

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

type ListCloudCertificationAuthorityLeafCertificatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudCertificationAuthorityLeafCertificate
}

type ListCloudCertificationAuthorityLeafCertificatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudCertificationAuthorityLeafCertificate
}

type ListCloudCertificationAuthorityLeafCertificatesOperationOptions struct {
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

func DefaultListCloudCertificationAuthorityLeafCertificatesOperationOptions() ListCloudCertificationAuthorityLeafCertificatesOperationOptions {
	return ListCloudCertificationAuthorityLeafCertificatesOperationOptions{}
}

func (o ListCloudCertificationAuthorityLeafCertificatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudCertificationAuthorityLeafCertificatesOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudCertificationAuthorityLeafCertificatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudCertificationAuthorityLeafCertificatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudCertificationAuthorityLeafCertificatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudCertificationAuthorityLeafCertificates - Get cloudCertificationAuthorityLeafCertificate from
// deviceManagement. Required OData property to expose leaf certificate API.
func (c CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient) ListCloudCertificationAuthorityLeafCertificates(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityLeafCertificatesOperationOptions) (result ListCloudCertificationAuthorityLeafCertificatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudCertificationAuthorityLeafCertificatesCustomPager{},
		Path:          fmt.Sprintf("%s/cloudCertificationAuthorityLeafCertificate", id.ID()),
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
		Values *[]beta.CloudCertificationAuthorityLeafCertificate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCloudCertificationAuthorityLeafCertificatesComplete retrieves all the results into a single object
func (c CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient) ListCloudCertificationAuthorityLeafCertificatesComplete(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityLeafCertificatesOperationOptions) (ListCloudCertificationAuthorityLeafCertificatesCompleteResult, error) {
	return c.ListCloudCertificationAuthorityLeafCertificatesCompleteMatchingPredicate(ctx, id, options, CloudCertificationAuthorityLeafCertificateOperationPredicate{})
}

// ListCloudCertificationAuthorityLeafCertificatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudCertificationAuthorityCloudCertificationAuthorityLeafCertificateClient) ListCloudCertificationAuthorityLeafCertificatesCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementCloudCertificationAuthorityId, options ListCloudCertificationAuthorityLeafCertificatesOperationOptions, predicate CloudCertificationAuthorityLeafCertificateOperationPredicate) (result ListCloudCertificationAuthorityLeafCertificatesCompleteResult, err error) {
	items := make([]beta.CloudCertificationAuthorityLeafCertificate, 0)

	resp, err := c.ListCloudCertificationAuthorityLeafCertificates(ctx, id, options)
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

	result = ListCloudCertificationAuthorityLeafCertificatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
