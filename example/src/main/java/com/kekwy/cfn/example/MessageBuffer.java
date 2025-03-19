package com.kekwy.cfn.example;

public interface MessageBuffer {

    void sendTask(Task task);
    Task readTask();

    void sendState(NodeState state);
    NodeState readState();

}
