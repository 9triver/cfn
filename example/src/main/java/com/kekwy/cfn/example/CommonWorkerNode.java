package com.kekwy.cfn.example;

import lombok.NoArgsConstructor;

import java.util.*;

@NoArgsConstructor
public class CommonWorkerNode implements WorkerNode {

    private String id;

    private Resource resource;

    private final Map<String, Task<?>> tasks = new HashMap<>();

    public CommonWorkerNode(int cpu, int mem, Collection<String> tags) {
        resource = new Resource(cpu, mem, tags);
    }

    @Override
    public void setId(String id) {
        this.id = id;
    }

    @Override
    public String getId() {
        return id;
    }


    private final Queue<Task<?>> taskQueue = new LinkedList<>();

    @Override
    public void start() {
        new Thread(this::run).start();
    }

    private void run() {
        while (true) {
            synchronized (taskQueue) {
                while (!taskQueue.isEmpty()) {
                    Task<?> task = taskQueue.poll();
                    // TODO: 创建一个容器、独立的上下文执行该任务
                    // 暂时使用开启一个线程的方式实现
                    new Thread(() -> runTask(task)).start();
                }
                try {
                    taskQueue.wait();
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }
            }
        }
    }

    private void runTask(Task<?> task) {
        task.setState("执行中");
        task.run();
        task.setState("已完成");
        synchronized (tasks) {
            resource.recycle(task.getRequiredResource());
            tasks.remove(task.getTid());
        }
        // TODO: 返回执行结果
    }

    @Override
    public void submitTask(Task<?> task) {
        synchronized (taskQueue) {
            if (!resource.check(task.getRequiredResource())) {
                // TODO: 向 Head 返回任务提交失败信息
                return;
            }
            synchronized (tasks) {
                // 预分配资源
                resource.allocate(task.getRequiredResource());
                tasks.put(task.getTid(), task);
            }
            taskQueue.offer(task);
            taskQueue.notify();
        }
    }

    @Override
    public Resource getResource() {
        return resource;
    }

    @Override
    public Collection<Task<?>> getTasks() {
        return tasks.values();
    }
}
