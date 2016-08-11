package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen specifies the listen state for access rights.
	Listen AccessRights = "Listen"
	// Manage specifies the manage state for access rights.
	Manage AccessRights = "Manage"
	// Send specifies the send state for access rights.
	Send AccessRights = "Send"
)

// AvailabilityStatus enumerates the values for availability status.
type AvailabilityStatus string

const (
	// Available specifies the available state for availability status.
	Available AvailabilityStatus = "Available"
	// Limited specifies the limited state for availability status.
	Limited AvailabilityStatus = "Limited"
	// Renaming specifies the renaming state for availability status.
	Renaming AvailabilityStatus = "Renaming"
	// Restoring specifies the restoring state for availability status.
	Restoring AvailabilityStatus = "Restoring"
	// Unknown specifies the unknown state for availability status.
	Unknown AvailabilityStatus = "Unknown"
)

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// EntityStatusActive specifies the entity status active state for entity
	// status.
	EntityStatusActive EntityStatus = "Active"
	// EntityStatusCreating specifies the entity status creating state for
	// entity status.
	EntityStatusCreating EntityStatus = "Creating"
	// EntityStatusDeleting specifies the entity status deleting state for
	// entity status.
	EntityStatusDeleting EntityStatus = "Deleting"
	// EntityStatusDisabled specifies the entity status disabled state for
	// entity status.
	EntityStatusDisabled EntityStatus = "Disabled"
	// EntityStatusReceiveDisabled specifies the entity status receive
	// disabled state for entity status.
	EntityStatusReceiveDisabled EntityStatus = "ReceiveDisabled"
	// EntityStatusRenaming specifies the entity status renaming state for
	// entity status.
	EntityStatusRenaming EntityStatus = "Renaming"
	// EntityStatusRestoring specifies the entity status restoring state for
	// entity status.
	EntityStatusRestoring EntityStatus = "Restoring"
	// EntityStatusSendDisabled specifies the entity status send disabled
	// state for entity status.
	EntityStatusSendDisabled EntityStatus = "SendDisabled"
	// EntityStatusUnknown specifies the entity status unknown state for
	// entity status.
	EntityStatusUnknown EntityStatus = "Unknown"
)

// Kind enumerates the values for kind.
type Kind string

const (
	// Messaging specifies the messaging state for kind.
	Messaging Kind = "Messaging"
)

// NamespaceState enumerates the values for namespace state.
type NamespaceState string

const (
	// NamespaceStateActivating specifies the namespace state activating state
	// for namespace state.
	NamespaceStateActivating NamespaceState = "Activating"
	// NamespaceStateActive specifies the namespace state active state for
	// namespace state.
	NamespaceStateActive NamespaceState = "Active"
	// NamespaceStateCreated specifies the namespace state created state for
	// namespace state.
	NamespaceStateCreated NamespaceState = "Created"
	// NamespaceStateCreating specifies the namespace state creating state for
	// namespace state.
	NamespaceStateCreating NamespaceState = "Creating"
	// NamespaceStateDisabled specifies the namespace state disabled state for
	// namespace state.
	NamespaceStateDisabled NamespaceState = "Disabled"
	// NamespaceStateDisabling specifies the namespace state disabling state
	// for namespace state.
	NamespaceStateDisabling NamespaceState = "Disabling"
	// NamespaceStateEnabling specifies the namespace state enabling state for
	// namespace state.
	NamespaceStateEnabling NamespaceState = "Enabling"
	// NamespaceStateFailed specifies the namespace state failed state for
	// namespace state.
	NamespaceStateFailed NamespaceState = "Failed"
	// NamespaceStateRemoved specifies the namespace state removed state for
	// namespace state.
	NamespaceStateRemoved NamespaceState = "Removed"
	// NamespaceStateRemoving specifies the namespace state removing state for
	// namespace state.
	NamespaceStateRemoving NamespaceState = "Removing"
	// NamespaceStateSoftDeleted specifies the namespace state soft deleted
	// state for namespace state.
	NamespaceStateSoftDeleted NamespaceState = "SoftDeleted"
	// NamespaceStateSoftDeleting specifies the namespace state soft deleting
	// state for namespace state.
	NamespaceStateSoftDeleting NamespaceState = "SoftDeleting"
	// NamespaceStateUnknown specifies the namespace state unknown state for
	// namespace state.
	NamespaceStateUnknown NamespaceState = "Unknown"
)

// Policykey enumerates the values for policykey.
type Policykey string

const (
	// PrimaryKey specifies the primary key state for policykey.
	PrimaryKey Policykey = "PrimaryKey"
	// SecondayKey specifies the seconday key state for policykey.
	SecondayKey Policykey = "SecondayKey"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic specifies the basic state for sku name.
	Basic SkuName = "Basic"
	// Premium specifies the premium state for sku name.
	Premium SkuName = "Premium"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "Standard"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic specifies the sku tier basic state for sku tier.
	SkuTierBasic SkuTier = "Basic"
	// SkuTierPremium specifies the sku tier premium state for sku tier.
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard specifies the sku tier standard state for sku tier.
	SkuTierStandard SkuTier = "Standard"
)

// MessageCountDetails is message Count Details.
type MessageCountDetails struct {
	ActiveMessageCount             *int64 `json:"ActiveMessageCount,omitempty"`
	DeadLetterMessageCount         *int64 `json:"DeadLetterMessageCount,omitempty"`
	ScheduledMessageCount          *int64 `json:"ScheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int64 `json:"TransferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int64 `json:"TransferMessageCount,omitempty"`
}

// NamespaceCreateOrUpdateParameters is parameters supplied to the
// CreateOrUpdate Namespace operation.
type NamespaceCreateOrUpdateParameters struct {
	Location   *string              `json:"location,omitempty"`
	Sku        *Sku                 `json:"sku,omitempty"`
	Tags       *map[string]*string  `json:"tags,omitempty"`
	Kind       Kind                 `json:"kind,omitempty"`
	Properties *NamespaceProperties `json:"properties,omitempty"`
}

// NamespaceListResult is the response of the List Namespace operation.
type NamespaceListResult struct {
	autorest.Response `json:"-"`
	Value             *[]NamespaceResource `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// NamespaceListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client NamespaceListResult) NamespaceListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// NamespaceProperties is properties of the Namespace.
type NamespaceProperties struct {
	ProvisioningState  *string        `json:"provisioningState,omitempty"`
	Status             NamespaceState `json:"status,omitempty"`
	CreatedAt          *date.Time     `json:"createdAt,omitempty"`
	UpdatedAt          *date.Time     `json:"updatedAt,omitempty"`
	ServiceBusEndpoint *string        `json:"serviceBusEndpoint,omitempty"`
	CreateACSNamespace *bool          `json:"createACSNamespace,omitempty"`
	Enabled            *bool          `json:"enabled,omitempty"`
}

// NamespaceResource is description of a Namespace resource.
type NamespaceResource struct {
	autorest.Response `json:"-"`
	ID                *string              `json:"id,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Type              *string              `json:"type,omitempty"`
	Location          *string              `json:"location,omitempty"`
	Tags              *map[string]*string  `json:"tags,omitempty"`
	Kind              Kind                 `json:"kind,omitempty"`
	Sku               *Sku                 `json:"sku,omitempty"`
	Properties        *NamespaceProperties `json:"properties,omitempty"`
}

// QueueCreateOrUpdateParameters is parameters supplied to the CreateOrUpdate
// Queue operation.
type QueueCreateOrUpdateParameters struct {
	Name       *string          `json:"name,omitempty"`
	Location   *string          `json:"location,omitempty"`
	Properties *QueueProperties `json:"properties,omitempty"`
}

// QueueListResult is the response of the List Queues operation.
type QueueListResult struct {
	autorest.Response `json:"-"`
	Value             *[]QueueResource `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// QueueListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client QueueListResult) QueueListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// QueueProperties is
type QueueProperties struct {
	AccessedAt                             *date.Time           `json:"AccessedAt,omitempty"`
	AutoDeleteOnIdle                       *string              `json:"AutoDeleteOnIdle,omitempty"`
	AvailabilityStatus                     AvailabilityStatus   `json:"AvailabilityStatus ,omitempty"`
	CreatedAt                              *date.Time           `json:"CreatedAt,omitempty"`
	DefaultMessageTimeToLive               *string              `json:"DefaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow    *string              `json:"DuplicateDetectionHistoryTimeWindow ,omitempty"`
	EnableBatchedOperations                *bool                `json:"EnableBatchedOperations,omitempty"`
	EnableDeadLetteringOnMessageExpiration *bool                `json:"EnableDeadLetteringOnMessageExpiration,omitempty"`
	EnableExpress                          *bool                `json:"EnableExpress,omitempty"`
	EnablePartitioning                     *bool                `json:"EnablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo          *string              `json:"ForwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                              *string              `json:"ForwardTo,omitempty"`
	IsAnonymousAccessible                  *bool                `json:"IsAnonymousAccessible,omitempty"`
	LockDuration                           *string              `json:"LockDuration ,omitempty"`
	MaxDeliveryCount                       *int32               `json:"MaxDeliveryCount ,omitempty"`
	MaxSizeInMegabytes                     *int64               `json:"MaxSizeInMegabytes ,omitempty"`
	MessageCount                           *int64               `json:"MessageCount ,omitempty"`
	MessageCountDetails                    *MessageCountDetails `json:"MessageCountDetails,omitempty"`
	Path                                   *string              `json:"Path,omitempty"`
	RequiresDuplicateDetection             *bool                `json:"RequiresDuplicateDetection,omitempty"`
	RequiresSession                        *bool                `json:"RequiresSession,omitempty"`
	SizeInBytes                            *int64               `json:"SizeInBytes ,omitempty"`
	Status                                 EntityStatus         `json:"Status,omitempty"`
	SupportOrdering                        *bool                `json:"SupportOrdering,omitempty"`
	UpdatedAt                              *date.Time           `json:"UpdatedAt,omitempty"`
	UserMetadata                           *string              `json:"UserMetadata,omitempty"`
}

// QueueResource is description of queue Resource.
type QueueResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *QueueProperties    `json:"properties,omitempty"`
}

// RegenerateKeysParameters is parameters supplied to the Regenerate Auth Rule.
type RegenerateKeysParameters struct {
	Policykey Policykey `json:"Policykey,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceListKeys is namespace/ServiceBus Connection String
type ResourceListKeys struct {
	autorest.Response         `json:"-"`
	PrimaryConnectionString   *string `json:"primaryConnectionString,omitempty"`
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`
	PrimaryKey                *string `json:"primaryKey,omitempty"`
	SecondaryKey              *string `json:"secondaryKey,omitempty"`
	KeyName                   *string `json:"keyName,omitempty"`
}

// SharedAccessAuthorizationRuleCreateOrUpdateParameters is parameters
// supplied to the CreateOrUpdate  AuthorizationRules.
type SharedAccessAuthorizationRuleCreateOrUpdateParameters struct {
	Location   *string                                  `json:"location,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Properties *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// SharedAccessAuthorizationRuleListResult is the response of the List
// Namespace operation.
type SharedAccessAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SharedAccessAuthorizationRuleResource `json:"value,omitempty"`
	NextLink          *string                                  `json:"nextLink,omitempty"`
}

// SharedAccessAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SharedAccessAuthorizationRuleListResult) SharedAccessAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SharedAccessAuthorizationRuleProperties is sharedAccessAuthorizationRule
// properties.
type SharedAccessAuthorizationRuleProperties struct {
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// SharedAccessAuthorizationRuleResource is description of a Namespace
// AuthorizationRules.
type SharedAccessAuthorizationRuleResource struct {
	autorest.Response `json:"-"`
	ID                *string                                  `json:"id,omitempty"`
	Name              *string                                  `json:"name,omitempty"`
	Type              *string                                  `json:"type,omitempty"`
	Location          *string                                  `json:"location,omitempty"`
	Tags              *map[string]*string                      `json:"tags,omitempty"`
	Properties        *SharedAccessAuthorizationRuleProperties `json:"properties,omitempty"`
}

// Sku is sku of the Namespace.
type Sku struct {
	Name     SkuName `json:"name,omitempty"`
	Tier     SkuTier `json:"tier,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
}

// SubscriptionCreateOrUpdateParameters is parameters supplied to the
// CreateOrUpdate Subscription operation.
type SubscriptionCreateOrUpdateParameters struct {
	Location   *string                 `json:"location,omitempty"`
	Type       *string                 `json:"type,omitempty"`
	Properties *SubscriptionProperties `json:"properties,omitempty"`
}

// SubscriptionListResult is the response of the List Subscriptions operation.
type SubscriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SubscriptionResource `json:"value,omitempty"`
	NextLink          *string                 `json:"nextLink,omitempty"`
}

// SubscriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SubscriptionListResult) SubscriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SubscriptionProperties is description of Subscription Resource.
type SubscriptionProperties struct {
	AccessedAt                                      *date.Time           `json:"AccessedAt,omitempty"`
	AutoDeleteOnIdle                                *string              `json:"AutoDeleteOnIdle,omitempty"`
	AvailabilityStatus                              AvailabilityStatus   `json:"AvailabilityStatus ,omitempty"`
	CreatedAt                                       *date.Time           `json:"CreatedAt,omitempty"`
	DefaultMessageTimeToLive                        *string              `json:"DefaultMessageTimeToLive,omitempty"`
	EnableBatchedOperations                         *bool                `json:"EnableBatchedOperations,omitempty"`
	EnableDeadLetteringOnFilterEvaluationExceptions *bool                `json:"EnableDeadLetteringOnFilterEvaluationExceptions,omitempty"`
	EnableDeadLetteringOnMessageExpiration          *bool                `json:"EnableDeadLetteringOnMessageExpiration,omitempty"`
	ForwardDeadLetteredMessagesTo                   *string              `json:"ForwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                       *string              `json:"ForwardTo,omitempty"`
	IsReadOnly                                      *bool                `json:"IsReadOnly,omitempty"`
	LockDuration                                    *string              `json:"LockDuration,omitempty"`
	MaxDeliveryCount                                *int32               `json:"MaxDeliveryCount,omitempty"`
	MessageCount                                    *int64               `json:"MessageCount,omitempty"`
	MessageCountDetails                             *MessageCountDetails `json:"MessageCountDetails,omitempty"`
	RequiresSession                                 *bool                `json:"RequiresSession,omitempty"`
	Status                                          EntityStatus         `json:"Status,omitempty"`
	TopicPath                                       *string              `json:"TopicPath,omitempty"`
	UpdatedAt                                       *date.Time           `json:"UpdatedAt,omitempty"`
	UserMetadata                                    *string              `json:"UserMetadata,omitempty"`
}

// SubscriptionResource is description of Subscription Resource.
type SubscriptionResource struct {
	autorest.Response `json:"-"`
	ID                *string                 `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Type              *string                 `json:"type,omitempty"`
	Location          *string                 `json:"location,omitempty"`
	Tags              *map[string]*string     `json:"tags,omitempty"`
	Properties        *SubscriptionProperties `json:"properties,omitempty"`
}

// TopicCreateOrUpdateParameters is parameters supplied to the CreateOrUpdate
// Topic operation.
type TopicCreateOrUpdateParameters struct {
	Name       *string          `json:"name,omitempty"`
	Location   *string          `json:"location,omitempty"`
	Properties *TopicProperties `json:"properties,omitempty"`
}

// TopicListResult is the response of the List Topics operation.
type TopicListResult struct {
	autorest.Response `json:"-"`
	Value             *[]TopicResource `json:"value,omitempty"`
	NextLink          *string          `json:"nextLink,omitempty"`
}

// TopicListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client TopicListResult) TopicListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// TopicProperties is
type TopicProperties struct {
	AccessedAt                              *date.Time           `json:"AccessedAt,omitempty"`
	AutoDeleteOnIdle                        *string              `json:"AutoDeleteOnIdle,omitempty"`
	AvailabilityStatus                      AvailabilityStatus   `json:"AvailabilityStatus ,omitempty"`
	CreatedAt                               *date.Time           `json:"CreatedAt,omitempty"`
	DefaultMessageTimeToLive                *string              `json:"DefaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow     *string              `json:"DuplicateDetectionHistoryTimeWindow ,omitempty"`
	EnableBatchedOperations                 *bool                `json:"EnableBatchedOperations,omitempty"`
	EnableExpress                           *bool                `json:"EnableExpress,omitempty"`
	EnableFilteringMessagesBeforePublishing *bool                `json:"EnableFilteringMessagesBeforePublishing,omitempty"`
	EnablePartitioning                      *bool                `json:"EnablePartitioning,omitempty"`
	IsAnonymousAccessible                   *bool                `json:"IsAnonymousAccessible,omitempty"`
	MaxSizeInMegabytes                      *int64               `json:"MaxSizeInMegabytes ,omitempty"`
	MessageCountDetails                     *MessageCountDetails `json:"MessageCountDetails,omitempty"`
	Path                                    *string              `json:"Path,omitempty"`
	RequiresDuplicateDetection              *bool                `json:"RequiresDuplicateDetection,omitempty"`
	SizeInBytes                             *int64               `json:"SizeInBytes ,omitempty"`
	Status                                  EntityStatus         `json:"Status,omitempty"`
	SubscriptionCount                       *int32               `json:"SubscriptionCount,omitempty"`
	SupportOrdering                         *bool                `json:"SupportOrdering,omitempty"`
	UpdatedAt                               *date.Time           `json:"UpdatedAt,omitempty"`
	UserMetadata                            *string              `json:"UserMetadata,omitempty"`
}

// TopicResource is description of topic Resource.
type TopicResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *TopicProperties    `json:"properties,omitempty"`
}
