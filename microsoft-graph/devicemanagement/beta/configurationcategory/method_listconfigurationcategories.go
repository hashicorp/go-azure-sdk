package configurationcategory

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

type ListConfigurationCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationCategory
}

type ListConfigurationCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationCategory
}

type ListConfigurationCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationCategories ...
func (c ConfigurationCategoryClient) ListConfigurationCategories(ctx context.Context) (result ListConfigurationCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationCategoriesCustomPager{},
		Path:       "/deviceManagement/configurationCategories",
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
		Values *[]beta.DeviceManagementConfigurationCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationCategoriesComplete retrieves all the results into a single object
func (c ConfigurationCategoryClient) ListConfigurationCategoriesComplete(ctx context.Context) (ListConfigurationCategoriesCompleteResult, error) {
	return c.ListConfigurationCategoriesCompleteMatchingPredicate(ctx, DeviceManagementConfigurationCategoryOperationPredicate{})
}

// ListConfigurationCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationCategoryClient) ListConfigurationCategoriesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementConfigurationCategoryOperationPredicate) (result ListConfigurationCategoriesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationCategory, 0)

	resp, err := c.ListConfigurationCategories(ctx)
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

	result = ListConfigurationCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
