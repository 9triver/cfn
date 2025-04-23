package com.github9triver.cfn.grpc;

import com.github9triver.cfn.manager.LocalResourceManager;
import com.github9triver.cfn.model.dto.ServerAddress;
import com.github9triver.cfn.proto.Common;
import com.github9triver.cfn.proto.data.FunctionServiceGrpc;
import com.github9triver.cfn.proto.data.Functions;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.stub.StreamObserver;
import org.lognet.springboot.grpc.GRpcService;
import org.springframework.beans.factory.annotation.Autowired;

@GRpcService
public class FunctionService extends FunctionServiceGrpc.FunctionServiceImplBase {

    private LocalResourceManager localResourceManager;

    @Autowired
    public void setLocalResourceManager(LocalResourceManager localResourceManager) {
        this.localResourceManager = localResourceManager;
    }

    @Override
    public void deployPyFunc(Functions.AppendPyFunc request, StreamObserver<Common.ServiceReplay> responseObserver) {
        System.out.println("Deploying PyFunc");

        ServerAddress address = localResourceManager.requestResources(request.getResource());

        // 创建 gRPC 通道
        ManagedChannel channel = ManagedChannelBuilder.forAddress(address.getHost(), address.getPort())
                .usePlaintext() // 如果没有 TLS 就加这个
                .build();

        // 创建 Stub（同步）
        FunctionServiceGrpc.FunctionServiceBlockingStub stub = FunctionServiceGrpc.newBlockingStub(channel);

        // 发送请求
        //noinspection ResultOfMethodCallIgnored
        stub.deployPyFunc(request);

        // 关闭通道
        channel.shutdown();

        responseObserver.onNext(Common.ServiceReplay.newBuilder().build());
        responseObserver.onCompleted();
    }

    @Override
    public void removePyFunc(Functions.AppendPyFunc request, StreamObserver<Common.ServiceReplay> responseObserver) {
        super.removePyFunc(request, responseObserver);
    }
}
