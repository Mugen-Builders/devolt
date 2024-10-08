package db

import (
	"fmt"

	"github.com/Mugen-Builders/devolt/internal/domain/entity"
	"gorm.io/gorm"
)

type StationRepositorySqlite struct {
	Db *gorm.DB
}

func NewStationRepositorySqlite(db *gorm.DB) *StationRepositorySqlite {
	return &StationRepositorySqlite{
		Db: db,
	}
}

func (r *StationRepositorySqlite) CreateStation(input *entity.Station) (*entity.Station, error) {
	err := r.Db.Create(input).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create station: %w", err)
	}
	return input, nil
}

func (r *StationRepositorySqlite) FindStationById(id uint) (*entity.Station, error) {
	var station entity.Station
	if err := r.Db.Preload("Orders").First(&station, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrStationNotFound
		}
		return nil, err
	}
	return &station, nil
}

func (r *StationRepositorySqlite) FindAllStations() ([]*entity.Station, error) {
	var stations []*entity.Station
	err := r.Db.Preload("Orders").Find(&stations).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find all stations: %w", err)
	}
	return stations, nil
}

func (r *StationRepositorySqlite) UpdateStation(input *entity.Station) (*entity.Station, error) {
	var station entity.Station
	err := r.Db.First(&station, "id = ?", input.Id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, entity.ErrStationNotFound
		}
		return nil, fmt.Errorf("failed to find station by ID: %w", err)
	}

	station.Consumption = input.Consumption
	station.Owner = input.Owner
	station.State = input.State
	station.PricePerCredit = input.PricePerCredit
	station.Latitude = input.Latitude
	station.Longitude = input.Longitude
	station.UpdatedAt = input.UpdatedAt

	res := r.Db.Save(&station)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to update station: %w", res.Error)
	}
	return &station, nil
}

func (r *StationRepositorySqlite) DeleteStation(id uint) error {
	res := r.Db.Delete(&entity.Station{}, "id = ?", id)
	if res.Error != nil {
		return fmt.Errorf("failed to delete station: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return entity.ErrStationNotFound
	}
	return nil
}
