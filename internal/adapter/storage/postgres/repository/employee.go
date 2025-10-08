package repository

import (
	"context"

	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-gin-jwt-oauth/internal/core/domain"
)

type EmployeeRepository struct {
	db *postgres.DB
}

func InitEmployeeRepository(db *postgres.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (er *EmployeeRepository) CreateEmployee(ctx context.Context, employee *domain.Employee) (*domain.Employee, error) {
	db := er.db.GetDB()
	if err := db.Create(&employee).Error; err != nil {
		return &domain.Employee{}, err
	}

	return employee, nil
}

func (er *EmployeeRepository) GetEmployees(ctx context.Context) ([]*domain.Employee, error) {
	db := er.db.GetDB()

	var emps []*domain.Employee
	if err := db.Find(&emps).Error; err != nil {
		return []*domain.Employee{}, err
	}

	return emps, nil
}
