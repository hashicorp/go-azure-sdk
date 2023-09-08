package meinformationprotectionbitlockerrecoverykey

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

type ListMeInformationProtectionBitlockerRecoveryKeysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.BitlockerRecoveryKeyCollectionResponse
}

type ListMeInformationProtectionBitlockerRecoveryKeysCompleteResult struct {
	Items []models.BitlockerRecoveryKeyCollectionResponse
}

// ListMeInformationProtectionBitlockerRecoveryKeys ...
func (c MeInformationProtectionBitlockerRecoveryKeyClient) ListMeInformationProtectionBitlockerRecoveryKeys(ctx context.Context) (result ListMeInformationProtectionBitlockerRecoveryKeysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/informationProtection/bitlocker/recoveryKeys",
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
		Values *[]models.BitlockerRecoveryKeyCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeInformationProtectionBitlockerRecoveryKeysComplete retrieves all the results into a single object
func (c MeInformationProtectionBitlockerRecoveryKeyClient) ListMeInformationProtectionBitlockerRecoveryKeysComplete(ctx context.Context) (ListMeInformationProtectionBitlockerRecoveryKeysCompleteResult, error) {
	return c.ListMeInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate(ctx, models.BitlockerRecoveryKeyCollectionResponseOperationPredicate{})
}

// ListMeInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeInformationProtectionBitlockerRecoveryKeyClient) ListMeInformationProtectionBitlockerRecoveryKeysCompleteMatchingPredicate(ctx context.Context, predicate models.BitlockerRecoveryKeyCollectionResponseOperationPredicate) (result ListMeInformationProtectionBitlockerRecoveryKeysCompleteResult, err error) {
	items := make([]models.BitlockerRecoveryKeyCollectionResponse, 0)

	resp, err := c.ListMeInformationProtectionBitlockerRecoveryKeys(ctx)
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

	result = ListMeInformationProtectionBitlockerRecoveryKeysCompleteResult{
		Items: items,
	}
	return
}
