-- Habilita suporte a UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Criação da tabela de permissões
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
