CREATE TABLE IF NOT EXISTS accounts (
  id INT NOT NULL,
  name varchar(250) NOT NULL,
  cpf varchar(11) NOT NULL,
  secret varchar(250) NOT NULL,
  balance numeric(19,4) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  UNIQUE(cpf),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS transfers (
  id SERIAL, --INT NOT NULL,
  account_origin_id INT NOT NULL,
  account_destination_id INT NOT NULL,
  amount numeric(19,4) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_dst
      FOREIGN KEY(account_origin_id)
	  REFERENCES accounts(id),
  CONSTRAINT fk_src
      FOREIGN KEY(account_destination_id)
	  REFERENCES accounts(id)
);
