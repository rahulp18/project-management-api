package organization

import "time"

type OrganizationInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	BannerUrl   string `json:"banner_url"`
}

type Organization struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	ImageUrl    string                `json:"image_url"`
	BannerUrl   string                `json:"banner_url"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
	Members     []OrganizationMembers `json:"members"`
}
type OrganizationMembers struct {
	ID     string           `json:"string"`
	UserID string           `json:"user_id"`
	Role   string           `json:"role"`
	User   OrganizationUser `json:"user"`
}
type OrganizationUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrganizationBasic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	BannerUrl   string    `json:"banner_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
