# Roadmap de Alto Nível

## Epics e Milestones

Vamos dividir o trabalho em epics (módulos grandes) e, dentro de cada epic, user stories específicas. O MVP (mínimo produto viável) deverá permitir:

1. **Epic 1: Infraestrutura e Skeleton**

    * **Objetivo**: Ter o esqueleto do aplicativo funcionando localmente (front + back).
    * **User Stories**:
        1.1. Configurar Repositório Git
        1.2. Inicializar Projeto Go + Wails
        1.3. Criar Frontend React + Tailwind Básico
        1.4. Estabelecer Estrutura de Pastas
        1.5. Servidor Go responde /ping e /api/version
        1.6. Build Local (Wails) e Execução mínima

2. **Epic 2: Módulo “Campaigns & Sessions”**
    * **Objetivo**: Permitir CRUD de “Campaigns” e “Sessions” via API e front React.
    * **User Stories**:
        2.1. Modelagem SQLite para Campaign
        2.2. Endpoints REST Campaign (GET, POST, PUT, DELETE)
        2.3. Front React: CampaignList e CampaignForm (lista + criação/edição)
        2.4. Modelagem SQLite para Session (vinculada à Campaign)
        2.5. Endpoints REST Session
        2.6. Front React: SessionList e SessionForm
3. **Epic 3: Lore-Keeping Básico (Markdown Editor + SSG)**
    * **Objetivo**: Permitir criar e editar páginas Markdown, gerar HTML local e exibir preview.
    * **User Stories**:
        3.1. Backend: Listar arquivos Markdown em /content/
        3.2. Backend: Salvar/Editar arquivo .md com frontmatter
        3.3. Frontend: LoreList com todas páginas e botão “Criar Nova”
        3.4. Frontend: LoreEditor (textarea + preview instantâneo via Goldmark no client)
        3.5. Função Go para converter Markdown → HTML (pasta /public/)
        3.6. Pipeline de Publicação: botão “Publicar Wiki” → git checkout gh-pages, sobrescrever /public, commit/push
4. **Epic 4: Autenticação & Conexão Mestre–Jogador**
    * **Objetivo**: Permitir que o Mestre inicie o servidor, descubra via mDNS ou inicie túnel ngrok, e que o Jogador se conecte.
    * **User Stories**:
        4.1. Implementar mDNS (announceService e discoverMasters)
        4.2. Configurar Tela de Settings: campo para Ngrok Token
        4.3. Lógica Go para executar ngrok http 3000 e capturar URL
        4.4. Endpoints Go: /ping (valida token) e /ws (aceita WS)
        4.5. Front React Mestre: botão “Iniciar Túnel Ngrok” e exibir URL
        4.6. Front React Jogador: DiscoverPanel (mostra hosts LAN) e input de URL ngrok
        4.7. Fluxo de Autenticação (JWT ou token simples)
5. **Epic 5: Módulo “Character Sheets” (D&D 5e)**
    * **Objetivo**: CRUD de fichas de personagem e exportação básica.
    * **User Stories**:
        5.1. Modelagem SQLite para Character (dados D&D 5e típicos)
        5.2. Endpoints REST Character
        5.3. Front React: CharacterList e CharacterForm (dynamic fields)
        5.4. Exportar PDF → Integração com gofpdf (exemplo: gera PDF simples)
        5.5. Exportar Markdown → Gera arquivo .md
6. **Epic 6: Ferramentas In-Session & DM Screen (Widget Canvas)**
    * **Objetivo**: Implementar o rolador de dados, tracker de iniciativa, PV/condições e widgets básicos no DM Screen.
    * **User Stories**:
        6.1. Rolador de Dados Go (função para parsear “1d20+5”)
        6.2. Endpoint Go: /api/roll que retorna resultado JSON
        6.3. Front React: componente RollButton que chama /api/roll e exibe resultado
        6.4. Tracker de Iniciativa: backend gera lista ordenada e envia via WS
        6.5. Front React DM Screen: lista de turno reordenável
        6.6. Widget PV/Condições: componente que envia via WS update de PV
        6.7. Implementar Canvas Drag & Drop com biblioteca leve (ex.: react-dnd ou react-canvas)
        6.8. Criar widgets básicos (Contador, Nota, Timer, Player de Música)
7. **Epic 7: Frontend Jogador (PWA Offline-First)**
    * **Objetivo**: Permitir ao Jogador usar o PWA para ver ficha, notificações e lore pública offline.
    * **User Stories**:
        7.1. Criar Pasta /frontend/jogador e segmentar componentes
        7.2. Configurar manifest.json e service-worker.js para cache básico
        7.3. Componente PlayerSheet (só leitura)
        7.4. Componente RollButtons (mesmo rolador, mas restrito)
        7.5. Componente Notifications (escuta WS e exibe alertas)
        7.6. Integrar iframe ou link para wiki público estático

---

## Milestone Atual

1. **Criar e Configurar Repositório GitHub**

- [X] Criar repositório público ou privado (a sua escolha).
- [X] Adicionar .gitignore para Go, Node, Wails.
- [X] Criar README.md inicial com instruções de setup.

2. **Configurar Skeleton Go + Wails**

- [X] Instalar Go (se ainda não tiver).
- [X] Instalar Wails localmente (go install github.com/wailsapp/wails/v2/cmd/wails@latest).
- [X] No diretório raiz, executar `wails init -n backend -t standard` (ou copiar template padrão). 
  - Na verdade, o comando foi `wails init -n backend -t vanilla`, porque o template `standard` não estava disponível.
- [X] Ajustar `go.mod` para `module codex-arcana/backend`.
- [X] Em `main.go`, criar handler para `/ping` e `/api/version`.
- [X] Executar `wails dev` para testar servidor+frontend integrado minimal. (Mas podemos separar: foco só no servidor Go, sem ainda empacotar no Wails).

3. **Configurar Frontend React Master**

- [X] No diretório `/frontend/master/`, executar `npm create vite@latest . -- --template react-ts`.
- [X] Instalar dependências: `npm install`, `npm install -D tailwindcss postcss autoprefixer`.
- [ ] Executar `npx tailwindcss init -p`.
  - Uma vez que esse comando não funcionou, criamos o arquivo `tailwind.config.cjs` manualmente.
  - Criamos também o arquivo `postcss.config.cjs` manualmente.
  - Editamos o `package.json` para incluir os scripts de build e dev do Tailwind.
  - Resetamos o `index.css` para incluir as diretivas do Tailwind (`@tailwind base; @tailwind components; @tailwind utilities;`).
  - O Tailwind parece não estar funcionando corretamente, então vamos investigar isso mais tarde. Os passos abaixo também foram feitos, mas o Tailwind não está aplicando os estilos como esperado.
- [X] Configurar `tailwind.config.cjs`:
```javascript
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  theme: { extend: {} },
  plugins: [],
};
```
- [X] No `src/index.css`, adicionar diretivas do Tailwind:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```
  - Retiramos essas diretivas e aplicaremos CSS vanilla por enquanto, já que o Tailwind não está funcionando corretamente.
- [X] Em `App.tsx`, remover conteúdo de exemplo e exibir “Hello Codex Arcana” dentro de uma `<div className="p-4 text-2xl">`.

1. **Ajustar CORS no Backend**

- [ ] Instalar pacote de middleware CORS (por exemplo: `github.com/rs/cors`).
- [ ] No `main.go`, antes de `http.ListenAndServe`, envolver o mux principal com o handler CORS:
```go
mux := http.NewServeMux()
mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
})
handler := cors.Default().Handler(mux)
log.Fatal(http.ListenAndServe(":3000", handler))
```
- [ ] Testar localmente: `curl localhost:3000/ping`.

5. **Conectar Frontend ao Backend**
- [ ] Em `App.tsx`, usar `useEffect(() => { fetch("http://localhost:3000/ping").then(...) }, [])`.
- [ ] Exibir no console devtools ou dentro de um `<p>` o texto “pong”.

6. **Commit & Push**
- [X] Criar branch `milestone1-skeleton` ou trabalhar diretamente em `main` se for conveniente.
- [ ] Fazer commits granulares (um para o skeleton Go, outro para o React + Tailwind, outro para a integração fetch).
- [ ] Push para o GitHub e abrir um PR (caso queira revisão automática).


