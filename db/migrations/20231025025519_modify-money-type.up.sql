-- 本来はダイレクトに変更しない。新しいカラムを追加して移行期間を経て不要なカラムを削除するが今回は
-- money からnumricへの変換は暗黙的キャストにより成功する
ALTER TABLE invoices ALTER COLUMN amount_billed TYPE numeric(19, 4);
ALTER TABLE invoices ALTER COLUMN total_amount TYPE numeric(19, 4);
ALTER TABLE invoices ALTER COLUMN commission TYPE numeric(19, 4);
ALTER TABLE invoices ALTER COLUMN consumption_tax TYPE numeric(19, 4);

