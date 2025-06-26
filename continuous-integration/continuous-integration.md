# Integração Contínua (CI) – Visão Geral

## Objetivo

O processo de integrar modificações no codebase de forma contínua e automatizada visa garantir **mais segurança e agilidade** no desenvolvimento. Com automações, conseguimos:

- Detectar erros e bugs mais cedo;
- Impedir que código problemático entre na base principal;
- Garantir qualidade, segurança e estabilidade do software.

## Exemplo Prático

Ao abrir um **Pull Request (PR)**, diversas verificações podem ser executadas automaticamente, como:

- Execução de testes automatizados;
- Análise de segurança;
- Verificação de código limpo (lint);
- Validação de padrões e boas práticas.

## Principais Processos da Integração Contínua

- ✅ **Execução de Testes Automatizados**
- 🔍 **Linter (Padronização de Código)**
- 📈 **Análise de Qualidade de Código**
- 🔐 **Verificação de Segurança**
- 📦 **Geração de Artefatos para Deploy**
- 📌 **Identificação da Próxima Versão (Versionamento Semântico)**
- 🏷️ **Criação de Tags e Releases**

## Status Checks

O CI pode impor **status checks obrigatórios**, ou seja, **impede o merge de um PR** até que ele passe por todos os processos definidos (testes, lint, etc.).

## Ferramentas Comuns de CI

- **Jenkins** – Popular e altamente configurável;
- **GitHub Actions** – Integração nativa com GitHub, baseada em arquivos YAML.

## Workflows

Um **workflow** é o conjunto de processos que define o pipeline de CI/CD. Pode haver múltiplos workflows por repositório. Cada workflow:

- É definido em arquivos \`.yml\` (localizados em \`.github/workflows/\`);
- Pode conter **um ou mais jobs**;
- É iniciado por **eventos** do GitHub ou **agendamentos** (cron).

### Exemplo de Estrutura de Workflow

Evento (on:push) -> Filtros (branch: master) -> Ambiente (runs-on> ubuntu) -> Ações (steps: - run: npm run prod)

## Actions

Uma **action** é a unidade de execução dentro de um step de um job. Pode ser:

- Escrita em **JavaScript**;
- Baseada em uma **imagem Docker**.

As actions realizam tarefas específicas dentro do workflow, como rodar testes, fazer deploy, analisar código, etc.
