package deponboardingsettingimportedappledeviceidentity

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

type ListDepOnboardingSettingImportedAppleDeviceIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedAppleDeviceIdentity
}

type ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedAppleDeviceIdentity
}

type ListDepOnboardingSettingImportedAppleDeviceIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingImportedAppleDeviceIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettingImportedAppleDeviceIdentities ...
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentities(ctx context.Context, id DeviceManagementDepOnboardingSettingId) (result ListDepOnboardingSettingImportedAppleDeviceIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDepOnboardingSettingImportedAppleDeviceIdentitiesCustomPager{},
		Path:       fmt.Sprintf("%s/importedAppleDeviceIdentities", id.ID()),
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
		Values *[]beta.ImportedAppleDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingImportedAppleDeviceIdentitiesComplete retrieves all the results into a single object
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentitiesComplete(ctx context.Context, id DeviceManagementDepOnboardingSettingId) (ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteResult, error) {
	return c.ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteMatchingPredicate(ctx, id, ImportedAppleDeviceIdentityOperationPredicate{})
}

// ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDepOnboardingSettingId, predicate ImportedAppleDeviceIdentityOperationPredicate) (result ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteResult, err error) {
	items := make([]beta.ImportedAppleDeviceIdentity, 0)

	resp, err := c.ListDepOnboardingSettingImportedAppleDeviceIdentities(ctx, id)
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

	result = ListDepOnboardingSettingImportedAppleDeviceIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
