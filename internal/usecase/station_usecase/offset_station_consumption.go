package station_usecase

import (
	"fmt"
	"math/big"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/rollmelette/rollmelette"
)

type OffSetStationConsumptionInputDTO struct {
	Id                uint               `json:"id"`
	CreditsToBeOffSet custom_type.BigInt `json:"credits_to_be_offset"`
}

type OffSetStationConsumptionOutputDTO struct {
	Id             uint                `json:"id"`
	Consumption    custom_type.BigInt  `json:"consumption"`
	Owner          custom_type.Address `json:"owner"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	State          string              `json:"state"`
	Latitude       float64             `json:"latitude"`
	Longitude      float64             `json:"longitude"`
	UpdatedAt      int64               `json:"updated_at"`
}

type OffSetStationConsumptionUseCase struct {
	StationRepository entity.StationRepository
}

func NewOffSetStationConsumptionUseCase(stationRepository entity.StationRepository) *OffSetStationConsumptionUseCase {
	return &OffSetStationConsumptionUseCase{
		StationRepository: stationRepository,
	}
}

func (u *OffSetStationConsumptionUseCase) Execute(input *OffSetStationConsumptionInputDTO, metadata rollmelette.Metadata) (*OffSetStationConsumptionOutputDTO, error) {
	station, err := u.StationRepository.FindStationById(input.Id)
	if err != nil {
		return nil, err
	}
	if station.Owner.Address != metadata.MsgSender {
		return nil, fmt.Errorf("can't offSet station consumption, because the station owner is not equal to the msg_sender address, expected: %v, got: %v", station.Owner.Address, metadata.MsgSender)
	}

	consumption := custom_type.NewBigInt(new(big.Int).Sub(station.Consumption.Int, input.CreditsToBeOffSet.Int))
	if consumption.Int.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("you are trying to offSet more credits than the station consumption, expected max: %v, got: %v", station.Consumption.Int, input.CreditsToBeOffSet.Int)
	}

	res, err := u.StationRepository.UpdateStation(&entity.Station{
		Id:             input.Id,
		Consumption:    consumption,
		Owner:          station.Owner,
		State:          station.State,
		Orders:         station.Orders,
		PricePerCredit: station.PricePerCredit,
		Latitude:       station.Latitude,
		Longitude:      station.Longitude,
		CreatedAt:      station.CreatedAt,
		UpdatedAt:      metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &OffSetStationConsumptionOutputDTO{
		Id:             res.Id,
		Consumption:    res.Consumption,
		Owner:          res.Owner,
		PricePerCredit: res.PricePerCredit,
		State:          string(res.State),
		Latitude:       res.Latitude,
		Longitude:      res.Longitude,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}
