package model

import (
	"time"

	"github.com/sugarshop/conv/timeconv"
)

const (
	ClerkWebhookEvent_USERCREATED = "user.created"
)

type User struct {
	Birthday              string                 `json:"birthday"`
	EmailAddresses        []EmailAddress         `json:"email_addresses"`
	ExternalAccounts      []interface{}          `json:"external_accounts"`
	ExternalID            string                 `json:"external_id"`
	FirstName             string                 `json:"first_name"`
	Gender                string                 `json:"gender"`
	ID                    string                 `json:"id"`
	ImageURL              string                 `json:"image_url"`
	LastName              string                 `json:"last_name"`
	Object                string                 `json:"object"`
	PasswordEnabled       bool                   `json:"password_enabled"`
	PhoneNumbers          []interface{}          `json:"phone_numbers"`
	PrimaryEmailAddressID string                 `json:"primary_email_address_id"`
	PrimaryPhoneNumberID  *string                `json:"primary_phone_number_id"`
	PrimaryWeb3WalletID   *string                `json:"primary_web3_wallet_id"`
	PrivateMetadata       map[string]interface{} `json:"private_metadata"`
	ProfileImageURL       string                 `json:"profile_image_url"`
	PublicMetadata        map[string]interface{} `json:"public_metadata"`
	TwoFactorEnabled      bool                   `json:"two_factor_enabled"`
	UnsafeMetadata        map[string]interface{} `json:"unsafe_metadata"`
	Username              *string                `json:"username"`
	Web3Wallets           []interface{}          `json:"web3_wallets"`
	LastSignInAt          int64                  `json:"last_sign_in_at"`
	UpdatedAt             int64                  `json:"updated_at"`
	CreatedAt             int64                  `json:"created_at"`
}

func (u *User) GetCreatedAt() time.Time {
	return timeconv.Int64ToTime(u.CreatedAt)
}

func (u *User) GetLastSignInAt() time.Time {
	return timeconv.Int64ToTime(u.LastSignInAt)
}

func (u *User) GetUpdatedAt() time.Time {
	return timeconv.Int64ToTime(u.UpdatedAt)
}

type EmailAddress struct {
	EmailAddress string        `json:"email_address"`
	ID           string        `json:"id"`
	LinkedTo     []interface{} `json:"linked_to"`
	Object       string        `json:"object"`
	Verification Verification  `json:"verification"`
}

type Verification struct {
	Status   string `json:"status"`
	Strategy string `json:"strategy"`
}

type ClerkEvent struct {
	Data   User   `json:"data"`
	Object string `json:"object"`
	Type   string `json:"type"`
}
