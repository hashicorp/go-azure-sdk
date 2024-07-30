package sitelistitemdocumentsetversion

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

type ListSiteListItemDocumentSetVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DocumentSetVersion
}

type ListSiteListItemDocumentSetVersionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DocumentSetVersion
}

type ListSiteListItemDocumentSetVersionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListItemDocumentSetVersionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListItemDocumentSetVersions ...
func (c SiteListItemDocumentSetVersionClient) ListSiteListItemDocumentSetVersions(ctx context.Context, id GroupIdSiteIdListIdItemId) (result ListSiteListItemDocumentSetVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteListItemDocumentSetVersionsCustomPager{},
		Path:       fmt.Sprintf("%s/documentSetVersions", id.ID()),
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
		Values *[]stable.DocumentSetVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListItemDocumentSetVersionsComplete retrieves all the results into a single object
func (c SiteListItemDocumentSetVersionClient) ListSiteListItemDocumentSetVersionsComplete(ctx context.Context, id GroupIdSiteIdListIdItemId) (ListSiteListItemDocumentSetVersionsCompleteResult, error) {
	return c.ListSiteListItemDocumentSetVersionsCompleteMatchingPredicate(ctx, id, DocumentSetVersionOperationPredicate{})
}

// ListSiteListItemDocumentSetVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListItemDocumentSetVersionClient) ListSiteListItemDocumentSetVersionsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdListIdItemId, predicate DocumentSetVersionOperationPredicate) (result ListSiteListItemDocumentSetVersionsCompleteResult, err error) {
	items := make([]stable.DocumentSetVersion, 0)

	resp, err := c.ListSiteListItemDocumentSetVersions(ctx, id)
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

	result = ListSiteListItemDocumentSetVersionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
