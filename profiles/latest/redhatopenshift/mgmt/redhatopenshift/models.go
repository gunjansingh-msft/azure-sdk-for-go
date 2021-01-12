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

package redhatopenshift

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/redhatopenshift/mgmt/2020-04-30/redhatopenshift"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ProvisioningState = original.ProvisioningState

const (
	AdminUpdating ProvisioningState = original.AdminUpdating
	Creating      ProvisioningState = original.Creating
	Deleting      ProvisioningState = original.Deleting
	Failed        ProvisioningState = original.Failed
	Succeeded     ProvisioningState = original.Succeeded
	Updating      ProvisioningState = original.Updating
)

type VMSize = original.VMSize

const (
	StandardD2sV3 VMSize = original.StandardD2sV3
	StandardD4sV3 VMSize = original.StandardD4sV3
	StandardD8sV3 VMSize = original.StandardD8sV3
)

type VMSize1 = original.VMSize1

const (
	VMSize1StandardD2sV3 VMSize1 = original.VMSize1StandardD2sV3
	VMSize1StandardD4sV3 VMSize1 = original.VMSize1StandardD4sV3
	VMSize1StandardD8sV3 VMSize1 = original.VMSize1StandardD8sV3
)

type Visibility = original.Visibility

const (
	Private Visibility = original.Private
	Public  Visibility = original.Public
)

type Visibility1 = original.Visibility1

const (
	Visibility1Private Visibility1 = original.Visibility1Private
	Visibility1Public  Visibility1 = original.Visibility1Public
)

type APIServerProfile = original.APIServerProfile
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ClusterProfile = original.ClusterProfile
type ConsoleProfile = original.ConsoleProfile
type Display = original.Display
type IngressProfile = original.IngressProfile
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type OpenShiftCluster = original.OpenShiftCluster
type OpenShiftClusterCredentials = original.OpenShiftClusterCredentials
type OpenShiftClusterList = original.OpenShiftClusterList
type OpenShiftClusterListIterator = original.OpenShiftClusterListIterator
type OpenShiftClusterListPage = original.OpenShiftClusterListPage
type OpenShiftClusterProperties = original.OpenShiftClusterProperties
type OpenShiftClusterUpdate = original.OpenShiftClusterUpdate
type OpenShiftClustersClient = original.OpenShiftClustersClient
type OpenShiftClustersCreateOrUpdateFuture = original.OpenShiftClustersCreateOrUpdateFuture
type OpenShiftClustersDeleteFuture = original.OpenShiftClustersDeleteFuture
type OpenShiftClustersUpdateFuture = original.OpenShiftClustersUpdateFuture
type Operation = original.Operation
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServicePrincipalProfile = original.ServicePrincipalProfile
type TrackedResource = original.TrackedResource
type WorkerProfile = original.WorkerProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOpenShiftClusterListIterator(page OpenShiftClusterListPage) OpenShiftClusterListIterator {
	return original.NewOpenShiftClusterListIterator(page)
}
func NewOpenShiftClusterListPage(cur OpenShiftClusterList, getNextPage func(context.Context, OpenShiftClusterList) (OpenShiftClusterList, error)) OpenShiftClusterListPage {
	return original.NewOpenShiftClusterListPage(cur, getNextPage)
}
func NewOpenShiftClustersClient(subscriptionID string) OpenShiftClustersClient {
	return original.NewOpenShiftClustersClient(subscriptionID)
}
func NewOpenShiftClustersClientWithBaseURI(baseURI string, subscriptionID string) OpenShiftClustersClient {
	return original.NewOpenShiftClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleVMSize1Values() []VMSize1 {
	return original.PossibleVMSize1Values()
}
func PossibleVMSizeValues() []VMSize {
	return original.PossibleVMSizeValues()
}
func PossibleVisibility1Values() []Visibility1 {
	return original.PossibleVisibility1Values()
}
func PossibleVisibilityValues() []Visibility {
	return original.PossibleVisibilityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
