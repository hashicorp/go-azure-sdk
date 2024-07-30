package importedwindowsautopilotdeviceidentity

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

type ListImportedWindowsAutopilotDeviceIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedWindowsAutopilotDeviceIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedWindowsAutopilotDeviceIdentities ...
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentities(ctx context.Context) (result ListImportedWindowsAutopilotDeviceIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListImportedWindowsAutopilotDeviceIdentitiesCustomPager{},
		Path:       "/deviceManagement/importedWindowsAutopilotDeviceIdentities",
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
		Values *[]beta.ImportedWindowsAutopilotDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImportedWindowsAutopilotDeviceIdentitiesComplete retrieves all the results into a single object
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentitiesComplete(ctx context.Context) (ListImportedWindowsAutopilotDeviceIdentitiesCompleteResult, error) {
	return c.ListImportedWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate(ctx, ImportedWindowsAutopilotDeviceIdentityOperationPredicate{})
}

// ListImportedWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate(ctx context.Context, predicate ImportedWindowsAutopilotDeviceIdentityOperationPredicate) (result ListImportedWindowsAutopilotDeviceIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedWindowsAutopilotDeviceIdentity, 0)

	resp, err := c.ListImportedWindowsAutopilotDeviceIdentities(ctx)
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

	result = ListImportedWindowsAutopilotDeviceIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
