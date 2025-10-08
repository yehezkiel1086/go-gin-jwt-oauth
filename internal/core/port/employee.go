package port

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *domain.Employee) (*domain.Employee, error)
	GetEmployees(ctx context.Context) ([]*domain.Employee, error)
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, emp *domain.Employee) (*domain.Employee, error)
	GetEmployees(ctx context.Context) ([]*domain.Employee, error)
}
