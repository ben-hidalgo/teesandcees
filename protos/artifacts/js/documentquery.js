/**
 * @fileoverview
 * @enhanceable
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.app.DocumentQuery');

goog.require('jspb.Message');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.app.DocumentQuery = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.app.DocumentQuery, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.app.DocumentQuery.displayName = 'proto.app.DocumentQuery';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.app.DocumentQuery.prototype.toObject = function(opt_includeInstance) {
  return proto.app.DocumentQuery.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.app.DocumentQuery} msg The msg instance to transform.
 * @return {!Object}
 */
proto.app.DocumentQuery.toObject = function(includeInstance, msg) {
  var f, obj = {
    page: msg.getPage(),
    size: msg.getSize()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Creates a deep clone of this proto. No data is shared with the original.
 * @return {!proto.app.DocumentQuery} The clone.
 */
proto.app.DocumentQuery.prototype.cloneMessage = function() {
  return /** @type {!proto.app.DocumentQuery} */ (jspb.Message.cloneMessage(this));
};


/**
 * optional int64 page = 1;
 * @return {number}
 */
proto.app.DocumentQuery.prototype.getPage = function() {
  return /** @type {number} */ (jspb.Message.getFieldProto3(this, 1, 0));
};


/** @param {number} value  */
proto.app.DocumentQuery.prototype.setPage = function(value) {
  jspb.Message.setField(this, 1, value);
};


/**
 * optional int64 size = 2;
 * @return {number}
 */
proto.app.DocumentQuery.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldProto3(this, 2, 0));
};


/** @param {number} value  */
proto.app.DocumentQuery.prototype.setSize = function(value) {
  jspb.Message.setField(this, 2, value);
};


