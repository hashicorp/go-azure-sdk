package directoryobject

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetAvailableExtensionPropertiesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultGetAvailableExtensionPropertiesOperationOptions() GetAvailableExtensionPropertiesOperationOptions {
	return GetAvailableExtensionPropertiesOperationOptions{}
}

func (o GetAvailableExtensionPropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAvailableExtensionPropertiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetAvailableExtensionPropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetAvailableExtensionProperties - Invoke action getAvailableExtensionProperties. Return all directory extension
// definitions that have been registered in a directory, including through multi-tenant apps. The following entities
// support extension properties:
func (c DirectoryObjectClient) GetAvailableExtensionProperties(ctx context.Context, input GetAvailableExtensionPropertiesRequest, options GetAvailableExtensionPropertiesOperationOptions) (result GetAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetAvailableExtensionPropertiesCustomPager{},
		Path:          "/directoryObjects/getAvailableExtensionProperties",
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
		Values *[]stable.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DirectoryObjectClient) GetAvailableExtensionPropertiesComplete(ctx context.Context, input GetAvailableExtensionPropertiesRequest, options GetAvailableExtensionPropertiesOperationOptions) (GetAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, options, ExtensionPropertyOperationPredicate{})
}

// GetAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryObjectClient) GetAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetAvailableExtensionPropertiesRequest, options GetAvailableExtensionPropertiesOperationOptions, predicate ExtensionPropertyOperationPredicate) (result GetAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetAvailableExtensionProperties(ctx, input, options)
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

	result = GetAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
