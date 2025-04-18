package com.github9triver.cfn;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.util.Map;
import java.util.Set;

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
//        HeadNode headNode1 = new CommonHeadNode();
//        headNode1.addWorkerNode(new CommonWorkerNode(1, 1024, Set.of()));
//        headNode1.addWorkerNode(new CommonWorkerNode(1, 512, Set.of("camera")));
//        headNode1.addWorkerNode(new CommonWorkerNode(2, 1024, Set.of()));
//
//        HeadNode headNode2 = new CommonHeadNode();
//        headNode2.addWorkerNode(new CommonWorkerNode(8, 4096, Set.of()));
//        headNode2.addWorkerNode(new CommonWorkerNode(16, 16384, Set.of()));
//
//        headNode1.addNeighborNode(headNode2);
//        headNode2.addNeighborNode(headNode1);
//
//        // 向监视器中注册结点，模拟带外感知
//        nodeMonitor.registerHeadNode(headNode1);
//        nodeMonitor.registerHeadNode(headNode2);
//        nodeMonitor.registerWorkerNodes(headNode1.getWorkerNodes());
//        nodeMonitor.registerWorkerNodes(headNode2.getWorkerNodes());
//
//        // 输出 headNode1 的 id 便于演示带外感知
//        System.out.println(headNode1.getId());
//
//        // 创建任务
//        Task<Integer> task = new Task<>(
//                (Map<String, Object> map) -> {
//                    System.out.println("任务开始执行...");
//                    int startTime = (int) (System.currentTimeMillis() / 1000);
//                    synchronized (Task.class) {
//                        try {
//                            Task.class.wait();
//                        } catch (InterruptedException e) {
//                            throw new RuntimeException(e);
//                        }
//                    }
//                    int endTime = (int) (System.currentTimeMillis() / 1000);
//                    return endTime - startTime;
//                },
//                null,
//                16, 512, Set.of()
//        );
//
//        // 提交任务
//        headNode1.submitTask(task);
//        System.out.println("任务已提交");
//
//        // 获取结果
//        System.out.println("任务用时：" + task.getResult() + "s");


    }
}
