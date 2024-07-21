package contract_usecase

import (
	"github.com/devolthq/devolt/internal/domain/entity"
	"github.com/rollmelette/rollmelette"
)

type UpdateContractInputDTO struct {
	Id      uint   `json:"id"`
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type UpdateContractOutputDTO struct {
	Id        uint   `json:"id"`
	Symbol    string `json:"symbol"`
	Address   string `json:"address"`
	UpdatedAt int64  `json:"updated_at"`
}

type UpdateContractUseCase struct {
	ContractReposiotry entity.ContractRepository
}

func NewUpdateContractUseCase(contractRepository entity.ContractRepository) *UpdateContractUseCase {
	return &UpdateContractUseCase{
		ContractReposiotry: contractRepository,
	}
}

func (s *UpdateContractUseCase) Execute(input *UpdateContractInputDTO, metadata rollmelette.Metadata) (*UpdateContractOutputDTO, error) {
	contract, err := s.ContractReposiotry.UpdateContract(&entity.Contract{
		Id:        input.Id,
		Address:   input.Address,
		Symbol:    input.Symbol,
		UpdatedAt: metadata.BlockTimestamp,
	})
	if err != nil {
		return nil, err
	}
	return &UpdateContractOutputDTO{
		Id:        contract.Id,
		Symbol:    contract.Symbol,
		Address:   contract.Address,
		UpdatedAt: contract.UpdatedAt,
	}, nil
}
