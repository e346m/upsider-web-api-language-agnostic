DROP INDEX IF EXISTS idx_due_date, idx_status;
ALTER TABLE client_bank_accounts DROP CONSTRAINT IF EXISTS client_bank_accounts_client_id_fkey;
DROP TABLE IF EXISTS invoices, client_bank_account, clients, members, organizations;
