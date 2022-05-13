package service

import (
	"github.com/ricardomaricato/banking/domain"
	"github.com/ricardomaricato/banking/dto"
	"github.com/ricardomaricato/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
	//if err := req.Validate(); err != nil {
	//	return nil, err
	//}
	//account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	//if newAccount, err := s.repo.Save(account); err != nil {
	//	return nil, err
	//} else {
	//	return newAccount.ToNewAccountResponseDto(), nil
	//}
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
