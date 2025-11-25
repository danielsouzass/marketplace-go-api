# Guia Rápido de Execução

## Passo a passo

1. Configure as variáveis de ambiente

   > Crie o arquivo `.env` com base no `.env.example`.

2. Suba o banco de dados

   ```bash
   docker compose up -d
   ```

3. Gere as migrations e o código go para lidar com o banco de dados

   ```bash
   go generate
   ```

4. Execute o projeto

   ```bash
   ...
   ```
