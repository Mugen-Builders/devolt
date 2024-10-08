package bid_usecase

import "github.com/Mugen-Builders/devolt/pkg/custom_type"

type FindBidOutputDTO struct {
	Id             uint                `json:"id"`
	AuctionId      uint                `json:"auction_id"`
	Bidder         custom_type.Address `json:"bidder"`
	Credits        custom_type.BigInt  `json:"credits"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	State          string              `json:"state"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      int64               `json:"updated_at"`
}
