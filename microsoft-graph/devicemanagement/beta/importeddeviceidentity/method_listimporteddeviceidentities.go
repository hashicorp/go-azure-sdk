package importeddeviceidentity

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

type ListImportedDeviceIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedDeviceIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedDeviceIdentities ...
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentities(ctx context.Context) (result ListImportedDeviceIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListImportedDeviceIdentitiesCustomPager{},
		Path:       "/deviceManagement/importedDeviceIdentities",
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
		Values *[]beta.ImportedDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImportedDeviceIdentitiesComplete retrieves all the results into a single object
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitiesComplete(ctx context.Context) (ListImportedDeviceIdentitiesCompleteResult, error) {
	return c.ListImportedDeviceIdentitiesCompleteMatchingPredicate(ctx, ImportedDeviceIdentityOperationPredicate{})
}

// ListImportedDeviceIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitiesCompleteMatchingPredicate(ctx context.Context, predicate ImportedDeviceIdentityOperationPredicate) (result ListImportedDeviceIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedDeviceIdentity, 0)

	resp, err := c.ListImportedDeviceIdentities(ctx)
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

	result = ListImportedDeviceIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
