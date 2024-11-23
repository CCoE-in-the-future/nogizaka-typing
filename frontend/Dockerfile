FROM node:latest

# アプリケーションのパスを指定
ENV APP_PATH=/frontend

# アプリケーションディレクトリを作成
RUN mkdir $APP_PATH

# 作業ディレクトリを設定
WORKDIR $APP_PATH

# package.json と package-lock.json をコピー
COPY package*.json ./

# 依存関係をインストール
RUN npm install

# ポート設定 (Next.js はデフォルトで 3000 ポートを使用)
EXPOSE 3000
