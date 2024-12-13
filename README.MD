# EVENT TRACKER

Projeto para capturar e registrar dados de interações realizados por usuários em uma página web.

## 📋 Pré-requisitos

Para executar o projeto é necessário apenas possuir o [Docker](https://www.docker.com/) e o [Docker Compose](https://docs.docker.com/compose/install/) instalado na sua máquina.

## ⚙️ Execução

Para executar o projeto, é necessário configurar o projeto criando um arquivo `.env` na raiz do projeto e definindo as variáveis de ambiente listadas no arquivo `.env.example`.

> ℹ️ Dica: Já existe um arquivo `.env` predefinido para facilitar a configuração.

Após definição das variáveis de ambiente, execute o seguinte comando na raiz do projeto:

```bash
docker compose up
```

## ▶️ Utilização

Para iniciar a captura de interações em uma página html, adicione a seguinte tag script, substituindo o `{FRONTEND_PORT}` pela porta definida no .env:

```html
<script src="http://localhost:{FRONTEND_PORT}/ws-sender.js"></script>
```

Dessa forma, as interações na página já estarão sendo registradas na base de dados.

> ℹ O projeto possui uma página de exemplo para testes, disponível em: http://localhost:{FRONTEND_PORT}/.

### 💻 Monitoramento

É possível visualizar os registros das interações em tempo real acessando http://localhost:{FRONTEND_PORT}/monitor.
