package configmanagercollection

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

type ListConfigManagerCollectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ConfigManagerCollection
}

type ListConfigManagerCollectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ConfigManagerCollection
}

type ListConfigManagerCollectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigManagerCollectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigManagerCollections ...
func (c ConfigManagerCollectionClient) ListConfigManagerCollections(ctx context.Context) (result ListConfigManagerCollectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigManagerCollectionsCustomPager{},
		Path:       "/deviceManagement/configManagerCollections",
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
		Values *[]beta.ConfigManagerCollection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigManagerCollectionsComplete retrieves all the results into a single object
func (c ConfigManagerCollectionClient) ListConfigManagerCollectionsComplete(ctx context.Context) (ListConfigManagerCollectionsCompleteResult, error) {
	return c.ListConfigManagerCollectionsCompleteMatchingPredicate(ctx, ConfigManagerCollectionOperationPredicate{})
}

// ListConfigManagerCollectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigManagerCollectionClient) ListConfigManagerCollectionsCompleteMatchingPredicate(ctx context.Context, predicate ConfigManagerCollectionOperationPredicate) (result ListConfigManagerCollectionsCompleteResult, err error) {
	items := make([]beta.ConfigManagerCollection, 0)

	resp, err := c.ListConfigManagerCollections(ctx)
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

	result = ListConfigManagerCollectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
