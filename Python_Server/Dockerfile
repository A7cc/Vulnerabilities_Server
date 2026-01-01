FROM python:3.9-slim

WORKDIR /app

# 安装系统依赖
RUN apt-get update && apt-get install -y \
    build-essential \
    python3-dev \
    default-libmysqlclient-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

# 复制整个项目
COPY ./src .

# 升级pip
RUN pip install --upgrade pip \
    && pip install --no-cache-dir -r requirements.txt

# MySQL 依赖
RUN apt-get update && apt-get install -y \
    build-essential \
    default-libmysqlclient-dev \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 8081

CMD ["python", "manage.py", "runserver", "0.0.0.0:8081"]
