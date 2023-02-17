
# Logic-API

this API is able to evaluate arbitrary logical expressions defined by the user.

## Getting starter

__Você precisará do Golang 1.19 instalado__

Para fazer o deploy desse projeto rode

```bash
  make local-start
```

após iniciado os container é necessário executar as migrations

```bash
  DATABASE_URL=postgres://postgres:postgres@localhost:15432/logic-api?sslmode=disable make migrate
```

Crie uma copia do arquivo local.env renomeando para .env
```bash
  cp local.env .env
```
Altere o valor da env __KEY_AUTH_API__ pelo token que deseja utilizar 

Após essas etapas a API está pronta para ser utilizada.

### Testes
Antes de executar os teste é necessário iniciar o container do postgres_test com o seguinte comando

```bash
  docker-compose up postgres_test
```
Então poderá executar os teste da seguinte forma 
```bash
  make test
```
## Stack utilizada

**Back-end:** Golang, Postgres, Docker


## Documentação da API

#### Autenticação

| Parâmetro   | Tipo       | Descrição                           |
| :---------- | :--------- | :---------------------------------- |
| `Authorization` | `string` | **Obrigatório**. Bearer token configurado na api |

#### Retorna lista de expressoes cadastradas

```http
  GET localhost:8383/expressions
```

#### Adiciona expressão nova

```http
  POST localhost:8383/expressions
```
payload:
```json
{
	"expression": "(x AND y) AND z"
}
```

| campo   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `expression`      | `string` | **Obrigatório**. A expressão que você quer |

#### Atualiza expressão

```http
  POST localhost:8383/expressions
```
payload:
```json
{
    "id": "b3d16a66-e711-4170-ae56-e503f4fe204a",
	"expression": "(x AND y) OR z"
}
```

| campo   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. referencia da expressão (caso não envie, será criada uma nova) |
| `expression`      | `string` | **Obrigatório**. A expressão que você quer |


#### Utiliza expressão salva
Passamos o id da expressão em questão e o valor (0 ou 1), referente a cada variável como parâmetros, obtendo o resultado da expressão na resposta da api.

expressão escolhida: 
```json
{
    "id": "b3d16a66-e711-4170-ae56-e503f4fe204a",
	"expression": "(x AND y) OR z"
}
```
request:
```http
  GET localhost:8383/evaluate/b3d16a66-e711-4170-ae56-e503f4fe204a?x=1&y=0&z=1
```
response:
```html
    true
```


#### Deletar expressão
Passamos o id da expressão em questão

expressão escolhida: 
```json
{
    "id": "b3d16a66-e711-4170-ae56-e503f4fe204a",
	"expression": "(x AND y) OR z"
}
```
request:
```http
  Delete localhost:8383/expressions/b3d16a66-e711-4170-ae56-e503f4fe204a?x=1&y=0&z=1
```
response:
```html
    204
```
