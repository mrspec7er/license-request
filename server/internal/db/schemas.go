package db

const FormSchema = `
CREATE TABLE Form (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(64),
    category VARCHAR(64)
);`

const SectionSchema = `
CREATE TABLE Section (
    id BIGSERIAL PRIMARY KEY,
    form_id BIGSERIAL,
    name VARCHAR(64),

    FOREIGN KEY (form_id) REFERENCES Form(id)
);`

const FieldSchema = `
CREATE TABLE Field (
    id BIGSERIAL PRIMARY KEY,
    section_id BIGSERIAL,
    label VARCHAR(256),
    type VARCHAR(64),
    field_order INTEGER,

    FOREIGN KEY (section_id) REFERENCES Section(id)
);`

const AccountSchema = `
CREATE TABLE Account (
    id BIGSERIAL PRIMARY KEY,
    uid VARCHAR(256),
    name VARCHAR(64),
    email VARCHAR(64),
    password TEXT,
    role VARCHAR(64)
);`

const ApplicationSchema = `
CREATE TABLE Application (
    number VARCHAR(64) PRIMARY KEY,
    form_id BIGSERIAL,
    account_id BIGSERIAL,

    FOREIGN KEY (form_id) REFERENCES Form(id),
    FOREIGN KEY (account_id) REFERENCES Account(id)
);`

const ResponseSchema = `
CREATE TABLE Response (
    id BIGSERIAL PRIMARY KEY,
    application_number VARCHAR(64),
    field_id BIGSERIAL,
    value TEXT,

    FOREIGN KEY (application_number) REFERENCES Application(number),
    FOREIGN KEY (field_id) REFERENCES Field(id)
);`
