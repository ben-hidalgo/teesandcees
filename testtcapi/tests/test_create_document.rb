#!/usr/bin/env ruby

require          'test/unit'
require          'grpc'
require_relative '../lib/tcapi_pb.rb'

class TestCreateDocument < Test::Unit::TestCase

    def test_create()
        assert_equal({}, {})
    end

    def test_main
      #stub = Tcapi::Stub.new('localhost:50051', :this_channel_is_insecure)
    end

end
