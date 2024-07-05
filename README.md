# Product API

Esta é uma API simples para cadastrar e visualizar produtos em um banco de dados PostgreSQL. A API utiliza Docker para a configuração do ambiente, `migrate` para gerenciamento de migrações de banco de dados e `make` para automatização de tarefas.

## Requisitos

- Docker e Docker Compose
- `migrate` (para gerenciar migrações de banco de dados)
- `make` (para automatizar comandos)

## Instalação das Dependências

### Docker e Docker Compose

Instale Docker e Docker Compose seguindo as instruções oficiais:

- [Docker](https://docs.docker.com/get-docker/)

### `migrate`

Instale `migrate` seguindo as instruções na documentação oficial:

- [migrate Documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### `make`

Instale `make` conforme necessário para o seu sistema operacional:

- [GNU Make Documentation](https://www.gnu.org/software/make/)

## Configuração do Projeto

Clone o repositório do projeto:

```sh
git clone https://github.com/Trovarelli/api-go.git
cd seu-repositorio
```

## Iniciar a API

Com o banco de dados configurado e as migrações aplicadas, você pode iniciar a API usando o comando:

```sh
docker-compose up --build
```

### Configuração do Banco de Dados

O projeto utiliza Docker Compose para configurar o banco de dados PostgreSQL. Use o comando abaixo para iniciar o banco de dados:

```sh
docker-compose up -d
```

Isso irá iniciar um contêiner Docker com PostgreSQL, usando as seguintes credenciais:

- Usuário: postgres
- Senha: 1234
- Banco de Dados: postgres
- Porta: 5432

### Comandos do make

Aplicar Migrações
Para aplicar migrações e gerar as tabelas no banco de dados, use o comando:

```sh
make gen_tables
```

Reverter Migrações
Para reverter as migrações e dropar as tabelas, use o comando:

```sh
make drop_tables
```

### Endpoints da API

#### Criar Produto

- Endpoint: /products
- Método: POST
- Corpo:

```json
{
  "descricao": "Nome do Produto",
  "preco": 10.99
}
```

#### Listar Produtos

- Endpoint: /products
- Método: GET
  Retorna uma lista de todos os produtos cadastrados.

#### Encontra Produto pelo Id

- Endpoint: /products/:id
- Método: GET
  Retorna o produto com o id escolhido.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.

## Licença

Este projeto está licenciado sob a [Licença MIT](https://opensource.org/licenses/MIT).
