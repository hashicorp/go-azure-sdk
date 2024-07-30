package devicecategory

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

type ListDeviceCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceCategory
}

type ListDeviceCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceCategory
}

type ListDeviceCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCategories ...
func (c DeviceCategoryClient) ListDeviceCategories(ctx context.Context) (result ListDeviceCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCategoriesCustomPager{},
		Path:       "/deviceManagement/deviceCategories",
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
		Values *[]stable.DeviceCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCategoriesComplete retrieves all the results into a single object
func (c DeviceCategoryClient) ListDeviceCategoriesComplete(ctx context.Context) (ListDeviceCategoriesCompleteResult, error) {
	return c.ListDeviceCategoriesCompleteMatchingPredicate(ctx, DeviceCategoryOperationPredicate{})
}

// ListDeviceCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCategoryClient) ListDeviceCategoriesCompleteMatchingPredicate(ctx context.Context, predicate DeviceCategoryOperationPredicate) (result ListDeviceCategoriesCompleteResult, err error) {
	items := make([]stable.DeviceCategory, 0)

	resp, err := c.ListDeviceCategories(ctx)
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

	result = ListDeviceCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
