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

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedAppleDeviceIdentityResult
}

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedAppleDeviceIdentityResult
}

type ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityLists ...
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityLists(ctx context.Context, id DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsRequest) (result ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCustomPager{},
		Path:       fmt.Sprintf("%s/importedAppleDeviceIdentities/importAppleDeviceIdentityList", id.ID()),
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
		Values *[]beta.ImportedAppleDeviceIdentityResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsComplete retrieves all the results into a single object
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsComplete(ctx context.Context, id DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsRequest) (ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteResult, error) {
	return c.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteMatchingPredicate(ctx, id, input, ImportedAppleDeviceIdentityResultOperationPredicate{})
}

// ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DepOnboardingSettingImportedAppleDeviceIdentityClient) ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementDepOnboardingSettingId, input ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsRequest, predicate ImportedAppleDeviceIdentityResultOperationPredicate) (result ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteResult, err error) {
	items := make([]beta.ImportedAppleDeviceIdentityResult, 0)

	resp, err := c.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityLists(ctx, id, input)
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

	result = ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
