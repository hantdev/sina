package sdk

import "github.com/hantdev/mitras/pkg/transformers/senml"

type PageRes struct {
	Total  uint64 `json:"total"`
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

// bootstrapsPage contains list of bootstrap configs in a page with proper metadata.
type BootstrapPage struct {
	Configs []BootstrapConfig `json:"configs"`
	PageRes
}

type SubscriptionPage struct {
	Subscriptions []Subscription `json:"subscriptions"`
	PageRes
}

type MessagesPage struct {
	Messages []senml.Message `json:"messages,omitempty"`
	PageRes
}