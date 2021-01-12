// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package databricks

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/databricks/mgmt/2018-04-01/databricks"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CustomParameterType = original.CustomParameterType

const (
	Bool   CustomParameterType = original.Bool
	Object CustomParameterType = original.Object
	String CustomParameterType = original.String
)

type KeySource = original.KeySource

const (
	Default           KeySource = original.Default
	MicrosoftKeyvault KeySource = original.MicrosoftKeyvault
)

type PeeringProvisioningState = original.PeeringProvisioningState

const (
	Deleting  PeeringProvisioningState = original.Deleting
	Failed    PeeringProvisioningState = original.Failed
	Succeeded PeeringProvisioningState = original.Succeeded
	Updating  PeeringProvisioningState = original.Updating
)

type PeeringState = original.PeeringState

const (
	Connected    PeeringState = original.Connected
	Disconnected PeeringState = original.Disconnected
	Initiated    PeeringState = original.Initiated
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted  ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateCreated   ProvisioningState = original.ProvisioningStateCreated
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted   ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateReady     ProvisioningState = original.ProvisioningStateReady
	ProvisioningStateRunning   ProvisioningState = original.ProvisioningStateRunning
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating  ProvisioningState = original.ProvisioningStateUpdating
)

type AddressSpace = original.AddressSpace
type BaseClient = original.BaseClient
type CreatedBy = original.CreatedBy
type Encryption = original.Encryption
type ErrorDetail = original.ErrorDetail
type ErrorInfo = original.ErrorInfo
type ErrorResponse = original.ErrorResponse
type ManagedIdentityConfiguration = original.ManagedIdentityConfiguration
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type VNetPeeringClient = original.VNetPeeringClient
type VNetPeeringCreateOrUpdateFuture = original.VNetPeeringCreateOrUpdateFuture
type VNetPeeringDeleteFuture = original.VNetPeeringDeleteFuture
type VirtualNetworkPeering = original.VirtualNetworkPeering
type VirtualNetworkPeeringList = original.VirtualNetworkPeeringList
type VirtualNetworkPeeringListIterator = original.VirtualNetworkPeeringListIterator
type VirtualNetworkPeeringListPage = original.VirtualNetworkPeeringListPage
type VirtualNetworkPeeringPropertiesFormat = original.VirtualNetworkPeeringPropertiesFormat
type VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork = original.VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork
type VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork = original.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork
type Workspace = original.Workspace
type WorkspaceCustomBooleanParameter = original.WorkspaceCustomBooleanParameter
type WorkspaceCustomObjectParameter = original.WorkspaceCustomObjectParameter
type WorkspaceCustomParameters = original.WorkspaceCustomParameters
type WorkspaceCustomStringParameter = original.WorkspaceCustomStringParameter
type WorkspaceEncryptionParameter = original.WorkspaceEncryptionParameter
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspaceProviderAuthorization = original.WorkspaceProviderAuthorization
type WorkspaceUpdate = original.WorkspaceUpdate
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateOrUpdateFuture = original.WorkspacesCreateOrUpdateFuture
type WorkspacesDeleteFuture = original.WorkspacesDeleteFuture
type WorkspacesUpdateFuture = original.WorkspacesUpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVNetPeeringClient(subscriptionID string) VNetPeeringClient {
	return original.NewVNetPeeringClient(subscriptionID)
}
func NewVNetPeeringClientWithBaseURI(baseURI string, subscriptionID string) VNetPeeringClient {
	return original.NewVNetPeeringClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkPeeringListIterator(page VirtualNetworkPeeringListPage) VirtualNetworkPeeringListIterator {
	return original.NewVirtualNetworkPeeringListIterator(page)
}
func NewVirtualNetworkPeeringListPage(cur VirtualNetworkPeeringList, getNextPage func(context.Context, VirtualNetworkPeeringList) (VirtualNetworkPeeringList, error)) VirtualNetworkPeeringListPage {
	return original.NewVirtualNetworkPeeringListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return original.NewWorkspaceListResultIterator(page)
}
func NewWorkspaceListResultPage(cur WorkspaceListResult, getNextPage func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)) WorkspaceListResultPage {
	return original.NewWorkspaceListResultPage(cur, getNextPage)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleCustomParameterTypeValues() []CustomParameterType {
	return original.PossibleCustomParameterTypeValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossiblePeeringProvisioningStateValues() []PeeringProvisioningState {
	return original.PossiblePeeringProvisioningStateValues()
}
func PossiblePeeringStateValues() []PeeringState {
	return original.PossiblePeeringStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
