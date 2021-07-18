FROM node:14.17 as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
# TODO fix env setting
ENV API_URL https://api.auto.lme.jp
ENV NUXT_ENV_WS_URL wss://chat.auto.lme.jp/subscribe
RUN npm run build

FROM nginx:1.19.1 as production-stage
ARG nginx_conf_path=nginx/default.conf
COPY ${nginx_conf_path} /etc/nginx/conf.d/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
