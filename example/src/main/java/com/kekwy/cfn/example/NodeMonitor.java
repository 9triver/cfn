package com.kekwy.cfn.example;


import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@RestController
public class NodeMonitor {

    private final Map<String, WorkerNode> workerNodeMap = new ConcurrentHashMap<>();
    private final Map<String, HeadNode> headNodeMap = new ConcurrentHashMap<>();

    public void registerWorkerNodes(Collection<WorkerNode> workerNodes) {
        for (WorkerNode workerNode : workerNodes) {
            workerNodeMap.put(workerNode.getId(), workerNode);
            workerNode.start(); // 启动工作结点
        }
    }

    public void registerHeadNode(HeadNode headNode) {
        headNodeMap.put(headNode.getId(), headNode);
        headNode.start();       // 启动头结点
    }

    @GetMapping("/head/{nodeId}")
    public HeadNodeState getHeadNodeState(@PathVariable("nodeId") String nodeId) {
        HeadNode headNode = headNodeMap.get(nodeId);
        if (headNode == null) {
            return null;
        }
        HeadNodeState headNodeState = new HeadNodeState();
        headNodeState.setId(headNode.getId());
        headNodeState.setWorkers(headNode.getWorkerNodes().stream().map(Node::getId).toList());
        return headNodeState;
    }

    @GetMapping("/worker/{nodeId}")
    public WorkerNodeState getWorkerNode(@PathVariable("nodeId") String nodeId) {
        WorkerNode workerNode = workerNodeMap.get(nodeId);
        if (workerNode == null) {
            return null;
        }
        WorkerNodeState workerNodeState = new WorkerNodeState();
        workerNodeState.setId(workerNode.getId());
        return workerNodeState;
    }

}
