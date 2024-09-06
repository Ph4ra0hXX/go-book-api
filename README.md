# Backend para Aplicação de TCC

Este repositório contém o código-fonte da API para a minha aplicação de Trabalho de Conclusão de Curso (TCC). O projeto é desenvolvido utilizando Golang com o framework Gin, e inclui funcionalidades de autenticação JWT, persistência de dados em PostgreSQL, e gerenciamento de contêineres com Docker. As configurações do projeto são gerenciadas através de arquivos TOML.

## Visão Geral

O objetivo deste projeto é fornecer uma API robusta e segura para suportar as funcionalidades do meu TCC. Utilizando o framework Gin, a aplicação é rápida e eficiente, enquanto o JWT garante a segurança na autenticação. O PostgreSQL é utilizado para o armazenamento confiável de dados, e o Docker facilita a implantação e gerenciamento do ambiente de execução.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação eficiente e moderna.
- **Gin**: Framework web para Go, focado em velocidade e flexibilidade.
- **JWT (JSON Web Tokens)**: Método seguro para autenticação de usuários.
- **PostgreSQL**: Sistema de gerenciamento de banco de dados relacional.
- **Docker**: Plataforma para desenvolver, enviar e executar aplicações em contêineres.
- **TOML**: Formato de arquivo para configurações, fácil de ler e editar.

## Instruções para Executar o Projeto Localmente

Para rodar o projeto localmente, siga as instruções abaixo:

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/Ph4ra0hXX/go-book-api.git
   ```

2. **Navegue até o diretório do projeto:**

   ```bash
   cd go-book-api
   ```

3. **Inicie o Docker para configurar o ambiente:**

   ```bash
   docker-compose up
   ```

4. **Execute a aplicação:**

   ```bash
   go run main.go
   ```
