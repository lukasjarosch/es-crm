/**
 * @fileoverview gRPC-Web generated client stub for crm.eventstore.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var crm_eventstore_v1_event_store_pb = require('../../../crm/eventstore/v1/event_store_pb.js')
const proto = {};
proto.crm = {};
proto.crm.eventstore = {};
proto.crm.eventstore.v1 = require('./event_store_api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.crm.eventstore.v1.EventStoreAPIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.crm.eventstore.v1.EventStoreAPIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.crm.eventstore.v1.GetEventsRequest,
 *   !proto.crm.eventstore.v1.GetEventsResponse>}
 */
const methodDescriptor_EventStoreAPI_GetEvents = new grpc.web.MethodDescriptor(
  '/crm.eventstore.v1.EventStoreAPI/GetEvents',
  grpc.web.MethodType.UNARY,
  proto.crm.eventstore.v1.GetEventsRequest,
  proto.crm.eventstore.v1.GetEventsResponse,
  /**
   * @param {!proto.crm.eventstore.v1.GetEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.crm.eventstore.v1.GetEventsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.crm.eventstore.v1.GetEventsRequest,
 *   !proto.crm.eventstore.v1.GetEventsResponse>}
 */
const methodInfo_EventStoreAPI_GetEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.crm.eventstore.v1.GetEventsResponse,
  /**
   * @param {!proto.crm.eventstore.v1.GetEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.crm.eventstore.v1.GetEventsResponse.deserializeBinary
);


/**
 * @param {!proto.crm.eventstore.v1.GetEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.crm.eventstore.v1.GetEventsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.crm.eventstore.v1.GetEventsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.crm.eventstore.v1.EventStoreAPIClient.prototype.getEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/crm.eventstore.v1.EventStoreAPI/GetEvents',
      request,
      metadata || {},
      methodDescriptor_EventStoreAPI_GetEvents,
      callback);
};


/**
 * @param {!proto.crm.eventstore.v1.GetEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.crm.eventstore.v1.GetEventsResponse>}
 *     A native promise that resolves to the response
 */
proto.crm.eventstore.v1.EventStoreAPIPromiseClient.prototype.getEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/crm.eventstore.v1.EventStoreAPI/GetEvents',
      request,
      metadata || {},
      methodDescriptor_EventStoreAPI_GetEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.crm.eventstore.v1.CreateEventRequest,
 *   !proto.crm.eventstore.v1.CreateEventResponse>}
 */
const methodDescriptor_EventStoreAPI_CreateEvent = new grpc.web.MethodDescriptor(
  '/crm.eventstore.v1.EventStoreAPI/CreateEvent',
  grpc.web.MethodType.UNARY,
  proto.crm.eventstore.v1.CreateEventRequest,
  proto.crm.eventstore.v1.CreateEventResponse,
  /**
   * @param {!proto.crm.eventstore.v1.CreateEventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.crm.eventstore.v1.CreateEventResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.crm.eventstore.v1.CreateEventRequest,
 *   !proto.crm.eventstore.v1.CreateEventResponse>}
 */
const methodInfo_EventStoreAPI_CreateEvent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.crm.eventstore.v1.CreateEventResponse,
  /**
   * @param {!proto.crm.eventstore.v1.CreateEventRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.crm.eventstore.v1.CreateEventResponse.deserializeBinary
);


/**
 * @param {!proto.crm.eventstore.v1.CreateEventRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.crm.eventstore.v1.CreateEventResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.crm.eventstore.v1.CreateEventResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.crm.eventstore.v1.EventStoreAPIClient.prototype.createEvent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/crm.eventstore.v1.EventStoreAPI/CreateEvent',
      request,
      metadata || {},
      methodDescriptor_EventStoreAPI_CreateEvent,
      callback);
};


/**
 * @param {!proto.crm.eventstore.v1.CreateEventRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.crm.eventstore.v1.CreateEventResponse>}
 *     A native promise that resolves to the response
 */
proto.crm.eventstore.v1.EventStoreAPIPromiseClient.prototype.createEvent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/crm.eventstore.v1.EventStoreAPI/CreateEvent',
      request,
      metadata || {},
      methodDescriptor_EventStoreAPI_CreateEvent);
};


module.exports = proto.crm.eventstore.v1;

