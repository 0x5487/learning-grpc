///
//  Generated code. Do not modify.
//  source: helloworld/proto/helloworld.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

import 'dart:async' as $async;

import 'package:grpc/grpc.dart';

import 'helloworld.pb.dart';
export 'helloworld.pb.dart';

class GreeterClient extends Client {
  static final _$sayHello = new ClientMethod<HelloRequest, HelloReply>(
      '/helloworld.Greeter/SayHello',
      (HelloRequest value) => value.writeToBuffer(),
      (List<int> value) => new HelloReply.fromBuffer(value));

  GreeterClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseFuture<HelloReply> sayHello(HelloRequest request,
      {CallOptions options}) {
    final call = $createCall(
        _$sayHello, new $async.Stream.fromIterable([request]),
        options: options);
    return new ResponseFuture(call);
  }
}

abstract class GreeterServiceBase extends Service {
  String get $name => 'helloworld.Greeter';

  GreeterServiceBase() {
    $addMethod(new ServiceMethod<HelloRequest, HelloReply>(
        'SayHello',
        sayHello_Pre,
        false,
        false,
        (List<int> value) => new HelloRequest.fromBuffer(value),
        (HelloReply value) => value.writeToBuffer()));
  }

  $async.Future<HelloReply> sayHello_Pre(
      ServiceCall call, $async.Future request) async {
    return sayHello(call, await request);
  }

  $async.Future<HelloReply> sayHello(ServiceCall call, HelloRequest request);
}

class ChatClient extends Client {
  static final _$bidStream = new ClientMethod<BidStreamRequest, BidStreamReply>(
      '/helloworld.Chat/BidStream',
      (BidStreamRequest value) => value.writeToBuffer(),
      (List<int> value) => new BidStreamReply.fromBuffer(value));

  ChatClient(ClientChannel channel, {CallOptions options})
      : super(channel, options: options);

  ResponseStream<BidStreamReply> bidStream(
      $async.Stream<BidStreamRequest> request,
      {CallOptions options}) {
    final call = $createCall(_$bidStream, request, options: options);
    return new ResponseStream(call);
  }
}

abstract class ChatServiceBase extends Service {
  String get $name => 'helloworld.Chat';

  ChatServiceBase() {
    $addMethod(new ServiceMethod<BidStreamRequest, BidStreamReply>(
        'BidStream',
        bidStream,
        true,
        true,
        (List<int> value) => new BidStreamRequest.fromBuffer(value),
        (BidStreamReply value) => value.writeToBuffer()));
  }

  $async.Stream<BidStreamReply> bidStream(
      ServiceCall call, $async.Stream<BidStreamRequest> request);
}
