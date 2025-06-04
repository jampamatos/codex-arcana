# Codex Arcana
Codex Arcana é um app desktop offline-first para Mestres de RPG (D&$D 5e no MVP), que combina editor de campanha, ferramentas in-session e DM screen interativa com widgets arrastáveis.

Esse aplicativo é um monorepo para aplicação desktop, API e documentação estática.

## Arquitetura

- **app/**: Electron + React + SQLite + isomorphic-git  
- **server/**: Express + Prisma + SQLite  
- **ssg/**: Docusaurus (TypeScript)

## Pré-requisitos

- Node.js ≥18  
- pnpm  
- Git

## Instalação

```bash
git clone <URL>
cd codex-arcana
pnpm install
```

## Desenvolvimento

Em terminais separados:

```bash
# UI React
pnpm --filter app ui

# Electron (desktop)
pnpm --filter app dev

# API (server)
pnpm --filter server dev

# Docs (Docusaurus)
pnpm --filter ssg start
```

## Build

```bash
pnpm --filter app build
pnpm --filter server build
pnpm --filter ssg build
```

## CI

Pipeline configurado em `.github/workflows/ci.yml` para lint, testes e build.