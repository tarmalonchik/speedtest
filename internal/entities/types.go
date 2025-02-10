// Package entities
// nolint
//
//go:generate go-enum -f=$GOFILE --nocase --sqlnullint
package entities

// PlatformType ENUM(invalid, ios, android, windows, mac_os)
type PlatformType int64

// ReviewType ENUM(invalid, review_for_current_users, review_for_old_users)
type ReviewType int

// ReviewForCurrentUsersValues ENUM(invalid, good, normal, bad_expensive, bad_difficult_to_connect, bad_difficult_to_pay, bad_vpn_works_bad, bad_support_is_bad, bad_something_else)
type ReviewForCurrentUsersValues int

// ReviewForOldUsers ENUM(invalid, expensive, dont_know_how_to_connect, dont_know_how_to_pay, vpn_bad_quality, dont_need_vpn, something_else)
type ReviewForOldUsers int

// UserEventType ENUM(invalid, demo1DayActivated, demo14DayActivated, subscriptionActivated, subscriptionDeactivated,
// promoActivated, deletedByAdmin, demo3DayActivated, demoActivated, clickedReferralButton)
type UserEventType int

// VpnProtocols ENUM(invalid, wireguard, xray)
type VpnProtocols string

// Actions ENUM(invalid, add, remove)
type Actions int8

// Protocols ENUM(vless_reality, shadowsocks)
type Protocols string

// SubscriptionHealthType ENUM(healthy, will_end_soon, has_ended, never_had_subscription)
type SubscriptionHealthType string

// UserSubscriptionType ENUM(MANUAL, APPLE_IN_APP, GOOGLE_IN_APP)
type UserSubscriptionType string

// UserNotificationType ENUM(TG_MSG, TG_CALL_BACK, TG_FILE, EMAIL)
type UserNotificationType string
