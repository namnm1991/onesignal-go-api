/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PlayerNotificationTarget struct for PlayerNotificationTarget
type PlayerNotificationTarget struct {
	// Specific playerids to send your notification to. _Does not require API Auth Key. Do not combine with other targeting parameters. Not compatible with any other targeting parameters. Example: [\"1dd608f2-c6a1-11e3-851d-000c2940e62c\"] Limit of 2,000 entries per REST API call 
	IncludePlayerIds []string `json:"include_player_ids,omitempty"`
	// Target specific devices by custom user IDs assigned via API. Not compatible with any other targeting parameters Example: [\"custom-id-assigned-by-api\"] REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call. Note: If targeting push, email, or sms subscribers with same ids, use with channel_for_external_user_ids to indicate you are sending a push or email or sms. 
	IncludeExternalUserIds []string `json:"include_external_user_ids,omitempty"`
	// Recommended for Sending Emails - Target specific email addresses. If an email does not correspond to an existing user, a new user will be created. Example: nick@catfac.ts Limit of 2,000 entries per REST API call 
	IncludeEmailTokens []string `json:"include_email_tokens,omitempty"`
	// Recommended for Sending SMS - Target specific phone numbers. The phone number should be in the E.164 format. Phone number should be an existing subscriber on OneSignal. Refer our docs to learn how to add phone numbers to OneSignal. Example phone number: +1999999999 Limit of 2,000 entries per REST API call 
	IncludePhoneNumbers []string `json:"include_phone_numbers,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using iOS device tokens. Warning: Only works with Production tokens. All non-alphanumeric characters must be removed from each token. If a token does not correspond to an existing user, a new user will be created. Example: ce777617da7f548fe7a9ab6febb56cf39fba6d38203... Limit of 2,000 entries per REST API call 
	IncludeIosTokens []string `json:"include_ios_tokens,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Windows URIs. If a token does not correspond to an existing user, a new user will be created. Example: http://s.notify.live.net/u/1/bn1/HmQAAACPaLDr-... Limit of 2,000 entries per REST API call 
	IncludeWpWnsUris []string `json:"include_wp_wns_uris,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Amazon ADM registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: amzn1.adm-registration.v1.XpvSSUk0Rc3hTVVV... Limit of 2,000 entries per REST API call 
	IncludeAmazonRegIds []string `json:"include_amazon_reg_ids,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome App registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call 
	IncludeChromeRegIds []string `json:"include_chrome_reg_ids,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome Web Push registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call 
	IncludeChromeWebRegIds []string `json:"include_chrome_web_reg_ids,omitempty"`
	// Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Android device registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call 
	IncludeAndroidRegIds []string `json:"include_android_reg_ids,omitempty"`
}

// NewPlayerNotificationTarget instantiates a new PlayerNotificationTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerNotificationTarget() *PlayerNotificationTarget {
	this := PlayerNotificationTarget{}
	return &this
}

// NewPlayerNotificationTargetWithDefaults instantiates a new PlayerNotificationTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerNotificationTargetWithDefaults() *PlayerNotificationTarget {
	this := PlayerNotificationTarget{}
	return &this
}

// GetIncludePlayerIds returns the IncludePlayerIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludePlayerIds() []string {
	if o == nil || o.IncludePlayerIds == nil {
		var ret []string
		return ret
	}
	return o.IncludePlayerIds
}

// GetIncludePlayerIdsOk returns a tuple with the IncludePlayerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludePlayerIdsOk() ([]string, bool) {
	if o == nil || o.IncludePlayerIds == nil {
		return nil, false
	}
	return o.IncludePlayerIds, true
}

// HasIncludePlayerIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludePlayerIds() bool {
	if o != nil && o.IncludePlayerIds != nil {
		return true
	}

	return false
}

// SetIncludePlayerIds gets a reference to the given []string and assigns it to the IncludePlayerIds field.
func (o *PlayerNotificationTarget) SetIncludePlayerIds(v []string) {
	o.IncludePlayerIds = v
}

// GetIncludeExternalUserIds returns the IncludeExternalUserIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeExternalUserIds() []string {
	if o == nil || o.IncludeExternalUserIds == nil {
		var ret []string
		return ret
	}
	return o.IncludeExternalUserIds
}

// GetIncludeExternalUserIdsOk returns a tuple with the IncludeExternalUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeExternalUserIdsOk() ([]string, bool) {
	if o == nil || o.IncludeExternalUserIds == nil {
		return nil, false
	}
	return o.IncludeExternalUserIds, true
}

// HasIncludeExternalUserIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeExternalUserIds() bool {
	if o != nil && o.IncludeExternalUserIds != nil {
		return true
	}

	return false
}

// SetIncludeExternalUserIds gets a reference to the given []string and assigns it to the IncludeExternalUserIds field.
func (o *PlayerNotificationTarget) SetIncludeExternalUserIds(v []string) {
	o.IncludeExternalUserIds = v
}

// GetIncludeEmailTokens returns the IncludeEmailTokens field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeEmailTokens() []string {
	if o == nil || o.IncludeEmailTokens == nil {
		var ret []string
		return ret
	}
	return o.IncludeEmailTokens
}

// GetIncludeEmailTokensOk returns a tuple with the IncludeEmailTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeEmailTokensOk() ([]string, bool) {
	if o == nil || o.IncludeEmailTokens == nil {
		return nil, false
	}
	return o.IncludeEmailTokens, true
}

// HasIncludeEmailTokens returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeEmailTokens() bool {
	if o != nil && o.IncludeEmailTokens != nil {
		return true
	}

	return false
}

// SetIncludeEmailTokens gets a reference to the given []string and assigns it to the IncludeEmailTokens field.
func (o *PlayerNotificationTarget) SetIncludeEmailTokens(v []string) {
	o.IncludeEmailTokens = v
}

// GetIncludePhoneNumbers returns the IncludePhoneNumbers field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludePhoneNumbers() []string {
	if o == nil || o.IncludePhoneNumbers == nil {
		var ret []string
		return ret
	}
	return o.IncludePhoneNumbers
}

// GetIncludePhoneNumbersOk returns a tuple with the IncludePhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludePhoneNumbersOk() ([]string, bool) {
	if o == nil || o.IncludePhoneNumbers == nil {
		return nil, false
	}
	return o.IncludePhoneNumbers, true
}

// HasIncludePhoneNumbers returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludePhoneNumbers() bool {
	if o != nil && o.IncludePhoneNumbers != nil {
		return true
	}

	return false
}

// SetIncludePhoneNumbers gets a reference to the given []string and assigns it to the IncludePhoneNumbers field.
func (o *PlayerNotificationTarget) SetIncludePhoneNumbers(v []string) {
	o.IncludePhoneNumbers = v
}

// GetIncludeIosTokens returns the IncludeIosTokens field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeIosTokens() []string {
	if o == nil || o.IncludeIosTokens == nil {
		var ret []string
		return ret
	}
	return o.IncludeIosTokens
}

// GetIncludeIosTokensOk returns a tuple with the IncludeIosTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeIosTokensOk() ([]string, bool) {
	if o == nil || o.IncludeIosTokens == nil {
		return nil, false
	}
	return o.IncludeIosTokens, true
}

// HasIncludeIosTokens returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeIosTokens() bool {
	if o != nil && o.IncludeIosTokens != nil {
		return true
	}

	return false
}

// SetIncludeIosTokens gets a reference to the given []string and assigns it to the IncludeIosTokens field.
func (o *PlayerNotificationTarget) SetIncludeIosTokens(v []string) {
	o.IncludeIosTokens = v
}

// GetIncludeWpWnsUris returns the IncludeWpWnsUris field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeWpWnsUris() []string {
	if o == nil || o.IncludeWpWnsUris == nil {
		var ret []string
		return ret
	}
	return o.IncludeWpWnsUris
}

// GetIncludeWpWnsUrisOk returns a tuple with the IncludeWpWnsUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeWpWnsUrisOk() ([]string, bool) {
	if o == nil || o.IncludeWpWnsUris == nil {
		return nil, false
	}
	return o.IncludeWpWnsUris, true
}

// HasIncludeWpWnsUris returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeWpWnsUris() bool {
	if o != nil && o.IncludeWpWnsUris != nil {
		return true
	}

	return false
}

// SetIncludeWpWnsUris gets a reference to the given []string and assigns it to the IncludeWpWnsUris field.
func (o *PlayerNotificationTarget) SetIncludeWpWnsUris(v []string) {
	o.IncludeWpWnsUris = v
}

// GetIncludeAmazonRegIds returns the IncludeAmazonRegIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeAmazonRegIds() []string {
	if o == nil || o.IncludeAmazonRegIds == nil {
		var ret []string
		return ret
	}
	return o.IncludeAmazonRegIds
}

// GetIncludeAmazonRegIdsOk returns a tuple with the IncludeAmazonRegIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeAmazonRegIdsOk() ([]string, bool) {
	if o == nil || o.IncludeAmazonRegIds == nil {
		return nil, false
	}
	return o.IncludeAmazonRegIds, true
}

// HasIncludeAmazonRegIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeAmazonRegIds() bool {
	if o != nil && o.IncludeAmazonRegIds != nil {
		return true
	}

	return false
}

// SetIncludeAmazonRegIds gets a reference to the given []string and assigns it to the IncludeAmazonRegIds field.
func (o *PlayerNotificationTarget) SetIncludeAmazonRegIds(v []string) {
	o.IncludeAmazonRegIds = v
}

// GetIncludeChromeRegIds returns the IncludeChromeRegIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeChromeRegIds() []string {
	if o == nil || o.IncludeChromeRegIds == nil {
		var ret []string
		return ret
	}
	return o.IncludeChromeRegIds
}

// GetIncludeChromeRegIdsOk returns a tuple with the IncludeChromeRegIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeChromeRegIdsOk() ([]string, bool) {
	if o == nil || o.IncludeChromeRegIds == nil {
		return nil, false
	}
	return o.IncludeChromeRegIds, true
}

// HasIncludeChromeRegIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeChromeRegIds() bool {
	if o != nil && o.IncludeChromeRegIds != nil {
		return true
	}

	return false
}

// SetIncludeChromeRegIds gets a reference to the given []string and assigns it to the IncludeChromeRegIds field.
func (o *PlayerNotificationTarget) SetIncludeChromeRegIds(v []string) {
	o.IncludeChromeRegIds = v
}

// GetIncludeChromeWebRegIds returns the IncludeChromeWebRegIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeChromeWebRegIds() []string {
	if o == nil || o.IncludeChromeWebRegIds == nil {
		var ret []string
		return ret
	}
	return o.IncludeChromeWebRegIds
}

// GetIncludeChromeWebRegIdsOk returns a tuple with the IncludeChromeWebRegIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeChromeWebRegIdsOk() ([]string, bool) {
	if o == nil || o.IncludeChromeWebRegIds == nil {
		return nil, false
	}
	return o.IncludeChromeWebRegIds, true
}

// HasIncludeChromeWebRegIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeChromeWebRegIds() bool {
	if o != nil && o.IncludeChromeWebRegIds != nil {
		return true
	}

	return false
}

// SetIncludeChromeWebRegIds gets a reference to the given []string and assigns it to the IncludeChromeWebRegIds field.
func (o *PlayerNotificationTarget) SetIncludeChromeWebRegIds(v []string) {
	o.IncludeChromeWebRegIds = v
}

// GetIncludeAndroidRegIds returns the IncludeAndroidRegIds field value if set, zero value otherwise.
func (o *PlayerNotificationTarget) GetIncludeAndroidRegIds() []string {
	if o == nil || o.IncludeAndroidRegIds == nil {
		var ret []string
		return ret
	}
	return o.IncludeAndroidRegIds
}

// GetIncludeAndroidRegIdsOk returns a tuple with the IncludeAndroidRegIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerNotificationTarget) GetIncludeAndroidRegIdsOk() ([]string, bool) {
	if o == nil || o.IncludeAndroidRegIds == nil {
		return nil, false
	}
	return o.IncludeAndroidRegIds, true
}

// HasIncludeAndroidRegIds returns a boolean if a field has been set.
func (o *PlayerNotificationTarget) HasIncludeAndroidRegIds() bool {
	if o != nil && o.IncludeAndroidRegIds != nil {
		return true
	}

	return false
}

// SetIncludeAndroidRegIds gets a reference to the given []string and assigns it to the IncludeAndroidRegIds field.
func (o *PlayerNotificationTarget) SetIncludeAndroidRegIds(v []string) {
	o.IncludeAndroidRegIds = v
}

func (o PlayerNotificationTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludePlayerIds != nil {
		toSerialize["include_player_ids"] = o.IncludePlayerIds
	}
	if o.IncludeExternalUserIds != nil {
		toSerialize["include_external_user_ids"] = o.IncludeExternalUserIds
	}
	if o.IncludeEmailTokens != nil {
		toSerialize["include_email_tokens"] = o.IncludeEmailTokens
	}
	if o.IncludePhoneNumbers != nil {
		toSerialize["include_phone_numbers"] = o.IncludePhoneNumbers
	}
	if o.IncludeIosTokens != nil {
		toSerialize["include_ios_tokens"] = o.IncludeIosTokens
	}
	if o.IncludeWpWnsUris != nil {
		toSerialize["include_wp_wns_uris"] = o.IncludeWpWnsUris
	}
	if o.IncludeAmazonRegIds != nil {
		toSerialize["include_amazon_reg_ids"] = o.IncludeAmazonRegIds
	}
	if o.IncludeChromeRegIds != nil {
		toSerialize["include_chrome_reg_ids"] = o.IncludeChromeRegIds
	}
	if o.IncludeChromeWebRegIds != nil {
		toSerialize["include_chrome_web_reg_ids"] = o.IncludeChromeWebRegIds
	}
	if o.IncludeAndroidRegIds != nil {
		toSerialize["include_android_reg_ids"] = o.IncludeAndroidRegIds
	}
	return json.Marshal(toSerialize)
}

type NullablePlayerNotificationTarget struct {
	value *PlayerNotificationTarget
	isSet bool
}

func (v NullablePlayerNotificationTarget) Get() *PlayerNotificationTarget {
	return v.value
}

func (v *NullablePlayerNotificationTarget) Set(val *PlayerNotificationTarget) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerNotificationTarget) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerNotificationTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerNotificationTarget(val *PlayerNotificationTarget) *NullablePlayerNotificationTarget {
	return &NullablePlayerNotificationTarget{value: val, isSet: true}
}

func (v NullablePlayerNotificationTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerNotificationTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


