package service

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/port"
)

type EmployeeService struct {
	repo port.EmployeeRepository
}

func InitEmployeeService(repo port.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (es *EmployeeService) CreateEmployee(ctx context.Context, emp *domain.Employee) (*domain.Employee, error) {
	res, err := es.repo.CreateEmployee(ctx, emp)
	if err != nil {
		return &domain.Employee{}, err
	}

	return res, nil
}

func (es *EmployeeService) GetEmployees(ctx context.Context) ([]*domain.Employee, error) {
	emps, err := es.repo.GetEmployees(ctx)
	if err != nil {
		return []*domain.Employee{}, err
	}

	return emps, nil
}
