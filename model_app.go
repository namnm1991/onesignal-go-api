/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.2
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
	"time"
)

// App struct for App
type App struct {
	Id string `json:"id"`
	// The name of your app, as displayed on your apps list on the dashboard.  This can be renamed.
	Name *string `json:"name,omitempty"`
	Players *int32 `json:"players,omitempty"`
	MessageablePlayers *int32 `json:"messageable_players,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Android: Your Google Project number.  Also known as Sender ID.
	AndroidGcmSenderId *string `json:"android_gcm_sender_id,omitempty"`
	// Android: Your Google Push Messaging Auth Key
	GcmKey *string `json:"gcm_key,omitempty"`
	// Chrome (All Browsers except Safari) (Recommended): The URL to your website.  This field is required if you wish to enable web push and specify other web push parameters.
	ChromeWebOrigin *string `json:"chrome_web_origin,omitempty"`
	// Not for web push.  Your Google Push Messaging Auth Key if you use Chrome Apps / Extensions.
	ChromeKey *string `json:"chrome_key,omitempty"`
	// Chrome (All Browsers except Safari): Your default notification icon. Should be 256x256 pixels, min 80x80.
	ChromeWebDefaultNotificationIcon *string `json:"chrome_web_default_notification_icon,omitempty"`
	// Chrome (All Browsers except Safari): A subdomain of your choice in order to support Web Push on non-HTTPS websites. This field must be set in order for the chrome_web_gcm_sender_id property to be processed.
	ChromeWebSubDomain *string `json:"chrome_web_sub_domain,omitempty"`
	// iOS: Either sandbox or production
	ApnsEnv *string `json:"apns_env,omitempty"`
	// iOS: Your apple push notification p12 certificate file, converted to a string and Base64 encoded.
	ApnsP12 *string `json:"apns_p12,omitempty"`
	// iOS: Required if using p12 certificate.  Password for the apns_p12 file.
	ApnsP12Password *string `json:"apns_p12_password,omitempty"`
	ApnsCertificates *string `json:"apns_certificates,omitempty"`
	SafariApnsCertificates *string `json:"safari_apns_certificates,omitempty"`
	// Safari: Your apple push notification p12 certificate file for Safari Push Notifications, converted to a string and Base64 encoded.
	SafariApnsP12 *string `json:"safari_apns_p12,omitempty"`
	// Safari: Password for safari_apns_p12 file
	SafariApnsP12Password *string `json:"safari_apns_p12_password,omitempty"`
	// Safari (Recommended): The hostname to your website including http(s)://
	SafariSiteOrigin *string `json:"safari_site_origin,omitempty"`
	SafariPushId *string `json:"safari_push_id,omitempty"`
	SafariIcon1616 *string `json:"safari_icon_16_16,omitempty"`
	SafariIcon3232 *string `json:"safari_icon_32_32,omitempty"`
	SafariIcon6464 *string `json:"safari_icon_64_64,omitempty"`
	SafariIcon128128 *string `json:"safari_icon_128_128,omitempty"`
	// Safari: A url for a 256x256 png notification icon. This is the only Safari icon URL you need to provide.
	SafariIcon256256 *string `json:"safari_icon_256_256,omitempty"`
	// All Browsers (Recommended): The Site Name. Requires both chrome_web_origin and safari_site_origin to be set to add or update it.
	SiteName *string `json:"site_name,omitempty"`
	BasicAuthKey *string `json:"basic_auth_key,omitempty"`
	// The Id of the Organization you would like to add this app to.
	OrganizationId *string `json:"organization_id,omitempty"`
	// iOS: Notification data (additional data) values will be added to the root of the apns payload when sent to the device.  Ignore if you're not using any other plugins, or not using OneSignal SDK methods to read the payload.
	AdditionalDataIsRootPayload *bool `json:"additional_data_is_root_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _App App

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(id string) *App {
	this := App{}
	this.Id = id
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetId returns the Id field value
func (o *App) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *App) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *App) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *App) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *App) SetName(v string) {
	o.Name = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *App) GetPlayers() int32 {
	if o == nil || o.Players == nil {
		var ret int32
		return ret
	}
	return *o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetPlayersOk() (*int32, bool) {
	if o == nil || o.Players == nil {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *App) HasPlayers() bool {
	if o != nil && o.Players != nil {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given int32 and assigns it to the Players field.
func (o *App) SetPlayers(v int32) {
	o.Players = &v
}

// GetMessageablePlayers returns the MessageablePlayers field value if set, zero value otherwise.
func (o *App) GetMessageablePlayers() int32 {
	if o == nil || o.MessageablePlayers == nil {
		var ret int32
		return ret
	}
	return *o.MessageablePlayers
}

// GetMessageablePlayersOk returns a tuple with the MessageablePlayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetMessageablePlayersOk() (*int32, bool) {
	if o == nil || o.MessageablePlayers == nil {
		return nil, false
	}
	return o.MessageablePlayers, true
}

// HasMessageablePlayers returns a boolean if a field has been set.
func (o *App) HasMessageablePlayers() bool {
	if o != nil && o.MessageablePlayers != nil {
		return true
	}

	return false
}

// SetMessageablePlayers gets a reference to the given int32 and assigns it to the MessageablePlayers field.
func (o *App) SetMessageablePlayers(v int32) {
	o.MessageablePlayers = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *App) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *App) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *App) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *App) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *App) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *App) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetAndroidGcmSenderId returns the AndroidGcmSenderId field value if set, zero value otherwise.
func (o *App) GetAndroidGcmSenderId() string {
	if o == nil || o.AndroidGcmSenderId == nil {
		var ret string
		return ret
	}
	return *o.AndroidGcmSenderId
}

// GetAndroidGcmSenderIdOk returns a tuple with the AndroidGcmSenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetAndroidGcmSenderIdOk() (*string, bool) {
	if o == nil || o.AndroidGcmSenderId == nil {
		return nil, false
	}
	return o.AndroidGcmSenderId, true
}

// HasAndroidGcmSenderId returns a boolean if a field has been set.
func (o *App) HasAndroidGcmSenderId() bool {
	if o != nil && o.AndroidGcmSenderId != nil {
		return true
	}

	return false
}

// SetAndroidGcmSenderId gets a reference to the given string and assigns it to the AndroidGcmSenderId field.
func (o *App) SetAndroidGcmSenderId(v string) {
	o.AndroidGcmSenderId = &v
}

// GetGcmKey returns the GcmKey field value if set, zero value otherwise.
func (o *App) GetGcmKey() string {
	if o == nil || o.GcmKey == nil {
		var ret string
		return ret
	}
	return *o.GcmKey
}

// GetGcmKeyOk returns a tuple with the GcmKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetGcmKeyOk() (*string, bool) {
	if o == nil || o.GcmKey == nil {
		return nil, false
	}
	return o.GcmKey, true
}

// HasGcmKey returns a boolean if a field has been set.
func (o *App) HasGcmKey() bool {
	if o != nil && o.GcmKey != nil {
		return true
	}

	return false
}

// SetGcmKey gets a reference to the given string and assigns it to the GcmKey field.
func (o *App) SetGcmKey(v string) {
	o.GcmKey = &v
}

// GetChromeWebOrigin returns the ChromeWebOrigin field value if set, zero value otherwise.
func (o *App) GetChromeWebOrigin() string {
	if o == nil || o.ChromeWebOrigin == nil {
		var ret string
		return ret
	}
	return *o.ChromeWebOrigin
}

// GetChromeWebOriginOk returns a tuple with the ChromeWebOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetChromeWebOriginOk() (*string, bool) {
	if o == nil || o.ChromeWebOrigin == nil {
		return nil, false
	}
	return o.ChromeWebOrigin, true
}

// HasChromeWebOrigin returns a boolean if a field has been set.
func (o *App) HasChromeWebOrigin() bool {
	if o != nil && o.ChromeWebOrigin != nil {
		return true
	}

	return false
}

// SetChromeWebOrigin gets a reference to the given string and assigns it to the ChromeWebOrigin field.
func (o *App) SetChromeWebOrigin(v string) {
	o.ChromeWebOrigin = &v
}

// GetChromeKey returns the ChromeKey field value if set, zero value otherwise.
func (o *App) GetChromeKey() string {
	if o == nil || o.ChromeKey == nil {
		var ret string
		return ret
	}
	return *o.ChromeKey
}

// GetChromeKeyOk returns a tuple with the ChromeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetChromeKeyOk() (*string, bool) {
	if o == nil || o.ChromeKey == nil {
		return nil, false
	}
	return o.ChromeKey, true
}

// HasChromeKey returns a boolean if a field has been set.
func (o *App) HasChromeKey() bool {
	if o != nil && o.ChromeKey != nil {
		return true
	}

	return false
}

// SetChromeKey gets a reference to the given string and assigns it to the ChromeKey field.
func (o *App) SetChromeKey(v string) {
	o.ChromeKey = &v
}

// GetChromeWebDefaultNotificationIcon returns the ChromeWebDefaultNotificationIcon field value if set, zero value otherwise.
func (o *App) GetChromeWebDefaultNotificationIcon() string {
	if o == nil || o.ChromeWebDefaultNotificationIcon == nil {
		var ret string
		return ret
	}
	return *o.ChromeWebDefaultNotificationIcon
}

// GetChromeWebDefaultNotificationIconOk returns a tuple with the ChromeWebDefaultNotificationIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetChromeWebDefaultNotificationIconOk() (*string, bool) {
	if o == nil || o.ChromeWebDefaultNotificationIcon == nil {
		return nil, false
	}
	return o.ChromeWebDefaultNotificationIcon, true
}

// HasChromeWebDefaultNotificationIcon returns a boolean if a field has been set.
func (o *App) HasChromeWebDefaultNotificationIcon() bool {
	if o != nil && o.ChromeWebDefaultNotificationIcon != nil {
		return true
	}

	return false
}

// SetChromeWebDefaultNotificationIcon gets a reference to the given string and assigns it to the ChromeWebDefaultNotificationIcon field.
func (o *App) SetChromeWebDefaultNotificationIcon(v string) {
	o.ChromeWebDefaultNotificationIcon = &v
}

// GetChromeWebSubDomain returns the ChromeWebSubDomain field value if set, zero value otherwise.
func (o *App) GetChromeWebSubDomain() string {
	if o == nil || o.ChromeWebSubDomain == nil {
		var ret string
		return ret
	}
	return *o.ChromeWebSubDomain
}

// GetChromeWebSubDomainOk returns a tuple with the ChromeWebSubDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetChromeWebSubDomainOk() (*string, bool) {
	if o == nil || o.ChromeWebSubDomain == nil {
		return nil, false
	}
	return o.ChromeWebSubDomain, true
}

// HasChromeWebSubDomain returns a boolean if a field has been set.
func (o *App) HasChromeWebSubDomain() bool {
	if o != nil && o.ChromeWebSubDomain != nil {
		return true
	}

	return false
}

// SetChromeWebSubDomain gets a reference to the given string and assigns it to the ChromeWebSubDomain field.
func (o *App) SetChromeWebSubDomain(v string) {
	o.ChromeWebSubDomain = &v
}

// GetApnsEnv returns the ApnsEnv field value if set, zero value otherwise.
func (o *App) GetApnsEnv() string {
	if o == nil || o.ApnsEnv == nil {
		var ret string
		return ret
	}
	return *o.ApnsEnv
}

// GetApnsEnvOk returns a tuple with the ApnsEnv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetApnsEnvOk() (*string, bool) {
	if o == nil || o.ApnsEnv == nil {
		return nil, false
	}
	return o.ApnsEnv, true
}

// HasApnsEnv returns a boolean if a field has been set.
func (o *App) HasApnsEnv() bool {
	if o != nil && o.ApnsEnv != nil {
		return true
	}

	return false
}

// SetApnsEnv gets a reference to the given string and assigns it to the ApnsEnv field.
func (o *App) SetApnsEnv(v string) {
	o.ApnsEnv = &v
}

// GetApnsP12 returns the ApnsP12 field value if set, zero value otherwise.
func (o *App) GetApnsP12() string {
	if o == nil || o.ApnsP12 == nil {
		var ret string
		return ret
	}
	return *o.ApnsP12
}

// GetApnsP12Ok returns a tuple with the ApnsP12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetApnsP12Ok() (*string, bool) {
	if o == nil || o.ApnsP12 == nil {
		return nil, false
	}
	return o.ApnsP12, true
}

// HasApnsP12 returns a boolean if a field has been set.
func (o *App) HasApnsP12() bool {
	if o != nil && o.ApnsP12 != nil {
		return true
	}

	return false
}

// SetApnsP12 gets a reference to the given string and assigns it to the ApnsP12 field.
func (o *App) SetApnsP12(v string) {
	o.ApnsP12 = &v
}

// GetApnsP12Password returns the ApnsP12Password field value if set, zero value otherwise.
func (o *App) GetApnsP12Password() string {
	if o == nil || o.ApnsP12Password == nil {
		var ret string
		return ret
	}
	return *o.ApnsP12Password
}

// GetApnsP12PasswordOk returns a tuple with the ApnsP12Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetApnsP12PasswordOk() (*string, bool) {
	if o == nil || o.ApnsP12Password == nil {
		return nil, false
	}
	return o.ApnsP12Password, true
}

// HasApnsP12Password returns a boolean if a field has been set.
func (o *App) HasApnsP12Password() bool {
	if o != nil && o.ApnsP12Password != nil {
		return true
	}

	return false
}

// SetApnsP12Password gets a reference to the given string and assigns it to the ApnsP12Password field.
func (o *App) SetApnsP12Password(v string) {
	o.ApnsP12Password = &v
}

// GetApnsCertificates returns the ApnsCertificates field value if set, zero value otherwise.
func (o *App) GetApnsCertificates() string {
	if o == nil || o.ApnsCertificates == nil {
		var ret string
		return ret
	}
	return *o.ApnsCertificates
}

// GetApnsCertificatesOk returns a tuple with the ApnsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetApnsCertificatesOk() (*string, bool) {
	if o == nil || o.ApnsCertificates == nil {
		return nil, false
	}
	return o.ApnsCertificates, true
}

// HasApnsCertificates returns a boolean if a field has been set.
func (o *App) HasApnsCertificates() bool {
	if o != nil && o.ApnsCertificates != nil {
		return true
	}

	return false
}

// SetApnsCertificates gets a reference to the given string and assigns it to the ApnsCertificates field.
func (o *App) SetApnsCertificates(v string) {
	o.ApnsCertificates = &v
}

// GetSafariApnsCertificates returns the SafariApnsCertificates field value if set, zero value otherwise.
func (o *App) GetSafariApnsCertificates() string {
	if o == nil || o.SafariApnsCertificates == nil {
		var ret string
		return ret
	}
	return *o.SafariApnsCertificates
}

// GetSafariApnsCertificatesOk returns a tuple with the SafariApnsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariApnsCertificatesOk() (*string, bool) {
	if o == nil || o.SafariApnsCertificates == nil {
		return nil, false
	}
	return o.SafariApnsCertificates, true
}

// HasSafariApnsCertificates returns a boolean if a field has been set.
func (o *App) HasSafariApnsCertificates() bool {
	if o != nil && o.SafariApnsCertificates != nil {
		return true
	}

	return false
}

// SetSafariApnsCertificates gets a reference to the given string and assigns it to the SafariApnsCertificates field.
func (o *App) SetSafariApnsCertificates(v string) {
	o.SafariApnsCertificates = &v
}

// GetSafariApnsP12 returns the SafariApnsP12 field value if set, zero value otherwise.
func (o *App) GetSafariApnsP12() string {
	if o == nil || o.SafariApnsP12 == nil {
		var ret string
		return ret
	}
	return *o.SafariApnsP12
}

// GetSafariApnsP12Ok returns a tuple with the SafariApnsP12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariApnsP12Ok() (*string, bool) {
	if o == nil || o.SafariApnsP12 == nil {
		return nil, false
	}
	return o.SafariApnsP12, true
}

// HasSafariApnsP12 returns a boolean if a field has been set.
func (o *App) HasSafariApnsP12() bool {
	if o != nil && o.SafariApnsP12 != nil {
		return true
	}

	return false
}

// SetSafariApnsP12 gets a reference to the given string and assigns it to the SafariApnsP12 field.
func (o *App) SetSafariApnsP12(v string) {
	o.SafariApnsP12 = &v
}

// GetSafariApnsP12Password returns the SafariApnsP12Password field value if set, zero value otherwise.
func (o *App) GetSafariApnsP12Password() string {
	if o == nil || o.SafariApnsP12Password == nil {
		var ret string
		return ret
	}
	return *o.SafariApnsP12Password
}

// GetSafariApnsP12PasswordOk returns a tuple with the SafariApnsP12Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariApnsP12PasswordOk() (*string, bool) {
	if o == nil || o.SafariApnsP12Password == nil {
		return nil, false
	}
	return o.SafariApnsP12Password, true
}

// HasSafariApnsP12Password returns a boolean if a field has been set.
func (o *App) HasSafariApnsP12Password() bool {
	if o != nil && o.SafariApnsP12Password != nil {
		return true
	}

	return false
}

// SetSafariApnsP12Password gets a reference to the given string and assigns it to the SafariApnsP12Password field.
func (o *App) SetSafariApnsP12Password(v string) {
	o.SafariApnsP12Password = &v
}

// GetSafariSiteOrigin returns the SafariSiteOrigin field value if set, zero value otherwise.
func (o *App) GetSafariSiteOrigin() string {
	if o == nil || o.SafariSiteOrigin == nil {
		var ret string
		return ret
	}
	return *o.SafariSiteOrigin
}

// GetSafariSiteOriginOk returns a tuple with the SafariSiteOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariSiteOriginOk() (*string, bool) {
	if o == nil || o.SafariSiteOrigin == nil {
		return nil, false
	}
	return o.SafariSiteOrigin, true
}

// HasSafariSiteOrigin returns a boolean if a field has been set.
func (o *App) HasSafariSiteOrigin() bool {
	if o != nil && o.SafariSiteOrigin != nil {
		return true
	}

	return false
}

// SetSafariSiteOrigin gets a reference to the given string and assigns it to the SafariSiteOrigin field.
func (o *App) SetSafariSiteOrigin(v string) {
	o.SafariSiteOrigin = &v
}

// GetSafariPushId returns the SafariPushId field value if set, zero value otherwise.
func (o *App) GetSafariPushId() string {
	if o == nil || o.SafariPushId == nil {
		var ret string
		return ret
	}
	return *o.SafariPushId
}

// GetSafariPushIdOk returns a tuple with the SafariPushId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariPushIdOk() (*string, bool) {
	if o == nil || o.SafariPushId == nil {
		return nil, false
	}
	return o.SafariPushId, true
}

// HasSafariPushId returns a boolean if a field has been set.
func (o *App) HasSafariPushId() bool {
	if o != nil && o.SafariPushId != nil {
		return true
	}

	return false
}

// SetSafariPushId gets a reference to the given string and assigns it to the SafariPushId field.
func (o *App) SetSafariPushId(v string) {
	o.SafariPushId = &v
}

// GetSafariIcon1616 returns the SafariIcon1616 field value if set, zero value otherwise.
func (o *App) GetSafariIcon1616() string {
	if o == nil || o.SafariIcon1616 == nil {
		var ret string
		return ret
	}
	return *o.SafariIcon1616
}

// GetSafariIcon1616Ok returns a tuple with the SafariIcon1616 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariIcon1616Ok() (*string, bool) {
	if o == nil || o.SafariIcon1616 == nil {
		return nil, false
	}
	return o.SafariIcon1616, true
}

// HasSafariIcon1616 returns a boolean if a field has been set.
func (o *App) HasSafariIcon1616() bool {
	if o != nil && o.SafariIcon1616 != nil {
		return true
	}

	return false
}

// SetSafariIcon1616 gets a reference to the given string and assigns it to the SafariIcon1616 field.
func (o *App) SetSafariIcon1616(v string) {
	o.SafariIcon1616 = &v
}

// GetSafariIcon3232 returns the SafariIcon3232 field value if set, zero value otherwise.
func (o *App) GetSafariIcon3232() string {
	if o == nil || o.SafariIcon3232 == nil {
		var ret string
		return ret
	}
	return *o.SafariIcon3232
}

// GetSafariIcon3232Ok returns a tuple with the SafariIcon3232 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariIcon3232Ok() (*string, bool) {
	if o == nil || o.SafariIcon3232 == nil {
		return nil, false
	}
	return o.SafariIcon3232, true
}

// HasSafariIcon3232 returns a boolean if a field has been set.
func (o *App) HasSafariIcon3232() bool {
	if o != nil && o.SafariIcon3232 != nil {
		return true
	}

	return false
}

// SetSafariIcon3232 gets a reference to the given string and assigns it to the SafariIcon3232 field.
func (o *App) SetSafariIcon3232(v string) {
	o.SafariIcon3232 = &v
}

// GetSafariIcon6464 returns the SafariIcon6464 field value if set, zero value otherwise.
func (o *App) GetSafariIcon6464() string {
	if o == nil || o.SafariIcon6464 == nil {
		var ret string
		return ret
	}
	return *o.SafariIcon6464
}

// GetSafariIcon6464Ok returns a tuple with the SafariIcon6464 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariIcon6464Ok() (*string, bool) {
	if o == nil || o.SafariIcon6464 == nil {
		return nil, false
	}
	return o.SafariIcon6464, true
}

// HasSafariIcon6464 returns a boolean if a field has been set.
func (o *App) HasSafariIcon6464() bool {
	if o != nil && o.SafariIcon6464 != nil {
		return true
	}

	return false
}

// SetSafariIcon6464 gets a reference to the given string and assigns it to the SafariIcon6464 field.
func (o *App) SetSafariIcon6464(v string) {
	o.SafariIcon6464 = &v
}

// GetSafariIcon128128 returns the SafariIcon128128 field value if set, zero value otherwise.
func (o *App) GetSafariIcon128128() string {
	if o == nil || o.SafariIcon128128 == nil {
		var ret string
		return ret
	}
	return *o.SafariIcon128128
}

// GetSafariIcon128128Ok returns a tuple with the SafariIcon128128 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariIcon128128Ok() (*string, bool) {
	if o == nil || o.SafariIcon128128 == nil {
		return nil, false
	}
	return o.SafariIcon128128, true
}

// HasSafariIcon128128 returns a boolean if a field has been set.
func (o *App) HasSafariIcon128128() bool {
	if o != nil && o.SafariIcon128128 != nil {
		return true
	}

	return false
}

// SetSafariIcon128128 gets a reference to the given string and assigns it to the SafariIcon128128 field.
func (o *App) SetSafariIcon128128(v string) {
	o.SafariIcon128128 = &v
}

// GetSafariIcon256256 returns the SafariIcon256256 field value if set, zero value otherwise.
func (o *App) GetSafariIcon256256() string {
	if o == nil || o.SafariIcon256256 == nil {
		var ret string
		return ret
	}
	return *o.SafariIcon256256
}

// GetSafariIcon256256Ok returns a tuple with the SafariIcon256256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSafariIcon256256Ok() (*string, bool) {
	if o == nil || o.SafariIcon256256 == nil {
		return nil, false
	}
	return o.SafariIcon256256, true
}

// HasSafariIcon256256 returns a boolean if a field has been set.
func (o *App) HasSafariIcon256256() bool {
	if o != nil && o.SafariIcon256256 != nil {
		return true
	}

	return false
}

// SetSafariIcon256256 gets a reference to the given string and assigns it to the SafariIcon256256 field.
func (o *App) SetSafariIcon256256(v string) {
	o.SafariIcon256256 = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *App) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *App) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *App) SetSiteName(v string) {
	o.SiteName = &v
}

// GetBasicAuthKey returns the BasicAuthKey field value if set, zero value otherwise.
func (o *App) GetBasicAuthKey() string {
	if o == nil || o.BasicAuthKey == nil {
		var ret string
		return ret
	}
	return *o.BasicAuthKey
}

// GetBasicAuthKeyOk returns a tuple with the BasicAuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetBasicAuthKeyOk() (*string, bool) {
	if o == nil || o.BasicAuthKey == nil {
		return nil, false
	}
	return o.BasicAuthKey, true
}

// HasBasicAuthKey returns a boolean if a field has been set.
func (o *App) HasBasicAuthKey() bool {
	if o != nil && o.BasicAuthKey != nil {
		return true
	}

	return false
}

// SetBasicAuthKey gets a reference to the given string and assigns it to the BasicAuthKey field.
func (o *App) SetBasicAuthKey(v string) {
	o.BasicAuthKey = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *App) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *App) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *App) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetAdditionalDataIsRootPayload returns the AdditionalDataIsRootPayload field value if set, zero value otherwise.
func (o *App) GetAdditionalDataIsRootPayload() bool {
	if o == nil || o.AdditionalDataIsRootPayload == nil {
		var ret bool
		return ret
	}
	return *o.AdditionalDataIsRootPayload
}

// GetAdditionalDataIsRootPayloadOk returns a tuple with the AdditionalDataIsRootPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetAdditionalDataIsRootPayloadOk() (*bool, bool) {
	if o == nil || o.AdditionalDataIsRootPayload == nil {
		return nil, false
	}
	return o.AdditionalDataIsRootPayload, true
}

// HasAdditionalDataIsRootPayload returns a boolean if a field has been set.
func (o *App) HasAdditionalDataIsRootPayload() bool {
	if o != nil && o.AdditionalDataIsRootPayload != nil {
		return true
	}

	return false
}

// SetAdditionalDataIsRootPayload gets a reference to the given bool and assigns it to the AdditionalDataIsRootPayload field.
func (o *App) SetAdditionalDataIsRootPayload(v bool) {
	o.AdditionalDataIsRootPayload = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Players != nil {
		toSerialize["players"] = o.Players
	}
	if o.MessageablePlayers != nil {
		toSerialize["messageable_players"] = o.MessageablePlayers
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.AndroidGcmSenderId != nil {
		toSerialize["android_gcm_sender_id"] = o.AndroidGcmSenderId
	}
	if o.GcmKey != nil {
		toSerialize["gcm_key"] = o.GcmKey
	}
	if o.ChromeWebOrigin != nil {
		toSerialize["chrome_web_origin"] = o.ChromeWebOrigin
	}
	if o.ChromeKey != nil {
		toSerialize["chrome_key"] = o.ChromeKey
	}
	if o.ChromeWebDefaultNotificationIcon != nil {
		toSerialize["chrome_web_default_notification_icon"] = o.ChromeWebDefaultNotificationIcon
	}
	if o.ChromeWebSubDomain != nil {
		toSerialize["chrome_web_sub_domain"] = o.ChromeWebSubDomain
	}
	if o.ApnsEnv != nil {
		toSerialize["apns_env"] = o.ApnsEnv
	}
	if o.ApnsP12 != nil {
		toSerialize["apns_p12"] = o.ApnsP12
	}
	if o.ApnsP12Password != nil {
		toSerialize["apns_p12_password"] = o.ApnsP12Password
	}
	if o.ApnsCertificates != nil {
		toSerialize["apns_certificates"] = o.ApnsCertificates
	}
	if o.SafariApnsCertificates != nil {
		toSerialize["safari_apns_certificates"] = o.SafariApnsCertificates
	}
	if o.SafariApnsP12 != nil {
		toSerialize["safari_apns_p12"] = o.SafariApnsP12
	}
	if o.SafariApnsP12Password != nil {
		toSerialize["safari_apns_p12_password"] = o.SafariApnsP12Password
	}
	if o.SafariSiteOrigin != nil {
		toSerialize["safari_site_origin"] = o.SafariSiteOrigin
	}
	if o.SafariPushId != nil {
		toSerialize["safari_push_id"] = o.SafariPushId
	}
	if o.SafariIcon1616 != nil {
		toSerialize["safari_icon_16_16"] = o.SafariIcon1616
	}
	if o.SafariIcon3232 != nil {
		toSerialize["safari_icon_32_32"] = o.SafariIcon3232
	}
	if o.SafariIcon6464 != nil {
		toSerialize["safari_icon_64_64"] = o.SafariIcon6464
	}
	if o.SafariIcon128128 != nil {
		toSerialize["safari_icon_128_128"] = o.SafariIcon128128
	}
	if o.SafariIcon256256 != nil {
		toSerialize["safari_icon_256_256"] = o.SafariIcon256256
	}
	if o.SiteName != nil {
		toSerialize["site_name"] = o.SiteName
	}
	if o.BasicAuthKey != nil {
		toSerialize["basic_auth_key"] = o.BasicAuthKey
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.AdditionalDataIsRootPayload != nil {
		toSerialize["additional_data_is_root_payload"] = o.AdditionalDataIsRootPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *App) UnmarshalJSON(bytes []byte) (err error) {
	varApp := _App{}

	if err = json.Unmarshal(bytes, &varApp); err == nil {
		*o = App(varApp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "players")
		delete(additionalProperties, "messageable_players")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "android_gcm_sender_id")
		delete(additionalProperties, "gcm_key")
		delete(additionalProperties, "chrome_web_origin")
		delete(additionalProperties, "chrome_key")
		delete(additionalProperties, "chrome_web_default_notification_icon")
		delete(additionalProperties, "chrome_web_sub_domain")
		delete(additionalProperties, "apns_env")
		delete(additionalProperties, "apns_p12")
		delete(additionalProperties, "apns_p12_password")
		delete(additionalProperties, "apns_certificates")
		delete(additionalProperties, "safari_apns_certificates")
		delete(additionalProperties, "safari_apns_p12")
		delete(additionalProperties, "safari_apns_p12_password")
		delete(additionalProperties, "safari_site_origin")
		delete(additionalProperties, "safari_push_id")
		delete(additionalProperties, "safari_icon_16_16")
		delete(additionalProperties, "safari_icon_32_32")
		delete(additionalProperties, "safari_icon_64_64")
		delete(additionalProperties, "safari_icon_128_128")
		delete(additionalProperties, "safari_icon_256_256")
		delete(additionalProperties, "site_name")
		delete(additionalProperties, "basic_auth_key")
		delete(additionalProperties, "organization_id")
		delete(additionalProperties, "additional_data_is_root_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


