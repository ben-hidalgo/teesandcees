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

      req = App::Document.new(name: 'Doc name')

      doc = stub.create_document(req)

      assert_equal('Doc name', doc.name)
      assert_not_nil(doc.uuid)
      assert_not_equal('', doc.uuid)
    end

end
