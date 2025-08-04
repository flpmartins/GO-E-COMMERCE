CREATE TABLE payment (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    method VARCHAR(100) NOT NULL, -- Ex: cartão, boleto, pix
    status VARCHAR(50) DEFAULT 'pendente', -- Ex: pendente, aprovado, recusado
    paid_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
