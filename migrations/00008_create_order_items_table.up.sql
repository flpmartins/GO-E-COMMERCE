CREATE TABLE order_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id),
    quantity INT NOT NULL,
    price NUMERIC(10,2) NOT NULL, -- preço no momento da compra
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
