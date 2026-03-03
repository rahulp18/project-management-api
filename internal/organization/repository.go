package organization

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateOrganization(ctx context.Context, input *OrganizationInput, userID string) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()
	var organizationID string
	query := `INSERT INTO organizations (name,description,image_url,banner_url)
	 VALUES ($1,$2,$3,$4)
	 RETURNING id
	 `
	err = tx.QueryRow(ctx, query, input.Name, input.Description, input.ImageUrl, input.BannerUrl).Scan(&organizationID)

	if err != nil {
		return err
	}

	memberQuery := `INSERT INTO organization_memberships (organization_id,user_id,role)
	VALUES ($1,$2,'OWNER')
	`
	_, err = tx.Exec(ctx, memberQuery, organizationID, userID)
	if err != nil {
		return nil
	}
	return tx.Commit(ctx)
}

// GET ORGANIZATION BY ID
func (r *Repository) GetOrganizationDetails(ctx context.Context, orgID string) (Organization, error) {
	var organization Organization
	orgQuery := `SELECT id,name,description,image_url,banner_url,created_at,updated_at
FROM organizations
WHERE id=$1`
	err := r.db.QueryRow(ctx, orgQuery, orgID).Scan(
		&organization.ID,
		&organization.Name,
		&organization.Description,
		&organization.ImageUrl,
		&organization.BannerUrl,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
	if err != nil {
		return Organization{}, err
	}
	// GET MEMBERS FROM ORGANIZATION ID
	members, err := r.getOrganizationMembers(ctx, orgID)
	if err != nil {
		return Organization{}, err
	}
	organization.Members = members
	return organization, nil
}

func (r *Repository) getOrganizationMembers(ctx context.Context,
	orgID string) ([]OrganizationMembers, error) {
	query := `
		SELECT 
			m.id,
			m.user_id,
			m.role,
			u.id,
			u.name,
			u.email
		FROM organization_memberships m
		JOIN users u ON u.id = m.user_id
		WHERE m.organization_id = $1
	`
	rows, err := r.db.Query(ctx, query, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []OrganizationMembers

	for rows.Next() {
		var member OrganizationMembers
		var user OrganizationUser

		err := rows.Scan(
			&member.ID,
			&member.UserID,
			&member.Role,
			&user.ID,
			&user.Name,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}

		member.User = user
		members = append(members, member)
	}
	return members, rows.Err()
}

func (r *Repository) GetAllOrganizations(ctx context.Context, userID string) ([]OrganizationBasic, error) {

	query := `
	SELECT o.id, o.name, o.description, o.image_url, o.banner_url
FROM organizations o
JOIN organization_memberships m 
    ON m.organization_id = o.id
WHERE m.user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var organizations []OrganizationBasic

	for rows.Next() {
		var org OrganizationBasic

		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Description,
			&org.ImageUrl,
			&org.BannerUrl,
		)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return organizations, nil
}

func (r *Repository) DeleteOrganization(
	ctx context.Context,
	orgID string,
	userID string,
) error {

	query := `
		DELETE FROM organizations o
		WHERE o.id = $1
		AND EXISTS (
			SELECT 1
			FROM organization_memberships m
			WHERE m.organization_id = o.id
			AND m.user_id = $2
			AND m.role = 'OWNER'
		)
	`

	cmdTag, err := r.db.Exec(ctx, query, orgID, userID)
	if err != nil {
		return err
	}

	// If no rows affected → either org doesn't exist
	// OR user is not OWNER
	if cmdTag.RowsAffected() == 0 {
		return errors.New("not authorized or organization not found")
	}

	return nil
}
