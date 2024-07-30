package sitepermission

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

type ListSitePermissionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Permission
}

type ListSitePermissionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Permission
}

type ListSitePermissionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePermissionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePermissions ...
func (c SitePermissionClient) ListSitePermissions(ctx context.Context, id GroupIdSiteId) (result ListSitePermissionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSitePermissionsCustomPager{},
		Path:       fmt.Sprintf("%s/permissions", id.ID()),
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
		Values *[]stable.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePermissionsComplete retrieves all the results into a single object
func (c SitePermissionClient) ListSitePermissionsComplete(ctx context.Context, id GroupIdSiteId) (ListSitePermissionsCompleteResult, error) {
	return c.ListSitePermissionsCompleteMatchingPredicate(ctx, id, PermissionOperationPredicate{})
}

// ListSitePermissionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePermissionClient) ListSitePermissionsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate PermissionOperationPredicate) (result ListSitePermissionsCompleteResult, err error) {
	items := make([]stable.Permission, 0)

	resp, err := c.ListSitePermissions(ctx, id)
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

	result = ListSitePermissionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
