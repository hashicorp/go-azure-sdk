package directoryprovisioning

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

type ListDirectoryProvisioningsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ProvisioningObjectSummary
}

type ListDirectoryProvisioningsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ProvisioningObjectSummary
}

type ListDirectoryProvisioningsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryProvisioningsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryProvisionings ...
func (c DirectoryProvisioningClient) ListDirectoryProvisionings(ctx context.Context) (result ListDirectoryProvisioningsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryProvisioningsCustomPager{},
		Path:       "/auditLogs/directoryProvisioning",
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
		Values *[]beta.ProvisioningObjectSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryProvisioningsComplete retrieves all the results into a single object
func (c DirectoryProvisioningClient) ListDirectoryProvisioningsComplete(ctx context.Context) (ListDirectoryProvisioningsCompleteResult, error) {
	return c.ListDirectoryProvisioningsCompleteMatchingPredicate(ctx, ProvisioningObjectSummaryOperationPredicate{})
}

// ListDirectoryProvisioningsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryProvisioningClient) ListDirectoryProvisioningsCompleteMatchingPredicate(ctx context.Context, predicate ProvisioningObjectSummaryOperationPredicate) (result ListDirectoryProvisioningsCompleteResult, err error) {
	items := make([]beta.ProvisioningObjectSummary, 0)

	resp, err := c.ListDirectoryProvisionings(ctx)
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

	result = ListDirectoryProvisioningsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
