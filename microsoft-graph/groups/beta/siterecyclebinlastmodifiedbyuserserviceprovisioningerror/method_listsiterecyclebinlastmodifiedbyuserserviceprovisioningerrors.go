package siterecyclebinlastmodifiedbyuserserviceprovisioningerror

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

type ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceProvisioningError
}

type ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceProvisioningError
}

type ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrors ...
func (c SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient) ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrors(ctx context.Context, id GroupIdSiteId) (result ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCustomPager{},
		Path:       fmt.Sprintf("%s/recycleBin/lastModifiedByUser/serviceProvisioningErrors", id.ID()),
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
		Values *[]beta.ServiceProvisioningError `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsComplete retrieves all the results into a single object
func (c SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient) ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsComplete(ctx context.Context, id GroupIdSiteId) (ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteResult, error) {
	return c.ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx, id, ServiceProvisioningErrorOperationPredicate{})
}

// ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient) ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteId, predicate ServiceProvisioningErrorOperationPredicate) (result ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteResult, err error) {
	items := make([]beta.ServiceProvisioningError, 0)

	resp, err := c.ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrors(ctx, id)
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

	result = ListSiteRecycleBinLastModifiedByUserServiceProvisioningErrorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
