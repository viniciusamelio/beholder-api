PRAGMA journal_mode=WAL;
PRAGMA synchronous = NORMAL;
PRAGMA temp_store = memory;
PRAGMA mmap_size = 30000000000;
PRAGMA page_size = 32768;
PRAGMA optimize;

CREATE TABLE organizations (
    id bigserial primary key not null,
    owner_id bigserial,
    name varchar not null,

    created_at timestamp default(CURRENT_TIMESTAMP),
    updated_at timestamp
);

CREATE TABLE users(
    id bigserial primary key not null,
    organization_id bigserial,
    name varchar not null,
    email varchar not null,
    status varchar not null default('active'),
    created_at timestamp default(CURRENT_TIMESTAMP),
    updated_at timestamp,

    constraint fk_users_organization_id
    foreign key(organization_id)
    references organizations(id)
);

CREATE TABLE projects(
    id bigserial primary key not null,
    user_id bigserial not null,
    organization_id bigserial not null,
    name varchar not null,
    created_at timestamp default(CURRENT_TIMESTAMP),
    updated_at timestamp,

    constraint fk_projects_user_id
    foreign key(user_id)
    references users(id),

    constraint fk_projects_organization_id
    foreign key(organization_id)
    references organizations(id)
);