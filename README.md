# Testando JWT no Golang com Fiber

### Descrição do Projeto

O projeto é uma aplicação em Golang que utiliza o framework Fiber para criar um servidor web com rotas para testar o uso de JSON Web Tokens (JWT). O objetivo é demonstrar a implementação de autenticação e autorização em rotas públicas e privadas usando o pacote JWT, além de utilizar o GORM para interagir com o banco de dados.

### Tecnologias Utilizadas

- Golang: Linguagem de programação utilizada para desenvolver a aplicação.
- Fiber: Um framework web rápido e fácil de usar em Golang para criação de servidores HTTP.
- GORM: Uma biblioteca de mapeamento objeto-relacional (ORM) em Golang para interagir com o banco de dados.
- JWT: JSON Web Tokens, um padrão de token de acesso que é usado para autenticação e compartilhamento de informações entre partes.

### Instalação

Para executar o projeto em sua máquina local, siga os passos abaixo:

1. Certifique-se de ter o Golang instalado em seu sistema. Caso ainda não tenha, faça o download e a instalação a partir do site oficial: https://golang.org/

2. Após instalar o Golang, você pode clonar o repositório do projeto usando o comando abaixo ou fazer o download do código-fonte manualmente:

  
```bash
   git clone https://github.com/KayoRonald/go-fiber-jwt-test.git
```   

3. Acesse o diretório do projeto:

  
```bash
   cd go-fiber-jwt-test
```   

4. Instale as dependências do projeto usando o comando:

  
```bash
   go get -u
```   

5. Configure as variáveis de ambiente necessárias para a aplicação, como as credenciais do banco de dados ou qualquer outra configuração específica que o projeto exija.

6. Agora, você pode executar a aplicação usando o seguinte comando:

  
```bash
   go run main.go
```   

7. O servidor estará em execução e poderá ser acessado através do navegador ou usando ferramentas como cURL ou Postman.

### Rotas

A aplicação possui as seguintes rotas:

1. Rota Pública:

   - Rota: /
   - Descrição: Rota pública que está disponível para todos os usuários, não exigindo autenticação.
   - Método: GET
   - Resposta: Retorna um array de usuário cadastrado, ou uma mensagem informando que não tem usuário.

2. Rota Privada - Informações do Usuário Logado:

   - Rota: /me
   - Descrição: Rota privada que requer autenticação com JWT. Apenas usuários autenticados podem acessá-la.
   - Método: GET
   - Resposta: Retorna informações do usuário logado, como seu nome de usuário, ID ou outras informações sensíveis.

3. Rota de Cadastro de Conta:

   - Rota: /signup
   - Descrição: Rota para criar uma nova conta de usuário.
   - Método: POST
   - Corpo da Requisição: Os dados do novo usuário a serem criados, como nome, email e senha.
   - Resposta: Retorna uma mensagem de sucesso ou falha na criação da conta.

4. Rota de Login:

   - Rota: /login
   - Descrição: Rota para autenticar um usuário e obter um token JWT válido.
   - Método: POST
   - Corpo da Requisição: As credenciais do usuário, como email e senha.
   - Resposta: Retorna um token JWT válido em caso de sucesso ou uma mensagem de falha na autenticação.
