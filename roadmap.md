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
        2.1. Modelagem SQLite para Campaign
       * 2.2. Endpoints REST Campaign (GET, POST, PUT, DELETE)
       * 2.3. Front React: CampaignList e CampaignForm (lista + criação/edição)
       * 2.4. Modelagem SQLite para Session (vinculada à Campaign)
       * 2.5. Endpoints REST Session
       * 2.6. Front React: SessionList e SessionForm
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

## Milestone Atual: Módulo “Campaigns & Sessions”

**Objetivo**: Permitir CRUD de “Campaigns” e “Sessions” via API em Go com SQLite e interfaces React para listagem, criação, edição e remoção.

1. **Modelagem SQLite para Campaign**

- [X] **Definir a struct Campaign no backend (Go)**
  - [X] Escolher nome da struct: `type Campaign struct { ID int; Name string; Description string; CreatedAt time.Time; UpdatedAt time.Time }`
  - [X] Incluir campos obrigatórios (ID autoincremental, nome, descrição, timestamps).
  - [ ] Avaliar metadados adicionais que possam ser necessários (por exemplo, “Owner” ou “Status” se quisermos expandir depois).
- [X] **Configurar conexão com SQLite**
  - [X] Verificar se já existe uma função ou pacote para inicializar o banco (ex.: `db, err := sql.Open("sqlite3", "./campaign.db")`).
  - [X] Garantir que, no momento de startup do servidor Go, a conexão seja aberta e armazenada numa variável global ou num objeto de “service”.
    - `var DB *sql.DB` em `database.go`.
- [X] **Criar a tabela campaigns no banco SQLite**
  - [X] Escrever comando SQL para criar a tabela, por exemplo:
  
  ```sql
  CREATE TABLE IF NOT EXISTS campaigns (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  description TEXT,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
  );
  ```
  - [X] Adicionar função de “migrations” ou “initDB()” que execute esse `CREATE TABLE` no momento de inicializar o servidor.
  - [X] Testar manualmente: ao iniciar o servidor, conferir se o arquivo `.db` é criado e se a tabela `campaigns` existe (usar DB Browser for SQLite ou CLI).
- [X] **Implementar métodos de acesso ao banco (Data Access Layer)**
  - [X] Criar função `func CreateCampaign(c Campaign) (Campaign, error)` que insere no SQLite e retorna o registro criado com ID.
  - [X] Criar função `func GetAllCampaigns() ([]Campaign, error)` que faz `SELECT * FROM campaigns`.
  - [X] Criar função `func GetCampaignByID(id int) (Campaign, error)` que faz `SELECT * FROM campaigns WHERE id = ?`.
  - [X] Criar função `func UpdateCampaign(c Campaign) error` que faz `UPDATE campaigns SET name = ?, description = ?, updated_at = ? WHERE id = ?`.
  - [X] Criar função `func DeleteCampaign(id int) error` que faz `DELETE FROM campaigns WHERE id = ?`.
  - [X] Testar cada função isoladamente (por exemplo, criando um `main_test.go` ou fazendo log no console para verificar se as queries são executadas).

2. **Endpoints REST Campaign (GET, POST, PUT, DELETE)**

- [X] **Definir rotas no servidor Go**
  - [X] Escolher framework de roteamento (por exemplo, `net/http` + `gorilla/mux` ou `chi`).
    - Estamos usando o `net/http` padrão do Go.
  - [X] Registrar rota `GET /api/campaigns` → handler que chama `GetAllCampaigns()` e retorna JSON.
  - [X] Registrar rota `GET /api/campaigns/{id}` → handler que chama `GetCampaignByID(id)`.
  - [X] Registrar rota `POST /api/campaigns` → handler que faz “decoder” de JSON para struct `Campaign`, chama `CreateCampaign` e retorna JSON do campaign criado.
  - [X] Registrar rota `PUT /api/campaigns/{id}` → handler que decoda JSON, define campo ID pelo parâmetro de rota, chama `UpdateCampaign`.
  - [X] Registrar rota `DELETE /api/campaigns/{id}` → handler que chama `DeleteCampaign(id)`.
- [X] **Implementar handlers HTTP**
  - [X] Em cada handler, ler e validar os campos obrigatórios (por exemplo, `name` não pode estar vazio).
  - [X] Tratar erros de banco (se não encontrar ID, retornar 404; se falhar inserção, retornar 500).
  - [X] Definir respostas HTTP adequadas:
    - `GET /api/campaigns` → 200 com array JSON de campaigns.
    - `GET /api/campaigns/{id}` → 200 se encontrado, 404 se não encontrado.
    - `POST /api/campaigns` → 201 com JSON do campaign criado.
    - `PUT /api/campaigns/{id}` → 200 com JSON do campaign atualizado, 400 se payload inválido, 404 se não encontrar registro.
    - `DELETE /api/campaigns/{id}` → 204 sem corpo, 404 se não encontrado.
- [X] **Testar rotas via cURL ou Postman**
  - [X] Testar `curl -X GET http://localhost:3000/api/campaigns`
  - [X] Testar `curl -X POST http://localhost:3000/api/campaigns -d '{"name":"Campanha 1","description":"Descrição"}' -H "Content-Type: application/json"`
  - [X] Testar `curl -X GET http://localhost:3000/api/campaigns/1`
  - [X] Testar `curl -X PUT http://localhost:3000/api/campaigns/1 -d '{"name":"Novo Nome","description":"Nova descrição"}' -H "Content-Type: application/json"`
  - [X] Testar `curl -X DELETE http://localhost:3000/api/campaigns/1`
  - [X] Verificar códigos HTTP e payloads de resposta. Anotar qualquer bug ou comportamento inesperado (por exemplo, erro 500).
    - Tudo testado via Postman, retornando os códigos corretos e payloads esperados.

3. **Front React: CampaignList e CampaignForm (lista + criação/edição)**
- [X] **Definir rotas de navegação no React (React Router)**
  - [X] Instalar e configurar `react-router-dom`.
  - [X] Criar página `CampaignListPage` (por exemplo, `/campanhas`).
  - [X] Criar página `CampaignFormPage` para criação e edição (por exemplo, `/campanhas/novo` e `/campanhas/:id/editar`).
  - [X] Ajustar `<App />` para renderizar esses componentes conforme a rota.
- [X] **Implementar CampaignList (para a página de listagem)**
  - [X] Criar componente `CampaignList.tsx` que serve de container: recupera array de campaigns da API ao montar (usando `useEffect`).
  - [X] Fazer fetch em `GET /api/campaigns` e armazenar no estado (por exemplo, `const [campaigns, setCampaigns] = useState<Campaign[]>([])`).
  - [X] Exibir lista de campaigns em tabela ou cards: cada linha deve mostrar nome e breve descrição, e botões “Editar” e “Excluir”.
  - [X] Botão “Nova Campanha” que leva à rota `/campanhas/novo`.
  - [X] Botão “Excluir” dispara `DELETE /api/campaigns/{id}`, em seguida refaz a listagem ou remove item do array local.
  - [X] Botão “Editar” leva à rota `/campanhas/{id}/editar`.
- [X] **Implementar CampaignForm (para criar ou editar)**
  - [X] Criar componente `CampaignForm.tsx` que recebe prop opcional `id` (caso edição) ou usa “novo”.
  - [X] Se houver `id`, fazer fetch em `GET /api/campaigns/{id}` e preencher campos (nome, descrição).
  - [X] Campos do formulário:
    - Input de texto para “Nome” (obrigatório).
    - Textarea para “Descrição” (opcional).
    - Outros campos de metadados (se decidirmos incluir, por exemplo, “Status”).
  - [X] Botão “Salvar”:
    - Se for “novo”, faz `POST /api/campaigns` com payload JSON.
    - Se for “edição”, faz `PUT /api/campaigns/{id}` com payload JSON.
    - Tratamento de feedback:
      - Enquanto aguarda resposta, mostrar spinner ou desabilitar botão.
      - Se sucesso, redirecionar para `/campanhas` e exibir mensagem “Campanha salva com sucesso”.
      - Se erro (nome vazio, falha no servidor), exibir mensagem de validação ou de erro genérico.
- [X] **Testar interações no navegador**
  - [X] Acessar `/campanhas` sem nenhuma campanha: lista vazia ou mensagem “Nenhuma campanha cadastrada”.
  - [X] Clicar em “Nova Campanha” → formulário aparece vazio. Preencher e salvar → retorna para listagem com a nova campanha visível.
  - [X] Clicar em “Editar” de uma campanha existente → formulário com dados carregados; alterar e salvar → listagem atualizada.
  - [X] Clicar em “Excluir” de uma campanha → confirmação (opcional) e remoção do item na lista.

4. **Modelagem SQLite para Session (vinculada à Campaign)**
- [X] **Definir a struct Session no backend (Go)**
  - [X] Estrutura básica: `type Session struct { ID int; CampaignID int; Title string; Date time.Time; Location string; Notes string; CreatedAt time.Time; UpdatedAt time.Time }`
  - [X] Verificar quais campos são obrigatórios (no mínimo, `CampaignID` e `Title`).
  - [X] Incluir relacionamento: `CampaignID` como campo estrangeiro que referencia `campaigns.id`.
- [X] **Criar a tabela sessions no banco SQLite**
  - [X] Escrever comando SQL para criar a tabela:
  
  ```sql
  CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    campaign_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    date DATETIME,
    location TEXT,
    notes TEXT,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (campaign_id) REFERENCES campaigns(id) ON DELETE CASCADE
  );
  ```
  - [X] Adicionar esse comando ao mesmo “migration” ou função `initDB()` do backend, logo após a criação da tabela `campaigns`.
  - [X] Testar manualmente: verificar no DB Browser se a tabela aparece com a referência correta (relacionamento).
- [X] **Implementar métodos de acesso ao banco (Data Access Layer) para Session**
  - [X] `func CreateSession(s Session) (Session, error)` → insere no SQLite retornando o session criado.
  - [X] `func GetSessionsByCampaign(campaignID int) ([]Session, error)` → lista todas as sessions de uma campanha específica (`SELECT * FROM sessions WHERE campaign_id = ?`).
  - [X] `func GetSessionByID(id int) (Session, error)` → recupera session por ID.
  - [X] `func UpdateSession(s Session) error` → faz `UPDATE sessions SET title=?, date=?, location=?, notes=?, updated_at=? WHERE id=?`.
  - [X] `func DeleteSession(id int) error` → faz `DELETE FROM sessions WHERE id = ?`.
  - [X] Testar cada função isoladamente (logs ou testes automatizados).

5. **Endpoints REST Session**
- [X] **Definir rotas no servidor Go para Session**
  - [X] `GET /api/campaigns/{campaignID}/sessions` → retorna todas as sessions daquela campanha (chama `GetSessionsByCampaign`).
  - [X] `GET /api/sessions/{id}` → retorna uma session pelo ID (chama `GetSessionByID`).
  - [X] `POST /api/campaigns/{campaignID}/sessions` → cria nova session em campanha específica.
  - [X] `PUT /api/sessions/{id}` → atualiza session existente.
  - [X] `DELETE /api/sessions/{id}` → remove session.
- [X] **Implementar handlers HTTP para Session**
  - [X] `GET /api/campaigns/{campaignID}/sessions`
    - Ler parâmetro `campaignID` da rota.
    - Chamar `GetSessionsByCampaign(campaignID)` e retornar JSON array.
    - Se `campaignID` inválido (não numérico), retornar 400; se não houver sessions, retornar array vazio.
  - [X] `GET /api/sessions/{id}`
    - Ler parâmetro `id`, chamar `GetSessionByID(id)`.
    - Se não encontrar, retornar 404; se encontrar, retornar JSON.
  - [X] `POST /api/campaigns/{campaignID}/sessions`
    - Ler `campaignID` da rota; decodificar JSON com campos da session (`title`, `date`, `location`, `notes`).
    - Definir `s.CampaignID = campaignID`; chamar `CreateSession(s)`.
    - Retornar 201 com JSON do registro criado.
  - [X] `PUT /api/sessions/{id}`
    - Ler `id` da rota; decodificar JSON para struct Session.
    - Garantir que o payload tenha `CampaignID` correto ou ignorar esse campo (mantém a campanha original).
    - Chamar `UpdateSession(s)`.
  - [X] `DELETE /api/sessions/{id}`
    - Ler `id` da rota; chamar `DeleteSession(id)`.
    - Retornar 204 se sucesso, 404 se não encontrado.
- [X] **Testar rotas de Session via cURL ou Postman**
  - [X] Criar uma campanha primeiro; depois:
    - `curl -X GET http://localhost:3000/api/campaigns/1/sessions`
    - `curl -X POST http://localhost:3000/api/campaigns/1/sessions -d '{"title":"Sessão 1","date":"2025-06-10","location":"Taverna","notes":"Notas livres"}' -H "Content-Type: application/json"`
    - `curl -X GET http://localhost:3000/api/sessions/1`
    - `curl -X PUT http://localhost:3000/api/sessions/1 -d '{"title":"Sessão 1 Atualizada","location":"Castelo"}' -H "Content-Type: application/json"`
    - `curl -X DELETE http://localhost:3000/api/sessions/1`
  - [X] Confirmar códigos de resposta e payloads.

6. **Front React: SessionList e SessionForm**
- [X] **Adicionar rotas e navegação para Sessions**
  - [X] Na configuração do `react-router-dom`, criar rota aninhada:
    - `/campanhas/:campaignID/sessoes` → renderiza `SessionListPage`.
    - `/campanhas/:campaignID/sessoes/novo` → renderiza `SessionFormPage` (criação).
    - `/campanhas/:campaignID/sessoes/:id/editar` → renderiza `SessionFormPage` (edição).
  - [X] Ajustar breadcrumbs ou link de volta para “Campanhas” quando estiver dentro das sessions.
- [X] **Implementar SessionList (para a página de listagem)**
  - [X] Criar componente `SessionList.tsx`:
    - Ler `campaignID` dos parâmetros de rota (usando `useParams`).
    - Fazer fetch em `GET /api/campaigns/{campaignID}/sessions` e armazenar em estado local.
    - Exibir lista de sessions em ordem cronológica (ordenar pelo campo `date`).
    - Cada linha deve mostrar `title`, `date` (formatada), `location`, e botões “Editar” e “Excluir”.
    - Botão “Nova Sessão” que leva para `/campanhas/{campaignID}/sessoes/novo`.
    - Botão “Excluir” dispara `DELETE /api/sessions/{id}`, e ao remover, atualiza lista localmente sem recarregar toda a página.
    - Botão “Editar” leva à rota `/campanhas/{campaignID}/sessoes/{id}/editar`.
- [X] **Implementar SessionForm (para criar ou editar)**
  - [X] Criar componente `SessionForm.tsx`:
    - Entrar em “modo criação” se não houver `id` nos parâmetros, ou “modo edição” se houver.
    - No modo edição, ao montar, fazer fetch em `GET /api/sessions/{id}` e preencher campos.
    - Campos do formulário:
      - Input para “Título” (obrigatório).
      - Input de data (`<input type="date">`) para “Data” (opcional).
      - Input de texto para “Local” (opcional).
      - Textarea para “Notas” (opcional).
    - Botão “Salvar”:
      - Se criação, faz `POST /api/campaigns/{campaignID}/sessions` com payload JSON.
      - Se edição, faz `PUT /api/sessions/{id}` com payload JSON.
      - Ao salvar com sucesso, redirecionar para `/campanhas/{campaignID}/sessoes` e exibir mensagem “Sessão salva com sucesso”.
      - Validações simples: `title` não pode estar vazio; se `date` informado, deve ser data válida.
- [X] **Testar interações no navegador (Modo Sessões)**
  - [X] Acessar `/campanhas/1/sessoes` sem nenhuma sessão → lista vazia ou mensagem “Nenhuma sessão cadastrada”.
  - [X] Clicar em “Nova Sessão” → formulário aparece em branco. Preencher e salvar → retorna para listagem com nova sessão.
  - [X] Clicar em “Editar” de uma sessão existente → formulário com dados carregados; editar e salvar → listagem atualizada.
  - [X] Clicar em “Excluir” de uma sessão → confirmação (opcional) e remoção imediata da lista.

# Checklist Geral do Epic 2
## 2.1. Modelagem SQLite para Campaign
- [X] Definir `struct Campaign` no Go
- [X] Configurar conexão SQLite no startup
- [X] Criar tabela `campaigns` via comando SQL
- [X] Implementar funções CRUD (`CreateCampaign`, `GetAllCampaigns`, `GetCampaignByID`, `UpdateCampaign`, `DeleteCampaign`)
## 2.2. Endpoints REST Campaign
- [X] Registrar rota `GET /api/campaigns`
- [X] Registrar rota `GET /api/campaigns/{id}`
- [X] Registrar rota `POST /api/campaigns`
- [X] Registrar rota `PUT /api/campaigns/{id}`
- [X] Registrar rota `DELETE /api/campaigns/{id}`
- [X] Implementar handlers com validações e status codes corretos
- [X] Testar cada rota (cURL/Postman)
## 2.3. Front React: CampaignList e CampaignForm
- [X] Configurar `react-router-dom` para `/campanhas`, `/campanhas/novo`, `/campanhas/:id/editar`
- [X] Implementar componente `CampaignList`:
  - [X] Fetch de `GET /api/campaigns`
  - [X] Exibir lista em tabela ou cards
  - [X] Botões “Editar” e “Excluir” para cada item
  - [X] Botão “Nova Campanha”
- [X] Implementar componente `CampaignForm`:
  - [X] Fetch de `GET /api/campaigns/{id}` se modo edição
  - [X] Inputs para “Nome” e “Descrição”
  - [X] Botão “Salvar” chamando POST ou PUT
  - [X] Feedback de loading, sucessos e erros
  - [X] Redirecionar e exibir mensagem ao salvar
## 2.4. Modelagem SQLite para Session
- [X] Definir `struct Session` (incluindo `CampaignID`)
- [X] Criar tabela `sessions` com foreign key para `campaigns`
- [X] Implementar funções CRUD (`CreateSession`, `GetSessionsByCampaign`, `GetSessionByID`, `UpdateSession`, `DeleteSession`)
## 2.5. Endpoints REST Session
- [X] Registrar rota `GET /api/campaigns/{campaignID}/sessions`
- [X] Registrar rota `GET /api/sessions/{id}`
- [X] Registrar rota `POST /api/campaigns/{campaignID}/sessions`
- [X] Registrar rota `PUT /api/sessions/{id}`
- [X] Registrar rota `DELETE /api/sessions/{id}`
- [X] Implementar handlers com validações e status codes corretos
- [X] Testar cada rota (cURL/Postman)
## 2.6. Front React: SessionList e SessionForm
- [X] Configurar rotas:
  - [X] `/campanhas/:campaignID/sessoes` (listar)
  - [X] `/campanhas/:campaignID/sessoes/novo` (criar)
  - [X] `/campanhas/:campaignID/sessoes/:id/editar` (editar)
- [X] Implementar componente `SessionList`:
  - [X] Fetch de `GET /api/campaigns/{campaignID}/sessions`
  - [X] Exibir lista ordenada por data
  - [X] Botões “Editar” e “Excluir” para cada item
  - [X] Botão “Nova Sessão”
- [X] Implementar componente `SessionForm`:
  - [X] Fetch de `GET /api/sessions/{id}` se modo edição
  - [X] Inputs para “Título”, “Data”, “Local” e “Notas”
  - [X] Botão “Salvar” chamando POST ou PUT
  - [X] Feedback de loading, sucessos e erros
  - [X] Redirecionar e exibir mensagem ao salvar

## Observações Finais

### 1. Separar Branches por Story ou Subdivisão
* Quando for iniciar cada Story (1, 2, etc.), crie uma branch específica: por exemplo, `feature/1-model-campaign` para não misturar tarefas.
* Ao terminar cada Story, abra um Pull Request e revisaremos juntos antes de mesclar.

### 2. Boas Práticas de Commit e Documentação
* Em cada commit, descreva claramente o que foi feito: ex.: “Adiciona tabela campaigns no SQLite e inicializa DB”.
* Se algum passo gerar dúvida (ex.: erro de SQL X ou CORS Y), registre no README ou em comentários de código para referência futura.

### 3. Testes Manuais & Registro de Bugs
* Antes de marcar cada tarefa como concluída, realize testes manuais básicos (via cURL/Postman no backend; navegação e formulários no frontend).
* Se aparecer um bug, abra um Issue específico (ex.: “Erro 500 ao criar campaign sem description”) e trate antes de prosseguir.

### 4. Iteração Contínua
* Assim que a listagem de Campaigns e Sessions estiver funcionando, podemos revisitar e criar testes unitários/metódicos para o Data Access Layer e handlers HTTP.
* No frontend, implementar feedback de carregamento (“Loading…”) nas chamadas assíncronas para melhorar UX.

Com este checklist detalhado, temos uma visão clara de tudo que precisa ser feito para concluir o Epic 2.
