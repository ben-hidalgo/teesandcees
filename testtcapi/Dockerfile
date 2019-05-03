FROM ruby:2.6.3-stretch

WORKDIR /usr/local/testtcapi

COPY run_all.rb run_all.rb
COPY entrypoint.sh entrypoint.sh
COPY tests/ tests/

CMD [ "bash", "entrypoint.sh" ]
