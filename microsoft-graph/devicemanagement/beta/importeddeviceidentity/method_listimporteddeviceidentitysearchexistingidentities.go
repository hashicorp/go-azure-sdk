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

type ListImportedDeviceIdentitySearchExistingIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedDeviceIdentity
}

type ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedDeviceIdentitySearchExistingIdentities ...
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentities(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest) (result ListImportedDeviceIdentitySearchExistingIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListImportedDeviceIdentitySearchExistingIdentitiesCustomPager{},
		Path:       "/deviceManagement/importedDeviceIdentities/searchExistingIdentities",
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

// ListImportedDeviceIdentitySearchExistingIdentitiesComplete retrieves all the results into a single object
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentitiesComplete(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest) (ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult, error) {
	return c.ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate(ctx, input, ImportedDeviceIdentityOperationPredicate{})
}

// ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedDeviceIdentityClient) ListImportedDeviceIdentitySearchExistingIdentitiesCompleteMatchingPredicate(ctx context.Context, input ListImportedDeviceIdentitySearchExistingIdentitiesRequest, predicate ImportedDeviceIdentityOperationPredicate) (result ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedDeviceIdentity, 0)

	resp, err := c.ListImportedDeviceIdentitySearchExistingIdentities(ctx, input)
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

	result = ListImportedDeviceIdentitySearchExistingIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
