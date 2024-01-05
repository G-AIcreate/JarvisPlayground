FROM nginx:alpine

# COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/conf.d/* /etc/nginx/conf.d/


# 复制Nuxt.js生成的文件到Nginx容器
# COPY frontend-services/GJarvis-frontend/.output/public/_nuxt/ /usr/share/nginx/html/gjarvis-website/

# 如果你有自定义的Nginx配置，也复制它