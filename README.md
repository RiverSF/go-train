# go-train

version: go1.20.3

## OpenSpec Skill 使用说明

本项目已接入 [OpenSpec](https://github.com/Fission-AI/OpenSpec)（spec-driven development），在 Cursor 中通过 slash 命令驱动 AI 按「提案 → 实现 → 归档」流程协作。

> 首次使用前请**重启 Cursor**，确保 `/opsx:*` 命令生效。

### 前置条件

- Node.js ≥ 20.19.0
- 已全局安装：`npm install -g @fission-ai/openspec@latest`
- 项目已执行：`openspec init --tools cursor`

升级 CLI 后，在项目内执行 `openspec update` 可刷新 skill / 命令文件。

### 可用命令（Slash Commands）

| 命令 | 对应 Skill | 用途 |
|------|------------|------|
| `/opsx:explore` | `openspec-explore` | 探索需求、权衡方案；只读代码，不写实现 |
| `/opsx:propose <想法>` | `openspec-propose` | 一次性生成 proposal / design / specs / tasks |
| `/opsx:apply [change]` | `openspec-apply-change` | 按 `tasks.md` 逐项实现 |
| `/opsx:update [change]` | `openspec-update-change` | 修订已有规划文档，保持产物一致；不改代码 |
| `/opsx:sync [change]` | `openspec-sync-specs` | 将变更中的 delta specs 合并进主 specs |
| `/opsx:archive [change]` | `openspec-archive-change` | 实现完成后归档该 change |

也可在对话中直接提及 skill 名称（如「用 openspec-propose」），效果等同于对应 slash 命令。

### 推荐工作流

```text
1. /opsx:explore          # 可选：先对齐问题与方案
2. /opsx:propose "..."    # 生成变更提案与任务清单
3. /opsx:apply            # 按 tasks 实现
4. /opsx:archive          # 归档；必要时先 /opsx:sync
```

已有提案但需改计划时，用 `/opsx:update`；不必等归档也可单独 `/opsx:sync`。

### 目录结构

```text
openspec/
  config.yaml          # schema 与项目上下文配置
  specs/               # 主规格（系统行为真相）
  changes/             # 进行中的变更
    <change-name>/
      proposal.md
      design.md
      tasks.md
      specs/           # 相对主 specs 的增量（delta）
    archive/           # 已归档变更

.cursor/
  commands/opsx-*.md   # Cursor slash 命令
  skills/openspec-*/   # 对应 Skill 定义
```

### 示例

```text
/opsx:explore
# → 讨论「要不要加缓存、放哪一层」

/opsx:propose add-redis-cache
# → 生成 openspec/changes/add-redis-cache/

/opsx:apply
# → 按 tasks.md 实现

/opsx:archive
# → 归档到 openspec/changes/archive/YYYY-MM-DD-add-redis-cache/
```

更多说明见 [OpenSpec 文档](https://github.com/Fission-AI/OpenSpec)。

## GitNexus Skill 使用说明

本项目已接入 [GitNexus](https://github.com/abhigyanpatwari/GitNexus)：将代码库索引为知识图谱，通过 MCP 工具与 Agent Skills 供 Cursor 做架构探索、影响分析、调试与重构。

> 首次配置后请**重启 Cursor**，确保 GitNexus MCP 生效。索引过期时在项目根目录执行 `gitnexus analyze`（或 `node .gitnexus/run.cjs analyze`）。

### 前置条件

- Node.js ≥ 22（推荐）；npm 11 建议全局安装：`npm install -g gitnexus@latest`
- 已执行：`gitnexus setup -c cursor`（写入 MCP 与 Cursor skills）
- 本仓库已索引：`gitnexus analyze`

### 可用 Skills

| 场景 | Skill | 用途 |
|------|-------|------|
| 理解架构 /「X 怎么工作？」 | `gitnexus-exploring` | 按集群与执行流浏览代码 |
| 影响面 /「改 X 会炸什么？」 | `gitnexus-impact-analysis` | 调用链与 blast radius |
| 排错 /「为什么 X 失败？」 | `gitnexus-debugging` | 沿执行流追踪问题 |
| 重命名 / 抽取 / 拆分 | `gitnexus-refactoring` | 基于图谱的安全重构 |
| 工具与 schema 参考 | `gitnexus-guide` | MCP tools / resources 速查 |
| 索引 / 状态 / 清理 / wiki | `gitnexus-cli` | CLI 运维命令 |

对话中直接说「用 gitnexus-impact-analysis」或描述对应场景，Agent 会读取对应 skill 并走其工作流。

项目内 skill 文件：`.claude/skills/gitnexus/`；Cursor 全局副本：`~/.cursor/skills/gitnexus-*`。

### 推荐工作流

```text
1. 读 gitnexus://repo/go-train/context   # 确认索引新鲜
2. 按任务选 skill（见上表）
3. 改符号前先 impact；提交前先 detect_changes
4. 大改后：gitnexus analyze              # 刷新索引
```

### 常用 MCP Resources

| Resource | 用途 |
|----------|------|
| `gitnexus://repo/go-train/context` | 概览与索引是否过期 |
| `gitnexus://repo/go-train/clusters` | 功能集群 |
| `gitnexus://repo/go-train/processes` | 执行流列表 |
| `gitnexus://repo/go-train/process/{name}` | 单条执行流逐步追踪 |

### 常用 CLI

```text
gitnexus analyze    # 索引 / 增量更新
gitnexus status     # 查看索引是否与当前 commit 一致
gitnexus clean      # 清除本仓库索引
gitnexus doctor     # 诊断 FTS / 扩展等问题
```

更完整的 Agent 约束见 `AGENTS.md` / `CLAUDE.md`（`<!-- gitnexus:* -->` 区块）。
