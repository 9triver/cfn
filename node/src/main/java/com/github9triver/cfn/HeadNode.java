package com.github9triver.cfn;

import java.util.Collection;

public interface HeadNode {

    void addWorkerNode(WorkerNode workerNode);

    Collection<WorkerNode> getWorkerNodes();

    void addNeighborNode(HeadNode headNode);

    Collection<HeadNode> getNeighborNodes();

    String getId();

    void submitTask(Task<?> task);

    void start();

    Resource getTotalResource();
}
