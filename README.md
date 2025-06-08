# Codex Arcana

Codex Arcana é um app desktop offline-first para Mestres de RPG (D&D 5e no MVP), combinando editor de campanha, ferramentas in-session e uma DM Screen interativa com widgets arrastáveis.

Este repositório é um monorepo contendo o backend em Go/Wails e o frontend em React. Consulte o arquivo [`roadmap.md`](roadmap.md) para o planejamento detalhado e próximos passos.

## Arquitetura

- `backend/` – Projeto [Wails](https://wails.io/) em Go. Além dos endpoints básicos (`/ping` e `/api/version`), inicializa o banco SQLite (`database.db`) com as tabelas `campaigns` e `sessions` e expõe a API REST de Campaign.
- `frontend/master/` – Aplicação React + Vite + Tailwind que consome esses endpoints. Servirá como interface do mestre.

Outros módulos (wiki estática, PWA para jogadores, etc.) serão adicionados conforme o roadmap.

## Status atual

O milestone inicial foi concluído e novas funcionalidades já foram adicionadas:

- Skeleton do backend em Go criado com Wails.
- Endpoints `/ping` e `/api/version` disponíveis e com CORS habilitado.
- **Persistência em SQLite** via arquivo `database.db` com modelos `Campaign` e `Session` (relacionados por chave estrangeira). Funções CRUD de `Campaign` já implementadas.
- **API REST de Campaign** com rotas:
  - `GET /api/campaigns` – lista todas as campanhas
  - `GET /api/campaigns/{id}` – retorna uma campanha específica
  - `POST /api/campaigns` – cria nova campanha
  - `PUT /api/campaigns/{id}` – atualiza campanha
  - `DELETE /api/campaigns/{id}` – remove campanha
- Frontend React configurado e integrado, exibindo a versão obtida do backend.
- **CRUD completo de Campaigns no frontend**: páginas React para listar, criar e
  editar campanhas (`/campaigns`, `/campaigns/new`, `/campaigns/:id/edit`) com
  opção de exclusão na listagem.
- **DAO de Session** com funções `CreateSession`, `GetSessionsByCampaign`, `GetSessionByID`, `UpdateSession` e `DeleteSession`.
- Testes básicos do DAO (Campaign e Session) podem ser executados com `go test ./...` dentro de `backend/`.

## Pré-requisitos

- Node.js ≥18
- Go ≥1.23
- npm ou pnpm
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

## Como rodar

```bash
git clone <URL>
cd codex-arcana

# Backend
cd backend
wails dev
# escuta em http://localhost:3000

# Frontend (novo terminal)
cd ../frontend/master
npm install
npm run dev
```

## Build

Para gerar o binário desktop e a build do frontend:

```bash
cd backend
wails build
# executáveis em backend/build/bin

cd ../frontend/master
npm run build
```

## Rodando via Docker com interface gráfica (Linux/X11)

Se quiser rodar o Codex Arcana em um container Docker **com interface gráfica (Wails/GTK)**, é necessário compartilhar o servidor X11 do seu host Linux com o container.  
**Atenção:** Isso só funciona em ambientes Linux com X11.

### Passos:

1. **Permita conexões X11 do Docker no host:**

   ```sh
   xhost +local:docker
   ```

2. **Suba o container normalmente:**

   ```sh
   docker compose up --build
   ```

   O app Wails abrirá a janela gráfica no seu desktop.

3. **(Opcional) Para revogar a permissão depois:**

   ```sh
   xhost -local:docker
   ```

### Observações

- Esse método exige configuração local e só funciona em Linux/X11.
- Para produção, distribua o binário gerado pelo Wails, não o container.
- Em outros sistemas operacionais (Windows, Mac, Wayland), o suporte a GUI via Docker é diferente e pode exigir soluções alternativas (VNC, Xpra, etc).

---

Para mais detalhes e próximas etapas consulte [`roadmap.md`](roadmap.md).
