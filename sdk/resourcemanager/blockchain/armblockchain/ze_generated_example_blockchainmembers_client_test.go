//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Get.json
func ExampleBlockchainMembersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockchainMember.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Create.json
func ExampleBlockchainMembersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		&armblockchain.BlockchainMembersBeginCreateOptions{BlockchainMember: &armblockchain.BlockchainMember{
			TrackedResource: armblockchain.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armblockchain.BlockchainMemberProperties{
				Consortium:                          to.StringPtr("<consortium>"),
				ConsortiumManagementAccountPassword: to.StringPtr("<consortium-management-account-password>"),
				Password:                            to.StringPtr("<password>"),
				ValidatorNodesSKU: &armblockchain.BlockchainMemberNodesSKU{
					Capacity: to.Int32Ptr(2),
				},
				Protocol: armblockchain.BlockchainProtocolQuorum.ToPtr(),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockchainMember.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Delete.json
func ExampleBlockchainMembersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_Update.json
func ExampleBlockchainMembersClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		&armblockchain.BlockchainMembersUpdateOptions{BlockchainMember: &armblockchain.BlockchainMemberUpdate{
			Properties: &armblockchain.BlockchainMemberPropertiesUpdate{
				TransactionNodePropertiesUpdate: armblockchain.TransactionNodePropertiesUpdate{
					FirewallRules: []*armblockchain.FirewallRule{},
					Password:      to.StringPtr("<password>"),
				},
				ConsortiumManagementAccountPassword: to.StringPtr("<consortium-management-account-password>"),
			},
			Tags: map[string]*string{},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BlockchainMember.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_List.json
func ExampleBlockchainMembersClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BlockchainMember.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListAll.json
func ExampleBlockchainMembersClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	pager := client.ListAll(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BlockchainMember.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListConsortiumMembers.json
func ExampleBlockchainMembersClient_ListConsortiumMembers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	pager := client.ListConsortiumMembers("<blockchain-member-name>",
		"<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListApiKeys.json
func ExampleBlockchainMembersClient_ListAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	_, err = client.ListAPIKeys(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/BlockchainMembers_ListRegenerateApiKeys.json
func ExampleBlockchainMembersClient_ListRegenerateAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewBlockchainMembersClient("<subscription-id>", cred, nil)
	_, err = client.ListRegenerateAPIKeys(ctx,
		"<blockchain-member-name>",
		"<resource-group-name>",
		&armblockchain.BlockchainMembersListRegenerateAPIKeysOptions{APIKey: &armblockchain.APIKey{
			KeyName: to.StringPtr("<key-name>"),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
