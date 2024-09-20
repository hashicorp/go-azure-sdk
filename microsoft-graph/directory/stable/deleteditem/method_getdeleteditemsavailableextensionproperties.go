package deleteditem

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

type GetDeletedItemsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetDeletedItemsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetDeletedItemsAvailableExtensionPropertiesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultGetDeletedItemsAvailableExtensionPropertiesOperationOptions() GetDeletedItemsAvailableExtensionPropertiesOperationOptions {
	return GetDeletedItemsAvailableExtensionPropertiesOperationOptions{}
}

func (o GetDeletedItemsAvailableExtensionPropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeletedItemsAvailableExtensionPropertiesOperationOptions) ToOData() *odata.Query {
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

func (o GetDeletedItemsAvailableExtensionPropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetDeletedItemsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDeletedItemsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDeletedItemsAvailableExtensionProperties - Invoke action getAvailableExtensionProperties. Return all directory
// extension definitions that have been registered in a directory, including through multi-tenant apps. The following
// entities support extension properties
func (c DeletedItemClient) GetDeletedItemsAvailableExtensionProperties(ctx context.Context, input GetDeletedItemsAvailableExtensionPropertiesRequest, options GetDeletedItemsAvailableExtensionPropertiesOperationOptions) (result GetDeletedItemsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetDeletedItemsAvailableExtensionPropertiesCustomPager{},
		Path:          "/directory/deletedItems/getAvailableExtensionProperties",
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

// GetDeletedItemsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDeletedItemsAvailableExtensionPropertiesComplete(ctx context.Context, input GetDeletedItemsAvailableExtensionPropertiesRequest, options GetDeletedItemsAvailableExtensionPropertiesOperationOptions) (GetDeletedItemsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, options, ExtensionPropertyOperationPredicate{})
}

// GetDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) GetDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetDeletedItemsAvailableExtensionPropertiesRequest, options GetDeletedItemsAvailableExtensionPropertiesOperationOptions, predicate ExtensionPropertyOperationPredicate) (result GetDeletedItemsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetDeletedItemsAvailableExtensionProperties(ctx, input, options)
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

	result = GetDeletedItemsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
