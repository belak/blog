FROM nginx:alpine
MAINTAINER Kaleb Elwert <belak@coded.io>

ENV HUGO_VERSION 0.29
ENV HUGO_ARCHIVE hugo_${HUGO_VERSION}_Linux-64bit.tar.gz

RUN mkdir /site
WORKDIR /site

RUN apk --no-cache add curl

# curl instead of ADD so we use the cache
RUN mkdir /usr/local/hugo \
  && curl -L https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_ARCHIVE} > /usr/local/${HUGO_ARCHIVE} \
  && tar xzf /usr/local/${HUGO_ARCHIVE} -C /usr/local/hugo \
  && ln -s /usr/local/hugo/hugo /usr/local/bin/hugo \
  && rm /usr/local/${HUGO_ARCHIVE}

# Add all our files
ADD . /site

RUN hugo --baseURL "https://coded.io" --appendPort=false --enableGitInfo --destination=/usr/share/nginx/html
