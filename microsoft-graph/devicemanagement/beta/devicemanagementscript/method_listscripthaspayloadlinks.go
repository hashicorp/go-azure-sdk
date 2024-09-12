package devicemanagementscript

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

type ListScriptHasPayloadLinksOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.HasPayloadLinkResultItem
}

type ListScriptHasPayloadLinksCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.HasPayloadLinkResultItem
}

type ListScriptHasPayloadLinksOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListScriptHasPayloadLinksOperationOptions() ListScriptHasPayloadLinksOperationOptions {
	return ListScriptHasPayloadLinksOperationOptions{}
}

func (o ListScriptHasPayloadLinksOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListScriptHasPayloadLinksOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListScriptHasPayloadLinksOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListScriptHasPayloadLinksCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListScriptHasPayloadLinksCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListScriptHasPayloadLinks - Invoke action hasPayloadLinks
func (c DeviceManagementScriptClient) ListScriptHasPayloadLinks(ctx context.Context, input ListScriptHasPayloadLinksRequest, options ListScriptHasPayloadLinksOperationOptions) (result ListScriptHasPayloadLinksOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListScriptHasPayloadLinksCustomPager{},
		Path:          "/deviceManagement/deviceManagementScripts/hasPayloadLinks",
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

// ListScriptHasPayloadLinksComplete retrieves all the results into a single object
func (c DeviceManagementScriptClient) ListScriptHasPayloadLinksComplete(ctx context.Context, input ListScriptHasPayloadLinksRequest, options ListScriptHasPayloadLinksOperationOptions) (ListScriptHasPayloadLinksCompleteResult, error) {
	return c.ListScriptHasPayloadLinksCompleteMatchingPredicate(ctx, input, options, HasPayloadLinkResultItemOperationPredicate{})
}

// ListScriptHasPayloadLinksCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementScriptClient) ListScriptHasPayloadLinksCompleteMatchingPredicate(ctx context.Context, input ListScriptHasPayloadLinksRequest, options ListScriptHasPayloadLinksOperationOptions, predicate HasPayloadLinkResultItemOperationPredicate) (result ListScriptHasPayloadLinksCompleteResult, err error) {
	items := make([]beta.HasPayloadLinkResultItem, 0)

	resp, err := c.ListScriptHasPayloadLinks(ctx, input, options)
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

	result = ListScriptHasPayloadLinksCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
