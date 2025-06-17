# Roadmap de Alto Nível

## Epics e Milestones

Vamos dividir o trabalho em epics (módulos grandes) e, dentro de cada epic, user stories específicas. O MVP (mínimo produto viável) deverá permitir:

1. **Epic 1: Infraestrutura e Skeleton**

    * **Objetivo**: Ter o esqueleto do aplicativo funcionando localmente (front + back).
    * **User Stories**:
      <s>*  1.1. Configurar Repositório Git
      *  1.2. Inicializar Projeto Go + Wails
      *  1.3. Criar Frontend React + Tailwind Básico
      *  1.4. Estabelecer Estrutura de Pastas
      *  1.5. Servidor Go responde /ping e /api/version
      *  1.6. Build Local (Wails) e Execução mínima</s>
      *  *Epic 1 já foi concluído, conforme o milestone atual.*

2. **Epic 2: Módulo “Campaigns & Sessions”**
    * **Objetivo**: Permitir CRUD de “Campaigns” e “Sessions” via API e front React.
    * **User Stories**:
       <s>* 2.1. Modelagem SQLite para Campaign
       * 2.2. Endpoints REST Campaign (GET, POST, PUT, DELETE)
       * 2.3. Front React: CampaignList e CampaignForm (lista + criação/edição)
       * 2.4. Modelagem SQLite para Session (vinculada à Campaign)
       * 2.5. Endpoints REST Session
       * 2.6. Front React: SessionList e SessionForm</s>
       * *Epic 2 já foi concluído, conforme o milestone atual.*
3. **Epic 3: Lore-Keeping Básico (Markdown Editor + SSG)**
    * **Objetivo**: Permitir criar e editar páginas Markdown, gerar HTML local e exibir preview.
    * **User Stories**:
        * 3.1. Backend: Listar arquivos Markdown em /content/
        * 3.2. Backend: Salvar/Editar arquivo .md com frontmatter
        * 3.3. Frontend: LoreList com todas páginas e botão “Criar Nova”
        * 3.4. Frontend: LoreEditor (textarea + preview instantâneo via Goldmark no client)
        * 3.5. Função Go para converter Markdown → HTML (pasta /public/)
        * 3.6. Pipeline de Publicação: botão “Publicar Wiki” → git checkout gh-pages, sobrescrever /public, commit/push
4. **Epic 4: Autenticação & Conexão Mestre–Jogador**
    * **Objetivo**: Permitir que o Mestre inicie o servidor, descubra via mDNS ou inicie túnel ngrok, e que o Jogador se conecte.
    * **User Stories**:
        * 4.1. Implementar mDNS (announceService e discoverMasters)
        * 4.2. Configurar Tela de Settings: campo para Ngrok Token
        * 4.3. Lógica Go para executar ngrok http 3000 e capturar URL
        * 4.4. Endpoints Go: /ping (valida token) e /ws (aceita WS)
        * 4.5. Front React Mestre: botão “Iniciar Túnel Ngrok” e exibir URL
        * 4.6. Front React Jogador: DiscoverPanel (mostra hosts LAN) e input de URL ngrok
        * 4.7. Fluxo de Autenticação (JWT ou token simples)
5. **Epic 5: Módulo “Character Sheets” (D&D 5e)**
    * **Objetivo**: CRUD de fichas de personagem e exportação básica.
    * **User Stories**:
        * 5.1. Modelagem SQLite para Character (dados D&D 5e típicos)
        * 5.2. Endpoints REST Character
        * 5.3. Front React: CharacterList e CharacterForm (dynamic fields)
        * 5.4. Exportar PDF → Integração com gofpdf (exemplo: gera PDF simples)
        * 5.5. Exportar Markdown → Gera arquivo .md
6. **Epic 6: Ferramentas In-Session & DM Screen (Widget Canvas)**
    * **Objetivo**: Implementar o rolador de dados, tracker de iniciativa, PV/condições e widgets básicos no DM Screen.
    * **User Stories**:
        * 6.1. Rolador de Dados Go (função para parsear “1d20+5”)
        * 6.2. Endpoint Go: /api/roll que retorna resultado JSON
        * 6.3. Front React: componente RollButton que chama /api/roll e exibe resultado
        * 6.4. Tracker de Iniciativa: backend gera lista ordenada e envia via WS
        * 6.5. Front React DM Screen: lista de turno reordenável
        * 6.6. Widget PV/Condições: componente que envia via WS update de PV
        * 6.7. Implementar Canvas Drag & Drop com biblioteca leve (ex.: react-dnd ou react-canvas)
        * 6.8. Criar widgets básicos (Contador, Nota, Timer, Player de Música)
7. **Epic 7: Frontend Jogador (PWA Offline-First)**
    * **Objetivo**: Permitir ao Jogador usar o PWA para ver ficha, notificações e lore pública offline.
    * **User Stories**:
        * 7.1. Criar Pasta /frontend/jogador e segmentar componentes
        * 7.2. Configurar manifest.json e service-worker.js para cache básico
        * 7.3. Componente PlayerSheet (só leitura)
        * 7.4. Componente RollButtons (mesmo rolador, mas restrito)
        * 7.5. Componente Notifications (escuta WS e exibe alertas)
        * 7.6. Integrar iframe ou link para wiki público estático

---

## Milestone Atual: Módulo “Lore-Keeping Básico”

**Objetivo**: Permitir criar e editar páginas Markdown, gerar HTML local e exibir preview no frontend.

**User Stories**:

* 3.1. Backend: Listar arquivos Markdown em `/content/`
* 3.2. Backend: Salvar/Editar arquivo `.md` com frontmatter
* 3.3. Frontend: `LoreList` com todas páginas e botão “Criar Nova”
* 3.4. Frontend: `LoreEditor` (textarea + preview instantâneo via Goldmark no client)
* 3.5. Função Go para converter Markdown → HTML (pasta `/public/`)
* 3.6. Pipeline de Publicação: botão “Publicar Wiki” → `git checkout gh-pages`, sobrescrever `/public`, commit/push

### User Story 3.1 – Backend: Listar arquivos Markdown em /content/

- [ ] Criar função em Go (por exemplo, ListLoreFiles() ([]LoreFile, error)) que:
  - Abre o diretório `./content/`
  - Filtra apenas arquivos com extensão `.md`
  - Para cada arquivo, extrai:
    - slug (nome do arquivo sem `.md`)
    - título (field `title:` do frontmatter, se existir; caso contrário, usa o slug)
    - visibilidade (field `visibility:` do frontmatter ou “master” por padrão)
  - Retorna um slice de structs `LoreFile{Slug, Title, Visibility}`
- [ ] Definir a struct `type LoreFile struct { Slug, Title, Visibility string }` em `backend/models`.
- [ ] Criar handler HTTP para `GET /api/lore` que:
  - Chama `ListLoreFiles()`
  - Retorna JSON no formato `[{ slug, title, visibility }, …]`
  - Se houver erro de leitura, retorna status 500
- [ ] Adicionar rota em `main.go`: `mux.HandleFunc("/api/lore", loreListHandler)`
- [ ] Testar manualmente:
  - Colocar alguns arquivos `.md` em `/content/` (com e sem frontmatter)
  - Executar `curl http://localhost:3000/api/lore` e verificar se lista corretamente slugs, titles e visibilities
- [ ] Criar testes básicos em `backend/lore_dao_test.go`:
  - Testar `ListLoreFiles()` num diretório de fixtures (ex.: criar pastinha temporária)
  - Verificar que retorna o número correto de arquivos e metadados esperados

### User Story 3.2 - Backend: Salvar/Editar arquivo .md com frontmatter

- [ ] Definir struct `LorePage` em `backend/models`:
  - Campos:
    - `Slug string`
    - `Title string`
    - `Visibility string` (valores possíveis: `"master"` ou `"player"`)
    - `Content string` (corpo Markdown sem frontmatter)
- [ ] Implementar função `SaveLoreFile(page LorePage) error` no DAO de lore:
  - Determinar caminho do arquivo: `./content/{page.Slug}.md`
  - Serializar frontmatter em YAML:
    ```yaml
    ---
    title: {{ page.Title }}
    visibility: {{ page.Visibility }}
    ---
    ```
  - Concatenar frontmatter + corpo (`page.Content`)
  - Escrever no disco (sobrescrevendo se já existir)
- [ ] Criar handler HTTP para criação e edição:
  - `POST /api/lore` (nova página):
    - Decodificar JSON com `{ slug, title, visibility, content }`
    - Validar:
      - `slug` não vazio e apenas letras, números, hífen ou underscore
      - `title` não vazio
      - `visibility` em `{ "master","player" }`
    - Chamar `SaveLoreFile()`
    - Retornar 201 Created com objeto `{ slug, title, visibility }`
  - `PUT /api/lore/{slug}` (edição):
    - Extrair slug da URL
    - Decodificar JSON com `{ title, visibility, content }`
    - Validar como acima (exceto slug)
    - Chamar `SaveLoreFile()`
    - Retornar `200 OK` com objeto atualizado
- [ ] Registrar rotas em `main.go`:
  - `mux.HandleFunc("/api/lore", loreCreateHandler)`
  - `mux.HandleFunc("/api/lore/", loreEditHandler)`
- [ ] Testar manualmente via cURL ou Postman:
  - Criar página:
    ```bash
    curl -X POST http://localhost:3000/api/lore \
      -H "Content-Type: application/json" \
      -d '{"slug":"introducao","title":"Introdução","visibility":"player","content":"# Bem-vindo"}'
    ```
    - Verificar que `content/introducao.md` foi criado com frontmatter e corpo.
- [ ] Editar página:
    ```bash
    curl -X PUT http://localhost:3000/api/lore/introducao \
      -H "Content-Type: application/json" \
      -d '{"title":"Intro","visibility":"master","content":"# Novo Conteúdo"}'
    ```
  - Confirmar que o arquivo foi atualizado corretamente.
- [ ] Escrever testes unitários em `backend/lore_dao_test.go`:
  - Criar diretório temporário como fixture
  - Chamar `SaveLoreFile()`
  - Ler o arquivo gerado e verificar:
    - Frontmatter formatado corretamente
    - Corpo Markdown preservado

### User Story 3.3 - Frontend: LoreList com todas páginas e botão “Criar Nova”

- [ ] Configurar rota no React Router
  - Definir rota `/lore` que renderiza o componente `LoreListPage.tsx.`
- [ ] Criar componente `LoreListPage.tsx`
  - Importar `useEffect`, `useState` e `useNavigate` (ou `useHistory`) do React Router.
  - Definir estado local:
    - `loreFiles: LoreFile[]` (lista de arquivos, onde `LoreFile` tem `{ slug, title, visibility }`)
    - `loading: boolean`
    - `error: string | null`
- [ ] Buscar dados da API
  - No `useEffect`, fazer `fetch("/api/lore")` ao montar o componente:
    - Enquanto aguarda, `loading = true`.
    - Se sucesso, popular `loreFiles` e `loading = false`.
    - Se erro, armazenar mensagem em `error` e `loading = false`.
- [ ] Renderizar lista de páginas
  - Se `loading` for `true`, mostrar “Carregando…”
  - Se `error` não for `null`, exibir a mensagem de erro
  - Caso contrário, iterar sobre `loreFiles` e renderizar cada item em uma tabela ou em cards, exibindo:
    - Título (`file.title`)
    - Visibilidade (`file.visibility`)
    - Botão “Editar” que navega para `/lore/{slug}/editar`
    - Botão “Excluir” que dispara `DELETE /api/lore/{slug}` e, em caso de sucesso, remove o item de `loreFiles`
- [ ] Adicionar botão “Criar Nova Página”
  - Posicionar no topo da lista
  - Ao clicar, navegar para `/lore/novo`
- [ ] Tratamento de ações
  - Ao clicar em “Excluir”, confirmar com o usuário (ex.: `window.confirm`) antes de chamar a API
  - Em caso de sucesso na exclusão, atualizar `loreFiles` sem recarregar a página
  - Em caso de falha na exclusão, exibir mensagem de erro
- [ ] Testar manualmente no navegador
  - Acessar `/lore` com nenhum arquivo no `/content/` → lista aparece vazia ou mensagem “Nenhuma página encontrada”.
  - Acessar com alguns arquivos criados via API → verificar se `title` e `visibility` estão corretos.
  - Clicar em “Criar Nova Página” → redireciona para o formulário de criação.
  - Clicar em “Editar” de uma página existente → redireciona para a rota de edição.
  - Clicar em “Excluir” → confirmar remoção e visualizar atualização imediata da lista.

### User Story 3.4 - Frontend: LoreEditor (textarea + preview instantâneo via Goldmark no client)

- [ ] Configurar rotas no React Router
  - Definir rota `/lore/novo` para criação e `/lore/:slug/editar` para edição, ambas apontando para `LoreEditorPage.tsx.`
- [ ] Instalar biblioteca de parsing Markdown no frontend
  - Adicionar dependência (por exemplo, `marked` ou similar) via `npm install marked`.
  - Atualizar `package.json` e rodar `npm install`.
- [ ] Criar componente `LoreEditorPage.tsx`
  - Importar React (`useState`, `useEffect`) e hooks do React Router (`useParams`, `useNavigate`).
  - Definir estados locais:
    - `title: string`
    - `visibility: "master" | "player"`
    - `content: string`
    - `loading: boolean` (no modo edição)
    - `error: string | null`
- [ ] Carregar dados no modo edição
  - Se `slug` existir no parâmetro de rota, no `useEffect` inicial:
    - `loading = true`
    - Fazer `GET /api/lore/{slug}` → retorna `{ title, visibility, content }`
    - Preencher estados (`setTitle`, `setVisibility`, `setContent`)
    - `loading = false`
    - Se erro, preencher `error`.
- [ ] Construir formulário de edição/criação
  - Input de texto para Título (obrigatório).
  - Select ou radio buttons para Visibilidade (`master` / `player`).
  - Textarea grande para Conteúdo Markdown (campo obrigatório).
  - Botão “Salvar”:
    - Desabilitado se `title` ou `content` estiverem vazios.
    - Ao clicar:
      - Mostrar indicador de “salvando”.
      - Chamar `POST /api/lore` (modo criação) ou `PUT /api/lore/{slug}` (modo edição) com JSON `{ slug, title, visibility, content }`.
      - Se sucesso, navegar de volta para `/lore` e mostrar toast ou alerta de confirmação.
      - Se falha, exibir mensagem de erro.
- [ ] Implementar preview instantâneo lado cliente
  - Importar a biblioteca `marked` (ou equivalente).
  - Sempre que `content` mudar:
    - Converter Markdown em HTML: `const html = marked(content)`
    - Renderizar dentro de um `<div>` com `dangerouslySetInnerHTML={{ __html: html }}`
  - Colocar o `<textarea>` e o painel de preview lado a lado (layout responsivo, switch ou abas mobile).
- [ ] Estilização básica
  - No CSS (Tailwind ou vanilla), definir:
    - Área de edição com borda e padding.
    - Área de preview com fundo claro/escuro contrastante.
    - Rolagem interna caso o conteúdo seja grande.
    - Garantir que links e headings no preview renderizem corretamente.
- [ ] Teste manual no navegador
  - Abrir `/lore/novo`: verificar formulário vazio e painel de preview inicialmente em branco.
  - Digitar Markdown: ver live preview refletindo título, listas, negrito, links etc.
  - Tentar salvar sem preencher título/conteúdo → botão permanece desabilitado.
  - Preencher tudo e salvar → novo arquivo criado no backend; ao voltar para `/lore`, a lista aparece atualizada.
  - Abrir `/lore/{slug}/editar`: verificar carregamento dos campos e do preview; editar e salvar → alterações refletidas no preview e na listagem.

### User Story 3.5 - Função Go para converter Markdown → HTML (pasta /public/)

- [ ] Estruturar diretórios de saída
  - Verificar se existe o diretório `/public/` na raiz do projeto.
  - Se não existir, criar o diretório `/public/` e subdiretórios necessários (como `/public/media/`).
  - Limpar o conteúdo antigo de `/public/` antes de gerar novos arquivos, mas sem remover o diretório `.git/`.
- [ ] Carregar templates HTML
  - Em `/templates/`, manter arquivos como `layout.html` (HTML base com `<head>`, `<body>{{ .Content }}</body>`) e `page.html` (marcações específicas para o corpo da página).
  - Criar função `loadTemplates()` que faça `template.ParseFiles("./templates/layout.html", "./templates/page.html")`.
- [ ] Listar todos os arquivos Markdown
  - Reaproveitar a função de 3.1 (`ListLoreFiles()`) para obter os slugs.
  - Para cada slug, montar caminho `./content/{slug}.md`.
- [ ] Para cada arquivo Markdown
  - Ler todo o conteúdo do arquivo.
  - Separar frontmatter (entre as linhas `---`) do corpo Markdown.
  - Extrair `title` e `visibility` (embora o template possa usar apenas o `title`).
  - Converter o corpo Markdown em HTML usando Goldmark:
    ```go
    var buf bytes.Buffer
    goldmark.Convert([]byte(markdownBody), &buf)
    htmlBody := buf.String()
    ```
  - Gerar contexto para o template Go:
    ```go
    data := struct {
      Title     string
      Visibility string
      Content   template.HTML
    }{ page.Title, page.Visibility, template.HTML(htmlBody) }
    ```
  - Executar template e escrever o resultado em `./public/{slug}.html`.
- [ ] Copiar recursos estáticos
  - Se existir `/media/`, copiar todo seu conteúdo recursivamente para `/public/media/` (manter hierarquia de pastas).
  - Tratar arquivos de imagem, CSS ou JS adicionais conforme necessidade.
- [ ] Implementar função `BuildWiki() error` que engloba todos os passos acima:
  - `loadTemplates()`
  - `ensureDir("./public/")` e `cleanDir("./public/")`
  - `copyStatic("./media", "./public/media")`
  - Para cada slug em `ListLoreFiles()`, gerar HTML.
- [ ] Expor endpoint HTTP (opcional neste estágio) ou chamar `BuildWiki()` diretamente quando o Mestre clicar em “Publicar Wiki” (será detalhado no 3.6).
- [ ] Testar manualmente
  - Criar ou editar algumas páginas em `/content/`.
  - No Go, chamar `BuildWiki()` (por exemplo, via função `main.go` em modo dev).
  - Verificar em `/public/` se foram gerados arquivos `{slug}.html` com o HTML esperado.
  - Abrir `file://…/public/{slug}.html` no navegador para conferir layout e formatação.
- [ ] Escrever testes unitários em `backend/wiki_builder_test.go` (pode usar diretório temporário):
  - Criar dois arquivos Markdown de teste em pasta temporária.
  - Chamar `BuildWiki()` direcionado a essa pasta.
  - Verificar se os arquivos HTML correspondentes existem e contêm elementos `<h1>` ou parágrafos esperados.

### User Story 3.6 - Pipeline de Publicação: botão “Publicar Wiki” → git checkout gh-pages, sobrescrever /public, commit/push

- [ ] **Primeiro passo (3.6.1): adicionar botão no Frontend**
  - [ ] Localizar componente onde ficará o botão (por exemplo, em `SettingsPanel.tsx` ou `TopBar.tsx`).
  - [ ] Inserir botão com label “Publicar Wiki” e handler `onClick={handlePublish}`.
  - [ ] Implementar função `handlePublish()` que:
    - Desabilita o botão enquanto a requisição está em andamento.
    - Faz `fetch("/api/publish", { method: "POST" })` para chamar o endpoint de publicação.
    - Se a resposta for bem-sucedida, exibe uma notificação de sucesso (pode usar `toast` ou `alert`).
- [ ] **Segundo passo (3.6.2): criar handler publishHandler para POST /api/publish**
  - [ ] Criar arquivo `publish_handlers.go` (ou inserir em `handlers.go`) e definir função `publishHandler(w http.ResponseWriter, r *http.Request)`.
  - [ ] Registrar rota no `main.go`: `mux.HandleFunc("/api/publish", publishHandler)`
  - [ ] Dentro de `publishHandler`:
  - [ ] Chamar `BuildWiki()` e tratar erro:
    - Se falhar, retornar HTTP 500 com JSON `{ "success": false, "message": "Erro ao gerar wiki: ..." }`.
  - [ ] Abrir repositório local usando go-git:
    - `repo, err := git.PlainOpen(".")`
    - Capturar diretório de trabalho (`worktree, err := repo.Worktree()`).
  - [ ] Checkout no branch `gh-pages`:
    - Se o branch não existir, criar *branch orphan* `gh-pages`:
      - `repo.CreateBranch(&config.Branch{Name: "gh-pages", Merge: "refs/heads/gh-pages"})`
      - `worktree.Checkout(&git.CheckoutOptions{Branch: "refs/heads/gh-pages", Create: true, Force: true})`
    - Se existir, apenas:
      - `worktree.Checkout(&git.CheckoutOptions{Branch: "refs/heads/gh-pages", Force: true})`
  - [ ] Limpar conteúdo antigo do working dir:
    - Iterar todos os arquivos e pastas no diretório raiz (exceto .git/) e remover.
  - [ ] Copiar recursivamente tudo que está em `./public/` para o diretório raiz do repositório.
  - [ ] Adicionar todas as mudanças:
    - `worktree.AddWithOptions(&git.AddOptions{All: true})`
  - [ ] Fazer commit com mensagem baseada em timestamp:
    - Exemplo de mensagem: `"Atualiza wiki em 2025-06-17 20:45"` (usar `time.Now().Format("2006-01-02 15:04")`).
  - [ ] Push das alterações para `origin gh-pages`:
    - `repo.Push(&git.PushOptions{RemoteName: "origin", RefSpecs: []config.RefSpec{"refs/heads/gh-pages:refs/heads/gh-pages"}})`
    - Tratar erros de autenticação ou conexão e retorná-los como HTTP 500.
  - [ ] Retornar HTTP 200 com JSON `{ success: true, url: "https://<usuário>.github.io/<repo>/" }`.
  - [ ] Testar manualmente:
    - [ ] Usar `curl -X POST http://localhost:3000/api/publish`
    - [ ] Verificar saída do servidor e logs de go-git (checkout, commit, push).
    - [ ] Confirmar no GitHub Pages que o branch `gh-pages` foi atualizado e o site reflete as mudanças.
- [ ] **Terceiro passo (3.6.3): Integração Front→Back no Botão “Publicar Wiki”**
  - [ ] Adicionar Estado no Frontend
    - No componente que contém o botão “Publicar Wiki” (ex.: `SettingsPanel.tsx` ou `TopBar.tsx`), criar estados locais:
      - `isPublishing: boolean` – indica se a requisição está em andamento.
      - `publishError: string | null` – armazena mensagem de erro, se houver.
      - `publishUrl: string | null` – recebe a URL pública do wiki, se sucesso.
  - [ ] Implementar função `handlePublish()` dentro do componente:
  ```tsx
  async function handlePublish() {
    setIsPublishing(true);
    setPublishError(null);
    try {
      const resp = await fetch("/api/publish", { method: "POST" });
      const data = await resp.json();
      if (!resp.ok || !data.success) {
        throw new Error(data.message || "Falha desconhecida");
      }
      setPublishUrl(data.url);
    } catch (err: any) {
      setPublishError(err.message);
    } finally {
      setIsPublishing(false);
    }
  }
  ```
  - [ ] Ajustar o JSX do botão
  ```tsx
  <button
    disabled={isPublishing}
    onClick={handlePublish}
    className="btn-primary"
  >
    {isPublishing ? "Publicando..." : "Publicar Wiki"}
  </button>
  ```
  - [ ] Exibir Feedback Após a Ação
    - Se `publishError` não for `null`, renderizar um alerta de erro abaixo do botão (ex.: `<p className="text-red-500">{publishError}</p>`).
    - Se `publishUrl` estiver definido, renderizar mensagem de sucesso com link, ex.:
  
  ```tsx
  <p className="text-green-600">
    Wiki publicada com sucesso! Acesse:{" "}
    <a href={publishUrl} target="_blank" rel="noopener noreferrer">
      {publishUrl}
    </a>
  </p>
  ```
  - [ ] Testar Manualmente
    - Acessar a tela com o botão “Publicar Wiki”.
    - Clicar em “Publicar Wiki”: o botão deve mudar para “Publicando…” e ficar desabilitado.
    - Se o backend responder sucesso, deve aparecer a mensagem com link clicável.
    - Se o backend retornar erro, deve aparecer a mensagem de erro apropriada.
    - Tentar rodar sem conexão de internet ou token inválido (caso tenha autenticação futura) para ver o erro.
- [ ] **Quarto passo (3.6.4): Testes Manuais Finais**
  - [ ] Cenário de Sucesso Completo
    - No frontend Mestre, acessar a tela com o botão “Publicar Wiki”.
    - Garantir que existam páginas Markdown em /content/ para publicar.
    - Clicar em “Publicar Wiki”:
      - O botão muda para “Publicando...” e fica desabilitado.
      - Ao finalizar, aparece a mensagem de sucesso com o link para o GitHub Pages.
      - Clicar no link e confirmar que o site está atualizado com as páginas geradas.
  - [ ] Cenário de Falha na Geração
    - Introduzir deliberadamente um erro no conteúdo Markdown (por ex. frontmatter mal formado).
    - Clicar em “Publicar Wiki”:
      - O botão volta ao estado normal.
      - Exibe mensagem de erro retornada pelo backend, indicando falha na geração.
  - [ ] Cenário de Falha no Git
    - Remover ou corromper as credenciais Git (ou desconectar da internet).
    - Clicar em “Publicar Wiki”:
      - Deve retornar erro ao tentar o `checkout`, `commit` ou `push`.
      - No frontend, exibir a mensagem de erro correspondente (por ex. “Erro ao publicar no GitHub: …”).
  - [ ] Verificar Idempotência Parcial
    - Publicar com sucesso uma vez.
    - Sem alterar nada, clicar em “Publicar Wiki” novamente.
    - Confirmar que não há duplicação de conteúdo e que o processo repete corretamente (`commit` novo com mesmo conteúdo atualizado no timestamp).
  - [ ] Limpeza Pós-Teste
    - Garantir que não existam artefatos temporários ou diretórios vazios remanescentes.
    - Verificar que o branch `gh-pages` local está em sincronia com `origin/gh-pages`.
    - Confirmar que o branch `main` continua intacto e não foi alterado pelo processo de publicação.

---

## Checklist Geral do Epic 3

### 3.1. Backend: Listar arquivos Markdown em /content/
- [ ] Criar função `ListLoreFiles()` que lê arquivos `.md` em `/content/`
- [ ] Definir struct `LoreFile` com campos `Slug`, `Title`, `Visibility`
- [ ] Implementar handler HTTP para `GET /api/lore` que retorna lista de arquivos
- [ ] Adicionar rota no `main.go` para o handler
- [ ] Testar manualmente com `curl` e arquivos de exemplo
- [ ] Criar testes unitários para `ListLoreFiles()` em `lore_dao_test.go`
### 3.2. Backend: Salvar/Editar arquivo .md com frontmatter
- [ ] Definir struct `LorePage` com campos `Slug`, `Title`, `Visibility`, `Content`
- [ ] Implementar função `SaveLoreFile(page LorePage) error` que salva arquivo `.md` com frontmatter
- [ ] Criar handler HTTP para `POST /api/lore` (nova página) e `PUT /api/lore/{slug}` (edição)
- [ ] Validar entrada no handler (slug, title, visibility)
- [ ] Registrar rotas no `main.go` para os handlers de criação e edição
- [ ] Testar manualmente via cURL ou Postman para criação e edição de páginas
- [ ] Escrever testes unitários em `lore_dao_test.go` para `SaveLoreFile()`
### 3.3. Frontend: LoreList com todas páginas e botão “Criar Nova”
- [ ] Configurar rota `/lore` no React Router
- [ ] Criar componente `LoreListPage.tsx` com estado para `loreFiles`, `loading` e `error`
- [ ] Fazer `fetch("/api/lore")` no `useEffect` para buscar dados
- [ ] Renderizar lista de páginas com título, visibilidade e botões de ação
- [ ] Adicionar botão “Criar Nova Página” que navega para `/lore/novo`
- [ ] Implementar ações de editar e excluir páginas com chamadas à API
- [ ] Testar manualmente no navegador para verificar carregamento, criação, edição e exclusão de páginas
### 3.4. Frontend: LoreEditor (textarea + preview instantâneo via Goldmark no client)
- [ ] Configurar rotas `/lore/novo` e `/lore/:slug/editar` no React Router
- [ ] Criar componente `LoreEditorPage.tsx` com estados para `title`, `visibility`, `content`, `loading` e `error`
- [ ] Carregar dados no modo edição usando `useParams` e `useEffect`
- [ ] Construir formulário com inputs para título, visibilidade e textarea para conteúdo Markdown
- [ ] Implementar preview instantâneo usando biblioteca de parsing Markdown (ex.: `marked`)
- [ ] Adicionar botão “Salvar” que chama API para criar ou editar página
- [ ] Testar manualmente no navegador para verificar criação, edição e preview de páginas Markdown
### 3.5. Função Go para converter Markdown → HTML (pasta /public/)
- [ ] Estruturar diretório `/public/` e garantir que está limpo antes de gerar novos arquivos
- [ ] Carregar templates HTML de `/templates/` (ex.: `layout.html`, `page.html`)
- [ ] Listar arquivos Markdown em `/content/` usando `ListLoreFiles()`
- [ ] Para cada arquivo, ler conteúdo, extrair frontmatter e converter Markdown em HTML usando Goldmark
- [ ] Gerar HTML usando templates Go e escrever em `/public/{slug}.html`
- [ ] Copiar recursos estáticos de `/media/` para `/public/media/`
- [ ] Implementar função `BuildWiki()` que engloba todos os passos acima
- [ ] Expor endpoint HTTP opcional ou chamar `BuildWiki()` ao clicar em “Publicar Wiki”
- [ ] Testar manualmente a geração de HTML e verificar arquivos em `/public/`
### 3.6. Pipeline de Publicação: botão “Publicar Wiki” → git checkout gh-pages, sobrescrever /public, commit/push
- [ ] Adicionar botão “Publicar Wiki” no frontend com handler `handlePublish()`
- [ ] Criar handler `publishHandler` para `POST /api/publish` que chama `BuildWiki()`
- [ ] Registrar rota no `main.go` para o handler de publicação
- [ ] Implementar lógica de checkout do branch `gh-pages` usando go-git
- [ ] Limpar conteúdo antigo do working dir e copiar arquivos de `/public/`
- [ ] Fazer commit com mensagem baseada em timestamp e push para `origin gh-pages`
- [ ] Retornar JSON com sucesso ou erro no handler de publicação
- [ ] Testar manualmente a publicação via cURL e verificar no GitHub Pages

---

## Observações Finais

### 1. Separar Branches por Story ou Subdivisão
* Quando for iniciar cada Story (1, 2, etc.), crie uma branch específica: por exemplo, `feature/1-model-campaign` para não misturar tarefas.
* Ao terminar cada Story, abra um Pull Request e revise antes de mesclar.

### 2. Boas Práticas de Commit e Documentação
* Em cada commit, descreva claramente o que foi feito: ex.: “Adiciona tabela campaigns no SQLite e inicializa DB”.
* Se algum passo gerar dúvida (ex.: erro de SQL X ou CORS Y), registre no README ou em comentários de código para referência futura.

### 3. Testes Manuais & Registro de Bugs
* Antes de marcar cada tarefa como concluída, realize testes manuais básicos (via cURL/Postman no backend; navegação e formulários no frontend).
* Se aparecer um bug, abra um Issue específico (ex.: “Erro 500 ao criar campaign sem description”) e trate antes de prosseguir.

### 4. Iteração Contínua
* Após concluir uma tarefa, revise o que foi feito e veja se há melhorias ou ajustes necessários.
* Se algo não estiver claro ou se precisar de ajustes, não hesite em abrir um Issue ou discutir com a equipe.
* Se necessário, crie uma nova branch para ajustes ou melhorias, seguindo o mesmo padrão de nomenclatura.

Com este checklist detalhado, temos uma visão clara de tudo que precisa ser feito para concluir o Epic 3.
