// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MPL-2.0 License. See NOTICE.txt in the project root for license information.

package odata_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

func TestODataId(t *testing.T) {
	type testCase struct {
		input    odata.Id
		expected []byte
	}

	testCases := []testCase{
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/00000000-0000-0000-0000-000000000000/directoryObjects/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`https://graph.microsoft.com/v1.0/00000000-0000-0000-0000-000000000000/directoryObjects/11111111-1111-1111-1111-111111111111`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/00000000-0000-0000-0000-000000000000/directoryObjects/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`https://graph.microsoft.com/v2/00000000-0000-0000-0000-000000000000/directoryObjects/11111111-1111-1111-1111-111111111111`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`https://graph.microsoft.com/v2/directoryObjects/11111111-1111-1111-1111-111111111111`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/directoryObjects/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`directoryObjects('11111111-1111-1111-1111-111111111111')`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/users/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`users('11111111-1111-1111-1111-111111111111')`),
		},
		{
			expected: []byte(`"https://graph.microsoft.com/v1.0/policies/claimsMappingPolicies/11111111-1111-1111-1111-111111111111"`),
			input:    odata.Id(`policies/claimsMappingPolicies('11111111-1111-1111-1111-111111111111')`),
		},
	}

	for n, c := range testCases {
		id, err := json.Marshal(c.input)
		if err != nil {
			t.Errorf("test case %d: JSON marshaling failed: %v", n, err)
			continue
		}
		if !reflect.DeepEqual(id, c.expected) {
			t.Errorf("test case %d: expected %#v, got %#v", n, string(c.expected), string(id))
		}
	}
}

func TestOData(t *testing.T) {
	type testCase struct {
		response string
		expected odata.OData
	}

	testCases := []testCase{
		{
			response: `{
				"@odata.context": "https://graph.microsoft.com/beta/$metadata#servicePrincipals",
				"@odata.nextLink": "https://graph.microsoft.com/beta/1564a4be-0377-4d9b-8aff-5a2b564e177c/servicePrincipals?$skiptoken=X%274453707402000100000035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D34336331643937353963313035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D3433633164393735396331300000000000000000000000%27",
				"value": [
					{
						"id": "00000000-0000-0000-0000-000000000000",
						"deletedDateTime": null,
						"accountEnabled": true,
						"createdDateTime": "2020-07-08T01:22:29Z"
					}
				]
			}`,
			expected: odata.OData{
				Context:  pointer.To("https://graph.microsoft.com/beta/$metadata#servicePrincipals"),
				NextLink: pointer.To(odata.Link("https://graph.microsoft.com/beta/1564a4be-0377-4d9b-8aff-5a2b564e177c/servicePrincipals?%24skiptoken=X%274453707402000100000035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D34336331643937353963313035536572766963655072696E636970616C5F31326430653134382D663634382D343233382D383566312D3433633164393735396331300000000000000000000000%27")),
				Value: []interface{}{map[string]interface{}{
					"id":              "00000000-0000-0000-0000-000000000000",
					"deletedDateTime": nil,
					"accountEnabled":  true,
					"createdDateTime": "2020-07-08T01:22:29Z",
				}},
			},
		},
		{
			response: `{
				"@odata.context": "https://graph.microsoft.us/beta/$metadata#identityGovernance/accessReviews/definitions",
				"@odata.count": 4,
				"value": [
					{
						"id": "00000000-0000-0000-0000-000000000000",
						"displayName": "test",
						"createdDateTime": "2020-07-08T01:22:29",
						"lastModifiedDateTime": "2020-07-08T01:22:29",
						"status": "InProgress",
						"createdBy": {
							"id": "11111111-0000-0000-0000-000000000000",
							"displayName": "tester",
							"type": null,
							"userPrincipalName": "tester@contoso.us"
						}
					}
				]
			}`,
			expected: odata.OData{
				Context: pointer.To("https://graph.microsoft.us/beta/$metadata#identityGovernance/accessReviews/definitions"),
				Count:   pointer.To(4),
				Value: []interface{}{map[string]interface{}{
					"id":                   "00000000-0000-0000-0000-000000000000",
					"displayName":          "test",
					"createdDateTime":      "2020-07-08T01:22:29",
					"lastModifiedDateTime": "2020-07-08T01:22:29",
					"status":               "InProgress",
					"createdBy": map[string]interface{}{
						"id":                "11111111-0000-0000-0000-000000000000",
						"displayName":       "tester",
						"type":              nil,
						"userPrincipalName": "tester@contoso.us",
					},
				}},
			},
		},
		{
			response: `{
				"@odata.context": "https://graph.microsoft.com/v1.0/$metadata#directoryObjects/$entity",
				"@odata.type": "#microsoft.graph.servicePrincipal",
				"@odata.id": "https://graph.microsoft.com/v2/1564a4be-0377-4d9b-8aff-5a2b564e177c/directoryObjects/11111111-0000-0000-0000-000000000000/Microsoft.DirectoryServices.ServicePrincipal",
				"@odata.editLink": "https://graph.microsoft.com/v2/1564a4be-0377-4d9b-8aff-5a2b564e177c/directoryObjects/11111111-0000-0000-0000-000000000000/Microsoft.DirectoryServices.ServicePrincipal",
				"id": "11111111-0000-0000-0000-000000000000"
			}`,
			expected: odata.OData{
				Context:  pointer.To("https://graph.microsoft.com/v1.0/$metadata#directoryObjects/$entity"),
				Type:     pointer.To(odata.TypeServicePrincipal),
				Id:       (*odata.Id)(pointer.To("https://graph.microsoft.com/v1.0/1564a4be-0377-4d9b-8aff-5a2b564e177c/directoryObjects/11111111-0000-0000-0000-000000000000/Microsoft.DirectoryServices.ServicePrincipal")),
				EditLink: (*odata.Link)(pointer.To("https://graph.microsoft.com/v1.0/1564a4be-0377-4d9b-8aff-5a2b564e177c/directoryObjects/11111111-0000-0000-0000-000000000000/Microsoft.DirectoryServices.ServicePrincipal")),
			},
		},
	}

	for n, c := range testCases {
		var o odata.OData
		err := json.Unmarshal([]byte(c.response), &o)
		if err != nil {
			t.Errorf("test case %d: failed to unmarshal JSON: %v", n, err)
			continue
		}
		if !reflect.DeepEqual(o, c.expected) {
			t.Error(spew.Sprintf("test case %d:\nexpected:\n  %#v\ngot:\n  %#v\n", n, c.expected, o))
		}
	}
}

func TestError(t *testing.T) {
	type testCase struct {
		response string
		expected string
	}

	testCases := []testCase{
		{
			response: `{
  "error": {
    "code": "Service_ServiceUnavailable",
    "message": "Service is temporarily unavailable. Please wait and retry again.",
    "innerError": {
      "date": "2021-01-24T14:52:27",
      "request-id": "7c974c85-e572-43ff-9633-f2dddf28725a",
      "client-request-id": "7c974c85-e572-43ff-9633-f2dddf28725a"
    }
  }
}`,
			expected: "Service_ServiceUnavailable: Service is temporarily unavailable. Please wait and retry again.",
		},
		{
			response: `{
  "odata.error": {
    "code": "Request_InvalidDataContractVersion",
    "message": {
      "lang": "en",
      "value": "The specified api-version is invalid. The value must exactly match a supported version."
    },
    "requestId": "e3a05e86-92ae-4e7e-9635-a3f62342da5b",
    "date": "2021-01-24T15:37:05"
  }
}`,
			expected: "Request_InvalidDataContractVersion: The specified api-version is invalid. The value must exactly match a supported version.",
		},
		{
			response: `{
  "error": {
    "code": "BadRequest",
    "message": "The server could not process the request because it is malformed or incorrect.",
    "innerError": {
      "message": "1034: Policy contains invalid applications: {\"499b84ac-1321-427f-aa17-267ca6975798\":\"ServicePrincipalNotFound\"}",
      "date": "2021-06-23T21:54:16",
      "request-id": "4486d728-c654-4a30-bf71-bd5035f008a4",
      "client-request-id": "4486d728-c654-4a30-bf71-bd5035f008a4"
    }
  }
}`,
			expected: "BadRequest: The server could not process the request because it is malformed or incorrect.: 1034: Policy contains invalid applications: {\"499b84ac-1321-427f-aa17-267ca6975798\":\"ServicePrincipalNotFound\"}",
		},
		{
			response: `{
	"error": {
		"code": "-2134347772",
		"message": "SyncGroup '/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449' does not exist.",
		"target": "InvalidResource",
		"details": {
			"httpErrorCode": "NotFound",
			"exceptionType": "FilesCommon.Diagnostics.BackendException",
			"hashedMessage": "SyncGroup '/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449' does not exist.",
			"httpMethod": "GET",
			"requestUri": "http://10.0.0.7:9081/2020-03-01/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449?api-version=2020-03-01"
		},
		"innererror": {
			"callStack": "   at StorageFiles.Foundation.Diagnostics.Tracer.TraceAndThrow(Exception ex, EventLevel level, String msg, UseInterpolatedStringOrStringDotFormat guard, String file, String member, Int32 line)\r\n   at FilesCommon.Diagnostics.ErrorHelper.CheckAndThrowException(Boolean condition, ErrorCodeCommon errorCode, ExceptionMessage errorMessage, String errorTarget, Exception innerException)\r\n   at Kailani.Hfs.V1.Actions.Services.SyncGroupProvider.InternalGet(String resourceId, String name, StorageSyncServiceResource parentResource, Boolean throwIfNotFound) in C:\\__w\\1\\s\\src\\Kailani\\DotNetCore\\Kailani.Afs.Services\\Services\\DataSetProvider.cs:line 347\r\n   at Kailani.Hfs.V1.Actions.Services.SyncGroupProvider.<>c__DisplayClass10_0.<<DoGet>b__1>d.MoveNext() in C:\\__w\\1\\s\\src\\Kailani\\DotNetCore\\Kailani.Afs.Services\\Services\\DataSetProvider.cs:line 235\r\n--- End of stack trace from previous location ---\r\n   at Kailani.Hfs.V1.Actions.Services.BaseResourceProvider'1.ExecuteGet[TParent](T resource, Func'3 getParentAction, Func'3 getResourceAction, Boolean bThrowIfNotFound) in C:\\__w\\1\\s\\src\\Kailani\\DotNetCore\\Kailani.Afs.Services\\Services\\BaseResourceProvider.cs:line 349\r\n   at Kailani.Hfs.V1.Actions.Services.SyncGroupProvider.DoGet(SyncGroupResource resource, Boolean bThrowIfNotFound) in C:\\__w\\1\\s\\src\\Kailani\\DotNetCore\\Kailani.Afs.Services\\Services\\DataSetProvider.cs:line 237\r\n   at Kailani.Hfs.V1.Actions.CloudActions.GetSyncGroupAction.ExecuteAction(SyncGroupResource syncGroup, AbstractOperation op) in C:\\__w\\1\\s\\src\\Kailani\\DotNetCore\\Kailani.Hfs.CloudActions\\CloudActions\\SyncGroup\\GetSyncGroupAction.cs:line 95\r\n   at FilesCommon.ActionFramework.GetTemplateAction'1.ExecuteAsync(TResource inputResource, AbstractOperation op)\r\n   at FilesCommon.ActionFramework.BaseCloudAction'1.ExecuteRequestAsync(AbstractOperation op)\r\n   at FilesCommon.ActionFramework.BaseCloudAction.ExecuteAsync()",
			"message": "SyncGroup '/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449' does not exist."
		}
	}
}`,
			expected: "-2134347772: SyncGroup '/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449' does not exist.: SyncGroup '/subscriptions/1a6092a6-137e-4025-9a7c-ef77f76f2c02/resourceGroups/acctestRG-SS-230605162007979449/providers/Microsoft.StorageSync/storageSyncServices/acctest-StorageSync-230605162007979449/syncGroups/acctest-StorageSyncGroup-230605162007979449' does not exist.",
		},
	}

	for n, c := range testCases {
		var o odata.OData
		err := json.Unmarshal([]byte(c.response), &o)
		if err != nil {
			t.Errorf("test case %d: JSON unmarshaling failed: %v", n, err)
			continue
		}
		if o.Error == nil {
			t.Errorf("test case %d: Error field was nil", n)
			continue
		}
		if s := o.Error.String(); s != c.expected {
			t.Errorf("test case %d: expected %q, got %q", n, c.expected, s)
		}
	}
}
