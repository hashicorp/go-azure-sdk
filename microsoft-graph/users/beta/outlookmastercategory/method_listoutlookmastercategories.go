package outlookmastercategory

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

type ListOutlookMasterCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookCategory
}

type ListOutlookMasterCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookCategory
}

type ListOutlookMasterCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookMasterCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookMasterCategories ...
func (c OutlookMasterCategoryClient) ListOutlookMasterCategories(ctx context.Context, id UserId) (result ListOutlookMasterCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutlookMasterCategoriesCustomPager{},
		Path:       fmt.Sprintf("%s/outlook/masterCategories", id.ID()),
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
		Values *[]beta.OutlookCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookMasterCategoriesComplete retrieves all the results into a single object
func (c OutlookMasterCategoryClient) ListOutlookMasterCategoriesComplete(ctx context.Context, id UserId) (ListOutlookMasterCategoriesCompleteResult, error) {
	return c.ListOutlookMasterCategoriesCompleteMatchingPredicate(ctx, id, OutlookCategoryOperationPredicate{})
}

// ListOutlookMasterCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookMasterCategoryClient) ListOutlookMasterCategoriesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate OutlookCategoryOperationPredicate) (result ListOutlookMasterCategoriesCompleteResult, err error) {
	items := make([]beta.OutlookCategory, 0)

	resp, err := c.ListOutlookMasterCategories(ctx, id)
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

	result = ListOutlookMasterCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
