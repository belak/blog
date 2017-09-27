FROM nginx
MAINTAINER Kaleb Elwert <belak@coded.io>

ENV HUGO_VERSION 0.29
ENV HUGO_ARCHIVE hugo_${HUGO_VERSION}_Linux-64bit.tar.gz
ENV HUGO_BINARY hugo_${HUGO_VERSION}_linux_amd64

RUN mkdir /site
WORKDIR /site

# curl instead of ADD so we use the cache
RUN mkdir /usr/local/hugo \
  && curl -L https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_ARCHIVE} > /usr/local/hugo/${HUGO_BINARY}.tar.gz \
  && tar xzf /usr/local/${HUGO_BINARY}.tar.gz -C /usr/local/hugo \
  && ln -s /usr/local/hugo/hugo /usr/local/bin/hugo \
  && rm /usr/local/hugo/${HUGO_BINARY}.tar.gz

# Add all our files
ADD . /site

CMD hugo build --baseURL "https://coded.io" --appendPort=false --enableGitInfo --destination /usr/share/nginx/html
