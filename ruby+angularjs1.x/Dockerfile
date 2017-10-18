FROM ruby:2.4-alpine3.6
MAINTAINER Jiang Wu "masterwujiang@gmail.com"

WORKDIR /myapp
COPY Gemfile /myapp/Gemfile
COPY Gemfile.lock /myapp/Gemfile.lock

RUN set -ex \
	\
	&& apk add --no-cache --virtual .ruby-builddeps \
		postgresql-dev \
		musl-dev \
    gcc \
    make \
  \
  && bundle \
	&& runDeps="$( \
  	scanelf --needed --nobanner --format '%n#p' --recursive /usr/local \
			| tr ',' '\n' \
			| sort -u \
			| awk 'system("[ -e /usr/local/lib/" $1 " ]") == 0 { next } { print "so:" $1 }' \
	)" \
	&& apk add --virtual .ruby-rundeps $runDeps \
	&& apk del .ruby-builddeps

CMD ["puma", "-e", "production", "-p", "9292"]
EXPOSE 9292
