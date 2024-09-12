package deviceenrollmentconfiguration

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

type ListDeviceEnrollmentConfigurationHasPayloadLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HasPayloadLinkResultItem
}

type ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HasPayloadLinkResultItem
}

type ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions() ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions {
	return ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions{}
}

func (o ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceEnrollmentConfigurationHasPayloadLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceEnrollmentConfigurationHasPayloadLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceEnrollmentConfigurationHasPayloadLinks - Invoke action hasPayloadLinks
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationHasPayloadLinks(ctx context.Context, id beta.UserId, input ListDeviceEnrollmentConfigurationHasPayloadLinksRequest, options ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions) (result ListDeviceEnrollmentConfigurationHasPayloadLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDeviceEnrollmentConfigurationHasPayloadLinksCustomPager{},
		Path:          fmt.Sprintf("%s/deviceEnrollmentConfigurations/hasPayloadLinks", id.ID()),
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

// ListDeviceEnrollmentConfigurationHasPayloadLinksComplete retrieves all the results into a single object
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationHasPayloadLinksComplete(ctx context.Context, id beta.UserId, input ListDeviceEnrollmentConfigurationHasPayloadLinksRequest, options ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions) (ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteResult, error) {
	return c.ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteMatchingPredicate(ctx, id, input, options, HasPayloadLinkResultItemOperationPredicate{})
}

// ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceEnrollmentConfigurationClient) ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteMatchingPredicate(ctx context.Context, id beta.UserId, input ListDeviceEnrollmentConfigurationHasPayloadLinksRequest, options ListDeviceEnrollmentConfigurationHasPayloadLinksOperationOptions, predicate HasPayloadLinkResultItemOperationPredicate) (result ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteResult, err error) {
	items := make([]beta.HasPayloadLinkResultItem, 0)

	resp, err := c.ListDeviceEnrollmentConfigurationHasPayloadLinks(ctx, id, input, options)
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

	result = ListDeviceEnrollmentConfigurationHasPayloadLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
