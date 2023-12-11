# 使用 Node.js 的官方镜像
FROM node:18-alpine

# 设置工作目录为 /app
WORKDIR /app

# 复制 package.json 和 package-lock.json 到工作目录
COPY package*.json ./

# 安装项目依赖
RUN npm install

# 复制其他源代码文件到工作目录
COPY . .

# 构建你的Nuxt应用
RUN npm run build

# 容器对外暴露的端口号。在 Dockerfile 中的 EXPOSE 指令只是一个文档/标记用途，它不会实际映射端口。
EXPOSE 3000

# 将构建结果复制到/dist目录
# RUN cp -R .output/public/_nuxt /dist/


# CMD ["npm", "start"]
# 使用 Nuxt 开发服务器启动应用
CMD ["npm", "run", "dev"]