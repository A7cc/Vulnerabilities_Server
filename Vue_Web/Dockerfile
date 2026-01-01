FROM node:18-alpine

# 设置工作目录
WORKDIR /app

# 拷贝项目文件
COPY . .

# 安装依赖
RUN npm install

# 暴露 dev 端口
EXPOSE 5173

# 默认命令：启动开发服务器
CMD ["npm", "run", "dev"]
