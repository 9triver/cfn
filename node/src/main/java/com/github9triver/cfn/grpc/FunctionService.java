package com.github9triver.cfn.grpc;

import com.github9triver.cfn.proto.Common;
import com.github9triver.cfn.proto.data.FunctionServiceGrpc;
import com.github9triver.cfn.proto.data.Functions;
import io.grpc.stub.StreamObserver;
import org.lognet.springboot.grpc.GRpcService;

@GRpcService
public class FunctionService extends FunctionServiceGrpc.FunctionServiceImplBase {

    @Override
    public void deployPyFunc(Functions.AppendPyFunc request, StreamObserver<Common.ServiceReplay> responseObserver) {
        System.out.println("Deploying PyFunc");
        responseObserver.onNext(Common.ServiceReplay.newBuilder().build());
        responseObserver.onCompleted();
    }
}
