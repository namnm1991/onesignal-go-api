# NotificationWithMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedSegments** | Pointer to **[]string** | The segment names you want to target. Users in these segments will receive a notification. This targeting parameter is only compatible with excluded_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**ExcludedSegments** | Pointer to **[]string** | Segment that will be excluded when sending. Users in these segments will not receive a notification, even if they were included in included_segments. This targeting parameter is only compatible with included_segments. Example: [\&quot;Active Users\&quot;, \&quot;Inactive Users\&quot;]  | [optional] 
**LastSession** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot; or \&quot;&lt;\&quot; hours_ago &#x3D; number of hours before or after the users last session. Example: \&quot;1.1\&quot;  | [optional] 
**FirstSession** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot; or \&quot;&lt;\&quot; hours_ago &#x3D; number of hours before or after the users first session. Example: \&quot;1.1\&quot;  | [optional] 
**SessionCount** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; number sessions. Example: \&quot;1\&quot;  | [optional] 
**SessionTime** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; Time in seconds the user has been in your app. Example: \&quot;3600\&quot;  | [optional] 
**AmountSpent** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, or \&quot;&#x3D;\&quot; value &#x3D; Amount in USD a user has spent on IAP (In App Purchases). Example: \&quot;0.99\&quot;  | [optional] 
**BoughtSku** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot; or \&quot;&#x3D;\&quot; key &#x3D; SKU purchased in your app as an IAP (In App Purchases). Example: \&quot;com.domain.100coinpack\&quot; value &#x3D; value of SKU to compare to. Example: \&quot;0.99\&quot;  | [optional] 
**Tag** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot;, \&quot;!&#x3D;\&quot;, \&quot;exists\&quot;, \&quot;not_exists\&quot;, \&quot;time_elapsed_gt\&quot; (paid plan only) or \&quot;time_elapsed_lt\&quot; (paid plan only) See Time Operators key &#x3D; Tag key to compare. value &#x3D; Tag value to compare. Not required for \&quot;exists\&quot; or \&quot;not_exists\&quot;. Example: See Formatting Filters  | [optional] 
**Language** | Pointer to **string** | relation &#x3D; \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; 2 character language code. Example: \&quot;en\&quot;. For a list of all language codes see Language &amp; Localization.  | [optional] 
**AppVersion** | Pointer to **string** | relation &#x3D; \&quot;&gt;\&quot;, \&quot;&lt;\&quot;, \&quot;&#x3D;\&quot; or \&quot;!&#x3D;\&quot; value &#x3D; app version. Example: \&quot;1.0.0\&quot;  | [optional] 
**Location** | Pointer to **string** | radius &#x3D; in meters lat &#x3D; latitude long &#x3D; longitude  | [optional] 
**Email** | Pointer to **string** | value &#x3D; email address Only for sending Push Notifications Use this for targeting push subscribers associated with an email set with all SDK setEmail methods To send emails to specific email addresses use include_email_tokens parameter  | [optional] 
**Country** | Pointer to **string** | relation &#x3D; \&quot;&#x3D;\&quot; value &#x3D; 2-digit Country code Example: \&quot;field\&quot;: \&quot;country\&quot;, \&quot;relation\&quot;: \&quot;&#x3D;\&quot;, \&quot;value\&quot;, \&quot;US\&quot;  | [optional] 
**IncludePlayerIds** | Pointer to **[]string** | Specific playerids to send your notification to. _Does not require API Auth Key. Do not combine with other targeting parameters. Not compatible with any other targeting parameters. Example: [\&quot;1dd608f2-c6a1-11e3-851d-000c2940e62c\&quot;] Limit of 2,000 entries per REST API call  | [optional] 
**IncludeExternalUserIds** | Pointer to **[]string** | Target specific devices by custom user IDs assigned via API. Not compatible with any other targeting parameters Example: [\&quot;custom-id-assigned-by-api\&quot;] REQUIRED: REST API Key Authentication Limit of 2,000 entries per REST API call. Note: If targeting push, email, or sms subscribers with same ids, use with channel_for_external_user_ids to indicate you are sending a push or email or sms.  | [optional] 
**IncludeEmailTokens** | Pointer to **[]string** | Recommended for Sending Emails - Target specific email addresses. If an email does not correspond to an existing user, a new user will be created. Example: nick@catfac.ts Limit of 2,000 entries per REST API call  | [optional] 
**IncludePhoneNumbers** | Pointer to **[]string** | Recommended for Sending SMS - Target specific phone numbers. The phone number should be in the E.164 format. Phone number should be an existing subscriber on OneSignal. Refer our docs to learn how to add phone numbers to OneSignal. Example phone number: +1999999999 Limit of 2,000 entries per REST API call  | [optional] 
**IncludeIosTokens** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using iOS device tokens. Warning: Only works with Production tokens. All non-alphanumeric characters must be removed from each token. If a token does not correspond to an existing user, a new user will be created. Example: ce777617da7f548fe7a9ab6febb56cf39fba6d38203... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeWpWnsUris** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Windows URIs. If a token does not correspond to an existing user, a new user will be created. Example: http://s.notify.live.net/u/1/bn1/HmQAAACPaLDr-... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAmazonRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Amazon ADM registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: amzn1.adm-registration.v1.XpvSSUk0Rc3hTVVV... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome App registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeChromeWebRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Chrome Web Push registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**IncludeAndroidRegIds** | Pointer to **[]string** | Not Recommended: Please consider using include_player_ids or include_external_user_ids instead. Target using Android device registration IDs. If a token does not correspond to an existing user, a new user will be created. Example: APA91bEeiUeSukAAUdnw3O2RB45FWlSpgJ7Ji_... Limit of 2,000 entries per REST API call  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **int32** |  | [optional] [readonly] 
**Aggregation** | Pointer to **string** |  | [optional] [readonly] 
**IsIos** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Apple iOS platform. | [optional] [default to true]
**IsAndroid** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Google Android platform. | [optional] 
**IsHuawei** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Huawei Android platform. | [optional] 
**IsAnyWeb** | Pointer to **NullableBool** | Indicates whether to send to all subscribed web browser users, including Chrome, Firefox, and Safari. You may use this instead as a combined flag instead of separately enabling isChromeWeb, isFirefox, and isSafari, though the three options are equivalent to this one.  | [optional] 
**IsChromeWeb** | Pointer to **NullableBool** | Indicates whether to send to all Google Chrome, Chrome on Android, and Mozilla Firefox users registered under your Chrome &amp; Firefox web push platform. | [optional] 
**IsFirefox** | Pointer to **NullableBool** | Indicates whether to send to all Mozilla Firefox desktop users registered under your Firefox web push platform. | [optional] 
**IsSafari** | Pointer to **NullableBool** | Does not support iOS Safari. Indicates whether to send to all Apple&#39;s Safari desktop users registered under your Safari web push platform. Read more iOS Safari | [optional] 
**IsWPWNS** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Windows platform. | [optional] 
**IsAdm** | Pointer to **NullableBool** | Indicates whether to send to all devices registered under your app&#39;s Amazon Fire platform. | [optional] 
**IsChrome** | Pointer to **NullableBool** | This flag is not used for web push Please see isChromeWeb for sending to web push users. This flag only applies to Google Chrome Apps &amp; Extensions. Indicates whether to send to all devices registered under your app&#39;s Google Chrome Apps &amp; Extension platform.  | [optional] 
**ChannelForExternalUserIds** | Pointer to **string** | Indicates if the message type when targeting with include_external_user_ids for cases where an email, sms, and/or push subscribers have the same external user id. Example: Use the string \&quot;push\&quot; to indicate you are sending a push notification or the string \&quot;email\&quot;for sending emails or \&quot;sms\&quot;for sending SMS.  | [optional] 
**AppId** | **string** | Required: Your OneSignal Application ID, which can be found in Keys &amp; IDs. It is a UUID and looks similar to 8250eaf6-1a58-489e-b136-7c74a864b434.  | 
**ExternalId** | Pointer to **string** | Correlation and idempotency key. A request received with this parameter will first look for another notification with the same external_id. If one exists, a notification will not be sent, and result of the previous operation will instead be returned. Therefore, if you plan on using this feature, it&#39;s important to use a good source of randomness to generate the UUID passed here. This key is only idempotent for 30 days. After 30 days, the notification could be removed from our system and a notification with the same external_id will be sent again.   See Idempotent Notification Requests for more details writeOnly: true  | [optional] 
**Contents** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Headings** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Subtitle** | Pointer to [**NullableStringMap**](StringMap.md) |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Huawei A custom map of data that is passed back to your app. Same as using Additional Data within the dashboard. Can use up to 2048 bytes of data. Example: {\&quot;abc\&quot;: 123, \&quot;foo\&quot;: \&quot;bar\&quot;, \&quot;event_performed\&quot;: true, \&quot;amount\&quot;: 12.1}  | [optional] 
**HuaweiMsgType** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Use \&quot;data\&quot; or \&quot;message\&quot; depending on the type of notification you are sending. More details in Data &amp; Background Notifications.  | [optional] 
**Url** | Pointer to **string** | Channel: Push Notifications Platform: All The URL to open in the browser when a user clicks on the notification. Note: iOS needs https or updated NSAppTransportSecurity in plist This field supports inline substitutions. Omit if including web_url or app_url Example: https://onesignal.com  | [optional] 
**WebUrl** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including Chrome, Firefox, Safari, Opera, etc. Example: https://onesignal.com  | [optional] 
**AppUrl** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Same as url but only sent to web push platforms. Including iOS, Android, macOS, Windows, ChromeApps, etc. Example: https://onesignal.com  | [optional] 
**IosAttachments** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ Adds media attachments to notifications. Set as JSON object, key as a media id of your choice and the value as a valid local filename or URL. User must press and hold on the notification to view. Do not set mutable_content to download attachments. The OneSignal SDK does this automatically Example: {\&quot;id1\&quot;: \&quot;https://domain.com/image.jpg\&quot;}  | [optional] 
**TemplateId** | Pointer to **string** | Channel: Push Notifications Platform: All Use a template you setup on our dashboard. The template_id is the UUID found in the URL when viewing a template on our dashboard. Example: be4a8044-bbd6-11e4-a581-000c2940e62c  | [optional] 
**ContentAvailable** | Pointer to **NullableBool** | Channel: Push Notifications Platform: iOS Sending true wakes your app from background to run custom native code (Apple interprets this as content-available&#x3D;1). Note: Not applicable if the app is in the \&quot;force-quit\&quot; state (i.e app was swiped away). Omit the contents field to prevent displaying a visible notification.  | [optional] 
**MutableContent** | Pointer to **bool** | Channel: Push Notifications Platform: iOS 10+ Always defaults to true and cannot be turned off. Allows tracking of notification receives and changing of the notification content in your app before it is displayed. Triggers didReceive(_:withContentHandler:) on your UNNotificationServiceExtension.  | [optional] 
**TargetContentIdentifier** | Pointer to **string** | Channel: Push Notifications Platform: iOS 13+ Use to target a specific experience in your App Clip, or to target your notification to a specific window in a multi-scene App.  | [optional] 
**BigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Android Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**HuaweiBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**AdmBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: Amazon Picture to display in the expanded view. Can be a drawable resource name or a URL.  | [optional] 
**ChromeBigPicture** | Pointer to **string** | Channel: Push Notifications Platform: ChromeApp Large picture to display below the notification text. Must be a local URL.  | [optional] 
**ChromeWebImage** | Pointer to **string** | Channel: Push Notifications Platform: Chrome 56+ Sets the web push notification&#39;s large image to be shown below the notification&#39;s title and text. Please see Web Push Notification Icons.  | [optional] 
**Buttons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: iOS 8.0+, Android 4.1+, and derivatives like Amazon Buttons to add to the notification. Icon only works for Android. Buttons show in reverse order of array position i.e. Last item in array shows as first button on device. Example: [{\&quot;id\&quot;: \&quot;id2\&quot;, \&quot;text\&quot;: \&quot;second button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_share\&quot;}, {\&quot;id\&quot;: \&quot;id1\&quot;, \&quot;text\&quot;: \&quot;first button\&quot;, \&quot;icon\&quot;: \&quot;ic_menu_send\&quot;}]  | [optional] 
**WebButtons** | Pointer to [**[]Button**](Button.md) | Channel: Push Notifications Platform: Chrome 48+ Add action buttons to the notification. The id field is required. Example: [{\&quot;id\&quot;: \&quot;like-button\&quot;, \&quot;text\&quot;: \&quot;Like\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/N8SN8ZS.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}, {\&quot;id\&quot;: \&quot;read-more-button\&quot;, \&quot;text\&quot;: \&quot;Read more\&quot;, \&quot;icon\&quot;: \&quot;http://i.imgur.com/MIxJp1L.png\&quot;, \&quot;url\&quot;: \&quot;https://yoursite.com\&quot;}]  | [optional] 
**IosCategory** | Pointer to **string** | Channel: Push Notifications Platform: iOS Category APS payload, use with registerUserNotificationSettings:categories in your Objective-C / Swift code. Example: calendar category which contains actions like accept and decline iOS 10+ This will trigger your UNNotificationContentExtension whose ID matches this category.  | [optional] 
**AndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**HuaweiChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Huawei The Android Oreo Notification Category to send the notification under. See the Category documentation on creating one and getting it&#39;s id.  | [optional] 
**ExistingAndroidChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Android Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**HuaweiExistingChannelId** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Use this if you have client side Android Oreo Channels you have already defined in your app with code.  | [optional] 
**AndroidBackgroundLayout** | Pointer to [**NotificationAllOfAndroidBackgroundLayout**](NotificationAllOfAndroidBackgroundLayout.md) |  | [optional] 
**SmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Android Icon shown in the status bar and on the top left of the notification. If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**HuaweiSmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Icon shown in the status bar and on the top left of the notification. Use an Android resource path (E.g. /drawable/small_icon). Defaults to your app icon if not set.  | [optional] 
**LargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Android Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**HuaweiLargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**AdmSmallIcon** | Pointer to **string** | Channel: Push Notifications Platform: Amazon If not set a bell icon will be used or ic_stat_onesignal_default if you have set this resource name. See: How to create small icons  | [optional] 
**AdmLargeIcon** | Pointer to **string** | Channel: Push Notifications Platform: Amazon If blank the small_icon is used. Can be a drawable resource name or a URL. See: How to create large icons  | [optional] 
**ChromeWebIcon** | Pointer to **string** | Channel: Push Notifications Platform: Chrome Sets the web push notification&#39;s icon. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices. Firefox will also use this icon, unless you specify firefox_icon.  | [optional] 
**ChromeWebBadge** | Pointer to **string** | Channel: Push Notifications Platform: Chrome Sets the web push notification icon for Android devices in the notification shade. Please see Web Push Notification Badge.  | [optional] 
**FirefoxIcon** | Pointer to **string** | Channel: Push Notifications Platform: Firefox Not recommended Few people need to set Firefox-specific icons. We recommend setting chrome_web_icon instead, which Firefox will also use. Sets the web push notification&#39;s icon for Firefox. An image URL linking to a valid image. Common image types are supported; GIF will not animate. We recommend 256x256 (at least 80x80) to display well on high DPI devices.  | [optional] 
**ChromeIcon** | Pointer to **string** | Channel: Push Notifications Platform: ChromeApp This flag is not used for web push For web push, please see chrome_web_icon instead. The local URL to an icon to use. If blank, the app icon will be used.  | [optional] 
**IosSound** | Pointer to **string** | Channel: Push Notifications Platform: iOS Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable vibration and sound for the notification. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidSound** | Pointer to **string** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**HuaweiSound** | Pointer to **string** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sound file that is included in your app to play instead of the default device notification sound. NOTE: Leave off file extension for and include the full path.  Example: \&quot;/res/raw/notification\&quot;  | [optional] 
**AdmSound** | Pointer to **string** | Channel: Push Notifications Platform: Amazon &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sound file that is included in your app to play instead of the default device notification sound. Pass nil to disable sound for the notification. NOTE: Leave off file extension for Android. Example: \&quot;notification\&quot;  | [optional] 
**WpWnsSound** | Pointer to **string** | Channel: Push Notifications Platform: Windows Sound file that is included in your app to play instead of the default device notification sound. Example: \&quot;notification.wav\&quot;  | [optional] 
**AndroidLedColor** | Pointer to **string** | Channel: Push Notifications Platform: Android &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. Sets the devices LED notification light if the device has one. ARGB Hex format. Example(Blue): \&quot;FF0000FF\&quot;  | [optional] 
**HuaweiLedColor** | Pointer to **string** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. Sets the devices LED notification light if the device has one. RGB Hex format. Example(Blue): \&quot;0000FF\&quot;  | [optional] 
**AndroidAccentColor** | Pointer to **string** | Channel: Push Notifications Platform: Android Sets the background color of the notification circle to the left of the notification text. Only applies to apps targeting Android API level 21+ on Android 5.0+ devices. Example(Red): \&quot;FFFF0000\&quot;  | [optional] 
**HuaweiAccentColor** | Pointer to **string** | Channel: Push Notifications Platform: Huawei Accent Color used on Action Buttons and Group overflow count. Uses RGB Hex value (E.g. #9900FF). Defaults to device&#39;s theme color if not set.  | [optional] 
**AndroidVisibility** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Android 5.0_ &amp;#9888;&amp;#65039;Deprecated, this field doesn&#39;t work on Android 8 (Oreo) and newer devices! Please use Notification Categories / Channels noted above instead to support ALL versions of Android. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**HuaweiVisibility** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Huawei &amp;#9888;&amp;#65039;Deprecated, this field ONLY works on EMUI 5 (Android 7 based) and older devices. Please also set Notification Categories / Channels noted above to support EMUI 8 (Android 8 based) devices. 1 &#x3D; Public (default) (Shows the full message on the lock screen unless the user has disabled all notifications from showing on the lock screen. Please consider the user and mark private if the contents are.) 0 &#x3D; Private (Hides message contents on lock screen if the user set \&quot;Hide sensitive notification content\&quot; in the system settings) -1 &#x3D; Secret (Notification does not show on the lock screen at all)  | [optional] 
**IosBadgeType** | Pointer to **string** | Channel: Push Notifications Platform: iOS Describes whether to set or increase/decrease your app&#39;s iOS badge count by the ios_badgeCount specified count. Can specify None, SetTo, or Increase. &#x60;None&#x60; leaves the count unaffected. &#x60;SetTo&#x60; directly sets the badge count to the number specified in ios_badgeCount. &#x60;Increase&#x60; adds the number specified in ios_badgeCount to the total. Use a negative number to decrease the badge count.  | [optional] 
**IosBadgeCount** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: iOS Used with ios_badgeType, describes the value to set or amount to increase/decrease your app&#39;s iOS badge count by. You can use a negative number to decrease the badge count when used with an ios_badgeType of Increase.  | [optional] 
**CollapseId** | Pointer to **string** | Channel: Push Notifications Platform: iOS 10+, Android Only one notification with the same id will be shown on the device. Use the same id to update an existing notification instead of showing a new one. Limit of 64 characters.  | [optional] 
**WebPushTopic** | Pointer to **string** | Channel: Push Notifications Platform: All Browsers Display multiple notifications at once with different topics.  | [optional] 
**ApnsAlert** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: iOS 10+ iOS can localize push notification messages on the client using special parameters such as loc-key. When using the Create Notification endpoint, you must include these parameters inside of a field called apns_alert. Please see Apple&#39;s guide on localizing push notifications to learn more.  | [optional] 
**SendAfter** | Pointer to **int64** | Unix timestamp indicating when notification delivery should begin. | [optional] 
**DelayedOption** | Pointer to **string** | Channel: All Possible values are: timezone (Deliver at a specific time-of-day in each users own timezone) last-active Same as Intelligent Delivery . (Deliver at the same time of day as each user last used your app). If send_after is used, this takes effect after the send_after time has elapsed.  | [optional] 
**DeliveryTimeOfDay** | Pointer to **string** | Channel: All Use with delayed_option&#x3D;timezone. Examples: \&quot;9:00AM\&quot; \&quot;21:45\&quot; \&quot;9:45:30\&quot;  | [optional] 
**Ttl** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: iOS, Android, Chrome, Firefox, Safari, ChromeWeb Time To Live - In seconds. The notification will be expired if the device does not come back online within this time. The default is 259,200 seconds (3 days). Max value to set is 2419200 seconds (28 days).  | [optional] 
**Priority** | Pointer to **NullableInt32** | Channel: Push Notifications Platform: Android, Chrome, ChromeWeb Delivery priority through the push server (example GCM/FCM). Pass 10 for high priority or any other integer for normal priority. Defaults to normal priority for Android and high for iOS. For Android 6.0+ devices setting priority to high will wake the device out of doze mode.  | [optional] 
**ApnsPushTypeOverride** | Pointer to **string** | Channel: Push Notifications Platform: iOS valid values: voip Set the value to voip for sending VoIP Notifications This field maps to the APNS header apns-push-type. Note: alert and background are automatically set by OneSignal  | [optional] 
**ThrottleRatePerMinute** | Pointer to **NullableInt32** | number of push notifications sent per minute. Paid Feature Only. If throttling is not enabled for the app or the notification, and for free accounts, null is returned. Refer to Throttling for more details. | [optional] 
**AndroidGroup** | Pointer to **string** | Channel: Push Notifications Platform: Android Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AndroidGroupMessage** | Pointer to **string** | Channel: Push Notifications Platform: Android Note: This only works for Android 6 and older. Android 7+ allows full expansion of all message. Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. Languages - The value of each key is the message that will be sent to users for that language. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. Read more: supported languages. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**AdmGroup** | Pointer to **string** | Channel: Push Notifications Platform: Amazon Notifications with the same group will be stacked together using Android&#39;s Notification Grouping feature.  | [optional] 
**AdmGroupMessage** | Pointer to **map[string]interface{}** | Channel: Push Notifications Platform: Amazon Summary message to display when 2+ notifications are stacked together. Default is \&quot;# new messages\&quot;. Include $[notif_count] in your message and it will be replaced with the current number. \&quot;en\&quot; (English) is required. The key of each hash is either a a 2 character language code or one of zh-Hans/zh-Hant for Simplified or Traditional Chinese. The value of each key is the message that will be sent to users for that language. Example: {\&quot;en\&quot;: \&quot;You have $[notif_count] new messages\&quot;}  | [optional] 
**ThreadId** | Pointer to **string** | Channel: Push Notifications Platform: iOS 12+ This parameter is supported in iOS 12 and above. It allows you to group related notifications together. If two notifications have the same thread-id, they will both be added to the same group.  | [optional] 
**SummaryArg** | Pointer to **string** | Channel: Push Notifications Platform: iOS 12+ When using thread_id to create grouped notifications in iOS 12+, you can also control the summary. For example, a grouped notification can say \&quot;12 more notifications from John Doe\&quot;. The summary_arg lets you set the name of the person/thing the notifications are coming from, and will show up as \&quot;X more notifications from summary_arg\&quot;  | [optional] 
**SummaryArgCount** | Pointer to **int32** | Channel: Push Notifications Platform: iOS 12+ When using thread_id, you can also control the count of the number of notifications in the group. For example, if the group already has 12 notifications, and you send a new notification with summary_arg_count &#x3D; 2, the new total will be 14 and the summary will be \&quot;14 more notifications from summary_arg\&quot;  | [optional] 
**EmailSubject** | Pointer to **string** | Channel: Email Required.  The subject of the email.  | [optional] 
**EmailBody** | Pointer to **string** | Channel: Email Required unless template_id is set. HTML suported The body of the email you wish to send. Typically, customers include their own HTML templates here. Must include [unsubscribe_url] in an &lt;a&gt; tag somewhere in the email. Note: any malformed HTML content will be sent to users. Please double-check your HTML is valid.  | [optional] 
**EmailFromName** | Pointer to **string** | Channel: Email The name the email is from. If not specified, will default to \&quot;from name\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**EmailFromAddress** | Pointer to **string** | Channel: Email The email address the email is from. If not specified, will default to \&quot;from email\&quot; set in the OneSignal Dashboard Email Settings.  | [optional] 
**SmsFrom** | Pointer to **string** | Channel: SMS Phone Number used to send SMS. Should be a registered Twilio phone number in E.164 format.  | [optional] 
**SmsMediaUrls** | Pointer to **[]string** | Channel: SMS URLs for the media files to be attached to the SMS content. Limit: 10 media urls with a total max. size of 5MBs.  | [optional] 
**Successful** | Pointer to **int32** | Number of notifications that were successfully delivered. | [optional] 
**Failed** | Pointer to **int32** | Number of notifications that could not be delivered due to those devices being unsubscribed. | [optional] 
**Errored** | Pointer to **int32** | Number of notifications that could not be delivered due to an error. You can find more information by viewing the notification in the dashboard. | [optional] 
**Converted** | Pointer to **int32** | Number of users who have clicked / tapped on your notification. | [optional] 
**Received** | Pointer to **NullableInt32** | Confirmed Deliveries number of devices that received the push notification. Paid Feature Only. Free accounts will see 0. | [optional] 
**Outcomes** | Pointer to [**[]OutcomeData**](OutcomeData.md) |  | [optional] 
**Remaining** | Pointer to **int32** | Number of notifications that have not been sent out yet. This can mean either our system is still processing the notification or you have delayed options set. | [optional] 
**QueuedAt** | Pointer to **int64** | Unix timestamp indicating when the notification was created. | [optional] 
**CompletedAt** | Pointer to **NullableInt64** | Unix timestamp indicating when notification delivery completed. The delivery duration from start to finish can be calculated with completed_at - send_after. | [optional] 
**PlatformDeliveryStats** | Pointer to [**PlatformDeliveryData**](PlatformDeliveryData.md) |  | [optional] 

## Methods

### NewNotificationWithMeta

`func NewNotificationWithMeta(appId string, ) *NotificationWithMeta`

NewNotificationWithMeta instantiates a new NotificationWithMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithMetaWithDefaults

`func NewNotificationWithMetaWithDefaults() *NotificationWithMeta`

NewNotificationWithMetaWithDefaults instantiates a new NotificationWithMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedSegments

`func (o *NotificationWithMeta) GetIncludedSegments() []string`

GetIncludedSegments returns the IncludedSegments field if non-nil, zero value otherwise.

### GetIncludedSegmentsOk

`func (o *NotificationWithMeta) GetIncludedSegmentsOk() (*[]string, bool)`

GetIncludedSegmentsOk returns a tuple with the IncludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegments

`func (o *NotificationWithMeta) SetIncludedSegments(v []string)`

SetIncludedSegments sets IncludedSegments field to given value.

### HasIncludedSegments

`func (o *NotificationWithMeta) HasIncludedSegments() bool`

HasIncludedSegments returns a boolean if a field has been set.

### GetExcludedSegments

`func (o *NotificationWithMeta) GetExcludedSegments() []string`

GetExcludedSegments returns the ExcludedSegments field if non-nil, zero value otherwise.

### GetExcludedSegmentsOk

`func (o *NotificationWithMeta) GetExcludedSegmentsOk() (*[]string, bool)`

GetExcludedSegmentsOk returns a tuple with the ExcludedSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegments

`func (o *NotificationWithMeta) SetExcludedSegments(v []string)`

SetExcludedSegments sets ExcludedSegments field to given value.

### HasExcludedSegments

`func (o *NotificationWithMeta) HasExcludedSegments() bool`

HasExcludedSegments returns a boolean if a field has been set.

### GetLastSession

`func (o *NotificationWithMeta) GetLastSession() string`

GetLastSession returns the LastSession field if non-nil, zero value otherwise.

### GetLastSessionOk

`func (o *NotificationWithMeta) GetLastSessionOk() (*string, bool)`

GetLastSessionOk returns a tuple with the LastSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSession

`func (o *NotificationWithMeta) SetLastSession(v string)`

SetLastSession sets LastSession field to given value.

### HasLastSession

`func (o *NotificationWithMeta) HasLastSession() bool`

HasLastSession returns a boolean if a field has been set.

### GetFirstSession

`func (o *NotificationWithMeta) GetFirstSession() string`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *NotificationWithMeta) GetFirstSessionOk() (*string, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSession

`func (o *NotificationWithMeta) SetFirstSession(v string)`

SetFirstSession sets FirstSession field to given value.

### HasFirstSession

`func (o *NotificationWithMeta) HasFirstSession() bool`

HasFirstSession returns a boolean if a field has been set.

### GetSessionCount

`func (o *NotificationWithMeta) GetSessionCount() string`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *NotificationWithMeta) GetSessionCountOk() (*string, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *NotificationWithMeta) SetSessionCount(v string)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *NotificationWithMeta) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetSessionTime

`func (o *NotificationWithMeta) GetSessionTime() string`

GetSessionTime returns the SessionTime field if non-nil, zero value otherwise.

### GetSessionTimeOk

`func (o *NotificationWithMeta) GetSessionTimeOk() (*string, bool)`

GetSessionTimeOk returns a tuple with the SessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTime

`func (o *NotificationWithMeta) SetSessionTime(v string)`

SetSessionTime sets SessionTime field to given value.

### HasSessionTime

`func (o *NotificationWithMeta) HasSessionTime() bool`

HasSessionTime returns a boolean if a field has been set.

### GetAmountSpent

`func (o *NotificationWithMeta) GetAmountSpent() string`

GetAmountSpent returns the AmountSpent field if non-nil, zero value otherwise.

### GetAmountSpentOk

`func (o *NotificationWithMeta) GetAmountSpentOk() (*string, bool)`

GetAmountSpentOk returns a tuple with the AmountSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSpent

`func (o *NotificationWithMeta) SetAmountSpent(v string)`

SetAmountSpent sets AmountSpent field to given value.

### HasAmountSpent

`func (o *NotificationWithMeta) HasAmountSpent() bool`

HasAmountSpent returns a boolean if a field has been set.

### GetBoughtSku

`func (o *NotificationWithMeta) GetBoughtSku() string`

GetBoughtSku returns the BoughtSku field if non-nil, zero value otherwise.

### GetBoughtSkuOk

`func (o *NotificationWithMeta) GetBoughtSkuOk() (*string, bool)`

GetBoughtSkuOk returns a tuple with the BoughtSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoughtSku

`func (o *NotificationWithMeta) SetBoughtSku(v string)`

SetBoughtSku sets BoughtSku field to given value.

### HasBoughtSku

`func (o *NotificationWithMeta) HasBoughtSku() bool`

HasBoughtSku returns a boolean if a field has been set.

### GetTag

`func (o *NotificationWithMeta) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *NotificationWithMeta) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *NotificationWithMeta) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *NotificationWithMeta) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetLanguage

`func (o *NotificationWithMeta) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *NotificationWithMeta) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *NotificationWithMeta) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *NotificationWithMeta) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAppVersion

`func (o *NotificationWithMeta) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *NotificationWithMeta) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *NotificationWithMeta) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *NotificationWithMeta) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetLocation

`func (o *NotificationWithMeta) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NotificationWithMeta) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NotificationWithMeta) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NotificationWithMeta) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetEmail

`func (o *NotificationWithMeta) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NotificationWithMeta) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NotificationWithMeta) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NotificationWithMeta) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCountry

`func (o *NotificationWithMeta) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *NotificationWithMeta) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *NotificationWithMeta) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *NotificationWithMeta) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIncludePlayerIds

`func (o *NotificationWithMeta) GetIncludePlayerIds() []string`

GetIncludePlayerIds returns the IncludePlayerIds field if non-nil, zero value otherwise.

### GetIncludePlayerIdsOk

`func (o *NotificationWithMeta) GetIncludePlayerIdsOk() (*[]string, bool)`

GetIncludePlayerIdsOk returns a tuple with the IncludePlayerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePlayerIds

`func (o *NotificationWithMeta) SetIncludePlayerIds(v []string)`

SetIncludePlayerIds sets IncludePlayerIds field to given value.

### HasIncludePlayerIds

`func (o *NotificationWithMeta) HasIncludePlayerIds() bool`

HasIncludePlayerIds returns a boolean if a field has been set.

### GetIncludeExternalUserIds

`func (o *NotificationWithMeta) GetIncludeExternalUserIds() []string`

GetIncludeExternalUserIds returns the IncludeExternalUserIds field if non-nil, zero value otherwise.

### GetIncludeExternalUserIdsOk

`func (o *NotificationWithMeta) GetIncludeExternalUserIdsOk() (*[]string, bool)`

GetIncludeExternalUserIdsOk returns a tuple with the IncludeExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExternalUserIds

`func (o *NotificationWithMeta) SetIncludeExternalUserIds(v []string)`

SetIncludeExternalUserIds sets IncludeExternalUserIds field to given value.

### HasIncludeExternalUserIds

`func (o *NotificationWithMeta) HasIncludeExternalUserIds() bool`

HasIncludeExternalUserIds returns a boolean if a field has been set.

### GetIncludeEmailTokens

`func (o *NotificationWithMeta) GetIncludeEmailTokens() []string`

GetIncludeEmailTokens returns the IncludeEmailTokens field if non-nil, zero value otherwise.

### GetIncludeEmailTokensOk

`func (o *NotificationWithMeta) GetIncludeEmailTokensOk() (*[]string, bool)`

GetIncludeEmailTokensOk returns a tuple with the IncludeEmailTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmailTokens

`func (o *NotificationWithMeta) SetIncludeEmailTokens(v []string)`

SetIncludeEmailTokens sets IncludeEmailTokens field to given value.

### HasIncludeEmailTokens

`func (o *NotificationWithMeta) HasIncludeEmailTokens() bool`

HasIncludeEmailTokens returns a boolean if a field has been set.

### GetIncludePhoneNumbers

`func (o *NotificationWithMeta) GetIncludePhoneNumbers() []string`

GetIncludePhoneNumbers returns the IncludePhoneNumbers field if non-nil, zero value otherwise.

### GetIncludePhoneNumbersOk

`func (o *NotificationWithMeta) GetIncludePhoneNumbersOk() (*[]string, bool)`

GetIncludePhoneNumbersOk returns a tuple with the IncludePhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePhoneNumbers

`func (o *NotificationWithMeta) SetIncludePhoneNumbers(v []string)`

SetIncludePhoneNumbers sets IncludePhoneNumbers field to given value.

### HasIncludePhoneNumbers

`func (o *NotificationWithMeta) HasIncludePhoneNumbers() bool`

HasIncludePhoneNumbers returns a boolean if a field has been set.

### GetIncludeIosTokens

`func (o *NotificationWithMeta) GetIncludeIosTokens() []string`

GetIncludeIosTokens returns the IncludeIosTokens field if non-nil, zero value otherwise.

### GetIncludeIosTokensOk

`func (o *NotificationWithMeta) GetIncludeIosTokensOk() (*[]string, bool)`

GetIncludeIosTokensOk returns a tuple with the IncludeIosTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIosTokens

`func (o *NotificationWithMeta) SetIncludeIosTokens(v []string)`

SetIncludeIosTokens sets IncludeIosTokens field to given value.

### HasIncludeIosTokens

`func (o *NotificationWithMeta) HasIncludeIosTokens() bool`

HasIncludeIosTokens returns a boolean if a field has been set.

### GetIncludeWpWnsUris

`func (o *NotificationWithMeta) GetIncludeWpWnsUris() []string`

GetIncludeWpWnsUris returns the IncludeWpWnsUris field if non-nil, zero value otherwise.

### GetIncludeWpWnsUrisOk

`func (o *NotificationWithMeta) GetIncludeWpWnsUrisOk() (*[]string, bool)`

GetIncludeWpWnsUrisOk returns a tuple with the IncludeWpWnsUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeWpWnsUris

`func (o *NotificationWithMeta) SetIncludeWpWnsUris(v []string)`

SetIncludeWpWnsUris sets IncludeWpWnsUris field to given value.

### HasIncludeWpWnsUris

`func (o *NotificationWithMeta) HasIncludeWpWnsUris() bool`

HasIncludeWpWnsUris returns a boolean if a field has been set.

### GetIncludeAmazonRegIds

`func (o *NotificationWithMeta) GetIncludeAmazonRegIds() []string`

GetIncludeAmazonRegIds returns the IncludeAmazonRegIds field if non-nil, zero value otherwise.

### GetIncludeAmazonRegIdsOk

`func (o *NotificationWithMeta) GetIncludeAmazonRegIdsOk() (*[]string, bool)`

GetIncludeAmazonRegIdsOk returns a tuple with the IncludeAmazonRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAmazonRegIds

`func (o *NotificationWithMeta) SetIncludeAmazonRegIds(v []string)`

SetIncludeAmazonRegIds sets IncludeAmazonRegIds field to given value.

### HasIncludeAmazonRegIds

`func (o *NotificationWithMeta) HasIncludeAmazonRegIds() bool`

HasIncludeAmazonRegIds returns a boolean if a field has been set.

### GetIncludeChromeRegIds

`func (o *NotificationWithMeta) GetIncludeChromeRegIds() []string`

GetIncludeChromeRegIds returns the IncludeChromeRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeRegIdsOk

`func (o *NotificationWithMeta) GetIncludeChromeRegIdsOk() (*[]string, bool)`

GetIncludeChromeRegIdsOk returns a tuple with the IncludeChromeRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeRegIds

`func (o *NotificationWithMeta) SetIncludeChromeRegIds(v []string)`

SetIncludeChromeRegIds sets IncludeChromeRegIds field to given value.

### HasIncludeChromeRegIds

`func (o *NotificationWithMeta) HasIncludeChromeRegIds() bool`

HasIncludeChromeRegIds returns a boolean if a field has been set.

### GetIncludeChromeWebRegIds

`func (o *NotificationWithMeta) GetIncludeChromeWebRegIds() []string`

GetIncludeChromeWebRegIds returns the IncludeChromeWebRegIds field if non-nil, zero value otherwise.

### GetIncludeChromeWebRegIdsOk

`func (o *NotificationWithMeta) GetIncludeChromeWebRegIdsOk() (*[]string, bool)`

GetIncludeChromeWebRegIdsOk returns a tuple with the IncludeChromeWebRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChromeWebRegIds

`func (o *NotificationWithMeta) SetIncludeChromeWebRegIds(v []string)`

SetIncludeChromeWebRegIds sets IncludeChromeWebRegIds field to given value.

### HasIncludeChromeWebRegIds

`func (o *NotificationWithMeta) HasIncludeChromeWebRegIds() bool`

HasIncludeChromeWebRegIds returns a boolean if a field has been set.

### GetIncludeAndroidRegIds

`func (o *NotificationWithMeta) GetIncludeAndroidRegIds() []string`

GetIncludeAndroidRegIds returns the IncludeAndroidRegIds field if non-nil, zero value otherwise.

### GetIncludeAndroidRegIdsOk

`func (o *NotificationWithMeta) GetIncludeAndroidRegIdsOk() (*[]string, bool)`

GetIncludeAndroidRegIdsOk returns a tuple with the IncludeAndroidRegIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAndroidRegIds

`func (o *NotificationWithMeta) SetIncludeAndroidRegIds(v []string)`

SetIncludeAndroidRegIds sets IncludeAndroidRegIds field to given value.

### HasIncludeAndroidRegIds

`func (o *NotificationWithMeta) HasIncludeAndroidRegIds() bool`

HasIncludeAndroidRegIds returns a boolean if a field has been set.

### GetId

`func (o *NotificationWithMeta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationWithMeta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationWithMeta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationWithMeta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *NotificationWithMeta) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NotificationWithMeta) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NotificationWithMeta) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *NotificationWithMeta) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAggregation

`func (o *NotificationWithMeta) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *NotificationWithMeta) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *NotificationWithMeta) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *NotificationWithMeta) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetIsIos

`func (o *NotificationWithMeta) GetIsIos() bool`

GetIsIos returns the IsIos field if non-nil, zero value otherwise.

### GetIsIosOk

`func (o *NotificationWithMeta) GetIsIosOk() (*bool, bool)`

GetIsIosOk returns a tuple with the IsIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIos

`func (o *NotificationWithMeta) SetIsIos(v bool)`

SetIsIos sets IsIos field to given value.

### HasIsIos

`func (o *NotificationWithMeta) HasIsIos() bool`

HasIsIos returns a boolean if a field has been set.

### SetIsIosNil

`func (o *NotificationWithMeta) SetIsIosNil(b bool)`

 SetIsIosNil sets the value for IsIos to be an explicit nil

### UnsetIsIos
`func (o *NotificationWithMeta) UnsetIsIos()`

UnsetIsIos ensures that no value is present for IsIos, not even an explicit nil
### GetIsAndroid

`func (o *NotificationWithMeta) GetIsAndroid() bool`

GetIsAndroid returns the IsAndroid field if non-nil, zero value otherwise.

### GetIsAndroidOk

`func (o *NotificationWithMeta) GetIsAndroidOk() (*bool, bool)`

GetIsAndroidOk returns a tuple with the IsAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAndroid

`func (o *NotificationWithMeta) SetIsAndroid(v bool)`

SetIsAndroid sets IsAndroid field to given value.

### HasIsAndroid

`func (o *NotificationWithMeta) HasIsAndroid() bool`

HasIsAndroid returns a boolean if a field has been set.

### SetIsAndroidNil

`func (o *NotificationWithMeta) SetIsAndroidNil(b bool)`

 SetIsAndroidNil sets the value for IsAndroid to be an explicit nil

### UnsetIsAndroid
`func (o *NotificationWithMeta) UnsetIsAndroid()`

UnsetIsAndroid ensures that no value is present for IsAndroid, not even an explicit nil
### GetIsHuawei

`func (o *NotificationWithMeta) GetIsHuawei() bool`

GetIsHuawei returns the IsHuawei field if non-nil, zero value otherwise.

### GetIsHuaweiOk

`func (o *NotificationWithMeta) GetIsHuaweiOk() (*bool, bool)`

GetIsHuaweiOk returns a tuple with the IsHuawei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHuawei

`func (o *NotificationWithMeta) SetIsHuawei(v bool)`

SetIsHuawei sets IsHuawei field to given value.

### HasIsHuawei

`func (o *NotificationWithMeta) HasIsHuawei() bool`

HasIsHuawei returns a boolean if a field has been set.

### SetIsHuaweiNil

`func (o *NotificationWithMeta) SetIsHuaweiNil(b bool)`

 SetIsHuaweiNil sets the value for IsHuawei to be an explicit nil

### UnsetIsHuawei
`func (o *NotificationWithMeta) UnsetIsHuawei()`

UnsetIsHuawei ensures that no value is present for IsHuawei, not even an explicit nil
### GetIsAnyWeb

`func (o *NotificationWithMeta) GetIsAnyWeb() bool`

GetIsAnyWeb returns the IsAnyWeb field if non-nil, zero value otherwise.

### GetIsAnyWebOk

`func (o *NotificationWithMeta) GetIsAnyWebOk() (*bool, bool)`

GetIsAnyWebOk returns a tuple with the IsAnyWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnyWeb

`func (o *NotificationWithMeta) SetIsAnyWeb(v bool)`

SetIsAnyWeb sets IsAnyWeb field to given value.

### HasIsAnyWeb

`func (o *NotificationWithMeta) HasIsAnyWeb() bool`

HasIsAnyWeb returns a boolean if a field has been set.

### SetIsAnyWebNil

`func (o *NotificationWithMeta) SetIsAnyWebNil(b bool)`

 SetIsAnyWebNil sets the value for IsAnyWeb to be an explicit nil

### UnsetIsAnyWeb
`func (o *NotificationWithMeta) UnsetIsAnyWeb()`

UnsetIsAnyWeb ensures that no value is present for IsAnyWeb, not even an explicit nil
### GetIsChromeWeb

`func (o *NotificationWithMeta) GetIsChromeWeb() bool`

GetIsChromeWeb returns the IsChromeWeb field if non-nil, zero value otherwise.

### GetIsChromeWebOk

`func (o *NotificationWithMeta) GetIsChromeWebOk() (*bool, bool)`

GetIsChromeWebOk returns a tuple with the IsChromeWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChromeWeb

`func (o *NotificationWithMeta) SetIsChromeWeb(v bool)`

SetIsChromeWeb sets IsChromeWeb field to given value.

### HasIsChromeWeb

`func (o *NotificationWithMeta) HasIsChromeWeb() bool`

HasIsChromeWeb returns a boolean if a field has been set.

### SetIsChromeWebNil

`func (o *NotificationWithMeta) SetIsChromeWebNil(b bool)`

 SetIsChromeWebNil sets the value for IsChromeWeb to be an explicit nil

### UnsetIsChromeWeb
`func (o *NotificationWithMeta) UnsetIsChromeWeb()`

UnsetIsChromeWeb ensures that no value is present for IsChromeWeb, not even an explicit nil
### GetIsFirefox

`func (o *NotificationWithMeta) GetIsFirefox() bool`

GetIsFirefox returns the IsFirefox field if non-nil, zero value otherwise.

### GetIsFirefoxOk

`func (o *NotificationWithMeta) GetIsFirefoxOk() (*bool, bool)`

GetIsFirefoxOk returns a tuple with the IsFirefox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirefox

`func (o *NotificationWithMeta) SetIsFirefox(v bool)`

SetIsFirefox sets IsFirefox field to given value.

### HasIsFirefox

`func (o *NotificationWithMeta) HasIsFirefox() bool`

HasIsFirefox returns a boolean if a field has been set.

### SetIsFirefoxNil

`func (o *NotificationWithMeta) SetIsFirefoxNil(b bool)`

 SetIsFirefoxNil sets the value for IsFirefox to be an explicit nil

### UnsetIsFirefox
`func (o *NotificationWithMeta) UnsetIsFirefox()`

UnsetIsFirefox ensures that no value is present for IsFirefox, not even an explicit nil
### GetIsSafari

`func (o *NotificationWithMeta) GetIsSafari() bool`

GetIsSafari returns the IsSafari field if non-nil, zero value otherwise.

### GetIsSafariOk

`func (o *NotificationWithMeta) GetIsSafariOk() (*bool, bool)`

GetIsSafariOk returns a tuple with the IsSafari field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafari

`func (o *NotificationWithMeta) SetIsSafari(v bool)`

SetIsSafari sets IsSafari field to given value.

### HasIsSafari

`func (o *NotificationWithMeta) HasIsSafari() bool`

HasIsSafari returns a boolean if a field has been set.

### SetIsSafariNil

`func (o *NotificationWithMeta) SetIsSafariNil(b bool)`

 SetIsSafariNil sets the value for IsSafari to be an explicit nil

### UnsetIsSafari
`func (o *NotificationWithMeta) UnsetIsSafari()`

UnsetIsSafari ensures that no value is present for IsSafari, not even an explicit nil
### GetIsWPWNS

`func (o *NotificationWithMeta) GetIsWPWNS() bool`

GetIsWPWNS returns the IsWPWNS field if non-nil, zero value otherwise.

### GetIsWPWNSOk

`func (o *NotificationWithMeta) GetIsWPWNSOk() (*bool, bool)`

GetIsWPWNSOk returns a tuple with the IsWPWNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWPWNS

`func (o *NotificationWithMeta) SetIsWPWNS(v bool)`

SetIsWPWNS sets IsWPWNS field to given value.

### HasIsWPWNS

`func (o *NotificationWithMeta) HasIsWPWNS() bool`

HasIsWPWNS returns a boolean if a field has been set.

### SetIsWPWNSNil

`func (o *NotificationWithMeta) SetIsWPWNSNil(b bool)`

 SetIsWPWNSNil sets the value for IsWPWNS to be an explicit nil

### UnsetIsWPWNS
`func (o *NotificationWithMeta) UnsetIsWPWNS()`

UnsetIsWPWNS ensures that no value is present for IsWPWNS, not even an explicit nil
### GetIsAdm

`func (o *NotificationWithMeta) GetIsAdm() bool`

GetIsAdm returns the IsAdm field if non-nil, zero value otherwise.

### GetIsAdmOk

`func (o *NotificationWithMeta) GetIsAdmOk() (*bool, bool)`

GetIsAdmOk returns a tuple with the IsAdm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdm

`func (o *NotificationWithMeta) SetIsAdm(v bool)`

SetIsAdm sets IsAdm field to given value.

### HasIsAdm

`func (o *NotificationWithMeta) HasIsAdm() bool`

HasIsAdm returns a boolean if a field has been set.

### SetIsAdmNil

`func (o *NotificationWithMeta) SetIsAdmNil(b bool)`

 SetIsAdmNil sets the value for IsAdm to be an explicit nil

### UnsetIsAdm
`func (o *NotificationWithMeta) UnsetIsAdm()`

UnsetIsAdm ensures that no value is present for IsAdm, not even an explicit nil
### GetIsChrome

`func (o *NotificationWithMeta) GetIsChrome() bool`

GetIsChrome returns the IsChrome field if non-nil, zero value otherwise.

### GetIsChromeOk

`func (o *NotificationWithMeta) GetIsChromeOk() (*bool, bool)`

GetIsChromeOk returns a tuple with the IsChrome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChrome

`func (o *NotificationWithMeta) SetIsChrome(v bool)`

SetIsChrome sets IsChrome field to given value.

### HasIsChrome

`func (o *NotificationWithMeta) HasIsChrome() bool`

HasIsChrome returns a boolean if a field has been set.

### SetIsChromeNil

`func (o *NotificationWithMeta) SetIsChromeNil(b bool)`

 SetIsChromeNil sets the value for IsChrome to be an explicit nil

### UnsetIsChrome
`func (o *NotificationWithMeta) UnsetIsChrome()`

UnsetIsChrome ensures that no value is present for IsChrome, not even an explicit nil
### GetChannelForExternalUserIds

`func (o *NotificationWithMeta) GetChannelForExternalUserIds() string`

GetChannelForExternalUserIds returns the ChannelForExternalUserIds field if non-nil, zero value otherwise.

### GetChannelForExternalUserIdsOk

`func (o *NotificationWithMeta) GetChannelForExternalUserIdsOk() (*string, bool)`

GetChannelForExternalUserIdsOk returns a tuple with the ChannelForExternalUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelForExternalUserIds

`func (o *NotificationWithMeta) SetChannelForExternalUserIds(v string)`

SetChannelForExternalUserIds sets ChannelForExternalUserIds field to given value.

### HasChannelForExternalUserIds

`func (o *NotificationWithMeta) HasChannelForExternalUserIds() bool`

HasChannelForExternalUserIds returns a boolean if a field has been set.

### GetAppId

`func (o *NotificationWithMeta) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *NotificationWithMeta) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *NotificationWithMeta) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetExternalId

`func (o *NotificationWithMeta) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NotificationWithMeta) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NotificationWithMeta) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NotificationWithMeta) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetContents

`func (o *NotificationWithMeta) GetContents() StringMap`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *NotificationWithMeta) GetContentsOk() (*StringMap, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *NotificationWithMeta) SetContents(v StringMap)`

SetContents sets Contents field to given value.

### HasContents

`func (o *NotificationWithMeta) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *NotificationWithMeta) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *NotificationWithMeta) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetHeadings

`func (o *NotificationWithMeta) GetHeadings() StringMap`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *NotificationWithMeta) GetHeadingsOk() (*StringMap, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *NotificationWithMeta) SetHeadings(v StringMap)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *NotificationWithMeta) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### SetHeadingsNil

`func (o *NotificationWithMeta) SetHeadingsNil(b bool)`

 SetHeadingsNil sets the value for Headings to be an explicit nil

### UnsetHeadings
`func (o *NotificationWithMeta) UnsetHeadings()`

UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
### GetSubtitle

`func (o *NotificationWithMeta) GetSubtitle() StringMap`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *NotificationWithMeta) GetSubtitleOk() (*StringMap, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *NotificationWithMeta) SetSubtitle(v StringMap)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *NotificationWithMeta) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *NotificationWithMeta) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *NotificationWithMeta) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetData

`func (o *NotificationWithMeta) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationWithMeta) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationWithMeta) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NotificationWithMeta) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHuaweiMsgType

`func (o *NotificationWithMeta) GetHuaweiMsgType() string`

GetHuaweiMsgType returns the HuaweiMsgType field if non-nil, zero value otherwise.

### GetHuaweiMsgTypeOk

`func (o *NotificationWithMeta) GetHuaweiMsgTypeOk() (*string, bool)`

GetHuaweiMsgTypeOk returns a tuple with the HuaweiMsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiMsgType

`func (o *NotificationWithMeta) SetHuaweiMsgType(v string)`

SetHuaweiMsgType sets HuaweiMsgType field to given value.

### HasHuaweiMsgType

`func (o *NotificationWithMeta) HasHuaweiMsgType() bool`

HasHuaweiMsgType returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationWithMeta) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationWithMeta) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationWithMeta) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationWithMeta) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebUrl

`func (o *NotificationWithMeta) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *NotificationWithMeta) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *NotificationWithMeta) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *NotificationWithMeta) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### GetAppUrl

`func (o *NotificationWithMeta) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *NotificationWithMeta) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *NotificationWithMeta) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.

### HasAppUrl

`func (o *NotificationWithMeta) HasAppUrl() bool`

HasAppUrl returns a boolean if a field has been set.

### GetIosAttachments

`func (o *NotificationWithMeta) GetIosAttachments() map[string]interface{}`

GetIosAttachments returns the IosAttachments field if non-nil, zero value otherwise.

### GetIosAttachmentsOk

`func (o *NotificationWithMeta) GetIosAttachmentsOk() (*map[string]interface{}, bool)`

GetIosAttachmentsOk returns a tuple with the IosAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosAttachments

`func (o *NotificationWithMeta) SetIosAttachments(v map[string]interface{})`

SetIosAttachments sets IosAttachments field to given value.

### HasIosAttachments

`func (o *NotificationWithMeta) HasIosAttachments() bool`

HasIosAttachments returns a boolean if a field has been set.

### GetTemplateId

`func (o *NotificationWithMeta) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *NotificationWithMeta) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *NotificationWithMeta) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *NotificationWithMeta) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetContentAvailable

`func (o *NotificationWithMeta) GetContentAvailable() bool`

GetContentAvailable returns the ContentAvailable field if non-nil, zero value otherwise.

### GetContentAvailableOk

`func (o *NotificationWithMeta) GetContentAvailableOk() (*bool, bool)`

GetContentAvailableOk returns a tuple with the ContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAvailable

`func (o *NotificationWithMeta) SetContentAvailable(v bool)`

SetContentAvailable sets ContentAvailable field to given value.

### HasContentAvailable

`func (o *NotificationWithMeta) HasContentAvailable() bool`

HasContentAvailable returns a boolean if a field has been set.

### SetContentAvailableNil

`func (o *NotificationWithMeta) SetContentAvailableNil(b bool)`

 SetContentAvailableNil sets the value for ContentAvailable to be an explicit nil

### UnsetContentAvailable
`func (o *NotificationWithMeta) UnsetContentAvailable()`

UnsetContentAvailable ensures that no value is present for ContentAvailable, not even an explicit nil
### GetMutableContent

`func (o *NotificationWithMeta) GetMutableContent() bool`

GetMutableContent returns the MutableContent field if non-nil, zero value otherwise.

### GetMutableContentOk

`func (o *NotificationWithMeta) GetMutableContentOk() (*bool, bool)`

GetMutableContentOk returns a tuple with the MutableContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutableContent

`func (o *NotificationWithMeta) SetMutableContent(v bool)`

SetMutableContent sets MutableContent field to given value.

### HasMutableContent

`func (o *NotificationWithMeta) HasMutableContent() bool`

HasMutableContent returns a boolean if a field has been set.

### GetTargetContentIdentifier

`func (o *NotificationWithMeta) GetTargetContentIdentifier() string`

GetTargetContentIdentifier returns the TargetContentIdentifier field if non-nil, zero value otherwise.

### GetTargetContentIdentifierOk

`func (o *NotificationWithMeta) GetTargetContentIdentifierOk() (*string, bool)`

GetTargetContentIdentifierOk returns a tuple with the TargetContentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContentIdentifier

`func (o *NotificationWithMeta) SetTargetContentIdentifier(v string)`

SetTargetContentIdentifier sets TargetContentIdentifier field to given value.

### HasTargetContentIdentifier

`func (o *NotificationWithMeta) HasTargetContentIdentifier() bool`

HasTargetContentIdentifier returns a boolean if a field has been set.

### GetBigPicture

`func (o *NotificationWithMeta) GetBigPicture() string`

GetBigPicture returns the BigPicture field if non-nil, zero value otherwise.

### GetBigPictureOk

`func (o *NotificationWithMeta) GetBigPictureOk() (*string, bool)`

GetBigPictureOk returns a tuple with the BigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigPicture

`func (o *NotificationWithMeta) SetBigPicture(v string)`

SetBigPicture sets BigPicture field to given value.

### HasBigPicture

`func (o *NotificationWithMeta) HasBigPicture() bool`

HasBigPicture returns a boolean if a field has been set.

### GetHuaweiBigPicture

`func (o *NotificationWithMeta) GetHuaweiBigPicture() string`

GetHuaweiBigPicture returns the HuaweiBigPicture field if non-nil, zero value otherwise.

### GetHuaweiBigPictureOk

`func (o *NotificationWithMeta) GetHuaweiBigPictureOk() (*string, bool)`

GetHuaweiBigPictureOk returns a tuple with the HuaweiBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiBigPicture

`func (o *NotificationWithMeta) SetHuaweiBigPicture(v string)`

SetHuaweiBigPicture sets HuaweiBigPicture field to given value.

### HasHuaweiBigPicture

`func (o *NotificationWithMeta) HasHuaweiBigPicture() bool`

HasHuaweiBigPicture returns a boolean if a field has been set.

### GetAdmBigPicture

`func (o *NotificationWithMeta) GetAdmBigPicture() string`

GetAdmBigPicture returns the AdmBigPicture field if non-nil, zero value otherwise.

### GetAdmBigPictureOk

`func (o *NotificationWithMeta) GetAdmBigPictureOk() (*string, bool)`

GetAdmBigPictureOk returns a tuple with the AdmBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmBigPicture

`func (o *NotificationWithMeta) SetAdmBigPicture(v string)`

SetAdmBigPicture sets AdmBigPicture field to given value.

### HasAdmBigPicture

`func (o *NotificationWithMeta) HasAdmBigPicture() bool`

HasAdmBigPicture returns a boolean if a field has been set.

### GetChromeBigPicture

`func (o *NotificationWithMeta) GetChromeBigPicture() string`

GetChromeBigPicture returns the ChromeBigPicture field if non-nil, zero value otherwise.

### GetChromeBigPictureOk

`func (o *NotificationWithMeta) GetChromeBigPictureOk() (*string, bool)`

GetChromeBigPictureOk returns a tuple with the ChromeBigPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeBigPicture

`func (o *NotificationWithMeta) SetChromeBigPicture(v string)`

SetChromeBigPicture sets ChromeBigPicture field to given value.

### HasChromeBigPicture

`func (o *NotificationWithMeta) HasChromeBigPicture() bool`

HasChromeBigPicture returns a boolean if a field has been set.

### GetChromeWebImage

`func (o *NotificationWithMeta) GetChromeWebImage() string`

GetChromeWebImage returns the ChromeWebImage field if non-nil, zero value otherwise.

### GetChromeWebImageOk

`func (o *NotificationWithMeta) GetChromeWebImageOk() (*string, bool)`

GetChromeWebImageOk returns a tuple with the ChromeWebImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebImage

`func (o *NotificationWithMeta) SetChromeWebImage(v string)`

SetChromeWebImage sets ChromeWebImage field to given value.

### HasChromeWebImage

`func (o *NotificationWithMeta) HasChromeWebImage() bool`

HasChromeWebImage returns a boolean if a field has been set.

### GetButtons

`func (o *NotificationWithMeta) GetButtons() []Button`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *NotificationWithMeta) GetButtonsOk() (*[]Button, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *NotificationWithMeta) SetButtons(v []Button)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *NotificationWithMeta) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetWebButtons

`func (o *NotificationWithMeta) GetWebButtons() []Button`

GetWebButtons returns the WebButtons field if non-nil, zero value otherwise.

### GetWebButtonsOk

`func (o *NotificationWithMeta) GetWebButtonsOk() (*[]Button, bool)`

GetWebButtonsOk returns a tuple with the WebButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebButtons

`func (o *NotificationWithMeta) SetWebButtons(v []Button)`

SetWebButtons sets WebButtons field to given value.

### HasWebButtons

`func (o *NotificationWithMeta) HasWebButtons() bool`

HasWebButtons returns a boolean if a field has been set.

### GetIosCategory

`func (o *NotificationWithMeta) GetIosCategory() string`

GetIosCategory returns the IosCategory field if non-nil, zero value otherwise.

### GetIosCategoryOk

`func (o *NotificationWithMeta) GetIosCategoryOk() (*string, bool)`

GetIosCategoryOk returns a tuple with the IosCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosCategory

`func (o *NotificationWithMeta) SetIosCategory(v string)`

SetIosCategory sets IosCategory field to given value.

### HasIosCategory

`func (o *NotificationWithMeta) HasIosCategory() bool`

HasIosCategory returns a boolean if a field has been set.

### GetAndroidChannelId

`func (o *NotificationWithMeta) GetAndroidChannelId() string`

GetAndroidChannelId returns the AndroidChannelId field if non-nil, zero value otherwise.

### GetAndroidChannelIdOk

`func (o *NotificationWithMeta) GetAndroidChannelIdOk() (*string, bool)`

GetAndroidChannelIdOk returns a tuple with the AndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidChannelId

`func (o *NotificationWithMeta) SetAndroidChannelId(v string)`

SetAndroidChannelId sets AndroidChannelId field to given value.

### HasAndroidChannelId

`func (o *NotificationWithMeta) HasAndroidChannelId() bool`

HasAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiChannelId

`func (o *NotificationWithMeta) GetHuaweiChannelId() string`

GetHuaweiChannelId returns the HuaweiChannelId field if non-nil, zero value otherwise.

### GetHuaweiChannelIdOk

`func (o *NotificationWithMeta) GetHuaweiChannelIdOk() (*string, bool)`

GetHuaweiChannelIdOk returns a tuple with the HuaweiChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiChannelId

`func (o *NotificationWithMeta) SetHuaweiChannelId(v string)`

SetHuaweiChannelId sets HuaweiChannelId field to given value.

### HasHuaweiChannelId

`func (o *NotificationWithMeta) HasHuaweiChannelId() bool`

HasHuaweiChannelId returns a boolean if a field has been set.

### GetExistingAndroidChannelId

`func (o *NotificationWithMeta) GetExistingAndroidChannelId() string`

GetExistingAndroidChannelId returns the ExistingAndroidChannelId field if non-nil, zero value otherwise.

### GetExistingAndroidChannelIdOk

`func (o *NotificationWithMeta) GetExistingAndroidChannelIdOk() (*string, bool)`

GetExistingAndroidChannelIdOk returns a tuple with the ExistingAndroidChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAndroidChannelId

`func (o *NotificationWithMeta) SetExistingAndroidChannelId(v string)`

SetExistingAndroidChannelId sets ExistingAndroidChannelId field to given value.

### HasExistingAndroidChannelId

`func (o *NotificationWithMeta) HasExistingAndroidChannelId() bool`

HasExistingAndroidChannelId returns a boolean if a field has been set.

### GetHuaweiExistingChannelId

`func (o *NotificationWithMeta) GetHuaweiExistingChannelId() string`

GetHuaweiExistingChannelId returns the HuaweiExistingChannelId field if non-nil, zero value otherwise.

### GetHuaweiExistingChannelIdOk

`func (o *NotificationWithMeta) GetHuaweiExistingChannelIdOk() (*string, bool)`

GetHuaweiExistingChannelIdOk returns a tuple with the HuaweiExistingChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiExistingChannelId

`func (o *NotificationWithMeta) SetHuaweiExistingChannelId(v string)`

SetHuaweiExistingChannelId sets HuaweiExistingChannelId field to given value.

### HasHuaweiExistingChannelId

`func (o *NotificationWithMeta) HasHuaweiExistingChannelId() bool`

HasHuaweiExistingChannelId returns a boolean if a field has been set.

### GetAndroidBackgroundLayout

`func (o *NotificationWithMeta) GetAndroidBackgroundLayout() NotificationAllOfAndroidBackgroundLayout`

GetAndroidBackgroundLayout returns the AndroidBackgroundLayout field if non-nil, zero value otherwise.

### GetAndroidBackgroundLayoutOk

`func (o *NotificationWithMeta) GetAndroidBackgroundLayoutOk() (*NotificationAllOfAndroidBackgroundLayout, bool)`

GetAndroidBackgroundLayoutOk returns a tuple with the AndroidBackgroundLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidBackgroundLayout

`func (o *NotificationWithMeta) SetAndroidBackgroundLayout(v NotificationAllOfAndroidBackgroundLayout)`

SetAndroidBackgroundLayout sets AndroidBackgroundLayout field to given value.

### HasAndroidBackgroundLayout

`func (o *NotificationWithMeta) HasAndroidBackgroundLayout() bool`

HasAndroidBackgroundLayout returns a boolean if a field has been set.

### GetSmallIcon

`func (o *NotificationWithMeta) GetSmallIcon() string`

GetSmallIcon returns the SmallIcon field if non-nil, zero value otherwise.

### GetSmallIconOk

`func (o *NotificationWithMeta) GetSmallIconOk() (*string, bool)`

GetSmallIconOk returns a tuple with the SmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallIcon

`func (o *NotificationWithMeta) SetSmallIcon(v string)`

SetSmallIcon sets SmallIcon field to given value.

### HasSmallIcon

`func (o *NotificationWithMeta) HasSmallIcon() bool`

HasSmallIcon returns a boolean if a field has been set.

### GetHuaweiSmallIcon

`func (o *NotificationWithMeta) GetHuaweiSmallIcon() string`

GetHuaweiSmallIcon returns the HuaweiSmallIcon field if non-nil, zero value otherwise.

### GetHuaweiSmallIconOk

`func (o *NotificationWithMeta) GetHuaweiSmallIconOk() (*string, bool)`

GetHuaweiSmallIconOk returns a tuple with the HuaweiSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSmallIcon

`func (o *NotificationWithMeta) SetHuaweiSmallIcon(v string)`

SetHuaweiSmallIcon sets HuaweiSmallIcon field to given value.

### HasHuaweiSmallIcon

`func (o *NotificationWithMeta) HasHuaweiSmallIcon() bool`

HasHuaweiSmallIcon returns a boolean if a field has been set.

### GetLargeIcon

`func (o *NotificationWithMeta) GetLargeIcon() string`

GetLargeIcon returns the LargeIcon field if non-nil, zero value otherwise.

### GetLargeIconOk

`func (o *NotificationWithMeta) GetLargeIconOk() (*string, bool)`

GetLargeIconOk returns a tuple with the LargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeIcon

`func (o *NotificationWithMeta) SetLargeIcon(v string)`

SetLargeIcon sets LargeIcon field to given value.

### HasLargeIcon

`func (o *NotificationWithMeta) HasLargeIcon() bool`

HasLargeIcon returns a boolean if a field has been set.

### GetHuaweiLargeIcon

`func (o *NotificationWithMeta) GetHuaweiLargeIcon() string`

GetHuaweiLargeIcon returns the HuaweiLargeIcon field if non-nil, zero value otherwise.

### GetHuaweiLargeIconOk

`func (o *NotificationWithMeta) GetHuaweiLargeIconOk() (*string, bool)`

GetHuaweiLargeIconOk returns a tuple with the HuaweiLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLargeIcon

`func (o *NotificationWithMeta) SetHuaweiLargeIcon(v string)`

SetHuaweiLargeIcon sets HuaweiLargeIcon field to given value.

### HasHuaweiLargeIcon

`func (o *NotificationWithMeta) HasHuaweiLargeIcon() bool`

HasHuaweiLargeIcon returns a boolean if a field has been set.

### GetAdmSmallIcon

`func (o *NotificationWithMeta) GetAdmSmallIcon() string`

GetAdmSmallIcon returns the AdmSmallIcon field if non-nil, zero value otherwise.

### GetAdmSmallIconOk

`func (o *NotificationWithMeta) GetAdmSmallIconOk() (*string, bool)`

GetAdmSmallIconOk returns a tuple with the AdmSmallIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSmallIcon

`func (o *NotificationWithMeta) SetAdmSmallIcon(v string)`

SetAdmSmallIcon sets AdmSmallIcon field to given value.

### HasAdmSmallIcon

`func (o *NotificationWithMeta) HasAdmSmallIcon() bool`

HasAdmSmallIcon returns a boolean if a field has been set.

### GetAdmLargeIcon

`func (o *NotificationWithMeta) GetAdmLargeIcon() string`

GetAdmLargeIcon returns the AdmLargeIcon field if non-nil, zero value otherwise.

### GetAdmLargeIconOk

`func (o *NotificationWithMeta) GetAdmLargeIconOk() (*string, bool)`

GetAdmLargeIconOk returns a tuple with the AdmLargeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmLargeIcon

`func (o *NotificationWithMeta) SetAdmLargeIcon(v string)`

SetAdmLargeIcon sets AdmLargeIcon field to given value.

### HasAdmLargeIcon

`func (o *NotificationWithMeta) HasAdmLargeIcon() bool`

HasAdmLargeIcon returns a boolean if a field has been set.

### GetChromeWebIcon

`func (o *NotificationWithMeta) GetChromeWebIcon() string`

GetChromeWebIcon returns the ChromeWebIcon field if non-nil, zero value otherwise.

### GetChromeWebIconOk

`func (o *NotificationWithMeta) GetChromeWebIconOk() (*string, bool)`

GetChromeWebIconOk returns a tuple with the ChromeWebIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebIcon

`func (o *NotificationWithMeta) SetChromeWebIcon(v string)`

SetChromeWebIcon sets ChromeWebIcon field to given value.

### HasChromeWebIcon

`func (o *NotificationWithMeta) HasChromeWebIcon() bool`

HasChromeWebIcon returns a boolean if a field has been set.

### GetChromeWebBadge

`func (o *NotificationWithMeta) GetChromeWebBadge() string`

GetChromeWebBadge returns the ChromeWebBadge field if non-nil, zero value otherwise.

### GetChromeWebBadgeOk

`func (o *NotificationWithMeta) GetChromeWebBadgeOk() (*string, bool)`

GetChromeWebBadgeOk returns a tuple with the ChromeWebBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeWebBadge

`func (o *NotificationWithMeta) SetChromeWebBadge(v string)`

SetChromeWebBadge sets ChromeWebBadge field to given value.

### HasChromeWebBadge

`func (o *NotificationWithMeta) HasChromeWebBadge() bool`

HasChromeWebBadge returns a boolean if a field has been set.

### GetFirefoxIcon

`func (o *NotificationWithMeta) GetFirefoxIcon() string`

GetFirefoxIcon returns the FirefoxIcon field if non-nil, zero value otherwise.

### GetFirefoxIconOk

`func (o *NotificationWithMeta) GetFirefoxIconOk() (*string, bool)`

GetFirefoxIconOk returns a tuple with the FirefoxIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefoxIcon

`func (o *NotificationWithMeta) SetFirefoxIcon(v string)`

SetFirefoxIcon sets FirefoxIcon field to given value.

### HasFirefoxIcon

`func (o *NotificationWithMeta) HasFirefoxIcon() bool`

HasFirefoxIcon returns a boolean if a field has been set.

### GetChromeIcon

`func (o *NotificationWithMeta) GetChromeIcon() string`

GetChromeIcon returns the ChromeIcon field if non-nil, zero value otherwise.

### GetChromeIconOk

`func (o *NotificationWithMeta) GetChromeIconOk() (*string, bool)`

GetChromeIconOk returns a tuple with the ChromeIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeIcon

`func (o *NotificationWithMeta) SetChromeIcon(v string)`

SetChromeIcon sets ChromeIcon field to given value.

### HasChromeIcon

`func (o *NotificationWithMeta) HasChromeIcon() bool`

HasChromeIcon returns a boolean if a field has been set.

### GetIosSound

`func (o *NotificationWithMeta) GetIosSound() string`

GetIosSound returns the IosSound field if non-nil, zero value otherwise.

### GetIosSoundOk

`func (o *NotificationWithMeta) GetIosSoundOk() (*string, bool)`

GetIosSoundOk returns a tuple with the IosSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosSound

`func (o *NotificationWithMeta) SetIosSound(v string)`

SetIosSound sets IosSound field to given value.

### HasIosSound

`func (o *NotificationWithMeta) HasIosSound() bool`

HasIosSound returns a boolean if a field has been set.

### GetAndroidSound

`func (o *NotificationWithMeta) GetAndroidSound() string`

GetAndroidSound returns the AndroidSound field if non-nil, zero value otherwise.

### GetAndroidSoundOk

`func (o *NotificationWithMeta) GetAndroidSoundOk() (*string, bool)`

GetAndroidSoundOk returns a tuple with the AndroidSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidSound

`func (o *NotificationWithMeta) SetAndroidSound(v string)`

SetAndroidSound sets AndroidSound field to given value.

### HasAndroidSound

`func (o *NotificationWithMeta) HasAndroidSound() bool`

HasAndroidSound returns a boolean if a field has been set.

### GetHuaweiSound

`func (o *NotificationWithMeta) GetHuaweiSound() string`

GetHuaweiSound returns the HuaweiSound field if non-nil, zero value otherwise.

### GetHuaweiSoundOk

`func (o *NotificationWithMeta) GetHuaweiSoundOk() (*string, bool)`

GetHuaweiSoundOk returns a tuple with the HuaweiSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiSound

`func (o *NotificationWithMeta) SetHuaweiSound(v string)`

SetHuaweiSound sets HuaweiSound field to given value.

### HasHuaweiSound

`func (o *NotificationWithMeta) HasHuaweiSound() bool`

HasHuaweiSound returns a boolean if a field has been set.

### GetAdmSound

`func (o *NotificationWithMeta) GetAdmSound() string`

GetAdmSound returns the AdmSound field if non-nil, zero value otherwise.

### GetAdmSoundOk

`func (o *NotificationWithMeta) GetAdmSoundOk() (*string, bool)`

GetAdmSoundOk returns a tuple with the AdmSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmSound

`func (o *NotificationWithMeta) SetAdmSound(v string)`

SetAdmSound sets AdmSound field to given value.

### HasAdmSound

`func (o *NotificationWithMeta) HasAdmSound() bool`

HasAdmSound returns a boolean if a field has been set.

### GetWpWnsSound

`func (o *NotificationWithMeta) GetWpWnsSound() string`

GetWpWnsSound returns the WpWnsSound field if non-nil, zero value otherwise.

### GetWpWnsSoundOk

`func (o *NotificationWithMeta) GetWpWnsSoundOk() (*string, bool)`

GetWpWnsSoundOk returns a tuple with the WpWnsSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpWnsSound

`func (o *NotificationWithMeta) SetWpWnsSound(v string)`

SetWpWnsSound sets WpWnsSound field to given value.

### HasWpWnsSound

`func (o *NotificationWithMeta) HasWpWnsSound() bool`

HasWpWnsSound returns a boolean if a field has been set.

### GetAndroidLedColor

`func (o *NotificationWithMeta) GetAndroidLedColor() string`

GetAndroidLedColor returns the AndroidLedColor field if non-nil, zero value otherwise.

### GetAndroidLedColorOk

`func (o *NotificationWithMeta) GetAndroidLedColorOk() (*string, bool)`

GetAndroidLedColorOk returns a tuple with the AndroidLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidLedColor

`func (o *NotificationWithMeta) SetAndroidLedColor(v string)`

SetAndroidLedColor sets AndroidLedColor field to given value.

### HasAndroidLedColor

`func (o *NotificationWithMeta) HasAndroidLedColor() bool`

HasAndroidLedColor returns a boolean if a field has been set.

### GetHuaweiLedColor

`func (o *NotificationWithMeta) GetHuaweiLedColor() string`

GetHuaweiLedColor returns the HuaweiLedColor field if non-nil, zero value otherwise.

### GetHuaweiLedColorOk

`func (o *NotificationWithMeta) GetHuaweiLedColorOk() (*string, bool)`

GetHuaweiLedColorOk returns a tuple with the HuaweiLedColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiLedColor

`func (o *NotificationWithMeta) SetHuaweiLedColor(v string)`

SetHuaweiLedColor sets HuaweiLedColor field to given value.

### HasHuaweiLedColor

`func (o *NotificationWithMeta) HasHuaweiLedColor() bool`

HasHuaweiLedColor returns a boolean if a field has been set.

### GetAndroidAccentColor

`func (o *NotificationWithMeta) GetAndroidAccentColor() string`

GetAndroidAccentColor returns the AndroidAccentColor field if non-nil, zero value otherwise.

### GetAndroidAccentColorOk

`func (o *NotificationWithMeta) GetAndroidAccentColorOk() (*string, bool)`

GetAndroidAccentColorOk returns a tuple with the AndroidAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidAccentColor

`func (o *NotificationWithMeta) SetAndroidAccentColor(v string)`

SetAndroidAccentColor sets AndroidAccentColor field to given value.

### HasAndroidAccentColor

`func (o *NotificationWithMeta) HasAndroidAccentColor() bool`

HasAndroidAccentColor returns a boolean if a field has been set.

### GetHuaweiAccentColor

`func (o *NotificationWithMeta) GetHuaweiAccentColor() string`

GetHuaweiAccentColor returns the HuaweiAccentColor field if non-nil, zero value otherwise.

### GetHuaweiAccentColorOk

`func (o *NotificationWithMeta) GetHuaweiAccentColorOk() (*string, bool)`

GetHuaweiAccentColorOk returns a tuple with the HuaweiAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiAccentColor

`func (o *NotificationWithMeta) SetHuaweiAccentColor(v string)`

SetHuaweiAccentColor sets HuaweiAccentColor field to given value.

### HasHuaweiAccentColor

`func (o *NotificationWithMeta) HasHuaweiAccentColor() bool`

HasHuaweiAccentColor returns a boolean if a field has been set.

### GetAndroidVisibility

`func (o *NotificationWithMeta) GetAndroidVisibility() int32`

GetAndroidVisibility returns the AndroidVisibility field if non-nil, zero value otherwise.

### GetAndroidVisibilityOk

`func (o *NotificationWithMeta) GetAndroidVisibilityOk() (*int32, bool)`

GetAndroidVisibilityOk returns a tuple with the AndroidVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidVisibility

`func (o *NotificationWithMeta) SetAndroidVisibility(v int32)`

SetAndroidVisibility sets AndroidVisibility field to given value.

### HasAndroidVisibility

`func (o *NotificationWithMeta) HasAndroidVisibility() bool`

HasAndroidVisibility returns a boolean if a field has been set.

### SetAndroidVisibilityNil

`func (o *NotificationWithMeta) SetAndroidVisibilityNil(b bool)`

 SetAndroidVisibilityNil sets the value for AndroidVisibility to be an explicit nil

### UnsetAndroidVisibility
`func (o *NotificationWithMeta) UnsetAndroidVisibility()`

UnsetAndroidVisibility ensures that no value is present for AndroidVisibility, not even an explicit nil
### GetHuaweiVisibility

`func (o *NotificationWithMeta) GetHuaweiVisibility() int32`

GetHuaweiVisibility returns the HuaweiVisibility field if non-nil, zero value otherwise.

### GetHuaweiVisibilityOk

`func (o *NotificationWithMeta) GetHuaweiVisibilityOk() (*int32, bool)`

GetHuaweiVisibilityOk returns a tuple with the HuaweiVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiVisibility

`func (o *NotificationWithMeta) SetHuaweiVisibility(v int32)`

SetHuaweiVisibility sets HuaweiVisibility field to given value.

### HasHuaweiVisibility

`func (o *NotificationWithMeta) HasHuaweiVisibility() bool`

HasHuaweiVisibility returns a boolean if a field has been set.

### SetHuaweiVisibilityNil

`func (o *NotificationWithMeta) SetHuaweiVisibilityNil(b bool)`

 SetHuaweiVisibilityNil sets the value for HuaweiVisibility to be an explicit nil

### UnsetHuaweiVisibility
`func (o *NotificationWithMeta) UnsetHuaweiVisibility()`

UnsetHuaweiVisibility ensures that no value is present for HuaweiVisibility, not even an explicit nil
### GetIosBadgeType

`func (o *NotificationWithMeta) GetIosBadgeType() string`

GetIosBadgeType returns the IosBadgeType field if non-nil, zero value otherwise.

### GetIosBadgeTypeOk

`func (o *NotificationWithMeta) GetIosBadgeTypeOk() (*string, bool)`

GetIosBadgeTypeOk returns a tuple with the IosBadgeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeType

`func (o *NotificationWithMeta) SetIosBadgeType(v string)`

SetIosBadgeType sets IosBadgeType field to given value.

### HasIosBadgeType

`func (o *NotificationWithMeta) HasIosBadgeType() bool`

HasIosBadgeType returns a boolean if a field has been set.

### GetIosBadgeCount

`func (o *NotificationWithMeta) GetIosBadgeCount() int32`

GetIosBadgeCount returns the IosBadgeCount field if non-nil, zero value otherwise.

### GetIosBadgeCountOk

`func (o *NotificationWithMeta) GetIosBadgeCountOk() (*int32, bool)`

GetIosBadgeCountOk returns a tuple with the IosBadgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosBadgeCount

`func (o *NotificationWithMeta) SetIosBadgeCount(v int32)`

SetIosBadgeCount sets IosBadgeCount field to given value.

### HasIosBadgeCount

`func (o *NotificationWithMeta) HasIosBadgeCount() bool`

HasIosBadgeCount returns a boolean if a field has been set.

### SetIosBadgeCountNil

`func (o *NotificationWithMeta) SetIosBadgeCountNil(b bool)`

 SetIosBadgeCountNil sets the value for IosBadgeCount to be an explicit nil

### UnsetIosBadgeCount
`func (o *NotificationWithMeta) UnsetIosBadgeCount()`

UnsetIosBadgeCount ensures that no value is present for IosBadgeCount, not even an explicit nil
### GetCollapseId

`func (o *NotificationWithMeta) GetCollapseId() string`

GetCollapseId returns the CollapseId field if non-nil, zero value otherwise.

### GetCollapseIdOk

`func (o *NotificationWithMeta) GetCollapseIdOk() (*string, bool)`

GetCollapseIdOk returns a tuple with the CollapseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseId

`func (o *NotificationWithMeta) SetCollapseId(v string)`

SetCollapseId sets CollapseId field to given value.

### HasCollapseId

`func (o *NotificationWithMeta) HasCollapseId() bool`

HasCollapseId returns a boolean if a field has been set.

### GetWebPushTopic

`func (o *NotificationWithMeta) GetWebPushTopic() string`

GetWebPushTopic returns the WebPushTopic field if non-nil, zero value otherwise.

### GetWebPushTopicOk

`func (o *NotificationWithMeta) GetWebPushTopicOk() (*string, bool)`

GetWebPushTopicOk returns a tuple with the WebPushTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPushTopic

`func (o *NotificationWithMeta) SetWebPushTopic(v string)`

SetWebPushTopic sets WebPushTopic field to given value.

### HasWebPushTopic

`func (o *NotificationWithMeta) HasWebPushTopic() bool`

HasWebPushTopic returns a boolean if a field has been set.

### GetApnsAlert

`func (o *NotificationWithMeta) GetApnsAlert() map[string]interface{}`

GetApnsAlert returns the ApnsAlert field if non-nil, zero value otherwise.

### GetApnsAlertOk

`func (o *NotificationWithMeta) GetApnsAlertOk() (*map[string]interface{}, bool)`

GetApnsAlertOk returns a tuple with the ApnsAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsAlert

`func (o *NotificationWithMeta) SetApnsAlert(v map[string]interface{})`

SetApnsAlert sets ApnsAlert field to given value.

### HasApnsAlert

`func (o *NotificationWithMeta) HasApnsAlert() bool`

HasApnsAlert returns a boolean if a field has been set.

### GetSendAfter

`func (o *NotificationWithMeta) GetSendAfter() int64`

GetSendAfter returns the SendAfter field if non-nil, zero value otherwise.

### GetSendAfterOk

`func (o *NotificationWithMeta) GetSendAfterOk() (*int64, bool)`

GetSendAfterOk returns a tuple with the SendAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAfter

`func (o *NotificationWithMeta) SetSendAfter(v int64)`

SetSendAfter sets SendAfter field to given value.

### HasSendAfter

`func (o *NotificationWithMeta) HasSendAfter() bool`

HasSendAfter returns a boolean if a field has been set.

### GetDelayedOption

`func (o *NotificationWithMeta) GetDelayedOption() string`

GetDelayedOption returns the DelayedOption field if non-nil, zero value otherwise.

### GetDelayedOptionOk

`func (o *NotificationWithMeta) GetDelayedOptionOk() (*string, bool)`

GetDelayedOptionOk returns a tuple with the DelayedOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedOption

`func (o *NotificationWithMeta) SetDelayedOption(v string)`

SetDelayedOption sets DelayedOption field to given value.

### HasDelayedOption

`func (o *NotificationWithMeta) HasDelayedOption() bool`

HasDelayedOption returns a boolean if a field has been set.

### GetDeliveryTimeOfDay

`func (o *NotificationWithMeta) GetDeliveryTimeOfDay() string`

GetDeliveryTimeOfDay returns the DeliveryTimeOfDay field if non-nil, zero value otherwise.

### GetDeliveryTimeOfDayOk

`func (o *NotificationWithMeta) GetDeliveryTimeOfDayOk() (*string, bool)`

GetDeliveryTimeOfDayOk returns a tuple with the DeliveryTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeOfDay

`func (o *NotificationWithMeta) SetDeliveryTimeOfDay(v string)`

SetDeliveryTimeOfDay sets DeliveryTimeOfDay field to given value.

### HasDeliveryTimeOfDay

`func (o *NotificationWithMeta) HasDeliveryTimeOfDay() bool`

HasDeliveryTimeOfDay returns a boolean if a field has been set.

### GetTtl

`func (o *NotificationWithMeta) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *NotificationWithMeta) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *NotificationWithMeta) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *NotificationWithMeta) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *NotificationWithMeta) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *NotificationWithMeta) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetPriority

`func (o *NotificationWithMeta) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NotificationWithMeta) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NotificationWithMeta) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NotificationWithMeta) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *NotificationWithMeta) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *NotificationWithMeta) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetApnsPushTypeOverride

`func (o *NotificationWithMeta) GetApnsPushTypeOverride() string`

GetApnsPushTypeOverride returns the ApnsPushTypeOverride field if non-nil, zero value otherwise.

### GetApnsPushTypeOverrideOk

`func (o *NotificationWithMeta) GetApnsPushTypeOverrideOk() (*string, bool)`

GetApnsPushTypeOverrideOk returns a tuple with the ApnsPushTypeOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnsPushTypeOverride

`func (o *NotificationWithMeta) SetApnsPushTypeOverride(v string)`

SetApnsPushTypeOverride sets ApnsPushTypeOverride field to given value.

### HasApnsPushTypeOverride

`func (o *NotificationWithMeta) HasApnsPushTypeOverride() bool`

HasApnsPushTypeOverride returns a boolean if a field has been set.

### GetThrottleRatePerMinute

`func (o *NotificationWithMeta) GetThrottleRatePerMinute() int32`

GetThrottleRatePerMinute returns the ThrottleRatePerMinute field if non-nil, zero value otherwise.

### GetThrottleRatePerMinuteOk

`func (o *NotificationWithMeta) GetThrottleRatePerMinuteOk() (*int32, bool)`

GetThrottleRatePerMinuteOk returns a tuple with the ThrottleRatePerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleRatePerMinute

`func (o *NotificationWithMeta) SetThrottleRatePerMinute(v int32)`

SetThrottleRatePerMinute sets ThrottleRatePerMinute field to given value.

### HasThrottleRatePerMinute

`func (o *NotificationWithMeta) HasThrottleRatePerMinute() bool`

HasThrottleRatePerMinute returns a boolean if a field has been set.

### SetThrottleRatePerMinuteNil

`func (o *NotificationWithMeta) SetThrottleRatePerMinuteNil(b bool)`

 SetThrottleRatePerMinuteNil sets the value for ThrottleRatePerMinute to be an explicit nil

### UnsetThrottleRatePerMinute
`func (o *NotificationWithMeta) UnsetThrottleRatePerMinute()`

UnsetThrottleRatePerMinute ensures that no value is present for ThrottleRatePerMinute, not even an explicit nil
### GetAndroidGroup

`func (o *NotificationWithMeta) GetAndroidGroup() string`

GetAndroidGroup returns the AndroidGroup field if non-nil, zero value otherwise.

### GetAndroidGroupOk

`func (o *NotificationWithMeta) GetAndroidGroupOk() (*string, bool)`

GetAndroidGroupOk returns a tuple with the AndroidGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroup

`func (o *NotificationWithMeta) SetAndroidGroup(v string)`

SetAndroidGroup sets AndroidGroup field to given value.

### HasAndroidGroup

`func (o *NotificationWithMeta) HasAndroidGroup() bool`

HasAndroidGroup returns a boolean if a field has been set.

### GetAndroidGroupMessage

`func (o *NotificationWithMeta) GetAndroidGroupMessage() string`

GetAndroidGroupMessage returns the AndroidGroupMessage field if non-nil, zero value otherwise.

### GetAndroidGroupMessageOk

`func (o *NotificationWithMeta) GetAndroidGroupMessageOk() (*string, bool)`

GetAndroidGroupMessageOk returns a tuple with the AndroidGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidGroupMessage

`func (o *NotificationWithMeta) SetAndroidGroupMessage(v string)`

SetAndroidGroupMessage sets AndroidGroupMessage field to given value.

### HasAndroidGroupMessage

`func (o *NotificationWithMeta) HasAndroidGroupMessage() bool`

HasAndroidGroupMessage returns a boolean if a field has been set.

### GetAdmGroup

`func (o *NotificationWithMeta) GetAdmGroup() string`

GetAdmGroup returns the AdmGroup field if non-nil, zero value otherwise.

### GetAdmGroupOk

`func (o *NotificationWithMeta) GetAdmGroupOk() (*string, bool)`

GetAdmGroupOk returns a tuple with the AdmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroup

`func (o *NotificationWithMeta) SetAdmGroup(v string)`

SetAdmGroup sets AdmGroup field to given value.

### HasAdmGroup

`func (o *NotificationWithMeta) HasAdmGroup() bool`

HasAdmGroup returns a boolean if a field has been set.

### GetAdmGroupMessage

`func (o *NotificationWithMeta) GetAdmGroupMessage() map[string]interface{}`

GetAdmGroupMessage returns the AdmGroupMessage field if non-nil, zero value otherwise.

### GetAdmGroupMessageOk

`func (o *NotificationWithMeta) GetAdmGroupMessageOk() (*map[string]interface{}, bool)`

GetAdmGroupMessageOk returns a tuple with the AdmGroupMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmGroupMessage

`func (o *NotificationWithMeta) SetAdmGroupMessage(v map[string]interface{})`

SetAdmGroupMessage sets AdmGroupMessage field to given value.

### HasAdmGroupMessage

`func (o *NotificationWithMeta) HasAdmGroupMessage() bool`

HasAdmGroupMessage returns a boolean if a field has been set.

### GetThreadId

`func (o *NotificationWithMeta) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *NotificationWithMeta) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *NotificationWithMeta) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *NotificationWithMeta) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetSummaryArg

`func (o *NotificationWithMeta) GetSummaryArg() string`

GetSummaryArg returns the SummaryArg field if non-nil, zero value otherwise.

### GetSummaryArgOk

`func (o *NotificationWithMeta) GetSummaryArgOk() (*string, bool)`

GetSummaryArgOk returns a tuple with the SummaryArg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArg

`func (o *NotificationWithMeta) SetSummaryArg(v string)`

SetSummaryArg sets SummaryArg field to given value.

### HasSummaryArg

`func (o *NotificationWithMeta) HasSummaryArg() bool`

HasSummaryArg returns a boolean if a field has been set.

### GetSummaryArgCount

`func (o *NotificationWithMeta) GetSummaryArgCount() int32`

GetSummaryArgCount returns the SummaryArgCount field if non-nil, zero value otherwise.

### GetSummaryArgCountOk

`func (o *NotificationWithMeta) GetSummaryArgCountOk() (*int32, bool)`

GetSummaryArgCountOk returns a tuple with the SummaryArgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryArgCount

`func (o *NotificationWithMeta) SetSummaryArgCount(v int32)`

SetSummaryArgCount sets SummaryArgCount field to given value.

### HasSummaryArgCount

`func (o *NotificationWithMeta) HasSummaryArgCount() bool`

HasSummaryArgCount returns a boolean if a field has been set.

### GetEmailSubject

`func (o *NotificationWithMeta) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *NotificationWithMeta) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *NotificationWithMeta) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *NotificationWithMeta) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetEmailBody

`func (o *NotificationWithMeta) GetEmailBody() string`

GetEmailBody returns the EmailBody field if non-nil, zero value otherwise.

### GetEmailBodyOk

`func (o *NotificationWithMeta) GetEmailBodyOk() (*string, bool)`

GetEmailBodyOk returns a tuple with the EmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailBody

`func (o *NotificationWithMeta) SetEmailBody(v string)`

SetEmailBody sets EmailBody field to given value.

### HasEmailBody

`func (o *NotificationWithMeta) HasEmailBody() bool`

HasEmailBody returns a boolean if a field has been set.

### GetEmailFromName

`func (o *NotificationWithMeta) GetEmailFromName() string`

GetEmailFromName returns the EmailFromName field if non-nil, zero value otherwise.

### GetEmailFromNameOk

`func (o *NotificationWithMeta) GetEmailFromNameOk() (*string, bool)`

GetEmailFromNameOk returns a tuple with the EmailFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromName

`func (o *NotificationWithMeta) SetEmailFromName(v string)`

SetEmailFromName sets EmailFromName field to given value.

### HasEmailFromName

`func (o *NotificationWithMeta) HasEmailFromName() bool`

HasEmailFromName returns a boolean if a field has been set.

### GetEmailFromAddress

`func (o *NotificationWithMeta) GetEmailFromAddress() string`

GetEmailFromAddress returns the EmailFromAddress field if non-nil, zero value otherwise.

### GetEmailFromAddressOk

`func (o *NotificationWithMeta) GetEmailFromAddressOk() (*string, bool)`

GetEmailFromAddressOk returns a tuple with the EmailFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFromAddress

`func (o *NotificationWithMeta) SetEmailFromAddress(v string)`

SetEmailFromAddress sets EmailFromAddress field to given value.

### HasEmailFromAddress

`func (o *NotificationWithMeta) HasEmailFromAddress() bool`

HasEmailFromAddress returns a boolean if a field has been set.

### GetSmsFrom

`func (o *NotificationWithMeta) GetSmsFrom() string`

GetSmsFrom returns the SmsFrom field if non-nil, zero value otherwise.

### GetSmsFromOk

`func (o *NotificationWithMeta) GetSmsFromOk() (*string, bool)`

GetSmsFromOk returns a tuple with the SmsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFrom

`func (o *NotificationWithMeta) SetSmsFrom(v string)`

SetSmsFrom sets SmsFrom field to given value.

### HasSmsFrom

`func (o *NotificationWithMeta) HasSmsFrom() bool`

HasSmsFrom returns a boolean if a field has been set.

### GetSmsMediaUrls

`func (o *NotificationWithMeta) GetSmsMediaUrls() []string`

GetSmsMediaUrls returns the SmsMediaUrls field if non-nil, zero value otherwise.

### GetSmsMediaUrlsOk

`func (o *NotificationWithMeta) GetSmsMediaUrlsOk() (*[]string, bool)`

GetSmsMediaUrlsOk returns a tuple with the SmsMediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMediaUrls

`func (o *NotificationWithMeta) SetSmsMediaUrls(v []string)`

SetSmsMediaUrls sets SmsMediaUrls field to given value.

### HasSmsMediaUrls

`func (o *NotificationWithMeta) HasSmsMediaUrls() bool`

HasSmsMediaUrls returns a boolean if a field has been set.

### GetSuccessful

`func (o *NotificationWithMeta) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *NotificationWithMeta) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *NotificationWithMeta) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *NotificationWithMeta) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetFailed

`func (o *NotificationWithMeta) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *NotificationWithMeta) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *NotificationWithMeta) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *NotificationWithMeta) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetErrored

`func (o *NotificationWithMeta) GetErrored() int32`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *NotificationWithMeta) GetErroredOk() (*int32, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *NotificationWithMeta) SetErrored(v int32)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *NotificationWithMeta) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### GetConverted

`func (o *NotificationWithMeta) GetConverted() int32`

GetConverted returns the Converted field if non-nil, zero value otherwise.

### GetConvertedOk

`func (o *NotificationWithMeta) GetConvertedOk() (*int32, bool)`

GetConvertedOk returns a tuple with the Converted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConverted

`func (o *NotificationWithMeta) SetConverted(v int32)`

SetConverted sets Converted field to given value.

### HasConverted

`func (o *NotificationWithMeta) HasConverted() bool`

HasConverted returns a boolean if a field has been set.

### GetReceived

`func (o *NotificationWithMeta) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *NotificationWithMeta) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *NotificationWithMeta) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *NotificationWithMeta) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### SetReceivedNil

`func (o *NotificationWithMeta) SetReceivedNil(b bool)`

 SetReceivedNil sets the value for Received to be an explicit nil

### UnsetReceived
`func (o *NotificationWithMeta) UnsetReceived()`

UnsetReceived ensures that no value is present for Received, not even an explicit nil
### GetOutcomes

`func (o *NotificationWithMeta) GetOutcomes() []OutcomeData`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *NotificationWithMeta) GetOutcomesOk() (*[]OutcomeData, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *NotificationWithMeta) SetOutcomes(v []OutcomeData)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *NotificationWithMeta) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### GetRemaining

`func (o *NotificationWithMeta) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *NotificationWithMeta) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *NotificationWithMeta) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *NotificationWithMeta) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetQueuedAt

`func (o *NotificationWithMeta) GetQueuedAt() int64`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *NotificationWithMeta) GetQueuedAtOk() (*int64, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *NotificationWithMeta) SetQueuedAt(v int64)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *NotificationWithMeta) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *NotificationWithMeta) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *NotificationWithMeta) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *NotificationWithMeta) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *NotificationWithMeta) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *NotificationWithMeta) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *NotificationWithMeta) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetPlatformDeliveryStats

`func (o *NotificationWithMeta) GetPlatformDeliveryStats() PlatformDeliveryData`

GetPlatformDeliveryStats returns the PlatformDeliveryStats field if non-nil, zero value otherwise.

### GetPlatformDeliveryStatsOk

`func (o *NotificationWithMeta) GetPlatformDeliveryStatsOk() (*PlatformDeliveryData, bool)`

GetPlatformDeliveryStatsOk returns a tuple with the PlatformDeliveryStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformDeliveryStats

`func (o *NotificationWithMeta) SetPlatformDeliveryStats(v PlatformDeliveryData)`

SetPlatformDeliveryStats sets PlatformDeliveryStats field to given value.

### HasPlatformDeliveryStats

`func (o *NotificationWithMeta) HasPlatformDeliveryStats() bool`

HasPlatformDeliveryStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


