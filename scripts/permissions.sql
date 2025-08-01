-- Inserção de permissões padrão
INSERT INTO permissions (id, name, value)
VALUES 
    (uuid_generate_v4(), 'admin', 'ADMIN'),
    (uuid_generate_v4(), 'client', 'CLIENT');
