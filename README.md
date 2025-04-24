# GoAPI (MC 架構)

這是一個使用 **Go + Gin + GORM + SQL Server** 建構的 RESTful API 專案，採用 MC（Model-Controller）架構，並整合 Swagger 做為 API 文件平台。

---

## 🚀 專案特色

- 🧱 **MVC 架構（無 View）**
- 🔄 自動建立資料庫與資料表（透過 GORM）
- 🔐 使用 `.env` 管理資料庫連線資訊
- 🧪 整合 [Swagger UI](http://localhost:8080/swagger/index.html)
- ⚙️ 支援 SQL Server（使用 `gorm.io/driver/sqlserver`）

---

## 📦 環境需求

- Go 1.24.2
- SQL Server（本地或遠端）
- Git

---

## 🔧 安裝與執行

```bash
# 1. 下載專案
git clone https://github.com/foylaou/goapi.git
cd goapi

# 2. 安裝相依套件
go mod tidy

# 3. 建立 .env 檔案
cp .env.example .env
# 並填入你的資料庫資訊
