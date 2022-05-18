/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	CsvFileUrl *string `json:"csv_file_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineResponse2008 InlineResponse2008

// NewInlineResponse2008 instantiates a new InlineResponse2008 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2008() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// NewInlineResponse2008WithDefaults instantiates a new InlineResponse2008 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2008WithDefaults() *InlineResponse2008 {
	this := InlineResponse2008{}
	return &this
}

// GetCsvFileUrl returns the CsvFileUrl field value if set, zero value otherwise.
func (o *InlineResponse2008) GetCsvFileUrl() string {
	if o == nil || o.CsvFileUrl == nil {
		var ret string
		return ret
	}
	return *o.CsvFileUrl
}

// GetCsvFileUrlOk returns a tuple with the CsvFileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2008) GetCsvFileUrlOk() (*string, bool) {
	if o == nil || o.CsvFileUrl == nil {
		return nil, false
	}
	return o.CsvFileUrl, true
}

// HasCsvFileUrl returns a boolean if a field has been set.
func (o *InlineResponse2008) HasCsvFileUrl() bool {
	if o != nil && o.CsvFileUrl != nil {
		return true
	}

	return false
}

// SetCsvFileUrl gets a reference to the given string and assigns it to the CsvFileUrl field.
func (o *InlineResponse2008) SetCsvFileUrl(v string) {
	o.CsvFileUrl = &v
}

func (o InlineResponse2008) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsvFileUrl != nil {
		toSerialize["csv_file_url"] = o.CsvFileUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineResponse2008) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2008 := _InlineResponse2008{}

	if err = json.Unmarshal(bytes, &varInlineResponse2008); err == nil {
		*o = InlineResponse2008(varInlineResponse2008)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "csv_file_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineResponse2008 struct {
	value *InlineResponse2008
	isSet bool
}

func (v NullableInlineResponse2008) Get() *InlineResponse2008 {
	return v.value
}

func (v *NullableInlineResponse2008) Set(val *InlineResponse2008) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2008) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2008) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2008(val *InlineResponse2008) *NullableInlineResponse2008 {
	return &NullableInlineResponse2008{value: val, isSet: true}
}

func (v NullableInlineResponse2008) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2008) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


