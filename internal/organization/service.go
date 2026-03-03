package organization

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrganization(ctx context.Context, inputData OrganizationInput, userID string) error {
	return s.repo.CreateOrganization(ctx, &inputData, userID)
}
func (s *Service) GetAllOrganizations(ctx context.Context, userID string) ([]OrganizationBasic, error) {
	return s.repo.GetAllOrganizations(ctx, userID)
}
func (s *Service) GetOrganizationDetail(ctx context.Context, orgID string) (Organization, error) {
	return s.repo.GetOrganizationDetails(ctx, orgID)
}
func (s *Service) DeleteOrganization(ctx context.Context, orgID string, userID string) error {
	return s.repo.DeleteOrganization(ctx, orgID, userID)
}
