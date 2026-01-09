package deviceconfiguration

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

type ListDeviceConfigurationHasPayloadLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HasPayloadLinkResultItem
}

type ListDeviceConfigurationHasPayloadLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HasPayloadLinkResultItem
}

type ListDeviceConfigurationHasPayloadLinksOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListDeviceConfigurationHasPayloadLinksOperationOptions() ListDeviceConfigurationHasPayloadLinksOperationOptions {
	return ListDeviceConfigurationHasPayloadLinksOperationOptions{}
}

func (o ListDeviceConfigurationHasPayloadLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDeviceConfigurationHasPayloadLinksOperationOptions) ToOData() *odata.Query {
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

func (o ListDeviceConfigurationHasPayloadLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDeviceConfigurationHasPayloadLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationHasPayloadLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationHasPayloadLinks - Invoke action hasPayloadLinks
func (c DeviceConfigurationClient) ListDeviceConfigurationHasPayloadLinks(ctx context.Context, input ListDeviceConfigurationHasPayloadLinksRequest, options ListDeviceConfigurationHasPayloadLinksOperationOptions) (result ListDeviceConfigurationHasPayloadLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListDeviceConfigurationHasPayloadLinksCustomPager{},
		Path:          "/deviceManagement/deviceConfigurations/hasPayloadLinks",
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
		Values *[]beta.HasPayloadLinkResultItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationHasPayloadLinksComplete retrieves all the results into a single object
func (c DeviceConfigurationClient) ListDeviceConfigurationHasPayloadLinksComplete(ctx context.Context, input ListDeviceConfigurationHasPayloadLinksRequest, options ListDeviceConfigurationHasPayloadLinksOperationOptions) (ListDeviceConfigurationHasPayloadLinksCompleteResult, error) {
	return c.ListDeviceConfigurationHasPayloadLinksCompleteMatchingPredicate(ctx, input, options, HasPayloadLinkResultItemOperationPredicate{})
}

// ListDeviceConfigurationHasPayloadLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationClient) ListDeviceConfigurationHasPayloadLinksCompleteMatchingPredicate(ctx context.Context, input ListDeviceConfigurationHasPayloadLinksRequest, options ListDeviceConfigurationHasPayloadLinksOperationOptions, predicate HasPayloadLinkResultItemOperationPredicate) (result ListDeviceConfigurationHasPayloadLinksCompleteResult, err error) {
	items := make([]beta.HasPayloadLinkResultItem, 0)

	resp, err := c.ListDeviceConfigurationHasPayloadLinks(ctx, input, options)
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

	result = ListDeviceConfigurationHasPayloadLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
