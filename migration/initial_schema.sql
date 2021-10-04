CREATE TABLE wallets (
    id SERIAL PRIMARY KEY,
    verified BOOLEAN NOT NULL,
    balance INT NOT NULL
);
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    date_time TIMESTAMPTZ DEFAULT Now(),
    wallet_id INT NOT NULL,
    summ INTEGER NOT NULL,
    FOREIGN KEY (wallet_id) REFERENCES wallets (id)
);