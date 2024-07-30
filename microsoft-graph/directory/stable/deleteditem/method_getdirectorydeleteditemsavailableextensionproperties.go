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

type GetDirectoryDeletedItemsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetDirectoryDeletedItemsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDirectoryDeletedItemsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDirectoryDeletedItemsAvailableExtensionProperties ...
func (c DeletedItemClient) GetDirectoryDeletedItemsAvailableExtensionProperties(ctx context.Context, input GetDirectoryDeletedItemsAvailableExtensionPropertiesRequest) (result GetDirectoryDeletedItemsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetDirectoryDeletedItemsAvailableExtensionPropertiesCustomPager{},
		Path:       "/directory/deletedItems/getAvailableExtensionProperties",
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

// GetDirectoryDeletedItemsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDirectoryDeletedItemsAvailableExtensionPropertiesComplete(ctx context.Context, input GetDirectoryDeletedItemsAvailableExtensionPropertiesRequest) (GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, ExtensionPropertyOperationPredicate{})
}

// GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetDirectoryDeletedItemsAvailableExtensionPropertiesRequest, predicate ExtensionPropertyOperationPredicate) (result GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetDirectoryDeletedItemsAvailableExtensionProperties(ctx, input)
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

	result = GetDirectoryDeletedItemsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
