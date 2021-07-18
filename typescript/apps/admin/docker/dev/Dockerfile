FROM node:16.5

WORKDIR /app

ENV NUXT_HOST 0.0.0.0

COPY package.json yarn.lock ./

RUN yarn install

COPY . .

EXPOSE 3000

CMD ["yarn", "run", "dev"]
