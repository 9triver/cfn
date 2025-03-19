package com.kekwy.cfn.example;

import java.util.Collection;

public interface HeadNode {

    void addWorkerNode(WorkerNode workerNode);

    Collection<WorkerNode> getWorkerNodes();

    void addNeighborNode(HeadNode headNode);

    String getId();
}
