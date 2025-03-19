package com.kekwy.cfn.example;

import java.util.Collection;

public interface WorkerNode extends Node {

    void start();

    void submitTask(Task<?> task);

    Resource getResource();
    Collection<Task<?>> getTasks();

}
