package station_usecase

import (
	"math/big"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"github.com/Mugen-Builders/devolt/pkg/custom_type"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rollmelette/rollmelette"
)

type UpdateStationInputDTO struct {
	Id             uint           `json:"id"`
	Consumption    *big.Int       `json:"consumption"`
	Owner          common.Address `json:"owner"`
	PricePerCredit *big.Int       `json:"price_per_credit"`
	State          string         `json:"state"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
}

type UpdateStationOutputDTO struct {
	Id             uint                `json:"id"`
	Consumption    custom_type.BigInt  `json:"consumption"`
	Owner          custom_type.Address `json:"owner"`
	PricePerCredit custom_type.BigInt  `json:"price_per_credit"`
	State          string              `json:"state"`
	Latitude       float64             `json:"latitude"`
	Longitude      float64             `json:"longitude"`
	CreatedAt      int64               `json:"created_at"`
	UpdatedAt      int64               `json:"updated_at"`
}

type UpdateStationUseCase struct {
	StationRepository entity.StationRepository
}

func NewUpdateStationUseCase(stationRepository entity.StationRepository) *UpdateStationUseCase {
	return &UpdateStationUseCase{
		StationRepository: stationRepository,
	}
}

func (u *UpdateStationUseCase) Execute(input *UpdateStationInputDTO, metadata rollmelette.Metadata) (*UpdateStationOutputDTO, error) {
	res, err := u.StationRepository.UpdateStation(&entity.Station{
		Id:             input.Id,
		Consumption:    custom_type.NewBigInt(input.Consumption),
		Owner:          custom_type.NewAddress(input.Owner),
		PricePerCredit: custom_type.NewBigInt(input.PricePerCredit),
		State:          entity.StationState(input.State),
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		UpdatedAt:      metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateStationOutputDTO{
		Id:             res.Id,
		Consumption:    res.Consumption,
		Owner:          res.Owner,
		PricePerCredit: res.PricePerCredit,
		State:          string(res.State),
		Latitude:       res.Latitude,
		Longitude:      res.Longitude,
		CreatedAt:      res.CreatedAt,
		UpdatedAt:      res.UpdatedAt,
	}, nil
}
