package auction_usecase

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type FinishAuctionOutputDTO struct {
	Id                  uint               `json:"id"`
	RequiredCredits     custom_type.BigInt `json:"required_credits,omitempty"`
	PriceLimitPerCredit custom_type.BigInt `json:"price_limit_per_credit,omitempty"`
	State               string             `json:"state,omitempty"`
	Bids                []*entity.Bid      `json:"bids,omitempty"`
	ExpiresAt           int64              `json:"expires_at,omitempty"`
	CreatedAt           int64              `json:"created_at,omitempty"`
	UpdatedAt           int64              `json:"updated_at,omitempty"`
}

type FinishAuctionUseCase struct {
	BidRepository     entity.BidRepository
	AuctionRepository entity.AuctionRepository
}

func NewFinishAuctionUseCase(auctionRepository entity.AuctionRepository, bidRepository entity.BidRepository) *FinishAuctionUseCase {
	return &FinishAuctionUseCase{
		AuctionRepository: auctionRepository,
		BidRepository:     bidRepository,
	}
}

func (u *FinishAuctionUseCase) Execute(metadata rollmelette.Metadata) (*FinishAuctionOutputDTO, error) {
	activeAuction, err := u.AuctionRepository.FindActiveAuction()
	if err != nil {
		return nil, err
	}

	if metadata.BlockTimestamp < activeAuction.ExpiresAt {
		return nil, fmt.Errorf("active auction not expired, you can't finish it yet")
	}

	bids, err := u.BidRepository.FindBidsByAuctionId(activeAuction.Id)
	if err != nil {
		return nil, err
	}

	if len(bids) == 0 {
		log.Println("no bids placed for active auction, finishing auction without bids")
	}

	sort.Slice(bids, func(i, j int) bool {
		return bids[i].PricePerCredit.Cmp(bids[j].PricePerCredit.Int) < 0
	})

	requiredCreditsRemaining := activeAuction.RequiredCredits

	for _, bid := range bids {
		if requiredCreditsRemaining.Cmp(big.NewInt(0)) == 0 {
			_, err := u.BidRepository.UpdateBid(&entity.Bid{
				Id:             bid.Id,
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        bid.Credits,
				PricePerCredit: bid.PricePerCredit,
				State:          "rejected",
				UpdatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
			continue
		}

		if requiredCreditsRemaining.Cmp(bid.Credits.Int) >= 0 {
			_, err := u.BidRepository.UpdateBid(&entity.Bid{
				Id:             bid.Id,
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        bid.Credits,
				PricePerCredit: bid.PricePerCredit,
				State:          "accepted",
				UpdatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}
			requiredCreditsRemaining.Sub(requiredCreditsRemaining.Int, bid.Credits.Int)
			continue
		}

		if bid.Credits.Int.Cmp(requiredCreditsRemaining.Int) == 1 {
			remainingCredits := new(big.Int).Set(requiredCreditsRemaining.Int)

			// Create the partially accepted bid
			res, err := u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(remainingCredits),
				PricePerCredit: bid.PricePerCredit,
				State:          "partially_accepted",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			// Create the rejected part of the bid
			_, err = u.BidRepository.CreateBid(&entity.Bid{
				AuctionId:      bid.AuctionId,
				Bidder:         bid.Bidder,
				Credits:        custom_type.NewBigInt(new(big.Int).Sub(bid.Credits.Int, remainingCredits)),
				PricePerCredit: bid.PricePerCredit,
				State:          "rejected",
				CreatedAt:      metadata.BlockTimestamp,
			})
			if err != nil {
				return nil, err
			}

			requiredCreditsRemaining.Sub(requiredCreditsRemaining.Int, res.Credits.Int)

			// Delete the original bid
			err = u.BidRepository.DeleteBid(bid.Id)
			if err != nil {
				return nil, err
			}
			continue
		}
	}

	if requiredCreditsRemaining.Cmp(big.NewInt(0)) > 0 {
		res, err := u.AuctionRepository.UpdateAuction(&entity.Auction{
			Id:                  activeAuction.Id,
			PriceLimitPerCredit: activeAuction.PriceLimitPerCredit,
			State:               "partially_awarded",
			ExpiresAt:           activeAuction.ExpiresAt,
			UpdatedAt:           metadata.BlockTimestamp,
		})
		if err != nil {
			return nil, err
		}
		return &FinishAuctionOutputDTO{
			Id:                  res.Id,
			RequiredCredits:     res.RequiredCredits,
			PriceLimitPerCredit: res.PriceLimitPerCredit,
			State:               string(res.State),
			Bids:                bids,
			ExpiresAt:           res.ExpiresAt,
			CreatedAt:           res.CreatedAt,
			UpdatedAt:           res.UpdatedAt,
		}, nil
	}

	res, err := u.AuctionRepository.UpdateAuction(&entity.Auction{
		Id:                  activeAuction.Id,
		PriceLimitPerCredit: activeAuction.PriceLimitPerCredit,
		State:               "finished",
		ExpiresAt:           activeAuction.ExpiresAt,
		UpdatedAt:           metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &FinishAuctionOutputDTO{
		Id:                  res.Id,
		RequiredCredits:     res.RequiredCredits,
		PriceLimitPerCredit: res.PriceLimitPerCredit,
		State:               string(res.State),
		Bids:                bids,
		ExpiresAt:           res.ExpiresAt,
		CreatedAt:           res.CreatedAt,
		UpdatedAt:           res.UpdatedAt,
	}, nil
}
