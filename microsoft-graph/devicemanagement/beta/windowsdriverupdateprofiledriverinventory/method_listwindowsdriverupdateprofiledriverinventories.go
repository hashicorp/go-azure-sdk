package windowsdriverupdateprofiledriverinventory

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

type ListWindowsDriverUpdateProfileDriverInventoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsDriverUpdateInventory
}

type ListWindowsDriverUpdateProfileDriverInventoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsDriverUpdateInventory
}

type ListWindowsDriverUpdateProfileDriverInventoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsDriverUpdateProfileDriverInventoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsDriverUpdateProfileDriverInventories ...
func (c WindowsDriverUpdateProfileDriverInventoryClient) ListWindowsDriverUpdateProfileDriverInventories(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId) (result ListWindowsDriverUpdateProfileDriverInventoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsDriverUpdateProfileDriverInventoriesCustomPager{},
		Path:       fmt.Sprintf("%s/driverInventories", id.ID()),
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
		Values *[]beta.WindowsDriverUpdateInventory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsDriverUpdateProfileDriverInventoriesComplete retrieves all the results into a single object
func (c WindowsDriverUpdateProfileDriverInventoryClient) ListWindowsDriverUpdateProfileDriverInventoriesComplete(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId) (ListWindowsDriverUpdateProfileDriverInventoriesCompleteResult, error) {
	return c.ListWindowsDriverUpdateProfileDriverInventoriesCompleteMatchingPredicate(ctx, id, WindowsDriverUpdateInventoryOperationPredicate{})
}

// ListWindowsDriverUpdateProfileDriverInventoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsDriverUpdateProfileDriverInventoryClient) ListWindowsDriverUpdateProfileDriverInventoriesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementWindowsDriverUpdateProfileId, predicate WindowsDriverUpdateInventoryOperationPredicate) (result ListWindowsDriverUpdateProfileDriverInventoriesCompleteResult, err error) {
	items := make([]beta.WindowsDriverUpdateInventory, 0)

	resp, err := c.ListWindowsDriverUpdateProfileDriverInventories(ctx, id)
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

	result = ListWindowsDriverUpdateProfileDriverInventoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
