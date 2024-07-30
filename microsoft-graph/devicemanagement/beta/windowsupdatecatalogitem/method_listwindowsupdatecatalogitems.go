package windowsupdatecatalogitem

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

type ListWindowsUpdateCatalogItemsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsUpdateCatalogItem
}

type ListWindowsUpdateCatalogItemsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsUpdateCatalogItem
}

type ListWindowsUpdateCatalogItemsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsUpdateCatalogItemsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsUpdateCatalogItems ...
func (c WindowsUpdateCatalogItemClient) ListWindowsUpdateCatalogItems(ctx context.Context) (result ListWindowsUpdateCatalogItemsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsUpdateCatalogItemsCustomPager{},
		Path:       "/deviceManagement/windowsUpdateCatalogItems",
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
		Values *[]beta.WindowsUpdateCatalogItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsUpdateCatalogItemsComplete retrieves all the results into a single object
func (c WindowsUpdateCatalogItemClient) ListWindowsUpdateCatalogItemsComplete(ctx context.Context) (ListWindowsUpdateCatalogItemsCompleteResult, error) {
	return c.ListWindowsUpdateCatalogItemsCompleteMatchingPredicate(ctx, WindowsUpdateCatalogItemOperationPredicate{})
}

// ListWindowsUpdateCatalogItemsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsUpdateCatalogItemClient) ListWindowsUpdateCatalogItemsCompleteMatchingPredicate(ctx context.Context, predicate WindowsUpdateCatalogItemOperationPredicate) (result ListWindowsUpdateCatalogItemsCompleteResult, err error) {
	items := make([]beta.WindowsUpdateCatalogItem, 0)

	resp, err := c.ListWindowsUpdateCatalogItems(ctx)
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

	result = ListWindowsUpdateCatalogItemsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
