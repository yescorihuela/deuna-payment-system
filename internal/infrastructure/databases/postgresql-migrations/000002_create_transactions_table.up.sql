CREATE TABLE transactions (
  id varchar(255) PRIMARY KEY,
  merchant_id VARCHAR(255) NOT NULL,
  amount DECIMAL(10,2) NOT NULL DEFAULT 0.0,
  status VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE transactions ADD CONSTRAINT fk_transactions_merchants FOREIGN KEY(merchant_id) REFERENCES merchants(merchant_code);