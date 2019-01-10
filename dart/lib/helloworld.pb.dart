///
//  Generated code. Do not modify.
//  source: helloworld/proto/helloworld.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, override;

import 'package:protobuf/protobuf.dart' as $pb;

class HelloRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('HelloRequest', package: const $pb.PackageName('helloworld'))
    ..aOS(1, 'name')
    ..hasRequiredFields = false
  ;

  HelloRequest() : super();
  HelloRequest.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  HelloRequest.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  HelloRequest clone() => new HelloRequest()..mergeFromMessage(this);
  HelloRequest copyWith(void Function(HelloRequest) updates) => super.copyWith((message) => updates(message as HelloRequest));
  $pb.BuilderInfo get info_ => _i;
  static HelloRequest create() => new HelloRequest();
  static $pb.PbList<HelloRequest> createRepeated() => new $pb.PbList<HelloRequest>();
  static HelloRequest getDefault() => _defaultInstance ??= create()..freeze();
  static HelloRequest _defaultInstance;
  static void $checkItem(HelloRequest v) {
    if (v is! HelloRequest) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get name => $_getS(0, '');
  set name(String v) { $_setString(0, v); }
  bool hasName() => $_has(0);
  void clearName() => clearField(1);
}

class HelloReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('HelloReply', package: const $pb.PackageName('helloworld'))
    ..aOS(1, 'message')
    ..hasRequiredFields = false
  ;

  HelloReply() : super();
  HelloReply.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  HelloReply.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  HelloReply clone() => new HelloReply()..mergeFromMessage(this);
  HelloReply copyWith(void Function(HelloReply) updates) => super.copyWith((message) => updates(message as HelloReply));
  $pb.BuilderInfo get info_ => _i;
  static HelloReply create() => new HelloReply();
  static $pb.PbList<HelloReply> createRepeated() => new $pb.PbList<HelloReply>();
  static HelloReply getDefault() => _defaultInstance ??= create()..freeze();
  static HelloReply _defaultInstance;
  static void $checkItem(HelloReply v) {
    if (v is! HelloReply) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get message => $_getS(0, '');
  set message(String v) { $_setString(0, v); }
  bool hasMessage() => $_has(0);
  void clearMessage() => clearField(1);
}

class BidStreamRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('BidStreamRequest', package: const $pb.PackageName('helloworld'))
    ..aOS(1, 'input')
    ..hasRequiredFields = false
  ;

  BidStreamRequest() : super();
  BidStreamRequest.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  BidStreamRequest.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  BidStreamRequest clone() => new BidStreamRequest()..mergeFromMessage(this);
  BidStreamRequest copyWith(void Function(BidStreamRequest) updates) => super.copyWith((message) => updates(message as BidStreamRequest));
  $pb.BuilderInfo get info_ => _i;
  static BidStreamRequest create() => new BidStreamRequest();
  static $pb.PbList<BidStreamRequest> createRepeated() => new $pb.PbList<BidStreamRequest>();
  static BidStreamRequest getDefault() => _defaultInstance ??= create()..freeze();
  static BidStreamRequest _defaultInstance;
  static void $checkItem(BidStreamRequest v) {
    if (v is! BidStreamRequest) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get input => $_getS(0, '');
  set input(String v) { $_setString(0, v); }
  bool hasInput() => $_has(0);
  void clearInput() => clearField(1);
}

class BidStreamReply extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('BidStreamReply', package: const $pb.PackageName('helloworld'))
    ..aOS(1, 'output')
    ..hasRequiredFields = false
  ;

  BidStreamReply() : super();
  BidStreamReply.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  BidStreamReply.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  BidStreamReply clone() => new BidStreamReply()..mergeFromMessage(this);
  BidStreamReply copyWith(void Function(BidStreamReply) updates) => super.copyWith((message) => updates(message as BidStreamReply));
  $pb.BuilderInfo get info_ => _i;
  static BidStreamReply create() => new BidStreamReply();
  static $pb.PbList<BidStreamReply> createRepeated() => new $pb.PbList<BidStreamReply>();
  static BidStreamReply getDefault() => _defaultInstance ??= create()..freeze();
  static BidStreamReply _defaultInstance;
  static void $checkItem(BidStreamReply v) {
    if (v is! BidStreamReply) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get output => $_getS(0, '');
  set output(String v) { $_setString(0, v); }
  bool hasOutput() => $_has(0);
  void clearOutput() => clearField(1);
}

