# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: tcapi.proto for package 'app'

require 'grpc'
require_relative 'tcapi_pb'

module App
  module Tcapi
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'app.Tcapi'

      rpc :SayHello, HelloRequest, HelloReply
      rpc :CreateDocument, Document, Document
      rpc :GetDocuments, DocumentQuery, DocumentList
    end

    Stub = Service.rpc_stub_class
  end
end
