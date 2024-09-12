package devicecompliancepolicy

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

type ListDeviceCompliancePolicyHasPayloadLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HasPayloadLinkResultItem
}

type ListDeviceCompliancePolicyHasPayloadLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HasPayloadLinkResultItem
}

type ListDeviceCompliancePolicyHasPayloadLinksOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListDeviceCompliancePolicyHasPayloadLinksOperationOptions() ListDeviceCompliancePolicyHasPayloadLinksOperationOptions {
	return ListDeviceCompliancePolicyHasPayloadLinksOperationOptions{}
}

func (o ListDeviceCompliancePolicyHasPayloadLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceCompliancePolicyHasPayloadLinksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDeviceCompliancePolicyHasPayloadLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceCompliancePolicyHasPayloadLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCompliancePolicyHasPayloadLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCompliancePolicyHasPayloadLinks - Invoke action hasPayloadLinks
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePolicyHasPayloadLinks(ctx context.Context, input ListDeviceCompliancePolicyHasPayloadLinksRequest, options ListDeviceCompliancePolicyHasPayloadLinksOperationOptions) (result ListDeviceCompliancePolicyHasPayloadLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDeviceCompliancePolicyHasPayloadLinksCustomPager{},
		Path:          "/deviceManagement/deviceCompliancePolicies/hasPayloadLinks",
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
		Values *[]beta.HasPayloadLinkResultItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCompliancePolicyHasPayloadLinksComplete retrieves all the results into a single object
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePolicyHasPayloadLinksComplete(ctx context.Context, input ListDeviceCompliancePolicyHasPayloadLinksRequest, options ListDeviceCompliancePolicyHasPayloadLinksOperationOptions) (ListDeviceCompliancePolicyHasPayloadLinksCompleteResult, error) {
	return c.ListDeviceCompliancePolicyHasPayloadLinksCompleteMatchingPredicate(ctx, input, options, HasPayloadLinkResultItemOperationPredicate{})
}

// ListDeviceCompliancePolicyHasPayloadLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCompliancePolicyClient) ListDeviceCompliancePolicyHasPayloadLinksCompleteMatchingPredicate(ctx context.Context, input ListDeviceCompliancePolicyHasPayloadLinksRequest, options ListDeviceCompliancePolicyHasPayloadLinksOperationOptions, predicate HasPayloadLinkResultItemOperationPredicate) (result ListDeviceCompliancePolicyHasPayloadLinksCompleteResult, err error) {
	items := make([]beta.HasPayloadLinkResultItem, 0)

	resp, err := c.ListDeviceCompliancePolicyHasPayloadLinks(ctx, input, options)
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

	result = ListDeviceCompliancePolicyHasPayloadLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
