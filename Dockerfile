FROM alpine
MAINTAINER Kaleb Elwert <belak@coded.io>

ENV HUGO_VERSION 0.29
ENV HUGO_ARCHIVE hugo_${HUGO_VERSION}_Linux-64bit.tar.gz
ENV HUGO_BINARY hugo_${HUGO_VERSION}_linux_amd64

RUN mkdir /site
WORKDIR /site

RUN apk --no-cache add curl

# curl instead of ADD so we use the cache
RUN curl -L https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_ARCHIVE} > /usr/local/${HUGO_BINARY}.tar.gz \
  && tar xzf /usr/local/${HUGO_BINARY}.tar.gz -C /usr/local/ \
  && ln -s /usr/local/${HUGO_BINARY}/${HUGO_BINARY} /usr/local/bin/hugo \
  && rm /usr/local/${HUGO_BINARY}.tar.gz

# Add all our files
ADD . /site

# for if we run hugo server, as is the default cmd
EXPOSE 1313

CMD hugo server --bind 0.0.0.0 -b "https://coded.io/" --appendPort=false --enableGitInfo --disableLiveReload
