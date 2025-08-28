CREATE TABLE IF NOT EXISTS inventory (
    item_id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    quantity INT NOT NULL
);

-- Exemplo de cadastro manual
INSERT INTO inventory (item_id, name, quantity) VALUES
('item1', 'Produto 1', 100),
('item2', 'Produto 2', 50),
('item3', 'Produto 3', 75)
ON CONFLICT (item_id) DO NOTHING;