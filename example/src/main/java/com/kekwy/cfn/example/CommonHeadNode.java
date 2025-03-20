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
        // TODO: 解决环形调度的情况
    }

    @Override
    public Collection<HeadNode> getNeighborNodes() {
        return neighbors.values();
    }

    @Override
    public String getId() {
        return id;
    }

    private final Queue<Task<?>> taskQueue = new LinkedList<>();

    @Override
    public void submitTask(Task<?> task) {
        synchronized (taskQueue) {
            taskQueue.add(task);
            taskQueue.notify();
        }
        // microphone
    }

    private void run() {
        while (true) {
            synchronized (taskQueue) {
                while (!taskQueue.isEmpty()) {
                    Task<?> task = taskQueue.poll();
                    // 查看当前结点是否可以处理该任务
                    if (getTotalResource().check(task.getRequiredResource())) {
                        // TODO: 优化路由算法
                        boolean flag = false;
                        for (WorkerNode node : workers.values()) {
                            if (node.getResource().check(task.getRequiredResource())) {
                                node.submitTask(task);
                                flag = true;
                                break;
                            }
                        }
                        if (flag) {
                            continue;
                        }
                    }

                    boolean flag = false;
                    // TODO: 也存在资源总量满足但没有单一设备资源满足的情况
                    // 若不资源不满足则转发给下一跳 Head 结点，暂时采用暴力遍历的方式
                    for (HeadNode node : neighbors.values()) {
                        if (node.getTotalResource().check(task.getRequiredResource())) {
                            node.submitTask(task);
                            flag = true;
                            break;
                        }
                    }

                    if (flag) {
                        continue;
                    }

                    // TODO: 根据所需资源、网络开销等运行一个路由算法，决定需要将该任务发给哪一个 head 结点

                    // TODO: 若任务路由失败需要向前一跳反馈
                }
                try {
                    taskQueue.wait();
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
            }
        }
    }

    @Override
    public void start() {
        new Thread(this::run).start();
    }

    @Override
    public Resource getTotalResource() {
        // TODO: 修改为从本地元数据中获取，避免与所有 worker 频繁通信
        int cpu = 0;
        int mem = 0;
        Set<String> tags = new HashSet<>();
        for (WorkerNode node : workers.values()) {
            Resource resource = node.getResource();
            cpu += resource.getCpu();
            mem += resource.getMemory();
            tags.addAll(resource.getTags());
        }
        return new Resource(cpu, mem, tags);
    }


}
