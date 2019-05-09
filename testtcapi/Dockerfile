FROM ruby:2.6.3-stretch

RUN apt-get update
RUN apt-get -q -y install vim

RUN gem install google-protobuf
RUN gem install grpc
RUN gem install grpc-tools

WORKDIR /usr/local/testtcapi

COPY run_all.rb run_all.rb
COPY entrypoint.sh entrypoint.sh
COPY tests/ tests/
COPY lib/ lib/

CMD [ "bash", "entrypoint.sh" ]
