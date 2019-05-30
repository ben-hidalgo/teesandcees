#!/usr/bin/env ruby

require          'test/unit'
require          'grpc'
require_relative '../lib/tcapi_pb.rb'
require_relative '../lib/tcapi_services_pb.rb'

class TestCreateDocument < Test::Unit::TestCase

    def test_create()
        assert_equal({}, {})
    end

    def test_main
      stub = App::Tcapi::Stub.new(ENV['STUB_ADDRESS'], :this_channel_is_insecure)

      req = App::HelloRequest.new(name: "My name")

      rep = stub.say_hello(req)
    end

end
