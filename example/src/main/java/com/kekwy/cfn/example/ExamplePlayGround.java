package com.kekwy.cfn.example;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
public class ExamplePlayGround implements CommandLineRunner {

    private NodeMonitor nodeMonitor;

    @Autowired
    public void setNodeMonitor(NodeMonitor nodeMonitor) {
        this.nodeMonitor = nodeMonitor;
    }

    @Override
    public void run(String... args) throws Exception {
        // 构建拓扑
        HeadNode headNode1 = new CommonHeadNode();
        headNode1.addWorkerNode(new CommonWorkerNode());
        headNode1.addWorkerNode(new CommonWorkerNode());
        headNode1.addWorkerNode(new CommonWorkerNode());

        HeadNode headNode2 = new CommonHeadNode();
        headNode2.addWorkerNode(new CommonWorkerNode());
        headNode2.addWorkerNode(new CommonWorkerNode());

        headNode1.addNeighborNode(headNode2);

        // 向监视器中注册结点，模拟带外感知
        nodeMonitor.registerHeadNode(headNode1);
        nodeMonitor.registerHeadNode(headNode2);
        nodeMonitor.registerWorkerNodes(headNode1.getWorkerNodes());
        nodeMonitor.registerWorkerNodes(headNode2.getWorkerNodes());

        // 创建任务


        // 提交任务


        // 获取结果


    }
}
