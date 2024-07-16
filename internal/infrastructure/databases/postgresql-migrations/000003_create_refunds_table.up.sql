CREATE TABLE refunds (
  id varchar(255) PRIMARY KEY,
  transaction_id VARCHAR(255) NOT NULL,
  merchant_id VARCHAR(255) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
  status VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX idx_merchant_code_transaction ON refunds(transaction_id, merchant_id);
ALTER TABLE refunds ADD CONSTRAINT fk_refunds_merchants FOREIGN KEY(merchant_id) REFERENCES merchants(merchant_code);
ALTER TABLE refunds ADD CONSTRAINT fk_refunds_transactions FOREIGN KEY(transaction_id) REFERENCES transactions(id);