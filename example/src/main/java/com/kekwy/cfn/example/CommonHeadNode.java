package com.kekwy.cfn.example;

import java.util.*;

public class CommonHeadNode implements HeadNode {

    /**
     * 暂时采用 uuid
     */
    private final String id = UUID.randomUUID().toString();

    /**
     * 用于工作结点的自增编号
     */
    private int index = 0;

    private final Map<String, WorkerNode> workers = new HashMap<>();

    private final Map<String, HeadNode> neighbors = new HashMap<>();



    @Override
    public void addWorkerNode(WorkerNode workerNode) {
        // 为 worker node 分配 id
        workerNode.setId(id + "." + index++);
        workers.put(workerNode.getId(), workerNode);
    }

    @Override
    public Collection<WorkerNode> getWorkerNodes() {
        return workers.values();
    }

    @Override
    public void addNeighborNode(HeadNode headNode) {
        neighbors.put(headNode.getId(), headNode);
    }

    @Override
    public String getId() {
        return id;
    }
}
