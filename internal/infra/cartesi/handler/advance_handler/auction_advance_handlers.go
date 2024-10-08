package advance_handler

import (
	"encoding/json"
	"fmt"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/internal/usecase/auction_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/bid_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/contract_usecase"
	"github.com/Mugen-Builders/devolt/internal/usecase/user_usecase"
	"github.com/rollmelette/rollmelette"
)

type AuctionAdvanceHandlers struct {
	BidRepository      entity.BidRepository
	UserRepository     entity.UserRepository
	OrderRepository    entity.OrderRepository
	AuctionRepository  entity.AuctionRepository
	ContractRepository entity.ContractRepository
}

func NewAuctionAdvanceHandlers(
	bidRepository entity.BidRepository,
	userRepository entity.UserRepository,
	OrderRepository entity.OrderRepository,
	auctionRepository entity.AuctionRepository,
	contractRepository entity.ContractRepository,
) *AuctionAdvanceHandlers {
	return &AuctionAdvanceHandlers{
		BidRepository:      bidRepository,
		UserRepository:     userRepository,
		OrderRepository:    OrderRepository,
		AuctionRepository:  auctionRepository,
		ContractRepository: contractRepository,
	}
}

func (h *AuctionAdvanceHandlers) CreateAuctionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	var input *auction_usecase.CreateAuctionInputDTO
	if err := json.Unmarshal(payload, &input); err != nil {
		return err
	}
	createAuction := auction_usecase.NewCreateAuctionUseCase(h.OrderRepository, h.AuctionRepository)
	res, err := createAuction.Execute(input, metadata)
	if err != nil {
		return err
	}
	auction, err := json.Marshal(res)
	if err != nil {
		return err
	}
	env.Notice(append([]byte("created auction - "), auction...))
	return nil
}

func (h *AuctionAdvanceHandlers) FinishAuctionHandler(env rollmelette.Env, metadata rollmelette.Metadata, deposit rollmelette.Deposit, payload []byte) error {
	finishAuction := auction_usecase.NewFinishAuctionUseCase(h.AuctionRepository, h.BidRepository)
	finishedAuction, err := finishAuction.Execute(metadata)
	if err != nil {
		return err
	}

	application, isDefined := env.AppAddress()
	if !isDefined {
		return fmt.Errorf("no application address defined yet, contact the DeVolt support")
	}

	findUserByRole := user_usecase.NewFindUserByRoleUseCase(h.UserRepository)
	auctioneer, err := findUserByRole.Execute(&user_usecase.FindUserByRoleInputDTO{Role: "auctioneer"})
	if err != nil {
		return err
	}

	findContractBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	volt, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "VOLT"})
	if err != nil {
		return err
	}

	stablecoin, err := findContractBySymbol.Execute(&contract_usecase.FindContractBySymbolInputDTO{Symbol: "STABLECOIN"})
	if err != nil {
		return err
	}

	findBidsByState := bid_usecase.NewFindBidsByStateUseCase(h.BidRepository)
	acceptedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "accepted",
	})
	if err != nil {
		return err
	}
	for _, bid := range acceptedBids {
		if err := env.ERC20Transfer(stablecoin.Address.Address, auctioneer.Address.Address, bid.Bidder.Address, bid.PricePerCredit.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	partialAcceptedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "partially_accepted",
	})
	if err != nil {
		return err
	}
	for _, bid := range partialAcceptedBids {
		if err := env.ERC20Transfer(stablecoin.Address.Address, auctioneer.Address.Address, bid.Bidder.Address, bid.PricePerCredit.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	rejectedBids, err := findBidsByState.Execute(&bid_usecase.FindBidsByStateInputDTO{
		AuctionId: finishedAuction.Id,
		State:     "rejected",
	})
	if err != nil {
		return err
	}
	for _, bid := range rejectedBids {
		if err := env.ERC20Transfer(volt.Address.Address, auctioneer.Address.Address, bid.Bidder.Address, bid.Credits.Int); err != nil {
			env.Report([]byte(err.Error()))
		}
	}

	profit := env.ERC20BalanceOf(stablecoin.Address.Address, auctioneer.Address.Address)
	if err := env.ERC20Transfer(stablecoin.Address.Address, auctioneer.Address.Address, application, profit); err != nil {
		env.Report([]byte(err.Error()))
	}

	env.Notice([]byte(fmt.Sprintf("finished auction with - id: %v, required credits: %v and price limit per credit: %v", finishedAuction.Id, finishedAuction.RequiredCredits.Int, finishedAuction.PriceLimitPerCredit.Int)))
	return nil
}
