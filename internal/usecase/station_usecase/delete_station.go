package station_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
)

type DeleteStationInputDTO struct {
	Id uint `json:"id"`
}

type DeleteStationUseCase struct {
	StationRepository entity.StationRepository
}

func NewDeleteStationUseCase(stationRepository entity.StationRepository) *DeleteStationUseCase {
	return &DeleteStationUseCase{
		StationRepository: stationRepository,
	}
}

func (c *DeleteStationUseCase) Execute(input *DeleteStationInputDTO) error {
	return c.StationRepository.DeleteStation(input.Id)
}
