CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id),
    address_id UUID REFERENCES address(id),
    status VARCHAR(50) DEFAULT 'pendente', -- Ex: pendente, pago, enviado, cancelado
    total NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
