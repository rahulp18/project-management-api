CREATE TYPE organization_role AS ENUM (
    'OWNER',
    'ADMIN',
    'MEMBER'
);

CREATE TABLE IF NOT EXISTS organization_memberships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,
    organization_id UUID NOT NULL,

    role organization_role NOT NULL DEFAULT 'MEMBER',

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_membership UNIQUE (user_id, organization_id),

    CONSTRAINT fk_membership_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_membership_organization
        FOREIGN KEY (organization_id)
        REFERENCES organizations(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_memberships_user_id
ON organization_memberships(user_id);

CREATE INDEX idx_memberships_org_id
ON organization_memberships(organization_id);