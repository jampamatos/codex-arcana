# Codex Arcana

Codex Arcana é um app desktop offline-first para Mestres de RPG (D&D 5e no MVP), combinando editor de campanha, ferramentas in-session e uma DM Screen interativa com widgets arrastáveis.

Este repositório é um monorepo contendo o backend em Go/Wails e o frontend em React. Consulte o arquivo [`roadmap.md`](roadmap.md) para o planejamento detalhado e próximos passos.

## Arquitetura

- `backend/` – Projeto [Wails](https://wails.io/) em Go. Expõe endpoints HTTP (como `/ping` e `/api/version`) e empacota a aplicação desktop.
- `frontend/master/` – Aplicação React + Vite + Tailwind que consome esses endpoints. Servirá como interface do mestre.

Outros módulos (wiki estática, PWA para jogadores, etc.) serão adicionados conforme o roadmap.

## Status atual

O milestone inicial está completo:

- Skeleton do backend em Go criado com Wails.
- Endpoints `/ping` e `/api/version` disponíveis e com CORS habilitado.
- Frontend React configurado e integrado, exibindo a versão obtida do backend.

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

Para mais detalhes e próximas etapas consulte [`roadmap.md`](roadmap.md).
