FROM ruby:2.7.4-alpine

RUN mkdir /db
COPY Gemfile /db
COPY Gemfile.lock /db
WORKDIR /db

RUN apk add --no-cache build-base mariadb-dev && \
    gem install bundler:2.2.13 && \
    bundle install

ENTRYPOINT ["rake"]
