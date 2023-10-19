CREATE TABLE IF NOT EXISTS organizations (
  id bytea PRIMARY KEY NOT NULL,
  name varchar NOT NULL,
  representative varchar NOT NULL,
  phone_number varchar NOT NULL,
  address varchar NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS members (
  id bytea PRIMARY KEY NOT NULL,
  organization_id bytea NOT NULL,
  full_name varchar NOT NULL,
  email varchar UNIQUE NOT NULL,
  password varchar NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  FOREIGN KEY(organization_id) REFERENCES organizations(id)
);
COMMENT ON COLUMN members.email IS 'Index will be created automatically';

CREATE TABLE IF NOT EXISTS clients (
  id bytea PRIMARY KEY NOT NULL,
  organization_id bytea NOT NULL,
  name varchar NOT NULL,
  representative varchar NOT NULL,
  phone_number varchar NOT NULL,
  address varchar NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  FOREIGN KEY(organization_id) REFERENCES organizations(id)
);
COMMENT ON TABLE clients IS 'Though it is similar data structre to organization for now, those data models will go different paths';

CREATE TABLE IF NOT EXISTS client_bank_accounts (
  id bytea PRIMARY KEY NOT NULL,
  client_id bytea NOT NULL,
  name varchar NOT NULL,
  branch_name varchar NOT NULL,
  account_number varchar NOT NULL,
  account_name varchar NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  FOREIGN KEY(client_id) REFERENCES clients(id)
);
COMMENT ON TABLE client_bank_accounts IS 'Using bank account instead of account because interpretations might differ depending on the person';

CREATE TABLE IF NOT EXISTS invoices (
  id bytea PRIMARY KEY NOT NULL,
  organization_id bytea NOT NULL,
  client_id bytea NOT NULL,
  issue_date timestamp NOT NULL,
  due_date timestamp NOT NULL,
  amount_billed money NOT NULL,
  total_amount money NOT NULL,
  commission money NOT NULL,
  commission_rate numeric(6, 4) NOT NULL,
  consumption_tax money,
  consumption_tax_rate numeric(6, 4),
  status smallint NOT NULL,

  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  FOREIGN KEY(organization_id) REFERENCES organizations(id),
  FOREIGN KEY(client_id) REFERENCES clients(id)
);

CREATE INDEX IF NOT EXISTS idx_due_date ON invoices (due_date);
CREATE INDEX IF NOT EXISTS idx_status ON invoices (status);

COMMENT ON COLUMN invoices.amount_billed IS 'system might accept dollar and save it as JPY';
COMMENT ON COLUMN invoices.total_amount IS 'system might accept dollar and save it as JPY';
COMMENT ON COLUMN invoices.commission IS 'system might accept dollar and save it as JPY';
COMMENT ON COLUMN invoices.commission_rate IS 'For the futuer flexisiblity 99.9999 - 0.0001';
COMMENT ON COLUMN invoices.consumption_tax IS 'Null is accepted because transactions between japan and outside country does not require consumption tax';
COMMENT ON COLUMN invoices.consumption_tax_rate IS 'For the futuer flexisiblity 99.9999 - 0.0001';
