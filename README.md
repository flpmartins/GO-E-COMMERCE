# 🛒 GO-E-COMMERCE

API RESTful para gerenciamento de um sistema e-commerce, construída em Go com PostgreSQL.

---

## 🚀 Tecnologias

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Chi](https://github.com/go-chi/chi) — roteador HTTP leve e rápido
- [GORM](https://gorm.io/) — ORM para Go
- [JWT](https://jwt.io/) — autenticação com token
- [Docker](https://www.docker.com/) — para ambiente de desenvolvimento

---

## 📦 Funcionalidades

- ✅ Autenticação de usuários (login com JWT)
- 🔐 Validações com `validator.v10`
- 📄 Respostas padronizadas (`httpx`)
- 🛍️ CRUD de produtos, categorias e permissões (em desenvolvimento)
- 🧾 Carrinho de compras (WIP) (em desenvolvimento)
- 📦 Pedidos e histórico (em desenvolvimento)
- 👤 Gestão de usuários e perfis (em desenvolvimento)


---

## 🛠️ Como rodar o projeto

### Pré-requisitos

- Go 1.21+
- PostgreSQL 13+
- [Make](https://www.gnu.org/software/make/) (opcional)

### 1. Clone o repositório

```bash
git clone https://github.com/seu-user/go-e-commerce.git
cd go-e-commerce
