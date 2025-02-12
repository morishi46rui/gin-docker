.PHONY: up upd down build buildn prune bash

# コンテナの起動
up:
	docker compose up

# コンテナの起動 (バックグラウンド)
upd:
	docker compose up -d

# コンテナの停止
down:
	docker compose down

# すべてのコンテナをビルド（キャッシュあり）
build:
	docker compose build

# すべてのコンテナをキャッシュなしでビルド
buildn:
	docker compose build --no-cache

# 未使用のコンテナ・イメージ・ネットワーク・ボリュームを削除
prune:
	docker system prune -f && docker volume prune -f && docker network prune -f

# backendコンテナのbashに入る（Alpineベースならsh）
bash:
	docker compose exec backend sh
