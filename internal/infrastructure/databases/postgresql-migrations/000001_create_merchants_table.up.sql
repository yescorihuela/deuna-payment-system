CREATE TABLE merchants (
  id bigserial PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  balance DECIMAL(10,2) DEFAULT 0.0,
  merchant_code VARCHAR(255) UNIQUE NOT NULL,
  notification_email VARCHAR(255) UNIQUE NOT NULL,
  enabled BOOLEAN DEFAULT false,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX idx_merchant_code ON merchants(merchant_code);