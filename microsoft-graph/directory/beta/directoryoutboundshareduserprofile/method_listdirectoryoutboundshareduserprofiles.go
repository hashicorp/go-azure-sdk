package directoryoutboundshareduserprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryOutboundSharedUserProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OutboundSharedUserProfileCollectionResponse
}

type ListDirectoryOutboundSharedUserProfilesCompleteResult struct {
	Items []models.OutboundSharedUserProfileCollectionResponse
}

// ListDirectoryOutboundSharedUserProfiles ...
func (c DirectoryOutboundSharedUserProfileClient) ListDirectoryOutboundSharedUserProfiles(ctx context.Context) (result ListDirectoryOutboundSharedUserProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/outboundSharedUserProfiles",
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
		Values *[]models.OutboundSharedUserProfileCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryOutboundSharedUserProfilesComplete retrieves all the results into a single object
func (c DirectoryOutboundSharedUserProfileClient) ListDirectoryOutboundSharedUserProfilesComplete(ctx context.Context) (ListDirectoryOutboundSharedUserProfilesCompleteResult, error) {
	return c.ListDirectoryOutboundSharedUserProfilesCompleteMatchingPredicate(ctx, models.OutboundSharedUserProfileCollectionResponseOperationPredicate{})
}

// ListDirectoryOutboundSharedUserProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryOutboundSharedUserProfileClient) ListDirectoryOutboundSharedUserProfilesCompleteMatchingPredicate(ctx context.Context, predicate models.OutboundSharedUserProfileCollectionResponseOperationPredicate) (result ListDirectoryOutboundSharedUserProfilesCompleteResult, err error) {
	items := make([]models.OutboundSharedUserProfileCollectionResponse, 0)

	resp, err := c.ListDirectoryOutboundSharedUserProfiles(ctx)
	if err != nil {
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

	result = ListDirectoryOutboundSharedUserProfilesCompleteResult{
		Items: items,
	}
	return
}
