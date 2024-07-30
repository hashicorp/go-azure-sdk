package devicemanagementpartner

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

type ListDeviceManagementPartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DeviceManagementPartner
}

type ListDeviceManagementPartnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DeviceManagementPartner
}

type ListDeviceManagementPartnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementPartnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementPartners ...
func (c DeviceManagementPartnerClient) ListDeviceManagementPartners(ctx context.Context) (result ListDeviceManagementPartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementPartnersCustomPager{},
		Path:       "/deviceManagement/deviceManagementPartners",
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
		Values *[]stable.DeviceManagementPartner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceManagementPartnersComplete retrieves all the results into a single object
func (c DeviceManagementPartnerClient) ListDeviceManagementPartnersComplete(ctx context.Context) (ListDeviceManagementPartnersCompleteResult, error) {
	return c.ListDeviceManagementPartnersCompleteMatchingPredicate(ctx, DeviceManagementPartnerOperationPredicate{})
}

// ListDeviceManagementPartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementPartnerClient) ListDeviceManagementPartnersCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementPartnerOperationPredicate) (result ListDeviceManagementPartnersCompleteResult, err error) {
	items := make([]stable.DeviceManagementPartner, 0)

	resp, err := c.ListDeviceManagementPartners(ctx)
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

	result = ListDeviceManagementPartnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
